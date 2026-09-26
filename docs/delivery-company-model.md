# Delivery company model (v36.1.0)

The current model omits courier API/manual classification and routing priority.
Each courier company's immutable, extensible `code` is its sole public identity.
A code does not prove that a verified provider implementation, coverage, or live
availability exists. This model contains no provider implementation, API routes,
booking or routing code. Pin
`github.com/Potato-Mart/Backend-Shared-Contract/v36 v36.1.0`.
See [v36 migration](v36-migration.md) for the market, authentication metadata,
provider settings and metadata changes and the service-owned configuration
migration.

## Ownership and projections

| Model | Owner / purpose |
| --- | --- |
| `supply/courier.DeliveryCompany` | Supply admin catalogue root with audit fields; code, name, credential requirements, enabled, instructions, derived dispatch capability, revision, countries, capabilities, slot source, connection, service areas, schedules, optional provider settings and inert custom metadata. No API/manual classification or provider selector. |
| `DeliveryCompanyRef` | Customer-safe snapshot: `code`, `name`, `revision`. No connection, instructions, coverage configuration or credentials. |
| `DeliveryCredentialRequirements` | Derived, non-sensitive API URL, connection-test address requirement, and optional supported authentication methods. It is read-only company metadata and is never accepted as a write field. |
| `DeliveryAuthenticationMethodRequirements` | Open method identifier and required credential field names, without values or an authentication flow. |
| `DeliveryConnection` | Sanitized backend observations: `credential_configured`, `health`, optional `last_checked_at`. |
| `DeliveryProviderCredentials` | Standalone privileged values: optional `sign_in_account`, `password`, `api_base_url`, `api_token`, and provider-specific credential extension. Only a separately authorized credential operation may serialize these sensitive values. |
| `DeliveryProviderSettings` | Versioned, non-secret JSON values interpreted only through Supply's registered schema for the parent company code. Unknown keys are preserved but not executed. |
| `DeliveryCompanyCustomMetadataEntry` | Non-secret `key`, open `value_type`, and typed JSON `value`; custom metadata is inert and is never sent to provider requests. |
| `DeliveryCapabilities` | Explicit booleans for booking, tracking, proof of delivery, refrigeration, provider coverage, provider slots and configured slots. Support does not imply current availability. |
| `DeliveryServiceArea` | Market/country-scoped exact postal filters, Orders zone ID reference and derived ISO subdivision state codes; independent of depot coverage and delivery area pricing. |
| `ShippingZoneRef` | ID-only reference to an Orders-owned `shipping.Zone`; it is not a copied zone snapshot. |
| `DeliveryServiceWindow` | Legacy recurring configured local-time window retained for JSON compatibility and deprecated for new availability configuration. |
| `orders/shipping.DeliverySelection` | Immutable accepted choice copied into Order, Supply's local job and OutboundShipment. |

Supply owns credential authorization, KMS-encrypted/versioned storage, secret
resource references, rotation and redaction. `DeliveryProviderCredentials` is a
standalone sensitive value model, never an embedded company or customer field;
provider extension values are write-only through independently authorized
credential operations. The shared envelope contains input values, not
ciphertext. Connection testing returns only sanitized status and time; raw
diagnostics, credentials, headers and provider payloads must not reach general
admin or customer projections. `DispatchCapable`, connection health and
configuration revision are derived server values, not editable authority flags.
Health values are `unknown`,
`healthy`, `unhealthy`; missing or stale evidence does not mean healthy.

`DeliveryCompany.CredentialRequirements` is derived, read-only metadata that is
always serialized. It is `null` when the exact company code has no verified
registered provider implementation. When present, `api_base_url` is that
implementation's verified public API origin/path, and
`connection_address_required` indicates whether Supply's connection-test
operation requires an address. These values are not credentials and are not
accepted on create/update operations. A non-null value does not claim current
coverage or live slots.

Optional `authentication_methods` advertises the registered implementation's
supported methods, each with an open `method` identifier (for example
`api_token`) and `required_fields` names. These contain no credential values.
An API URL supplied by derived metadata may already satisfy that prerequisite.
Missing or empty methods make no support claim; older responses may omit them.
The privileged credential model's account/password fields do not establish an
account authentication protocol or permission to fall back to one.

`Code` is the immutable, extensible identity for one independent company. It is
an open string, not a fixed company list or provider selector, so other markets
can add company codes. Supply privately registers and selects verified provider
behavior by the exact `Code`; the shared model contains no provider identity.
Adding a company record does not install or verify provider behavior and does
not imply quote, booking, dispatch or live-slot readiness. A company revision
is positive and increases when its effective routing, connection configuration,
capabilities, schedules, provider settings, custom metadata or Supply-managed
credential binding changes. Supply owns atomic revision checks and history.
Routine health checks need not change configuration revision.

## Provider settings and custom metadata

`DeliveryCompany.provider_settings` is optional and has the shape
`{"schema_version": 1, "values": {"client_id": "example"}}`. Its values
are typed JSON carried by the shared `metadata.Metadata` value map; the version
identifies the Supply-registered provider settings schema.
`DeliveryCompany.Code` selects the implementation and implicitly keys the
settings envelope. Supply validates settings and only registered, understood
keys may affect provider requests.
Unknown keys must survive read-modify-write but remain inert until a compatible
provider schema recognizes them.

