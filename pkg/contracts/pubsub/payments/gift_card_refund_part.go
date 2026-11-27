package payments

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// GiftCardRefundPart is the money a completed refund put back on the gift
// cards held by one Pricing benefit reservation. BenefitReservationID names
// the reservation and GiftCardRefundAmount how much of it the refund
// restored.
//
// Producers must send a non-empty BenefitReservationID and a positive
// GiftCardRefundAmount. Every part of one event uses the same currency, and
// a reservation appears at most once per event. Consumers must reject an
// event that breaks these rules rather than guess.
type GiftCardRefundPart struct {
	BenefitReservationID string      `json:"benefit_reservation_id"`
	GiftCardRefundAmount money.Money `json:"gift_card_refund_amount"`
}
