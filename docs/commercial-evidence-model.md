# Commercial evidence models

Money is represented in integer minor units with an explicit currency. Pricing owns commercial calculation; Orders freezes accepted outcomes and Payments records settlement facts.

Order subtotal is goods before promotional discounts. `discount_amount` is promotional discount, not points or voucher tender. Total is the commercial obligation including shipping, applicable tax, tip and surcharge; inclusive tax must not be counted twice. Points, vouchers and gift cards settle that obligation without changing Total. `order.paid.amount_paid` is settlement evidence, not qualifying goods value.

`quote_key` and `quote_fingerprint` preserve the accepted quote key and opaque Pricing revision exactly. Do not parse, truncate, rehash or numerically convert the fingerprint. Numeric `quote_revision` cannot independently verify an opaque revision. If both representations are supplied, the owner verifies consistency.

`deferred_payment = {actor, authorized_at, reason?}` is Orders-stamped staff authorization evidence. Absence grants no unpaid-fulfilment authority. It does not settle payment or disable checkout expiry.

Tier thresholds are configured minor-unit values. `progress_basis_points` describes progress toward the next threshold; `tier_progress_basis_points` describes progress between current and next thresholds. Both are capped at 10000; at maximum tier they are 10000. Invalid or absent thresholds omit optional interval progress. Pricing performs the calculation.

Tax categories and effective-dated rules live under `pricing/tax`. Rates use a numerator and positive denominator, with an inclusion basis, market, revision and flat audit fields. Pricing owns interval validation and tax decisions.

`ShippingArrivalRule.delivery_company_code` scopes arrival rules to one delivery company. `EXPIRY_HOLD` identifies expired-stock storage; Supply owns allocation eligibility. `PurchaseReceiptItem.date_mark` is optional lot evidence and does not belong on customer receipt lines.

Notification preferences are channel-first: `channels[{channel, topics:[{topic_code, enabled, destination_codes?}]}]`. Choices do not authorize delivery; required topics, destination consent and unsubscribe policy remain owner-enforced.

Executable model fixtures include `pkg/test/testdata/configurable_contract.json`. This module contains no calculations, owner decisions or runtime workflows.
