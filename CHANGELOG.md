## 3.17.0 (Unreleased)

[2.x to 3.x upgrade guide](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/guides/version-3-upgrade)

BREAKING CHANGES:

* resource/cloudflare_access_rule: `configuration` is now a `TypeList` instead of a `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_custom_ssl: `custom_ssl_options` is now a `TypeList` instead of `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_load_balancer: `fixed_response` is now a `TypeList` instead of a `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_load_balancer: fixed_response.status_code` is now a `TypeInt` instead of a `TypeString` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_record: `data` is now a `TypeList` instead of a `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_ruleset: rename `mitigation_expression` to `counting_expression` ([#1477](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1477))

NOTES:

* provider: Golang version has been upgraded to 1.17 ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* provider: HTTP user agent is now "terraform/:version terraform-plugin-sdk/:version terraform-provider-cloudflare/:version" ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* provider: Minimum Terraform core version is now 0.14 ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* provider: Update to cloudflare-go v0.22.0 ([#1184](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1184))
* provider: cloudflare-go has been upgraded to v0.25.0 ([#1236](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1236))
* provider: internally swapped to using `diag.Diagnostics` for CRUD return types and using `context.Context` passed in from the provider itself instead of instantiating our own in each operation ([#1592](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1592))
* provider: split schema definition from resource CRUD operations ([#1321](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1321))
* provider: swap internal logging mechanism to use `tflog` ([#1638](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1638))
* provider: terraform-plugin-sdk has been upgraded to 2.x ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* provider: updated internal package structure of repository ([#1636](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1636))
* resource/cloudflare_api_token: revert swap from TypeList to TypeSet due to broken migration ([#1455](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1455))
* resource/cloudflare_byo_ip_prefix: now requires an explicit `account_id` parameter instead of implicitly relying on `client.AccountID` ([#1563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1563))
* resource/cloudflare_healthcheck: `notification_suspended` and `notification_email_addresses` attributes are being deprecated in favour of `cloudflare_notification_policy` resource instead. ([#1529](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1529))
* resource/cloudflare_ip_list: no longer sets `client.AccountID` internally for resources ([#1563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1563))
* resource/cloudflare_magic_firewall_ruleset: no longer sets `client.AccountID` internally for resources ([#1563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1563))
* resource/cloudflare_origin_ca_certificate: `requested_validity` no longer decrements until the `expires_on` value but is now the amount of days the certificate was requested for. ([#1502](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1502))
* resource/cloudflare_static_route: no longer sets `client.AccountID` internally for resources ([#1563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1563))
* resource/cloudflare_worker_cron_trigger: now requires an explicit `account_id` parameter instead of implicitly relying on `client.AccountID` ([#1563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1563))

FEATURES:

* **New Data Source:** `cloudflare_access_identity_provider` ([#1300](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1300))
* **New Data Source:** `cloudflare_account_roles` ([#1238](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1238))
* **New Data Source:** `cloudflare_devices` ([#1453](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1453))
* **New Data Source:** `cloudflare_zone` ([#1213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1213))
* **New Resource:** `cloudflare_access_bookmark` ([#1539](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1539))
* **New Resource:** `cloudflare_access_keys_configuration` ([#1186](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1186))
* **New Resource:** `cloudflare_device_posture_integration` ([#1340](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1340))
* **New Resource:** `cloudflare_fallback_domain` ([#1356](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1356))
* **New Resource:** `cloudflare_gre_tunnel` ([#1423](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1423))
* **New Resource:** `cloudflare_ipsec_tunnel` ([#1404](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1404))
* **New Resource:** `cloudflare_split_tunnel` ([#1207](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1207))
* **New Resource:** `cloudflare_teams_account` ([#1173](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1173))
* **New Resource:** `cloudflare_teams_proxy_endpoint` ([#1517](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1517))
* **New Resource:** `cloudflare_teams_rule` ([#1173](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1173))
* **New Resource:** `cloudflare_tunnel_route` ([#1572](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1572))
* **New Resource:** `cloudflare_waiting_room_event` ([#1509](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1509))
* **New Resource:** `cloudflare_zone_cache_variants` ([#1444](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1444))

ENHANCEMENTS:

* cloudflare_ruleset: add support for "managed_challenge" action ([#1442](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1442))
* datasource/cloudflare_zones: allow filtering by account_id ([#1401](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1401))
* provider: add support for debugging via debuggers (like delve) ([#1217](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1217))
* provider: add the ability to configure a different hostname and base path for the API client ([#1270](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1270))
* resource/certificate_pack: adds `validation_errors` and `validation_records` with same format as custom hostnames. ([#1424](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1424))
* resource/cloudflare_access_application: Add service_auth_401_redirect field. ([#1540](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1540))
* resource/cloudflare_access_application: add bookmark type to apptypes ([#1343](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1343))
* resource/cloudflare_access_application: add support for 'SameSite' and 'HttpOnly' cookie attributes ([#1241](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1241))
* resource/cloudflare_access_application: add support for 'skip_interstitial' and 'logo_url' properties ([#1262](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1262))
* resource/cloudflare_access_application: add support for `app_launcher_visible` to the schema ([#1303](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1303))
* resource/cloudflare_access_group: add support for external evaluation as a new access group rule ([#1623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1623))
* resource/cloudflare_access_policy: add support for approval_required flag ([#1230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1230))
* resource/cloudflare_access_policy: add support for purpose justification and approvals ([#1199](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1199))
* resource/cloudflare_access_rule: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))
* resource/cloudflare_access_rule: add support for managed_challenge action ([#1457](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1457))
* resource/cloudflare_argo_tunnel: add `cname` as exported attribute ([#1259](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1259))
* resource/cloudflare_argo_tunnel: add `tunnel_token` support ([#1590](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1590))
* resource/cloudflare_certificate_pack: add support for new option (`wait_for_active_status`) to block creation until certificate pack is active ([#1567](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1567))
* resource/cloudflare_cloudflare_teams_rules: Add `check_session` and `add_headers` attributes to settings ([#1402](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1402))
* resource/cloudflare_cloudflare_teams_rules: Add `disable_download`, `disable_keyboard`, and `disable_upload` attributes to `BISOAdminControls` ([#1402](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1402))
* resource/cloudflare_custom_hostname: `settings.ciphers` is now a `TypeSet` internally to handle suppress ordering changes. Schema representation remains the same ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_custom_hostname: `settings` is now `Optional`/`Computed` to reflect the stricter schema validation introduced in terraform-plugin-sdk v2 ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_custom_hostname: `status` is now `Computed` as the value isn't managed by an end user ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_custom_hostname: add `settings.early_hints` to ssl schema ([#1286](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1286))
* resource/cloudflare_custom_hostname: adds support for custom_origin_sni ([#1482](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1482))
* resource/cloudflare_custom_pages: add support for managed_challenge action ([#1478](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1478))
* resource/cloudflare_custom_ssl: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))
* resource/cloudflare_device_policy_certificates: add support for device policy certificate settings ([#1467](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1467))
* resource/cloudflare_device_posture_rule: Add `expiration` to device posture rule ([#1585](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1585))
* resource/cloudflare_firewall_rule: add support for managed_challenge action ([#1378](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1378))
* resource/cloudflare_load_balancer: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))
* resource/cloudflare_load_balancer_monitor: added support for smtp, icmp_ping, and udp_icmp monitors ([#1371](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1371))
* resource/cloudflare_load_balancer_pool: add support for origin steering ([#1240](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1240))
* resource/cloudflare_logpush_job: add support for account-level logpush jobs ([#1311](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1311))
* resource/cloudflare_logpush_job: add support for managing `dns_logs` ([#1400](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1400))
* resource/cloudflare_logpush_job: add support for managing `network_analytics_logs` ([#1627](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1627))
* resource/cloudflare_logpush_job: add support for specifying `frequency` ([#1634](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1634))
* resource/cloudflare_logpush_job: allow r2 logpush destinations without ownership validation ([#1597](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1597))
* resource/cloudflare_logpush_ownership_challenge: add support for account-level logpush ownership challenges ([#1311](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1311))
* resource/cloudflare_notification_policy: Add `slo` to notification policy filters ([#1573](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1573))
* resource/cloudflare_page_rule: add support for `actions.disable_zaraz` ([#1523](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1523))
* resource/cloudflare_record: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))
* resource/cloudflare_ruleset: Setting description to `Optional` to better reflect API requirements ([#1412](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1412))
* resource/cloudflare_ruleset: Setting description to `Optional` to better reflect API requirements ([#1556](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1556))
* resource/cloudflare_ruleset: add skip support for `products` and `phases` ([#1391](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1391))
* resource/cloudflare_ruleset: add support for 'Action' and 'Enabled' action_parameters > overrides attributes ([#1249](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1249))
* resource/cloudflare_ruleset: add support for HTTP rate limiting ([#1179](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1179))
* resource/cloudflare_ruleset: add support for Transform Rules ([#1169](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1169))
* resource/cloudflare_ruleset: add support for WAF payload logging ([#1174](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1174))
* resource/cloudflare_ruleset: add support for `action_parameters.response` to control the response when triggering a WAF rule ([#1507](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1507))
* resource/cloudflare_ruleset: add support for `ratelimit.requests_to_origin` ([#1507](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1507))
* resource/cloudflare_ruleset: add support for custom fields logging ([#1630](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1630))
* resource/cloudflare_ruleset: add support for ddos_l7 configuration ([#1212](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1212))
* resource/cloudflare_ruleset: add support for exposed credential checks ([#1263](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1263))
* resource/cloudflare_ruleset: add support for more complex skip ruleset configurations ([#1201](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1201))
* resource/cloudflare_ruleset: add support for rewriting HTTP response headers ([#1339](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1339))
* resource/cloudflare_ruleset: add support for rule `logging` ([#1538](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1538))
* resource/cloudflare_ruleset: smoother handling of UI/API collisions during migrations ([#1393](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1393))
* resource/cloudflare_teams_accounts: Add the `fips` field for configuring FIPS-compliant TLS. ([#1380](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1380))
* resource/cloudflare_teams_list: Add support for IP type ([#1550](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1550))
* resource/cloudflare_teams_rules: Add `insecure_disable_dnssec_validation` option to settings ([#1469](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1469))
* resource/cloudflare_teams_rules: GATE-2273: Adds support for device posture gateway rules ([#1353](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1353))
* resource/cloudflare_waiting_room: Add default_template_language field. ([#1651](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1651))
* resource/cloudflare_zone: add support for partner rate plans ([#1464](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1464))
* resource/cloudflare_zone: support changing `type` values ([#1301](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1301))
* resource/cloudflare_zone_setting_override: add support for overriding `binary_ast` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/cloudflare_zone_setting_override: add support for overriding `early_hints` ([#1285](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1285))
* resource/cloudflare_zone_setting_override: add support for overriding `filter_logs_to_cloudflare` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/cloudflare_zone_setting_override: add support for overriding `log_to_cloudflare` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/cloudflare_zone_setting_override: add support for overriding `orange_to_orange` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/cloudflare_zone_setting_override: add support for overriding `proxy_read_timeout` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/cloudflare_zone_setting_override: add support for overriding `visitor_ip` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/custom_hostname: also adds missing `validation_errors`, and `certificate_authority` ([#1424](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1424))
* resource/custom_hostname: validation tokens are now an array (`validation_records`) instead of a top level, but the only top level record that was previously here was for cname validation, txt/http/email were entirely missing. ([#1424](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1424))
* resource/logpush_job: Add `filter` field support ([#1664](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1664))
* resource/ruleset: add support for `origin` and `host_header` attributes ([#1620](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1620))

BUG FIXES:

* cloudflare_argo_tunnel: conditionally fetch settings based on the provided configuration ([#1451](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1451))
* cloudflare_tunnel_routes: Fix reads matching routers with larger CIDRs ([#1581](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1581))
* resource/cloudflare_access_application: Fix inability to update `http_only_cookie_attribute` to false ([#1602](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1602))
* resource/cloudflare_access_group: allow github access groups to be created without a list of teams ([#1589](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1589))
* resource/cloudflare_access_group: fix mapping error for AzureAD ([#1341](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1341))
* resource/cloudflare_access_policy: handle empty `nil` values for building policies ([#1237](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1237))
* resource/cloudflare_access_rule: Fix lifecycle of access_rule update ([#1601](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1601))
* resource/cloudflare_access_rule: allow "ip6" to be a padded or unpadded value and compare correctly ([#1294](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1294))
* resource/cloudflare_account_member: handle role changes made in the dashboard ([#1202](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1202))
* resource/cloudflare_api_token: ignore ordering changes in `permission_groups` ([#1545](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1545))
* resource/cloudflare_api_token: modified_on is now read correctly ([#1368](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1368))
* resource/cloudflare_argo: call `Read` for `Import` operations ([#1295](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1295))
* resource/cloudflare_argo_tunnel: fix import mechanism ([#1329](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1329))
* resource/cloudflare_argo_tunnel: update CNAME to use `cfargotunnel.com` ([#1293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1293))
* resource/cloudflare_device_posture_integration: remove superfluous `id` from schema ([#1504](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1504))
* resource/cloudflare_fallback_domain: default entries are now restored on delete. ([#1399](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1399))
* resource/cloudflare_load_balancer: handle empty `rules` for `resourceCloudflareLoadBalancerStateUpgradeV1` ([#1257](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1257))
* resource/cloudflare_logpush_job: make ownership challenge check for https not required ([#1588](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1588))
* resource/cloudflare_notification_policy: Fix unexpected crashes when using cloudflare_notification_policy with a filters attribute ([#1542](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1542))
* resource/cloudflare_origin_ca_certificate: ignore `requested_validity` changes due to the value decreasing but still store it ([#1214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1214))
* resource/cloudflare_origin_ca_certificate: reintroduce `DiffSuppressFunc` for `requested_validity` changes to handle all schema/SDK combinations ([#1289](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1289))
* resource/cloudflare_record: handle `Update`s for records with `data` blocks ([#1229](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1229))
* resource/cloudflare_record: no need to pass the resourceCloudflareRecordUpdate to the NonRetryable handler ([#1496](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1496))
* resource/cloudflare_ruleset: allow action parameter override `enabled` to be true/false or uninitialised ([#1275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1275))
* resource/cloudflare_ruleset: allow setting `uri` and `path` action parmeters together in a single rule ([#1271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1271))
* resource/cloudflare_ruleset: conditionally set action parameter "version" ([#1388](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1388))
* resource/cloudflare_ruleset: don't attempt to update "custom" rulesets using the phase entrypoint ([#1245](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1245))
* resource/cloudflare_ruleset: don't attempt to upgrade ratelimit if it isn't set ([#1501](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1501))
* resource/cloudflare_ruleset: fix handling of `false` values for category/rule overrides ([#1405](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1405))
* resource/cloudflare_ruleset: fix state handling for terraform-plugin-sdk v2 ([#1183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1183))
* resource/cloudflare_spectrum_application: Fix 'edge_ip_connectivity' state persistence ([#1515](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1515))
* resource/cloudflare_spectrum_application: prevent panic when configuration does not include `edge_ips.connectivity` ([#1599](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1599))
* resource/cloudflare_split_tunnel: import now works by specifying accountId/mode ([#1313](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1313))
* resource/cloudflare_split_tunnel: import will now use correct import function ([#1345](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1345))
* resource/cloudflare_teams_list: ignore `items` ordering ([#1338](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1338))
* resource/cloudflare_teams_rule: fixed detection of deleted teams rules ([#1622](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1622))
* resource/cloudflare_tunnel_route: Fix importing resource ([#1580](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1580))
* resource/cloudflare_waiting_room_event: handle time pointer for nullable struct member ([#1648](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1648))
* resource/cloudflare_workers_kv: handle invalid id during terraform import ([#1635](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1635))
* resource/cloudflare_zone: don't get stuck in endless loop for partner zone rate plans ([#1547](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1547))
* resource/cloudflare_zone: update plan identifier for professional rate plans ([#1583](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1583))
* resource/cloudflare_zone_dnssec: don't try to enable DNSSEC when state is "pending" ([#1530](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1530))
* resource/cloudflare_zone_settings_override: remap `zero_rtt` => `0rtt` for resource delete ([#1175](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1175))

DEPENDENCIES:

* `github.com/cloudflare/cloudflare-go` v0.29.0 => v0.30.0 ([#1379](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1379))

## 1.18.1 (August 29, 2019)

**Fixes:**

- `resource/cloudflare_load_balancer`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_page_rule`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_rate_limit`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_waf_rule`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_worker_route`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_worker_script`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_zone_lockdown`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))

## 1.18.0 (August 27, 2019)

**Fixes:**

- `resource/cloudflare_page_rule`: Fix a logic condition where setting `edge_cache_ttl` action but then not updating it in subsequent `apply` runs causes it to be blown away ([#453](https://github.com/cloudflare/terraform-provider-cloudflare/issues/453))

**Improvements:**

- provider: You can now use API tokens to authenticate instead of user email and key ([#450](https://github.com/cloudflare/terraform-provider-cloudflare/issues/450))
- `resource/cloudflare_zone_lockdown`: `priority` can now be set on the resource ([#445](https://github.com/cloudflare/terraform-provider-cloudflare/issues/445))
- `resource/cloudflare_custom_ssl`: Updated website documentation navigation to include link for resource ([#442)](https://github.com/cloudflare/terraform-provider-cloudflare/issues/442))

**Deprecations:**

- `resource/cloudflare_access_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_filter`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_firewall_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_load_balancer`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_page_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_rate_limit`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_waf_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_worker_route`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_worker_script`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_zone_lockdown`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))

