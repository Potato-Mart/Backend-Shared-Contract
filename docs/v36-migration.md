# v36 contract migration

Pin `github.com/Potato-Mart/Backend-Shared-Contract/v36 v36.0.0` after its
immutable release is available. Update imports and each adopting service's
`go.mod` and CI `contract_version` together. A local `go.work` does not prove
the committed dependency pin; run the owning service's gates with `GOWORK=off`.

## Courier areas

`DeliveryServiceArea` removes the Go field `RoutingPriority` and JSON field
`routing_priority`, and adds `MarketCode string` as `market_code` (always
serialized). The remaining fields retain their names and types:

```json
{
  "code": "EXAMPLE-ZONE",
  "country_code": "AU",
  "market_code": "example-market",
  "shipping_zone": {"id": "example-zone-id"},
  "state_codes": ["AU-VIC"],
  "postal_code_mode": "include_only",
  "include_postal_codes": ["0800"],
  "enabled": true,
  "provider_coverage_required": false
}
```

This is synthetic shape evidence, not approved geographic/provider coverage.
Supply validates the Pricing market and country and resolves the live Orders
zone. State codes are a derived zone projection. All-country selection
materializes the existing active zones into the same rows used for a subset;
it does not cover future zones automatically. There is no wildcard scope.
Code normalization and collision handling stay in Supply.

Legacy JSON containing `routing_priority` still decodes with standard Go JSON,
but v36 does not serialize it. A missing market decodes as an empty string;
Supply must obtain an explicit mapping rather than infer one from country.
Legacy missing zone references do not grant coverage. Services must replace
priority-based routing with validated unambiguous ownership and reject invalid
new configurations. Renaming area codes must update any legacy
`DeliveryServiceWindow.service_area_codes` references. Accepted old bookings
may need private historical configuration decoding; never rewrite frozen
`DeliveryCompanyRef` or `DeliverySelection` to infer a new choice.

## Authentication requirement metadata

`DeliveryCredentialRequirements` adds optional
`AuthenticationMethods []DeliveryAuthenticationMethodRequirements` with JSON
key `authentication_methods`. Each record has `Method string` (`method`) and
`RequiredFields []string` (`required_fields`):

```json
{
  "api_base_url": "https://provider.example.test/api/v1",
  "connection_address_required": false,
  "authentication_methods": [
    {"method": "api_token", "required_fields": ["api_base_url", "api_token"]}
  ]
}
```

Supply derives these open identifiers from verified implementation support.
Field names are configuration prerequisites, not credential values, and some
may already be supplied by derived metadata. The unchanged privileged
`DeliveryProviderCredentials` shape does not advertise method support merely
because it has account/password fields. Missing/empty method metadata makes no
authentication-method claim; older producers may omit it. Do not invent an
account fallback. Unimplemented company codes retain null credential
requirements. This release implements no authentication protocol.

## Offer audience

Both `Promotion.Controls` and `Coupon.Controls` use `PromotionControls`, which
adds `Audience *audience.Audience` as `audience,omitempty`. It reuses the
existing campaign customer-type and platform enums:

```json
{"customer_type": "retail", "platform": "web"}
```

The object lives at `controls.audience`. Nil omits it; `{}` and omitted
dimensions add no audience restriction. Customer values remain `guest`,
`retail`, `wholesale`; platform values remain `web`, `mobile`. Platform is
separate from the existing order `channels`. Pricing must enforce these fields
against trusted buyer/platform evidence; serialization alone grants nothing.

Campaign prefill is an editable offer snapshot. Services validate compatibility
against the current campaign without silently overwriting offers after a
campaign edit. Explicit geography/audience selection and valid start/end dates
are create-command requirements, not new shared enum values or globally
required `PromotionPeriod` fields. Existing product scope groups suffice;
services must preserve paired category selectors and avoid widening child
scopes. Platform context and coupon category evaluation require owning-service
work before restricted offers can be enabled.

## Preserved shapes and rollout

Packing, conversion, substitution, allocation and container shapes are
unchanged. Loose case-sized quantities retain EACH composition; intact cases
retain CASE composition. Packing commands, conservation of totals, sequence
allocation, conversion and placement receipts remain Supply-owned.

No event schema/version changes accompany this release. Publish and verify the
contract release first, then update backend pins/imports, service-local data
migrations, validation and OpenAPI, and finally their frontend consumers.
The contract release alone does not prove runtime or provider readiness.
