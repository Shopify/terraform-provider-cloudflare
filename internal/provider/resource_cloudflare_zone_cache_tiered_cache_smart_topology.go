package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZoneCacheTieredCacheSmartTopology() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareZoneCacheTieredCacheSmartTopologySchema(),
		CreateContext: resourceCloudflareZoneCacheTieredCacheSmartTopologyUpdate,
		ReadContext:   resourceCloudflareZoneCacheTieredCacheSmartTopologyRead,
		UpdateContext: resourceCloudflareZoneCacheTieredCacheSmartTopologyUpdate,
		DeleteContext: resourceCloudflareZoneCacheTieredCacheSmartTopologyDelete,
	}
}

func resourceCloudflareZoneCacheTieredCacheSmartTopologyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	tflog.Info(ctx, fmt.Sprintf("Reading Zone Cache TieredCacheSmartTopology %q", d.Id()))

	topology, err := client.TieredCacheSmartTopology(ctx, d.Id())

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			tflog.Info(ctx, fmt.Sprintf("Zone Cache TieredCacheSmartTopology for zone %q not found", d.Id()))
			d.SetId("")
			return nil
		} else {
			return diag.FromErr(fmt.Errorf("error reading cache TieredCacheSmartTopology for zone %q: %w", d.Id(), err))
		}
	}

	value := topology.Value

	if err := d.Set("value", value); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set value: %w", err))
	}

	return nil
}

func parseValue(value string) (cloudflare.TieredCacheSmartTopologyValue, error) {
	switch value {
	case "on":
		return cloudflare.TieredCacheSmartTopologyOn, nil
	case "off":
		return cloudflare.TieredCacheSmartTopologyOff, nil
	default:
		return cloudflare.TieredCacheSmartTopologyOff, errors.New("invalid value for TieredCacheSmartTopologyValue")
	}
}

func resourceCloudflareZoneCacheTieredCacheSmartTopologyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get("zone_id").(string)
	d.SetId(zoneID)

	value, err := parseValue(d.Get("value").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, fmt.Sprintf("Setting Zone Cache TieredCacheSmartTopology to struct: %+v for zone ID: %q", value, d.Id()))

	_, err = client.UpdateTieredCacheSmartTopology(ctx, d.Id(), cloudflare.TieredCacheSmartTopologyUpdateOptions{Value: value})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error setting cache TieredCacheSmartTopology for zone %q: %w", d.Id(), err))
	}

	return resourceCloudflareZoneCacheTieredCacheSmartTopologyRead(ctx, d, meta)
}

func resourceCloudflareZoneCacheTieredCacheSmartTopologyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Info(ctx, fmt.Sprintf("Deleting Zone Cache TieredCacheSmartTopology for zone ID: %q", d.Id()))
	d.Set("value", cloudflare.TieredCacheSmartTopologyOff)

	return resourceCloudflareZoneCacheTieredCacheSmartTopologyUpdate(ctx, d, meta)
}
