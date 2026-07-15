package enums_test

import (
	"testing"

	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/wallet"
)

func TestWalletEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "walletenum.WalletInstrumentType", valid: []stringEnum{walletenum.WalletInstrumentTypePoints, walletenum.WalletInstrumentTypeGiftCard, walletenum.WalletInstrumentTypeVoucher, walletenum.WalletInstrumentTypeCoupon, walletenum.WalletInstrumentTypeReward}, invalid: walletenum.WalletInstrumentType("__invalid__")},
		{name: "walletenum.GiftCardStatus", valid: []stringEnum{walletenum.GiftCardStatusActive, walletenum.GiftCardStatusPartiallyRedeemed, walletenum.GiftCardStatusDepleted, walletenum.GiftCardStatusExpired, walletenum.GiftCardStatusVoid}, invalid: walletenum.GiftCardStatus("__invalid__")},
		{name: "walletenum.GiftCardTransactionReason", valid: []stringEnum{walletenum.GiftCardTransactionReasonIssue, walletenum.GiftCardTransactionReasonRedeem, walletenum.GiftCardTransactionReasonRefund, walletenum.GiftCardTransactionReasonTopUp, walletenum.GiftCardTransactionReasonExpire, walletenum.GiftCardTransactionReasonAdjust}, invalid: walletenum.GiftCardTransactionReason("__invalid__")},
		{name: "walletenum.VoucherStatus", valid: []stringEnum{walletenum.VoucherStatusIssued, walletenum.VoucherStatusReserved, walletenum.VoucherStatusRedeemed, walletenum.VoucherStatusExpired, walletenum.VoucherStatusVoid}, invalid: walletenum.VoucherStatus("__invalid__")},
		{name: "walletenum.CheckoutBenefitReservationStatus", valid: []stringEnum{walletenum.CheckoutBenefitReservationStatusReserved, walletenum.CheckoutBenefitReservationStatusCommitted, walletenum.CheckoutBenefitReservationStatusCancelled, walletenum.CheckoutBenefitReservationStatusExpired, walletenum.CheckoutBenefitReservationStatusPartiallyRefunded, walletenum.CheckoutBenefitReservationStatusRefunded}, invalid: walletenum.CheckoutBenefitReservationStatus("__invalid__")},
	})
}
