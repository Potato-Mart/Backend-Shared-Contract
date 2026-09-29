package notification

import "time"

// GiftCardDeliveredEvent records successful recipient delivery after issuance.
// Notification owns this fact for every issuance source. Orders correlates
// IssuanceID with its actual purchase, if any; delivery of a reward or refund
// replacement does not imply completion of a purchase order.
// Delivery material, gift codes, PINs and recipient contact data remain protected
// service-owned data. Claiming the gift card is an independent customer action.
type GiftCardDeliveredEvent struct {
	IssuanceID  string    `json:"issuance_id"`
	DeliveredAt time.Time `json:"delivered_at"`
}
