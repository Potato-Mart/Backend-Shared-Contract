package enums_test

import (
	"testing"

	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/payment"
)

func TestPaymentEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "paymentenum.PaymentMethod", valid: []stringEnum{paymentenum.PaymentMethodCard, paymentenum.PaymentMethodCash, paymentenum.PaymentMethodQR, paymentenum.PaymentMethodBankTransfer, paymentenum.PaymentMethodLinePay, paymentenum.PaymentMethodApplePay, paymentenum.PaymentMethodGooglePay, paymentenum.PaymentMethodECPay, paymentenum.PaymentMethodManual, paymentenum.PaymentMethodGiftCard, paymentenum.PaymentMethodEFTPOS, paymentenum.PaymentMethodMOTO, paymentenum.PaymentMethodCashout}, invalid: paymentenum.PaymentMethod("__invalid__")},
		{name: "paymentenum.PaymentRecordStatus", valid: []stringEnum{paymentenum.PaymentRecordStatusPending, paymentenum.PaymentRecordStatusProcessing, paymentenum.PaymentRecordStatusCompleted, paymentenum.PaymentRecordStatusFailed, paymentenum.PaymentRecordStatusCancelled, paymentenum.PaymentRecordStatusRefunded, paymentenum.PaymentRecordStatusAwaitingAction, paymentenum.PaymentRecordStatusUnknown}, invalid: paymentenum.PaymentRecordStatus("__invalid__")},
		{name: "paymentenum.PaymentStatus", valid: []stringEnum{paymentenum.PaymentStatusUnknown, paymentenum.PaymentStatusUnpaid, paymentenum.PaymentStatusPending, paymentenum.PaymentStatusPaid, paymentenum.PaymentStatusPartiallyPaid, paymentenum.PaymentStatusRefunded, paymentenum.PaymentStatusPartialRefunded}, invalid: paymentenum.PaymentStatus("__invalid__")},
		{name: "paymentenum.RecoveryDecision", valid: []stringEnum{paymentenum.RecoveryDecisionPending, paymentenum.RecoveryDecisionApproved, paymentenum.RecoveryDecisionDeclined}, invalid: paymentenum.RecoveryDecision("__invalid__")},
		{name: "paymentenum.SettlementType", valid: []stringEnum{paymentenum.SettlementTypeSettlement, paymentenum.SettlementTypeEnquiry}, invalid: paymentenum.SettlementType("__invalid__")},
		{name: "paymentenum.TerminalConnectionMode", valid: []stringEnum{paymentenum.TerminalConnectionModeCloudSync, paymentenum.TerminalConnectionModeCloudAsync, paymentenum.TerminalConnectionModeLocal}, invalid: paymentenum.TerminalConnectionMode("__invalid__")},
		{name: "paymentenum.TerminalProvider", valid: []stringEnum{paymentenum.TerminalProviderMx51}, invalid: paymentenum.TerminalProvider("__invalid__")},
		{name: "paymentenum.TerminalRefundType", valid: []stringEnum{paymentenum.TerminalRefundTypeReferenced, paymentenum.TerminalRefundTypeUnreferenced}, invalid: paymentenum.TerminalRefundType("__invalid__")},
		{name: "paymentenum.TerminalStatus", valid: []stringEnum{paymentenum.TerminalStatusRegistered, paymentenum.TerminalStatusActive, paymentenum.TerminalStatusDeregistered, paymentenum.TerminalStatusExpired, paymentenum.TerminalStatusError}, invalid: paymentenum.TerminalStatus("__invalid__")},
		{name: "paymentenum.TerminalTxFinancialStatus", valid: []stringEnum{paymentenum.TerminalTxFinancialStatusApproved, paymentenum.TerminalTxFinancialStatusDeclined, paymentenum.TerminalTxFinancialStatusCancelled, paymentenum.TerminalTxFinancialStatusUnknown}, invalid: paymentenum.TerminalTxFinancialStatus("__invalid__")},
		{name: "paymentenum.TerminalTxStatus", valid: []stringEnum{paymentenum.TerminalTxStatusUnknown, paymentenum.TerminalTxStatusPending, paymentenum.TerminalTxStatusAwaitingAction, paymentenum.TerminalTxStatusFinalised, paymentenum.TerminalTxStatusOverridePending, paymentenum.TerminalTxStatusOverrideResolved}, invalid: paymentenum.TerminalTxStatus("__invalid__")},
		{name: "paymentenum.TerminalTxType", valid: []stringEnum{paymentenum.TerminalTxTypePurchase, paymentenum.TerminalTxTypeRefund, paymentenum.TerminalTxTypeReversal, paymentenum.TerminalTxTypeCashout, paymentenum.TerminalTxTypePurchaseWithCashout, paymentenum.TerminalTxTypeMOTO, paymentenum.TerminalTxTypeSettlement, paymentenum.TerminalTxTypeSettlementEnquiry}, invalid: paymentenum.TerminalTxType("__invalid__")},
	})
}
