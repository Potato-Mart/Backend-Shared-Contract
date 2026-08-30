package reservation

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/wallet_enums"
)

// VoucherBenefitReservation is an exclusive checkout voucher hold.
type VoucherBenefitReservation struct {
	VoucherCode   string                                        `json:"voucher_code"`
	AppliedAmount money.Money                                   `json:"applied_amount"`
	Status        wallet_enums.CheckoutBenefitReservationStatus `json:"status"`
}
