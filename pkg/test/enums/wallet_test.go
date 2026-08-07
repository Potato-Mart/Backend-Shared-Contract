package enums_test

import (
	"testing"

	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/pricing/wallet"
)

func TestWalletEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "walletenum.WalletPassPlatform", valid: []stringEnum{walletenum.WalletPassPlatformGoogle, walletenum.WalletPassPlatformApple}, invalid: walletenum.WalletPassPlatform("__invalid__")},
		{name: "walletenum.WalletPassBarcodeFormat", valid: []stringEnum{walletenum.WalletPassBarcodeFormatCode128}, invalid: walletenum.WalletPassBarcodeFormat("__invalid__")},
		{name: "walletenum.WalletInstrumentType", valid: []stringEnum{walletenum.WalletInstrumentTypePoints, walletenum.WalletInstrumentTypeGiftCard, walletenum.WalletInstrumentTypeVoucher, walletenum.WalletInstrumentTypeCoupon, walletenum.WalletInstrumentTypeReward}, invalid: walletenum.WalletInstrumentType("__invalid__")},
		{name: "walletenum.GiftCardStatus", valid: []stringEnum{walletenum.GiftCardStatusActive, walletenum.GiftCardStatusPartiallyRedeemed, walletenum.GiftCardStatusDepleted, walletenum.GiftCardStatusExpired, walletenum.GiftCardStatusVoid}, invalid: walletenum.GiftCardStatus("__invalid__")},
		{name: "walletenum.GiftCardTransactionReason", valid: []stringEnum{walletenum.GiftCardTransactionReasonIssue, walletenum.GiftCardTransactionReasonRedeem, walletenum.GiftCardTransactionReasonRefund, walletenum.GiftCardTransactionReasonTopUp, walletenum.GiftCardTransactionReasonExpire, walletenum.GiftCardTransactionReasonAdjust}, invalid: walletenum.GiftCardTransactionReason("__invalid__")},
		{name: "walletenum.VoucherStatus", valid: []stringEnum{walletenum.VoucherStatusIssued, walletenum.VoucherStatusReserved, walletenum.VoucherStatusRedeemed, walletenum.VoucherStatusExpired, walletenum.VoucherStatusVoid}, invalid: walletenum.VoucherStatus("__invalid__")},
		{name: "walletenum.CheckoutBenefitReservationStatus", valid: []stringEnum{walletenum.CheckoutBenefitReservationStatusReserved, walletenum.CheckoutBenefitReservationStatusCommitted, walletenum.CheckoutBenefitReservationStatusCancelled, walletenum.CheckoutBenefitReservationStatusExpired, walletenum.CheckoutBenefitReservationStatusPartiallyRefunded, walletenum.CheckoutBenefitReservationStatusRefunded}, invalid: walletenum.CheckoutBenefitReservationStatus("__invalid__")},
	})
	if got := walletenum.WalletPassPlatformGoogle.String(); got != "google_wallet" {
		t.Fatalf("google wallet platform wire value = %q", got)
	}
	if got := walletenum.WalletPassPlatformApple.String(); got != "apple_wallet" {
		t.Fatalf("apple wallet platform wire value = %q", got)
	}
	if got := walletenum.WalletPassBarcodeFormatCode128.String(); got != "code_128" {
		t.Fatalf("Code 128 wallet-pass wire value = %q", got)
	}
}
