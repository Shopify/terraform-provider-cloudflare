package sdkv2provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudflareLoadBalancerPoolUpdate(t *testing.T) {
	tests := []struct {
		name           string
		before         map[string]interface{}
		changes        map[string]interface{}
		wantBody       string
		changedOrigins bool
	}{
		{
			name:     "description leaves externally managed origins untouched",
			changes:  map[string]interface{}{"description": "updated description"},
			wantBody: `{"description":"updated description"}`,
		},
		{
			name: "Terraform managed origins still update",
			changes: map[string]interface{}{
				"origins": []interface{}{map[string]interface{}{
					"name": "managed-origin", "address": "192.0.2.3", "enabled": false, "weight": 0.5,
				}},
			},
			wantBody:       `{"origins":[{"name":"managed-origin","address":"192.0.2.3","enabled":false,"weight":0.5,"header":{}}]}`,
			changedOrigins: true,
		},
		{
			name:     "disabled pool and zero minimum origins",
			changes:  map[string]interface{}{"enabled": false, "minimum_origins": 0},
			wantBody: `{"enabled":false,"minimum_origins":0}`,
		},
		{
			name: "clear optional strings",
			before: map[string]interface{}{
				"monitor": "monitor-id", "notification_email": "alerts@example.com",
			},
			changes: map[string]interface{}{
				"description": "", "monitor": "", "notification_email": "",
			},
			wantBody: `{"description":"","monitor":"","notification_email":""}`,
		},
		{
			name:     "rename pool",
			changes:  map[string]interface{}{"name": "renamed-pool"},
			wantBody: `{"name":"renamed-pool"}`,
		},
		{
			name:     "latitude zero includes unchanged longitude",
			before:   map[string]interface{}{"latitude": 12.5, "longitude": 45.5},
			changes:  map[string]interface{}{"latitude": 0.0},
			wantBody: `{"latitude":0,"longitude":45.5}`,
		},
		{
			name:     "longitude zero includes unchanged latitude",
			before:   map[string]interface{}{"latitude": 12.5, "longitude": 45.5},
			changes:  map[string]interface{}{"longitude": 0.0},
			wantBody: `{"latitude":12.5,"longitude":0}`,
		},
		{
			name:     "change check regions",
			before:   map[string]interface{}{"check_regions": []interface{}{"WNAM"}},
			changes:  map[string]interface{}{"check_regions": []interface{}{"WEU"}},
			wantBody: `{"check_regions":["WEU"]}`,
		},
		{
			name:     "clear check regions",
			before:   map[string]interface{}{"check_regions": []interface{}{"WNAM"}},
			changes:  map[string]interface{}{"check_regions": []interface{}{}},
			wantBody: `{"check_regions":[]}`,
		},
		{
			name: "clear load shedding percentages",
			before: map[string]interface{}{
				"load_shedding": []interface{}{map[string]interface{}{
					"default_percent": 10.0, "default_policy": "random", "session_percent": 20.0, "session_policy": "hash",
				}},
			},
			changes: map[string]interface{}{
				"load_shedding": []interface{}{map[string]interface{}{
					"default_percent": 0.0, "default_policy": "", "session_percent": 0.0, "session_policy": "",
				}},
			},
			wantBody: `{"load_shedding":{"default_percent":0,"session_percent":0}}`,
		},
		{
			name: "load shedding percentages without policies",
			changes: map[string]interface{}{
				"load_shedding": []interface{}{map[string]interface{}{"default_percent": 10.0}},
			},
			wantBody: `{"load_shedding":{"default_percent":10,"session_percent":0}}`,
		},
		{
			name: "load shedding with explicit policies",
			changes: map[string]interface{}{
				"load_shedding": []interface{}{map[string]interface{}{
					"default_percent": 10.0, "default_policy": "random", "session_percent": 20.0, "session_policy": "hash",
				}},
			},
			wantBody: `{"load_shedding":{"default_percent":10,"default_policy":"random","session_percent":20,"session_policy":"hash"}}`,
		},
		{
			name: "remove load shedding",
			before: map[string]interface{}{
				"load_shedding": []interface{}{map[string]interface{}{"default_percent": 10.0}},
			},
			changes:  map[string]interface{}{"load_shedding": []interface{}{}},
			wantBody: `{"load_shedding":null}`,
		},
		{
			name: "change origin steering",
			before: map[string]interface{}{
				"origin_steering": []interface{}{map[string]interface{}{"policy": "random"}},
			},
			changes: map[string]interface{}{
				"origin_steering": []interface{}{map[string]interface{}{"policy": "hash"}},
			},
			wantBody: `{"origin_steering":{"policy":"hash"}}`,
		},
		{
			name: "remove origin steering",
			before: map[string]interface{}{
				"origin_steering": []interface{}{map[string]interface{}{"policy": "hash"}},
			},
			changes:  map[string]interface{}{"origin_steering": []interface{}{}},
			wantBody: `{"origin_steering":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := testLoadBalancerPoolUpdateConfig()
			for key, value := range tt.before {
				before[key] = value
			}
			after := testLoadBalancerPoolUpdateConfig()
			for key, value := range before {
				after[key] = value
			}
			for key, value := range tt.changes {
				after[key] = value
			}

			// The remote origin changed after the plan was made. Its port is
			// deliberately absent from this version of the SDK's origin model.
			var remote map[string]interface{}
			require.NoError(t, json.Unmarshal([]byte(`{
				"id":"pool-id","name":"example-pool","description":"original description",
				"enabled":true,"minimum_origins":1,
				"created_on":"2024-01-01T00:00:00Z","modified_on":"2024-01-01T00:00:00Z",
				"origins":[{"name":"external-origin","address":"192.0.2.2","enabled":true,"weight":1,"port":30000}]
			}`), &remote))
			var methods []string
			client := testLoadBalancerPoolUpdateClient(t, func(request *http.Request) (*http.Response, error) {
				methods = append(methods, request.Method)
				assert.Equal(t, "/client/v4/accounts/account-id/load_balancers/pools/pool-id", request.URL.Path)
				switch len(methods) {
				case 1:
					assert.Equal(t, http.MethodPatch, request.Method)
					body, err := io.ReadAll(request.Body)
					require.NoError(t, err)
					assert.JSONEq(t, tt.wantBody, string(body))
					var patch map[string]interface{}
					require.NoError(t, json.Unmarshal(body, &patch))
					for key, value := range patch {
						remote[key] = value
					}
				case 2:
					assert.Equal(t, http.MethodGet, request.Method)
				default:
					return nil, fmt.Errorf("unexpected request: %s %s", request.Method, request.URL)
				}
				body, err := json.Marshal(map[string]interface{}{"success": true, "result": remote})
				require.NoError(t, err)
				return testLoadBalancerPoolUpdateResponse(http.StatusOK, string(body)), nil
			})

			state, diagnostics := testApplyLoadBalancerPoolUpdate(context.Background(), t, before, after, client)
			require.False(t, diagnostics.HasError(), "%+v", diagnostics)
			assert.Equal(t, []string{http.MethodPatch, http.MethodGet}, methods)
			require.NotNil(t, state)
			assert.Equal(t, "pool-id", state.ID)
			data := resourceCloudflareLoadBalancerPool().Data(state)
			origins := data.Get("origins").(*schema.Set).List()
			require.Len(t, origins, 1)
			if tt.changedOrigins {
				assert.Equal(t, "192.0.2.3", origins[0].(map[string]interface{})["address"])
			} else {
				assert.Equal(t, "192.0.2.2", origins[0].(map[string]interface{})["address"])
				remoteOrigins, err := json.Marshal(remote["origins"])
				require.NoError(t, err)
				assert.JSONEq(t, `[{"name":"external-origin","address":"192.0.2.2","enabled":true,"weight":1,"port":30000}]`, string(remoteOrigins))
			}
		})
	}
}

func TestCloudflareLoadBalancerPoolUpdateError(t *testing.T) {
	before := testLoadBalancerPoolUpdateConfig()
	after := testLoadBalancerPoolUpdateConfig()
	after["description"] = "updated description"
	requests := 0
	client := testLoadBalancerPoolUpdateClient(t, func(request *http.Request) (*http.Response, error) {
		requests++
		assert.Equal(t, http.MethodPatch, request.Method)
		return testLoadBalancerPoolUpdateResponse(http.StatusBadRequest, `{"success":false,"errors":[{"code":1000,"message":"invalid pool update"}]}`), nil
	})

	_, diagnostics := testApplyLoadBalancerPoolUpdate(context.Background(), t, before, after, client)
	require.True(t, diagnostics.HasError())
	require.Len(t, diagnostics, 1)
	assert.Contains(t, diagnostics[0].Summary, "error updating load balancer pool")
	assert.Contains(t, diagnostics[0].Summary, "invalid pool update")
	assert.Equal(t, 1, requests, "failed update must not refresh state")
}

func testLoadBalancerPoolUpdateConfig() map[string]interface{} {
	return map[string]interface{}{
		"account_id":  "account-id",
		"name":        "example-pool",
		"description": "original description",
		"origins": []interface{}{map[string]interface{}{
			"name": "planned-origin", "address": "192.0.2.1", "enabled": true, "weight": 1.0,
		}},
	}
}

func testApplyLoadBalancerPoolUpdate(ctx context.Context, t *testing.T, before, after map[string]interface{}, client *cloudflare.API) (*terraform.InstanceState, diag.Diagnostics) {
	t.Helper()
	resource := resourceCloudflareLoadBalancerPool()
	data := schema.TestResourceDataRaw(t, resource.Schema, before)
	data.SetId("pool-id")
	state := data.State()
	diff, err := resource.Diff(ctx, state, terraform.NewResourceConfigRaw(after), nil)
	require.NoError(t, err)
	require.NotNil(t, diff)
	return resource.Apply(ctx, state, diff, client)
}

func testLoadBalancerPoolUpdateClient(t *testing.T, transport loadBalancerPoolUpdateTransport) *cloudflare.API {
	t.Helper()
	client, err := cloudflare.NewWithAPIToken("test-token", cloudflare.HTTPClient(&http.Client{Transport: transport}), cloudflare.UsingRetryPolicy(0, 0, 0), cloudflare.UsingRateLimit(1000))
	require.NoError(t, err)
	return client
}

func testLoadBalancerPoolUpdateResponse(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body))}
}

type loadBalancerPoolUpdateTransport func(*http.Request) (*http.Response, error)

func (transport loadBalancerPoolUpdateTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := transport(request)
	if response != nil {
		response.Request = request
	}
	return response, err
}
