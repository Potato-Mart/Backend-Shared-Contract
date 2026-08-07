package enums_test

import (
	"testing"

	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/payments/payment"
	settlementenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/payments/settlement"
	terminalenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/payments/terminal"
)

func TestPaymentEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "paymentenum.PaymentMethod", valid: []stringEnum{paymentenum.PaymentMethodCard, paymentenum.PaymentMethodCash, paymentenum.PaymentMethodQR, paymentenum.PaymentMethodBankTransfer, paymentenum.PaymentMethodLinePay, paymentenum.PaymentMethodApplePay, paymentenum.PaymentMethodGooglePay, paymentenum.PaymentMethodECPay, paymentenum.PaymentMethodManual, paymentenum.PaymentMethodGiftCard, paymentenum.PaymentMethodEFTPOS, paymentenum.PaymentMethodMOTO, paymentenum.PaymentMethodCashout}, invalid: paymentenum.PaymentMethod("__invalid__")},
		{name: "paymentenum.PaymentRecordStatus", valid: []stringEnum{paymentenum.PaymentRecordStatusPending, paymentenum.PaymentRecordStatusProcessing, paymentenum.PaymentRecordStatusCompleted, paymentenum.PaymentRecordStatusFailed, paymentenum.PaymentRecordStatusCancelled, paymentenum.PaymentRecordStatusRefunded, paymentenum.PaymentRecordStatusAwaitingAction, paymentenum.PaymentRecordStatusUnknown}, invalid: paymentenum.PaymentRecordStatus("__invalid__")},
		{name: "paymentenum.PaymentStatus", valid: []stringEnum{paymentenum.PaymentStatusUnknown, paymentenum.PaymentStatusUnpaid, paymentenum.PaymentStatusPending, paymentenum.PaymentStatusPaid, paymentenum.PaymentStatusPartiallyPaid, paymentenum.PaymentStatusRefunded, paymentenum.PaymentStatusPartialRefunded}, invalid: paymentenum.PaymentStatus("__invalid__")},
		{name: "paymentenum.RecoveryDecision", valid: []stringEnum{paymentenum.RecoveryDecisionPending, paymentenum.RecoveryDecisionApproved, paymentenum.RecoveryDecisionDeclined}, invalid: paymentenum.RecoveryDecision("__invalid__")},
		{name: "settlementenum.SettlementType", valid: []stringEnum{settlementenum.SettlementTypeSettlement, settlementenum.SettlementTypeEnquiry}, invalid: settlementenum.SettlementType("__invalid__")},
		{name: "terminalenum.TerminalConnectionMode", valid: []stringEnum{terminalenum.TerminalConnectionModeCloudSync, terminalenum.TerminalConnectionModeCloudAsync, terminalenum.TerminalConnectionModeLocal}, invalid: terminalenum.TerminalConnectionMode("__invalid__")},
		{name: "terminalenum.TerminalProvider", valid: []stringEnum{terminalenum.TerminalProviderMx51}, invalid: terminalenum.TerminalProvider("__invalid__")},
		{name: "terminalenum.TerminalRefundType", valid: []stringEnum{terminalenum.TerminalRefundTypeReferenced, terminalenum.TerminalRefundTypeUnreferenced}, invalid: terminalenum.TerminalRefundType("__invalid__")},
		{name: "terminalenum.TerminalStatus", valid: []stringEnum{terminalenum.TerminalStatusRegistered, terminalenum.TerminalStatusActive, terminalenum.TerminalStatusDeregistered, terminalenum.TerminalStatusExpired, terminalenum.TerminalStatusError}, invalid: terminalenum.TerminalStatus("__invalid__")},
		{name: "terminalenum.TerminalTxFinancialStatus", valid: []stringEnum{terminalenum.TerminalTxFinancialStatusApproved, terminalenum.TerminalTxFinancialStatusDeclined, terminalenum.TerminalTxFinancialStatusCancelled, terminalenum.TerminalTxFinancialStatusUnknown}, invalid: terminalenum.TerminalTxFinancialStatus("__invalid__")},
		{name: "terminalenum.TerminalTxStatus", valid: []stringEnum{terminalenum.TerminalTxStatusUnknown, terminalenum.TerminalTxStatusPending, terminalenum.TerminalTxStatusAwaitingAction, terminalenum.TerminalTxStatusFinalised, terminalenum.TerminalTxStatusOverridePending, terminalenum.TerminalTxStatusOverrideResolved}, invalid: terminalenum.TerminalTxStatus("__invalid__")},
		{name: "terminalenum.TerminalTxType", valid: []stringEnum{terminalenum.TerminalTxTypePurchase, terminalenum.TerminalTxTypeRefund, terminalenum.TerminalTxTypeReversal, terminalenum.TerminalTxTypeCashout, terminalenum.TerminalTxTypePurchaseWithCashout, terminalenum.TerminalTxTypeMOTO, terminalenum.TerminalTxTypeSettlement, terminalenum.TerminalTxTypeSettlementEnquiry}, invalid: terminalenum.TerminalTxType("__invalid__")},
	})
}