`custom_metadata` is a separate optional array of
`{"key":"dispatch_group","value_type":"string","value":"north"}`
entries. It supports company-specific non-secret metadata absent from common
fields. It is never interpreted as provider configuration or passed to an
adapter. Consumers should preserve unknown value types and values. Supply owns
key uniqueness, type validation, size limits and explicit update/removal
semantics.

Both namespaces may appear under one Additional settings UI, but they are
separate in storage and semantics. On partial updates, omission preserves the
stored namespace; explicit replacement or clearing is defined by Supply's
service-owned request DTOs. Older writers that replace a full record must be
upgraded or protected by server-side merge behavior before extension data is
activated. Existing country codes, service areas, market binding and shipping
zone references retain their v36.0.0 meanings.

The v35 major release removes `integration` from `DeliveryCompany`,
`DeliveryCompanyRef` and nested frozen delivery references. The legacy
`adapter` field was already removed in v34 and remains absent. A stored v34 JSON
object may contain `integration`; older v33 data may also contain `adapter`.
Standard Go JSON decoding ignores these unknown fields, and v35 serialization
omits them. Consumers must preserve the existing company `code` and must not
derive or rewrite it from the display name. Each service owns migration of
accepted historical selections and private provider state; this shared model
does not define credential, operation-receipt or provider-storage migration.

## Coverage semantics and backend validation

1. Require an enabled company, enabled area and a destination country present in
   both the company `country_codes` and the area's `country_code`. Empty country
   or service-area lists grant no coverage. Countries are ISO 3166-1 alpha-2.
   Require a valid `market_code` whose Pricing-owned country matches the area;
   legacy missing markets require explicit migration, never country inference.
   `shipping_zone` contains only the Orders zone ID. Supply
   resolves the current Orders zone before using its active state, country,
   administrative areas or postcodes; copied names or coverage are not routing
   authority. `state_codes` is derived from the zone and uses ISO 3166-2
   `geography.SubdivisionCode` values. Selecting all country zones materializes
   existing active zone rows; it never grants coverage to future zones. Legacy
   areas may omit optional references, but missing data grants no wildcard.
   Supply owns new-write validation and migration.
2. Normalize postal strings per destination country before exact comparison;
   preserve leading zeros. No numeric ranges, prefix expansion or inferred city
   boundary is implied by the model.
3. `postal_code_mode=include_only` requires an explicit include set; an empty
   set matches nothing. Exclusions always win, including when the same value is
   included. `all_except` requires an empty include set and
   `provider_coverage_required=true`. It only identifies candidates outside the
   exclusion set within the resolved zone; it is not proof of coverage for an
   entire country.
4. When `provider_coverage_required=true`, authoritative provider coverage must
   also confirm the destination. Unknown, unavailable or failed coverage checks
   do not produce available slots. A verified, administratively maintained exact
   coverage list can instead use `include_only`; its provenance/freshness policy
   remains backend-owned. Provider marketing claims are not a postcode matrix.
5. `routing_priority` is removed. Supply validates unambiguous market/zone
   ownership across enabled slot-capable providers and rejects ambiguous
   matches; array order, map iteration and display names are not precedence.
   Reject duplicate area/window codes and invalid modes at configuration writes.
   Invalid/unknown persisted modes fail closed at reads.

No authoritative Australian postcode set is supplied by this release. `AU-VIC`
identifies a subdivision; it does not imply coverage of all Victoria or Melbourne.
Supply must resolve an active Orders zone and validate its configured postal
codes before routing. Detrack live slots remain inactive until exact coverage and
its pre-booking endpoint are verified. BeCool fallback uses its own live Orders
zone and verified coverage; it must not inherit Detrack's coverage. Country-neutral
types support later markets; this release does not enable any market.

New or migrated provider records stay disabled/unavailable until explicit,
verified coverage and connection settings exist. Existing Orders postcode/rate
data is not evidence of Supply carrier coverage. Provider endpoints, immutable
credential-version selection, compare-and-swap updates and rollback remain
backend/DevOps responsibilities; the shared projection exposes none of them.

## Configured schedules and provider offers

`slot_source` retains `none`, `configured_schedule` and `provider` for wire
compatibility. Manual `schedules` and `DeliveryServiceWindow` are deprecated for
new availability configuration, but remain readable on existing records.
Configured windows do not establish live provider availability. Provider offers
require actual provider evidence; `none` supplies no selectable windows.

Each configured window has a stable `code`, destination `country_code`, optional
`service_area_codes`, `days_of_week` (0 Sunday through 6 Saturday), `start_time`
and `end_time` (`HH:MM`), IANA `timezone`, optional label and local excluded dates,
inclusive effective date bounds, non-negative `minimum_lead_time_minutes`, optional
`fee` and enabled flag. Empty weekdays produces no windows. Empty area codes
applies to all enabled areas of that country; explicit codes must reference
existing same-country areas. Require start before end within one local date;
overnight windows must be expressed separately. Validate effective bounds,
calendar dates and timezone identifiers; never use the host's timezone.

