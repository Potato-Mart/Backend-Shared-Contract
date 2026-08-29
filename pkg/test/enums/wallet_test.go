package enums_test

import (
	"testing"

	membership_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/wallet_enums"
)

func TestWalletEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "walletenum.CouponSource", valid: []stringEnum{wallet_enums.CouponSourceManual, wallet_enums.CouponSourceRFMComeback, wallet_enums.CouponSourceBirthday, wallet_enums.CouponSourceReferral, wallet_enums.CouponSourceSignupBonus, wallet_enums.CouponSourceCampaign}, invalid: wallet_enums.CouponSource("__invalid__")},
		{name: "walletenum.WalletPassPlatform", valid: []stringEnum{membership_enums.WalletPassPlatformGoogle, membership_enums.WalletPassPlatformApple}, invalid: membership_enums.WalletPassPlatform("__invalid__")},
		{name: "walletenum.WalletPassBarcodeFormat", valid: []stringEnum{membership_enums.WalletPassBarcodeFormatCode128}, invalid: membership_enums.WalletPassBarcodeFormat("__invalid__")},
		{name: "walletenum.WalletInstrumentType", valid: []stringEnum{wallet_enums.WalletInstrumentTypePoints, wallet_enums.WalletInstrumentTypeGiftCard, wallet_enums.WalletInstrumentTypeVoucher, wallet_enums.WalletInstrumentTypeCoupon, wallet_enums.WalletInstrumentTypeReward}, invalid: wallet_enums.WalletInstrumentType("__invalid__")},
		{name: "walletenum.GiftCardStatus", valid: []stringEnum{wallet_enums.GiftCardStatusActive, wallet_enums.GiftCardStatusPartiallyRedeemed, wallet_enums.GiftCardStatusDepleted, wallet_enums.GiftCardStatusExpired, wallet_enums.GiftCardStatusVoid}, invalid: wallet_enums.GiftCardStatus("__invalid__")},
		{name: "walletenum.GiftCardTransactionReason", valid: []stringEnum{wallet_enums.GiftCardTransactionReasonIssue, wallet_enums.GiftCardTransactionReasonRedeem, wallet_enums.GiftCardTransactionReasonRefund, wallet_enums.GiftCardTransactionReasonTopUp, wallet_enums.GiftCardTransactionReasonExpire, wallet_enums.GiftCardTransactionReasonAdjust}, invalid: wallet_enums.GiftCardTransactionReason("__invalid__")},
		{name: "walletenum.VoucherStatus", valid: []stringEnum{wallet_enums.VoucherStatusIssued, wallet_enums.VoucherStatusReserved, wallet_enums.VoucherStatusRedeemed, wallet_enums.VoucherStatusExpired, wallet_enums.VoucherStatusVoid}, invalid: wallet_enums.VoucherStatus("__invalid__")},
		{name: "walletenum.CheckoutBenefitReservationStatus", valid: []stringEnum{wallet_enums.CheckoutBenefitReservationStatusReserved, wallet_enums.CheckoutBenefitReservationStatusCommitted, wallet_enums.CheckoutBenefitReservationStatusCancelled, wallet_enums.CheckoutBenefitReservationStatusExpired, wallet_enums.CheckoutBenefitReservationStatusPartiallyRefunded, wallet_enums.CheckoutBenefitReservationStatusRefunded}, invalid: wallet_enums.CheckoutBenefitReservationStatus("__invalid__")},
		{name: "walletenum.PointLedgerReason", valid: []stringEnum{wallet_enums.PointLedgerReasonOrder, wallet_enums.PointLedgerReasonBirthday, wallet_enums.PointLedgerReasonRedeem, wallet_enums.PointLedgerReasonRefund, wallet_enums.PointLedgerReasonRewardRedeem, wallet_enums.PointLedgerReasonAdminAdjust, wallet_enums.PointLedgerReasonExpired, wallet_enums.PointLedgerReasonReferral, wallet_enums.PointLedgerReasonSignupBonus, wallet_enums.PointLedgerReasonTierUpgrade, wallet_enums.PointLedgerReasonDebtIncurred, wallet_enums.PointLedgerReasonDebtRepaid, wallet_enums.PointLedgerReasonManual, wallet_enums.PointLedgerReasonRewardRedeemReversal}, invalid: wallet_enums.PointLedgerReason("__invalid__")},
		{name: "walletenum.PointAwardStatus", valid: []stringEnum{wallet_enums.PointAwardStatusIneligible, wallet_enums.PointAwardStatusDisabled, wallet_enums.PointAwardStatusPending, wallet_enums.PointAwardStatusAwarded, wallet_enums.PointAwardStatusFailed}, invalid: wallet_enums.PointAwardStatus("__invalid__")},
		{name: "walletenum.PointRedemptionType", valid: []stringEnum{wallet_enums.PointRedemptionTypeCheckoutDiscount, wallet_enums.PointRedemptionTypeRewardCatalog}, invalid: wallet_enums.PointRedemptionType("__invalid__")},
		{name: "walletenum.RewardRedemptionStatus", valid: []stringEnum{wallet_enums.RewardRedemptionStatusReserved, wallet_enums.RewardRedemptionStatusRedeemed, wallet_enums.RewardRedemptionStatusCancelled, wallet_enums.RewardRedemptionStatusExpired}, invalid: wallet_enums.RewardRedemptionStatus("__invalid__")},
		{name: "walletenum.RewardType", valid: []stringEnum{wallet_enums.RewardTypeOrderDiscount, wallet_enums.RewardTypeProduct, wallet_enums.RewardTypeFreeShipping, wallet_enums.RewardTypeVoucher, wallet_enums.RewardTypeCoupon, wallet_enums.RewardTypeGiftCard, wallet_enums.RewardTypeExternal}, invalid: wallet_enums.RewardType("__invalid__")},
		{name: "walletenum.ExternalRewardFulfilmentStatus", valid: []stringEnum{wallet_enums.ExternalRewardFulfilmentStatusPending, wallet_enums.ExternalRewardFulfilmentStatusProvisioned, wallet_enums.ExternalRewardFulfilmentStatusFailed, wallet_enums.ExternalRewardFulfilmentStatusRevoked}, invalid: wallet_enums.ExternalRewardFulfilmentStatus("__invalid__")},
		{name: "walletenum.PointReservationStatus", valid: []stringEnum{wallet_enums.PointReservationStatusReserved, wallet_enums.PointReservationStatusCommitted, wallet_enums.PointReservationStatusCancelled, wallet_enums.PointReservationStatusExpired}, invalid: wallet_enums.PointReservationStatus("__invalid__")},
	})
	if got := membership_enums.WalletPassPlatformGoogle.String(); got != "google_wallet" {
		t.Fatalf("google wallet platform wire value = %q", got)
	}
	if got := membership_enums.WalletPassPlatformApple.String(); got != "apple_wallet" {
		t.Fatalf("apple wallet platform wire value = %q", got)
	}
	if got := membership_enums.WalletPassBarcodeFormatCode128.String(); got != "code_128" {
		t.Fatalf("Code 128 wallet-pass wire value = %q", got)
	}
}
