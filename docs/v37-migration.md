# v37 migration

Pin `github.com/Potato-Mart/Backend-Shared-Contract/v37 v37.0.0` after the
GitHub release is published. This major changes the Go import path and removes
three fixed courier credential fields. HTTP route versions and runtime rollout
remain owned by their services.

## Courier credentials

The privileged shared credential shape is now:

```json
{
  "api_base_url": "https://provider.example.test",
  "provider_extension": {
    "schema_version": 1,
    "values": {"provider_auth_field": "sensitive-value"}
  }
}
```

`sign_in_account`, `password`, and `api_token` are absent from the v37 type.
Supply must keep credential values encrypted and provider extension keys intact.
It may retain legacy v1 courier HTTP DTOs and persisted-value readers while
its v2 provider routes use the v37 model. The module introduces no data
migration or backfill.

## Promotion applications

`PromotionApplication.resolved_target_package_allocations` freezes the order
item, SKU, package option, applied purchased-package count, optional complete
group ordinal and discount. One order item can aggregate multiple packages.
Compare an allocation's applied count with the existing requested component
count and remaining fully paid eligible quantity on the matching frozen order
item. Do not count the order row, base units, or substituted/picked/packed
composition as purchased packages. Pricing owns
same-versus-mixed matching, repeats, leftovers, volume-price exclusivity and
cart-discount allocation.

## Order edits and physical work

`Order.fulfillment_generation`, `PickingList.fulfillment_generation`, and
`OrderPackingProgress.fulfillment_generation` bind physical evidence to one
requested order composition. Orders advances the generation after a committed
product edit. Supply preserves old records for audit but treats older or
unknown-generation picking and packing confirmations as non-authorizing for
the revised order, then rechecks the whole order. The existing
`fulfilment.packing_updated` payload may gain this optional nested generation
without changing its event version; consumers must tolerate the field before
Supply emits it. Legacy payloads omit the field.

`order.edited` is a new `order-events` fact with `event_version` `v1`. Orders
publishes it only after the edit commits and includes one stable `edit_id`,
previous/new generation, safe whole-order requested item snapshots and
previous/revised totals. Consumers deploy before the producer emits it, dedupe
delivery by the envelope `event_id`, and dedupe the customer notification by
`edit_id`. The event is not a payment decision or dispatch authorization.

Order edit preview/commit/detail routes, Pricing preview, payment adjustment,
packing and picking print history, provider settings validation, legacy courier
adapters and persistence remain in their owning backend services. Frontends
follow those services' HTTP schemas; the shared `/v37` path alone does not
change `/v1` or `/v2` HTTP behavior.
