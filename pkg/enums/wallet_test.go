package enums_test

import (
	"testing"

	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/enums/wallet"
)

func TestWalletEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "walletenum.WalletInstrumentType", valid: []stringEnum{walletenum.WalletInstrumentTypePoints, walletenum.WalletInstrumentTypeGiftCard, walletenum.WalletInstrumentTypeVoucher, walletenum.WalletInstrumentTypeCoupon, walletenum.WalletInstrumentTypeReward}, invalid: walletenum.WalletInstrumentType("__invalid__")},
		{name: "walletenum.WalletExportFormat", valid: []stringEnum{walletenum.WalletExportFormatJSON, walletenum.WalletExportFormatCSVZip}, invalid: walletenum.WalletExportFormat("__invalid__")},
		{name: "walletenum.WalletExportStatus", valid: []stringEnum{walletenum.WalletExportStatusPending, walletenum.WalletExportStatusRunning, walletenum.WalletExportStatusCompleted, walletenum.WalletExportStatusFailed, walletenum.WalletExportStatusExpired}, invalid: walletenum.WalletExportStatus("__invalid__")},
		{name: "walletenum.GiftCardStatus", valid: []stringEnum{walletenum.GiftCardStatusActive, walletenum.GiftCardStatusPartiallyRedeemed, walletenum.GiftCardStatusDepleted, walletenum.GiftCardStatusExpired, walletenum.GiftCardStatusVoid}, invalid: walletenum.GiftCardStatus("__invalid__")},
		{name: "walletenum.GiftCardTransactionReason", valid: []stringEnum{walletenum.GiftCardTransactionReasonIssue, walletenum.GiftCardTransactionReasonRedeem, walletenum.GiftCardTransactionReasonRefund, walletenum.GiftCardTransactionReasonTopUp, walletenum.GiftCardTransactionReasonExpire, walletenum.GiftCardTransactionReasonAdjust}, invalid: walletenum.GiftCardTransactionReason("__invalid__")},
	})
}
