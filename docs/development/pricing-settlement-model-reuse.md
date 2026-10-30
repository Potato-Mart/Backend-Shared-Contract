# Existing pricing and settlement evidence models

This reference describes the shared models available in `v53.8.0`. It adds no
types, fields, wire representations or release version. It does not establish
that the asynchronous exact-pricing or POS settlement protocols are complete.

## Promotions and quote evidence

| Information | Existing models | Reuse boundary |
| --- | --- | --- |
| Authored promotion rules | [Promotion](../../pkg/contracts/pricing/promotion/promotion.go), [PromotionRelation](../../pkg/contracts/pricing/promotion/promotion_relation.go), [PromotionTerm](../../pkg/contracts/pricing/promotion/promotion_term.go) | Retain revisioned definitions, qualifier/target scopes and typed terms. |
| Frozen application result | [PromotionApplication](../../pkg/contracts/pricing/promotion/promotion_application.go), [PromotionTargetPackageAllocation](../../pkg/contracts/pricing/promotion/promotion_target_package_allocation.go) | Retain the applied revision, resolved terms, money and purchased-package allocations. This is not a complete solver trace. |
| Frozen line and source evidence | [PriceSnapshot](../../pkg/contracts/pricing/quote/price_snapshot.go), [PriceSourceAllocation](../../pkg/contracts/pricing/quote/price_source_allocation.go), [PriceUnitAllocation](../../pkg/contracts/pricing/quote/price_unit_allocation.go) | Preserve the existing source ranges, package composition, tax and immutable quote identity. |
| Promotion participation and arithmetic | [PricePromotionAllocation](../../pkg/contracts/pricing/quote/price_promotion_allocation.go), [PricePromotionStageRun](../../pkg/contracts/pricing/quote/price_promotion_stage_run.go), [PromotionGroupSpan](../../pkg/contracts/pricing/quote/promotion_group_span.go) | Reuse unit ranges, roles, exact stage values, package/group ordinals and settled contributions. |
| Rounding and retained identity | [PromotionStageRoundingPool](../../pkg/contracts/pricing/quote/promotion_stage_rounding_pool.go), [RoundingEvidence](../../pkg/contracts/pricing/quote/rounding_evidence.go), [PromotionStageLineage](../../pkg/contracts/pricing/quote/promotion_stage_lineage.go) | Preserve original rounding awards and the root quote ID, fingerprint, line and unit interval. |

The existing lineage binds a run to an independently authenticated root. It
does not provide a standalone owner lookup or authenticate the predecessor
chain. Pricing's amendment source and Orders' accepted quote history already
carry quote lookup coordinates; adding the same coordinates to every shared
lineage is not currently required.

These models do not expose a complete candidate population, every authored
requirement, or every intermediate book/net event for independent downstream
recomputation. Pricing's detailed manifests, search checkpoints and calculation
audits remain service-owned. A future shared proof addition needs a concrete
producer/consumer requirement that the existing representations cannot meet.
The existing snapshot does not, by itself, prove global optimality.

## POS sale, tender and inventory evidence

| Information | Existing models | Reuse boundary |
| --- | --- | --- |
| POS context | [POSAttribution](../../pkg/contracts/orders/order/pos_attribution.go) | Reuse depot, register, terminal, session and operator attribution. |
| External payment identity | [Payment](../../pkg/contracts/payments/payment/payment.go) | Use the Payments-owned record and typed provider reference when authenticating a payment; copied status is not a durable fence. |
| Original gift-card and points debits | [GiftCardRedemptionSnapshot](../../pkg/contracts/orders/order/gift_card_redemption_snapshot.go), [PointRedemptionSnapshot](../../pkg/contracts/orders/order/point_redemption_snapshot.go) | Preserve the original reservation and wallet/ledger transaction identifiers and money. A snapshot alone does not establish that the debit remains valid. |
| Cash adjustment | [CashRoundingSnapshot](../../pkg/contracts/payments/receipt/cash_rounding_snapshot.go) | Reuse exact due, cash due, adjustment and rounding operands without changing their existing semantics. |
| Frozen reservation sources | [StockReservationAllocation](../../pkg/contracts/supply/inventory/stock_reservation_allocation.go), [OrderAllocationEvidence](../../pkg/contracts/supply/fulfilment/order_allocation_evidence.go) | Retain native source variants, revision and package composition; the existing evidence also binds a reservation/allocation to its order item. |
| Recorded inventory mutation | [StockMovement](../../pkg/contracts/supply/inventory/stock_movement.go) | Reuse the Supply-owned movement record. A movement alone is not a complete per-sale terminal receipt. |

The existing [CustomerPaymentSummary](../../pkg/contracts/payments/payment/customer_payment_summary.go)
and [CustomerPaymentAllocation](../../pkg/contracts/payments/payment/customer_payment_allocation.go)
are customer-facing projections, not tender authority. The Payments
[Settlement](../../pkg/contracts/payments/settlement/settlement.go) model describes
terminal/batch reconciliation, not an order's stock settlement.

## Additions deferred pending consumer contracts

Existing models supply reusable facts, but do not yet define the immutable
obligation tying a paid order generation to its quote, source allocations and
complete tender, or the corresponding durable owner holds and terminal outcome.
Those bindings still require provider/consumer design and qualification.

Compose the existing facts in the owning service first. Do not promote duplicate
tender unions, reservation wrappers, refund barriers, hold observations or abort
commands into this module merely because multiple services participate in the
workflow. A shared record needs a named domain owner and at least two independent
consumers requiring identical semantics. Routes, transport DTOs, authorization,
canonical digest execution and workflow transitions remain service-owned.

Preserve existing JSON names, optionality, array semantics and exact money.
Optional fields can still change serialized evidence and its digest when set;
`omitempty` alone does not establish compatibility with consumers that verify
hashes. Any later shared addition needs explicit wire/digest fixtures and an
appropriately versioned release before consumers adopt it.
