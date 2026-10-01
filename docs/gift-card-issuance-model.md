# Gift-card issuance evidence

Pricing owns `pubsub/pricing.GiftCardIssuedEventV2`. It identifies a committed issuance using `issuance_id`, `source`, `issued_value`, optional `market_code` and `country_code`, and `issued_at`.

Sources are `purchase`, `membership_reward` and `refund_replacement`. `issued_value` is the original initial credited value, including any purchase bonus. It is not buyer charge, later spendable balance, refund payment amount, points spent or denomination evidence. A card charged at AUD 50 with an AUD 5 bonus records AUD 55 as issued value.

Issuance identity and source association are immutable owner facts. Missing geography supplies no authority to infer a market or country. Card codes, PINs, customer identity and contact details are excluded from the event.

Notification retrieves protected delivery material through an authorized Pricing-private lookup. `notification.gift_card_delivered` carries issuance identity and delivery time. Orders completes purchase fulfilment only for an issuance genuinely correlated to that purchase; reward or replacement delivery does not establish purchase fulfilment. Claim is independent of delivery.

Pricing owns issuance, credentials, ledger consistency and idempotency. Notification owns delivery. No private transport DTO or runtime behavior is implemented by this module.
