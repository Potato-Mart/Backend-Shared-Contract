package notification

import "time"

// GiftCardDeliveredEvent records successful recipient delivery after issuance.
// Notification owns this fact; Orders correlates IssuanceID with its purchase.
// Delivery material, gift codes, PINs and recipient contact data remain protected
// service-owned data. Claiming the gift card is an independent customer action.
type GiftCardDeliveredEvent struct {
	IssuanceID  string    `json:"issuance_id"`
	DeliveredAt time.Time `json:"delivered_at"`
}
