# Net tier-refund proof adoption

Pin `github.com/Potato-Mart/Backend-Shared-Contract/v39 v39.3.0` and the matching CI pin before dependent source changes. The model release does not supply runtime reader/preparation/release interfaces, grants, deployment, verified lineage or selection authority. Orders, Pricing and Payments agreed the projection; remaining source workflows and cross-service acceptance are held separately.

`RefundCompletedEvent.tier_refund_proof` is optional. Its type is `pricing/membership.TierRefundProof`:

| JSON field | Go type | Meaning |
|---|---|---|
| basis | membership_enums.TierRefundProofBasis | Required quoted_credit or excluded_gift_purchase |
| proof_id | string | Immutable native Pricing reservation identity |
| proof_fingerprint | string | Opaque digest of native proof binding and canonical selection |
| allocation_id | string | Original credited or witnessed exclusion allocation lineage |
| allocation_version | int64 | Positive immutable Pricing allocation format/algorithm version, initially 1 |
| allocation_fingerprint | string | Exact frozen original allocation evidence fingerprint |
| quote_key | string, omitempty | Exact original credited quote key for quoted_credit |
| quote_fingerprint | string, omitempty | Exact original opaque quote revision for quoted_credit |
| net_qualifying_spend_reversal | money.Money | Required nonnegative incremental NET tier amount, including currency-bearing verified zero |

Opaque strings are compared exactly without trimming, conversion, aliasing or rehashing. AllocationVersion is not a numeric quote revision. Standard encoding/json preserves these data fields but does not validate their business evidence.

For `quoted_credit`, both quote strings are nonempty, including coupon-zero grants. For `excluded_gift_purchase`, both are omitted and Money is verified zero; the native record must witness the protected ORIGINAL dedicated gift purchase identity/classification. Mutable SKU/tags, absent quote/credit, or a settlement response alone cannot establish exclusion. No sentinel quote or future arbitrary exclusion is introduced.

Missing proof is missing/unverified evidence, not zero. Guest/wholesale tier-not-applicable decisions require protected native Orders original buyer/version/selection and exact refund intent; services may omit proof for that decision. Neither absent customer nor absent proof establishes applicability. Unquoted or ambiguous retail remains pending.

Pricing's immutable native reservation is authority: every projected field and outer refund_id/order/customer/market/currency binding must match. Local lookup supplies independent selection and cumulative/replay verification, so no cumulative amount is duplicated here. Net evidence remains distinct from `qualifying_spend_reversal` (released gross earned-point clawback), `points_to_restore` (redeemed points) and refund settlement `amount`. None is a fallback for missing net authority.

Orders must preserve original accepted version/item/component/package/ordinal lineage and original gift classification prospectively, with no historical reconstruction or SKU/array-position joins. Pricing must verify retained native quote evidence against the protected Orders accepted version, durably seal original net allocation with actual paid credit (including witnessed zero), and retain original-grant lineage across amendments. Newly added units do not inherit credit from repricing. Pricing confirmed its native quote records are durable and insert-only; no new pre-effect pin callback is required merely for quote lifetime. Allocation arithmetic, unit selection and service DTOs remain Pricing/Orders-owned.

Payments must durably establish RefundID/request/selection authority before network preparation, call outside retrying database transactions, and freeze returned proof/complete selection before provider, gift/wallet or terminal effects. Ambiguous responses retry identical authority. Native Pricing reservations fence selected original units and cumulative completed/reserved credit against the original grant. Completion consumes proof and writes tier reversal atomically in Pricing; completion time selects its reversal year.

Confirmed pre-dispatch cancellation or final provider failure needs durable RefundID/selection tombstones, including an ambiguous prepare whose proof ID was not recovered. Completed proof cannot release. Timeout, pending/unknown provider status or transport failure never release. These requirements do not publish a route or runtime permission.

Remaining amount-only, edited-order and multi-capture selection/batch assignment designs are service-local and fail closed until independently verified. FullOrderRefund or remaining captured cash alone is not all-unit proof; no cash-proportional unit guess. Model publication does not clear source-ticket or production activation holds. Already-started legacy financial truth is recorded without another provider effect or reconstructed historical net proof.

Fixtures: `pkg/contracts/pubsub/payments/testdata/tier_refund_proof.json` and `tier_refund_proof_json_test.go` show quoted positive/zero, witnessed gift zero and unchanged legacy omission/gross values. Runtime event versions remain owner source authority; this optional addition causes no version bump.