## 1.17.1 (August 09, 2019)

**Fixes:**

- Partially revert [[#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421)] deprecation messages

## 1.17.0 (August 09, 2019)

**Removals:**

- `resource/cloudflare_zone_settings_override`: `sha1_support` has been removed due to Cloudflare no longer supporting SHA1 certificates or the API endpoint ([#415](https://github.com/cloudflare/terraform-provider-cloudflare/issues/415))

**Deprecations:**

- `resource/cloudflare_zone_settings_override`: `tls_1_2_only` has been superseded by using `min_tls_version` instead ([#405](https://github.com/cloudflare/terraform-provider-cloudflare/issues/405))
- `resource/cloudflare_access_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_filter`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_firewall_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_load_balancer`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_page_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_rate_limit`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_waf_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_worker_route`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_worker_script`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_zone_lockdown`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))

**Improvements:**

- **New Resource:** `cloudflare_custom_ssl` ([#418](https://github.com/cloudflare/terraform-provider-cloudflare/issues/418))
- `resource/cloudflare_filter`: Strip all surrounding whitespace from filter expressions to match API responses ([#361](https://github.com/cloudflare/terraform-provider-cloudflare/issues/361))
- `resource/cloudflare_zone`: Support unicode zone name values ([#412](https://github.com/cloudflare/terraform-provider-cloudflare/issues/412))
- `resource/cloudflare_page_rule`: Allow setting `origin_pull` for SSL ([#430](https://github.com/cloudflare/terraform-provider-cloudflare/issues/430))
- `resource/cloudflare_load_balancer_monitor`: Add TCP support for load balancer monitor ([#428](https://github.com/cloudflare/terraform-provider-cloudflare/issues/428))

**Fixes:**

- `resource/cloudflare_logpush_job`: Update documentation ([#395](https://github.com/cloudflare/terraform-provider-cloudflare/issues/395))
- `resource/cloudflare_zone_lockdown`: Fix: examples in documentation ([#407](https://github.com/cloudflare/terraform-provider-cloudflare/issues/407))
- `resource/cloudflare_page_rule`: Set nil on changed string-based Page Rule actions

## 1.16.1 (June 27, 2019)

**Fixes:**

- `resource/cloudflare_page_rule`: Fix regression in `browser_cache_ttl` where the value was sent as a string instead of an integer to the remote ([#390](https://github.com/cloudflare/terraform-provider-cloudflare/issues/390))

## 1.16.0 (June 20, 2019)

**Improvements:**

- `resource/cloudflare_zone_settings_override`: Add support for `h2_prioritization` and `image_resizing` ([#381](https://github.com/cloudflare/terraform-provider-cloudflare/issues/381))
- `resource/cloudflare_load_balancer_pool`: Update IP range for tests to not use reserved ranges ([#369](https://github.com/cloudflare/terraform-provider-cloudflare/issues/369))

**Fixes:**

- `resource/cloudflare_page_rule`: Fix issues with `browser_cache_ttl` defaults and when value is `0` (for Enterprise users) ([#379](https://github.com/cloudflare/terraform-provider-cloudflare/issues/379))

## 1.15.0 (May 24, 2019)

- The provider is now compatible with Terraform v0.12, while retaining compatibility with prior versions. ([#309](https://github.com/cloudflare/terraform-provider-cloudflare/issues/309))

## 1.14.0 (May 15, 2019)

**Improvements:**

- **New Resource:** `cloudflare_argo` Manage Argo features ([#304](https://github.com/cloudflare/terraform-provider-cloudflare/issues/304))
- `cloudflare_zone`: Support management of partial zones ([#303](https://github.com/cloudflare/terraform-provider-cloudflare/issues/303))
- `cloudflare_rate_limit`: Update `modes` documentation ([#293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/212))
- `cloudflare_load_balancer`: Allow steering policy of "random" ([#329](https://github.com/cloudflare/terraform-provider-cloudflare/issues/329))

**Fixes:**

- `cloudflare_page_rule` - Allow setting `browser_cache_ttl` to 0 ([#293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/291))
- `cloudflare_page_rule` - Swap to completely replacing rules ([#338](https://github.com/cloudflare/terraform-provider-cloudflare/issues/338))

## 1.13.0 (April 12, 2019)

**Improvements**

- **New Resource:** `cloudflare_logpush_job` ([#287](https://github.com/cloudflare/terraform-provider-cloudflare/issues/287))
- `cloudflare_zone_settings` - Remove option to toggle `always_on_ddos` ([#253](https://github.com/cloudflare/terraform-provider-cloudflare/issues/253))
- `cloudflare_page_rule` - Update documentation to clarify "0" usage
- `cloudflare_zones` - Return zone ID and zone name ([#275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/275))
- `cloudflare_load_balancer` - Add `enabled` field ([#208](https://github.com/cloudflare/terraform-provider-cloudflare/issues/208))
- `cloudflare_record` - validators: Allow PTR DNS records ([#283](https://github.com/cloudflare/terraform-provider-cloudflare/issues/283))

**Fixes:**

- `cloudflare_custom_pages` - Use correct casing for `zone_id` lookups
- `cloudflare_rate_limit` - Make `correlate` optional and not flap in state management ([#271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/271))
- `cloudflare_spectrum_application` - Fixed integration tests to work ([#275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/275))
- `cloudflare_page_rule` - Better track field changes in `actions` resource. ([#107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/107))

## 1.12.0 (March 07, 2019)

**Improvements:**

- provider: Enable request/response logging ([#212](https://github.com/cloudflare/terraform-provider-cloudflare/issues/212))
- resource/cloudflare_load_balancer_monitor: Add validation for `port` ([#213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/213))
- resource/cloudflare_load_balancer_monitor: Add `allow_insecure` and `follow_redirects` ([#205](https://github.com/cloudflare/terraform-provider-cloudflare/issues/205))
- resource/cloudflare_page_rule: Updated available actions documentation to match what is available ([#228](https://github.com/cloudflare/terraform-provider-cloudflare/issues/228))
- provider: Swap to using go modules for dependency management ([#230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/230))
- provider: Minimum Go version for development is now 1.11 ([#230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/230))

**Fixes:**

- resource/cloudflare_record: Read `data` back from API correctly ([#217](https://github.com/cloudflare/terraform-provider-cloudflare/issues/217))
- resource/cloudflare_rate_limit: Read `correlate` back from API correctly ([#204](https://github.com/cloudflare/terraform-provider-cloudflare/issues/204))
- resource/cloudflare_load_balancer_monitor: Fix incorrect type cast for `port` ([#213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/213))
- resource/cloudflare_load_balancer: Make `steering_policy` computed to avoid spurious diffs ([#214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/214))
- resource/cloudflare_load_balancer: Read `session_affinity` back from API to make import work & detects drifts ([#214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/214))

## 1.11.0 (January 11, 2019)

**Improvements:**

- **New Resource:** `cloudflare_spectrum_app` ([#156](https://github.com/cloudflare/terraform-provider-cloudflare/issues/156))
- **New Data Source:** `cloudflare_zones` ([#168](https://github.com/cloudflare/terraform-provider-cloudflare/issues/168))
- `cloudflare_load_balancer_monitor` - Add optional `port` parameter ([#179](https://github.com/cloudflare/terraform-provider-cloudflare/issues/179))
- `cloudflare_page_rule` - Improved documentation for `priority` attribute ([#182](https://github.com/cloudflare/terraform-provider-cloudflare/issues/182)], missing `explicit_cache_control` [[#185](https://github.com/cloudflare/terraform-provider-cloudflare/issues/185))
- `cloudflare_rate_limit` - Add `challenge` and `js_challenge` rate-limit modes ([#172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/172))

**Fixes:**

- `cloudflare_page_rule` - Page rule `zone` attribute change to trigger new resource ([#183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/183))

## 1.10.0 (December 18, 2018)

**Improvements:**

- `cloudflare_zone_settings_override` - Add `opportunistic_onion` zone setting support ([#170](https://github.com/cloudflare/terraform-provider-cloudflare/issues/170))
- `cloudflare_zone` - Add ability to set zone plan ([#160](https://github.com/cloudflare/terraform-provider-cloudflare/issues/160))

**Fixes:**

- `cloudflare_zone` - Allow zones to be properly imported ([#157](https://github.com/cloudflare/terraform-provider-cloudflare/issues/157))
- `cloudflare_access_policy` - Match access_policy argument requisites with reality ([#158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/158))
- `cloudflare_filter` - Allow `zone_id` to set `zone` and vice versa ([#162](https://github.com/cloudflare/terraform-provider-cloudflare/issues/162))
- `cloudflare_firewall_rule` - Allow `zone_id` to set `zone` and vice versa ([#174](https://github.com/cloudflare/terraform-provider-cloudflare/issues/174))
- `cloudflare_access_rule` - Ensure `zone` and `zone_id` are always set ([#175](https://github.com/cloudflare/terraform-provider-cloudflare/issues/175))
- Minor documentation fixes

## 1.9.0 (November 15, 2018)

**Improvements:**

- **New Resource:** `cloudflare_access_application` ([#145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/145))
- **New Resource:** `cloudflare_access_policy` ([#145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/145))
- `cloudflare_load_balancer` - Add steering policy support ([#147](https://github.com/cloudflare/terraform-provider-cloudflare/issues/147))
- `cloudflare_load_balancer` - Support `session_affinity` ([#153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/153))
- `cloudflare_load_balancer_pool` - Support `weight` ([#153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/153))

**Fixes:**

- `cloudflare_record` - Compare name without the zone name ([#151](https://github.com/cloudflare/terraform-provider-cloudflare/issues/151))
- Minor documentation fixes ([#149](https://github.com/cloudflare/terraform-provider-cloudflare/issues/149)] [[#152](https://github.com/cloudflare/terraform-provider-cloudflare/issues/152))

## 1.8.0 (November 05, 2018)

**Improvements:**

- **New Resource:** `cloudflare_zone` ([#58](https://github.com/cloudflare/terraform-provider-cloudflare/issues/58))
- **New Resource:** `cloudflare_custom_pages` ([#132](https://github.com/cloudflare/terraform-provider-cloudflare/issues/132))
- `cloudflare_zone_settings_override` - Allow setting SSL level to Strict (SSL-Only Origin Pull) ([#122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/122))
- Update provider usage/build docs and how to update a dependency ([#138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/138))
- Improve `Building The Provider` instructions ([#143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/143))
- `cloudflare_access_rule` - Make importable for all rule types ([#141](https://github.com/cloudflare/terraform-provider-cloudflare/issues/141))
- `cloudflare_load_balancer_pool` - Implement `Update` ([#140](https://github.com/cloudflare/terraform-provider-cloudflare/issues/140))

**Fixes:**

- `cloudflare_rate_limit` - Documentation fixes for markdown where \_ALL\_ is italicized ([#125](https://github.com/cloudflare/terraform-provider-cloudflare/issues/125))
- `cloudflare_worker_route` - Correctly set `multi_script` on Enterprise worker imports ([#124](https://github.com/cloudflare/terraform-provider-cloudflare/issues/124))
- `account_member` - Ignore role ID ordering ([#128](https://github.com/cloudflare/terraform-provider-cloudflare/issues/128))
- `cloudflare_rate_limit` - Origin traffic isn't default anymore ([#130](https://github.com/cloudflare/terraform-provider-cloudflare/issues/130))
- `cloudflare_rate_limit` - Update rate limit validation to allow `1` ([#129](https://github.com/cloudflare/terraform-provider-cloudflare/issues/129))
- `cloudflare_record` - Add validation to ensure TTL is not set while `proxied` is true ([#127](https://github.com/cloudflare/terraform-provider-cloudflare/issues/127))
- Updated code for provider version in User-Agent
- `cloudflare_zone_lockdown` - Fix import of zone lockdowns ([#135](https://github.com/cloudflare/terraform-provider-cloudflare/issues/135))

## 1.7.0 (October 09, 2018)

**Improvements:**

- **New Resource:** `cloudflare_account_member` ([#78](https://github.com/cloudflare/terraform-provider-cloudflare/issues/78))

## 1.6.0 (October 05, 2018)

**Improvements:**

- **New Resource:** `cloudflare_filter`
- **New Resource:** `cloudflare_firewall_rule`

## 1.5.0 (September 21, 2018)

**Improvements:**

- **New Resource:** `cloudflare_zone_lockdown` ([#115](https://github.com/cloudflare/terraform-provider-cloudflare/issues/115))

**Fixes:**

- Send User-Agent header with name and version when contacting API
- `cloudflare_page_rule` - Fix page rule polish (off, lossless or lossy) ([#116](https://github.com/cloudflare/terraform-provider-cloudflare/issues/116))

## 1.4.0 (September 11, 2018)

**Improvements:**

- **New Resource:** `cloudflare_worker_route` ([#110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/110))
- **New Resource:** `cloudflare_worker_script` ([#110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/110))

## 1.3.0 (September 04, 2018)

**Improvements:**

- **New Resource:** `cloudflare_access_rule` ([#64](https://github.com/cloudflare/terraform-provider-cloudflare/issues/64))

**Fixes:**

- `cloudflare_zone_settings_override` - Change Zone Settings Override to use GetOkExists ([#107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/107))

## 1.2.0 (August 13, 2018)

**Improvements:**

- **New Resource:** `cloudflare_waf_rule` ([#98](https://github.com/cloudflare/terraform-provider-cloudflare/issues/98))
- `cloudflare_zone_settings_override` - Add `off` as Security Level setting ([#99](https://github.com/cloudflare/terraform-provider-cloudflare/issues/99))
- `resource_cloudflare_rate_limit` - Add nat support ([#96](https://github.com/cloudflare/terraform-provider-cloudflare/issues/96))
- `resource_cloudflare_zone_settings_override` - Add `zrt` as a value for the `tls_1_3` setting ([#106](https://github.com/cloudflare/terraform-provider-cloudflare/issues/106))
- Minor documentation improvements

**Fixes:**

- `cloudflare_record` - Setting a DNS record's `proxied` flag to false stopped working ([#103](https://github.com/cloudflare/terraform-provider-cloudflare/issues/103))

## 1.1.0 (July 25, 2018)

FIXES:

- `cloudflare_ip_ranges` - IPv6 CIDR blocks should return IPv6 addresses ([#51](https://github.com/cloudflare/terraform-provider-cloudflare/issues/51))
- `cloudflare_zone_settings_override` - Allow `0` for `browser_cache_ttl` ([#71](https://github.com/cloudflare/terraform-provider-cloudflare/issues/71))
- `cloudflare_page_rule` - `forwarding_urls` in page rules are lists ([#79](https://github.com/cloudflare/terraform-provider-cloudflare/issues/79))
- `cloudflare_page_rule` - The API supports `active` and `disabled`, not `paused` ([#84](https://github.com/cloudflare/terraform-provider-cloudflare/issues/84))

IMPROVEMENTS:

- `cloudflare_zone_settings_override` - Add support for `min_tls_version` ([#72](https://github.com/cloudflare/terraform-provider-cloudflare/issues/72))
- `cloudflare_page_rule` - Add support for more settings: `bypass_cache_on_cookie`, `cache_by_device_type`, `cache_deception_armor`, `cache_on_cookie`, `host_header_override`, `polish`, `explicit_cache_control`, `origin_error_page_pass_thru`, `sort_query_string_for_cache`, `resolve_override`, `respect_strong_etag`, `response_buffering`, `true_client_ip_header`, `mirage`, `disable_railgun`, `cache_key`, `waf`, `rocket_loader`, `cname_flattening` ([#68](https://github.com/cloudflare/terraform-provider-cloudflare/issues/68)], [[#81](https://github.com/cloudflare/terraform-provider-cloudflare/issues/81)], [[#85](https://github.com/cloudflare/terraform-provider-cloudflare/issues/85))
- `cloudflare_page_rule` - Add `off` setting to `security_level` ([#81](https://github.com/cloudflare/terraform-provider-cloudflare/issues/81))
- `cloudflare_record` - DNS Record improvements ([#97](https://github.com/cloudflare/terraform-provider-cloudflare/issues/97))
- Various documentation improvements

## 1.0.0 (April 06, 2018)

BACKWARDS INCOMPATIBILITIES / NOTES:

- resource/cloudflare_record: Changing `name` or `domain` now force a recreation
  of the record ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))

FEATURES:

- **New Resource:** `cloudflare_rate_limit` ([#30](https://github.com/cloudflare/terraform-provider-cloudflare/issues/30))
- **New Resource:** `cloudflare_page_rule` ([#38](https://github.com/cloudflare/terraform-provider-cloudflare/issues/38))
- **New Resource:** `cloudflare_load_balancer` ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))
- **New Resource:** `cloudflare_load_balancer_pool` ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))
- **New Resource:** `cloudflare_zone_settings_override` ([#41](https://github.com/cloudflare/terraform-provider-cloudflare/issues/41))
- **New Resource:** `cloudflare_load_balancer_monitor` ([#42](https://github.com/cloudflare/terraform-provider-cloudflare/issues/42))
- **New Data Source:** `cloudflare_ip_ranges` ([#28](https://github.com/cloudflare/terraform-provider-cloudflare/issues/28))

IMPROVEMENTS:

- resource/cloudflare_record: Validate `TXT` records ([#14](https://github.com/cloudflare/terraform-provider-cloudflare/issues/14))
- resource/cloudflare_record: Add `data` input to suppport SRV, LOC records
  ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))
- resource/cloudflare_record: Add computed attributes `created_on`,
  `modified_on`, `proxiable`, and `metadata` to records ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))
- resource/cloudflare_record: Support import of existing records ([#36](https://github.com/cloudflare/terraform-provider-cloudflare/issues/36))
- New Provider configuration options for API rate limiting ([#43](https://github.com/cloudflare/terraform-provider-cloudflare/issues/43))
- New Provider configuration options for using Organizations ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))

## 0.1.0 (June 20, 2017)

NOTES:

- Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider
  Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
