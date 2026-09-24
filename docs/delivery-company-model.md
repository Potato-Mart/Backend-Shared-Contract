# Delivery company model (v33.8.0)

This additive release defines shared data for Supply's delivery-company catalogue
and Orders-to-Supply delivery selection. It contains no postcode defaults,
provider availability claims, API routes, adapters, booking or routing code.
Pin `github.com/Potato-Mart/Backend-Shared-Contract/v33 v33.8.0`.

## Ownership and projections

| Model | Owner / purpose |
| --- | --- |
| `supply/courier.DeliveryCompany` | Supply admin catalogue root with audit fields; code, name, integration, adapter, enabled, instructions, derived dispatch capability, revision, countries, capabilities, slot source, connection, service areas and schedules. |
| `DeliveryCompanyRef` | Customer-safe snapshot: `code`, `name`, `integration`, optional `adapter`, `revision`. No connection, instructions, coverage configuration or credentials. |
| `DeliveryConnection` | Sanitized backend observations: `credential_configured`, `health`, optional `last_checked_at`. |
| `DeliveryProviderCredentials` | Standalone privileged values: optional `sign_in_account`, `password`, `api_base_url`, `api_token`. Only an independently authorized credential write or writer-only detail operation may serialize this model. |
| `DeliveryCapabilities` | Explicit booleans for booking, tracking, proof of delivery, refrigeration, provider coverage, provider slots and configured slots. Support does not imply current availability. |
| `DeliveryServiceArea` | Country-scoped exact postal filters, optional Orders zone ID reference, ISO subdivision state codes and priority; independent of depot coverage and delivery area pricing. |
| `ShippingZoneRef` | ID-only reference to an Orders-owned `shipping.Zone`; it is not a copied zone snapshot. |
| `DeliveryServiceWindow` | Legacy recurring configured local-time window retained for JSON compatibility and deprecated for new availability configuration. |
| `orders/shipping.DeliverySelection` | Immutable accepted choice copied into Order, Supply's local job and OutboundShipment. |

Supply owns credential authorization, KMS-encrypted/versioned storage, secret
resource references, rotation and redaction. `DeliveryProviderCredentials` is a
standalone sensitive value model, never an embedded company or customer field;
only separately authorized credential write and writer-only detail operations
may serialize its values. Connection testing returns only sanitized status and
time; raw diagnostics, credentials, headers and provider payloads must not reach
general admin or customer projections. `DispatchCapable`, connection health and
configuration revision are derived server values, not editable authority flags.
Health values are `unknown`,
`healthy`, `unhealthy`; missing or stale evidence does not mean healthy.

`Code` identifies a company instance; `Integration` preserves the existing
service-owned `api` or `manual` mode. `Adapter` is a separate optional open string
identifying a registered backend adapter, such as `detrack` or `bcrc`. Multiple
instance codes may select the same adapter. Adding a record does not install an
adapter. Supply retains its immutable lowercase company-code convention and
validates supported modes/adapters. Manual mode never gains automatic dispatch
by supplying an adapter. A company revision is positive and increases when its effective routing,
connection configuration, capabilities or schedules change. Routine health checks
need not change configuration revision.

For legacy `api` records missing adapter, only Supply's explicit migration may
resolve exact known company codes `detrack`/`bcrc` to registered adapters. Custom
API company codes require an explicit supported adapter; never guess from the
display name or substitute an arbitrary provider. Missing adapter remains absent
in legacy JSON, while new API references and frozen selections capture the
resolved adapter before acceptance. Manual records can omit it. The optional
wire field preserves compatibility; it does not grant dispatch readiness.

## Coverage semantics and backend validation

1. Require an enabled company, enabled area and a destination country present in
   both the company `country_codes` and the area's `country_code`. Empty country
   or service-area lists grant no coverage. Countries are ISO 3166-1 alpha-2.
   `shipping_zone`, when configured, contains only the Orders zone ID. Supply
   resolves the current Orders zone before using its active state, country,
   administrative areas or postcodes; copied names or coverage are not routing
   authority. `state_codes` uses ISO 3166-2 `geography.SubdivisionCode` values.
   Legacy areas may omit these optional fields; Supply owns the requirements for
   new configuration writes and migration.
2. Normalize postal strings per destination country before exact comparison;
   preserve leading zeros. No numeric ranges, prefix expansion or inferred city
   boundary is implied by the model.
3. `postal_code_mode=include_only` requires an explicit include set; an empty
   set matches nothing. Exclusions always win, including when the same value is
   included. `all_except` requires an empty include set and
   `provider_coverage_required=true`. It only identifies candidates outside the
   exclusion set; it is not proof of coverage for an entire country.
4. When `provider_coverage_required=true`, authoritative provider coverage must
   also confirm the destination. Unknown, unavailable or failed coverage checks
   do not produce available slots. A verified, administratively maintained exact
   coverage list can instead use `include_only`; its provenance/freshness policy
   remains backend-owned. Provider marketing claims are not a postcode matrix.
5. Lower `routing_priority` wins among eligible rules. Reject overlapping rules
   that tie ambiguously; do not rely on array order, map iteration or display
   names. Reject duplicate area/window codes and invalid modes at configuration
   writes. Invalid/unknown persisted modes fail closed at reads.

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
| `delivery_company` | Safe code/name/integration/adapter/config revision snapshot. |
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
- Existing company fields (`code`, `name`, `integration`, `enabled`,
  `default_instructions`, audit timestamps and `dispatch_capable`) are represented.
  Supply maps its local records/DTOs deliberately and migrates country, revision,
  coverage and schedule configuration before enabling new routing.
- Orders currently has service-local delivery-option DTOs with different required
  label/fee behavior. Do not blindly alias them to shared models; preserve API
  compatibility through explicit mappings. Frontend types follow backend OpenAPI.
- New fields on existing records are optional. A missing selection on a legacy
  order means unknown; it grants no permission to infer an end time, timezone or
  alternate carrier. Service policy must distinguish legacy completion from new
  orders requiring a frozen selection.
- Shared tests cover legacy JSON, explicit postal strings/modes, no implicit enum
  defaults, customer-safe references, nil versus zero fees, offer metadata and
  Order-to-Shipment snapshot round trips. Runtime coverage, ambiguous priority,
  stale offers, destination changes, daylight-saving transitions, adapter failure,
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
