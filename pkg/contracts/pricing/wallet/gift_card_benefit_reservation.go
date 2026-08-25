package wallet

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/wallet/wallet_enums"
)

// GiftCardBenefitReservation is one ordered gift-card allocation.
type GiftCardBenefitReservation struct {
	GiftCardCode   string                                        `json:"gift_card_code"`
	AppliedAmount  money.Money                                   `json:"applied_amount"`
	RefundedAmount money.Money                                   `json:"refunded_amount"`
	Status         wallet_enums.CheckoutBenefitReservationStatus `json:"status"`
}
