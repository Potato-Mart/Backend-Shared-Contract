package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/wallet/wallet_enums"
)

func TestWalletEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "walletenum.CouponSource", valid: []stringEnum{wallet_enums.CouponSourceManual, wallet_enums.CouponSourceRFMComeback, wallet_enums.CouponSourceBirthday, wallet_enums.CouponSourceReferral, wallet_enums.CouponSourceSignupBonus, wallet_enums.CouponSourceCampaign}, invalid: wallet_enums.CouponSource("__invalid__")},
		{name: "walletenum.WalletPassPlatform", valid: []stringEnum{wallet_enums.WalletPassPlatformGoogle, wallet_enums.WalletPassPlatformApple}, invalid: wallet_enums.WalletPassPlatform("__invalid__")},
		{name: "walletenum.WalletPassBarcodeFormat", valid: []stringEnum{wallet_enums.WalletPassBarcodeFormatCode128}, invalid: wallet_enums.WalletPassBarcodeFormat("__invalid__")},
		{name: "walletenum.WalletInstrumentType", valid: []stringEnum{wallet_enums.WalletInstrumentTypePoints, wallet_enums.WalletInstrumentTypeGiftCard, wallet_enums.WalletInstrumentTypeVoucher, wallet_enums.WalletInstrumentTypeCoupon, wallet_enums.WalletInstrumentTypeReward}, invalid: wallet_enums.WalletInstrumentType("__invalid__")},
		{name: "walletenum.GiftCardStatus", valid: []stringEnum{wallet_enums.GiftCardStatusActive, wallet_enums.GiftCardStatusPartiallyRedeemed, wallet_enums.GiftCardStatusDepleted, wallet_enums.GiftCardStatusExpired, wallet_enums.GiftCardStatusVoid}, invalid: wallet_enums.GiftCardStatus("__invalid__")},
		{name: "walletenum.GiftCardTransactionReason", valid: []stringEnum{wallet_enums.GiftCardTransactionReasonIssue, wallet_enums.GiftCardTransactionReasonRedeem, wallet_enums.GiftCardTransactionReasonRefund, wallet_enums.GiftCardTransactionReasonTopUp, wallet_enums.GiftCardTransactionReasonExpire, wallet_enums.GiftCardTransactionReasonAdjust}, invalid: wallet_enums.GiftCardTransactionReason("__invalid__")},
		{name: "walletenum.VoucherStatus", valid: []stringEnum{wallet_enums.VoucherStatusIssued, wallet_enums.VoucherStatusReserved, wallet_enums.VoucherStatusRedeemed, wallet_enums.VoucherStatusExpired, wallet_enums.VoucherStatusVoid}, invalid: wallet_enums.VoucherStatus("__invalid__")},
		{name: "walletenum.CheckoutBenefitReservationStatus", valid: []stringEnum{wallet_enums.CheckoutBenefitReservationStatusReserved, wallet_enums.CheckoutBenefitReservationStatusCommitted, wallet_enums.CheckoutBenefitReservationStatusCancelled, wallet_enums.CheckoutBenefitReservationStatusExpired, wallet_enums.CheckoutBenefitReservationStatusPartiallyRefunded, wallet_enums.CheckoutBenefitReservationStatusRefunded}, invalid: wallet_enums.CheckoutBenefitReservationStatus("__invalid__")},
	})
	if got := wallet_enums.WalletPassPlatformGoogle.String(); got != "google_wallet" {
		t.Fatalf("google wallet platform wire value = %q", got)
	}
	if got := wallet_enums.WalletPassPlatformApple.String(); got != "apple_wallet" {
		t.Fatalf("apple wallet platform wire value = %q", got)
	}
	if got := wallet_enums.WalletPassBarcodeFormatCode128.String(); got != "code_128" {
		t.Fatalf("Code 128 wallet-pass wire value = %q", got)
	}
}