Backends construct UTC start/end for each local date using timezone rules and
apply a documented policy for ambiguous/nonexistent daylight-saving times.
Lead time is elapsed time before the resulting start instant. A missing fee is
unknown/inherited, not free; explicit free delivery is Money with amount_minor=0
and a valid currency. Pricing, capacity and quote authorization remain service-owned.

`DeliverySchedule` retains all old fields and adds optional `delivery_company`
and `expires_at`. A top-level company describes every slot; mixed-company results
omit it and supply a company on each slot. If both are supplied they must agree.
The existing schedule revision identifies the offered schedule view.
`DeliverySlot` adds optional `delivery_company`, `source`, `schedule_code`,
`unavailable_reason`, `expires_at`; existing fee, times and availability retain
their meanings. Reason codes are sanitized service-owned values. Quote IDs and
API envelopes stay service-local. The earliest applicable schedule/slot expiry
bounds acceptance of a new choice, not fulfilment of an already accepted order.

## Frozen order and shipment handoff

`PreferredDeliverySlot.delivery_company` optionally identifies the selected
company/config revision. Treat every client-supplied reference, time and revision
as untrusted. The backend resolves the slot ID and validates destination,
availability, fee, expiry and current revisions before accepting the choice.

The resulting `DeliverySelection` contains these required fields:

| JSON field | Meaning |
| --- | --- |
| `delivery_company` | Safe code/name/config revision snapshot. |
| `schedule_code`, `schedule_revision` | Window/provider schedule identity and positive accepted schedule-view revision. |
| `slot_id` | Opaque dated offer identity, scoped to the accepted company/config and schedule. |
| `date` | Local delivery calendar date, `YYYY-MM-DD`. |
| `start_at`, `end_at` | Increasing nonzero UTC RFC3339 instants for the accepted window. |
| `timezone` | IANA timezone used to interpret the local date. |
| `source` | `configured_schedule` or `provider`; never `none` for an accepted slot. |

Store it on `Order.delivery_selection`, copy it through every order-create/saga
mapping, persist it before provider side effects, and copy it into
`OutboundShipment.delivery_selection`. Reject incomplete snapshots, non-positive
revisions, conflicting legacy values and mismatched local date/timezone before
booking. A config change must not silently re-route an accepted choice. Services
must retain or resolve the accepted configuration revision, or require an explicit
validated replacement choice if that configuration can no longer be honoured.
Retries retain the same accepted snapshot. Provider receipts, claims,
idempotency and retry state stay service-local. Provider booking acceptance must
not itself mark an order physically dispatched.

## Migration and verification

- Preserve existing `DeliverySchedule.carrier` and `Order.outsourced_carrier`
  display-name meaning. `OutboundShipment.carrier` remains the company code.
  Use the new reference code for identity; never reassign old scalar meanings.
- Preserve every existing company `code` exactly; codes are open/extensible and
  must not be mapped from the display name or restricted to the currently
  registered providers. Supply maps its local records deliberately. The removed
  `integration` field is ignored when decoding v34 JSON; legacy v33 data may
  also contain `adapter`. Both are absent from v35 serialization. Each consuming
  service owns historical selection compatibility.
- Supply privately selects a verified provider implementation by exact company
  code. An unverified code does not establish provider support, quote/booking
  eligibility or live slots; consumer surfaces must not imply those capabilities.
  Supply maps country, revision, coverage and schedule configuration before
  enabling new routing.
- Orders currently has service-local delivery-option DTOs with different required
  label/fee behavior. Do not blindly alias them to shared models; preserve API
  compatibility through explicit mappings. Frontend types follow backend OpenAPI.
- `credential_requirements` is always present in current company responses but may
  be `null`; consumers reading stored v33 company records should treat an absent
  value like `null`. Other legacy optional fields may remain absent. A missing
  selection on a legacy order means unknown; it grants no permission to infer an
  end time, timezone or alternate carrier. Service policy must distinguish
  legacy completion from new orders requiring a frozen selection.
- Shared tests cover legacy JSON, explicit postal strings/modes, no implicit enum
  defaults, customer-safe references, nil versus zero fees, offer metadata and
  Order-to-Shipment snapshot round trips. Runtime coverage, ambiguous ownership,
  stale offers, destination changes, daylight-saving transitions, provider failure,
  configuration retention, retry deduplication and booking-versus-dispatch require
  backend integration tests. Admin must show configured versus provider sources
  and unavailable/unknown states accurately.

## Provider reference boundary

The [official Detrack job fields](https://help.detrack.com/en/articles/6126913-detrack-api-job-fields-and-descriptions)
describe postal and delivery-time fields; these are not evidence of selectable
slot discovery. [BeCool's public site](https://becoolrefrigeratedcouriers.com.au/)
does not provide a complete postcode coverage matrix for this contract. Provider
requirements and operational defaults remain Supply-owned. The standalone
`DeliveryProviderCredentials` type defines only the privileged value shape;
actual credential values must never be included in source, tests, or
documentation examples.
