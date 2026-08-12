package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/payments/settlement/settlement_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/payments/terminal/terminal_enums"
)

func TestPaymentEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "paymentenum.PaymentMethod", valid: []stringEnum{payment_enums.PaymentMethodCard, payment_enums.PaymentMethodCash, payment_enums.PaymentMethodQR, payment_enums.PaymentMethodBankTransfer, payment_enums.PaymentMethodLinePay, payment_enums.PaymentMethodApplePay, payment_enums.PaymentMethodGooglePay, payment_enums.PaymentMethodECPay, payment_enums.PaymentMethodManual, payment_enums.PaymentMethodGiftCard, payment_enums.PaymentMethodEFTPOS, payment_enums.PaymentMethodMOTO, payment_enums.PaymentMethodCashout}, invalid: payment_enums.PaymentMethod("__invalid__")},
		{name: "paymentenum.PaymentRecordStatus", valid: []stringEnum{payment_enums.PaymentRecordStatusPending, payment_enums.PaymentRecordStatusProcessing, payment_enums.PaymentRecordStatusCompleted, payment_enums.PaymentRecordStatusFailed, payment_enums.PaymentRecordStatusCancelled, payment_enums.PaymentRecordStatusRefunded, payment_enums.PaymentRecordStatusAwaitingAction, payment_enums.PaymentRecordStatusUnknown}, invalid: payment_enums.PaymentRecordStatus("__invalid__")},
		{name: "paymentenum.PaymentStatus", valid: []stringEnum{payment_enums.PaymentStatusUnknown, payment_enums.PaymentStatusUnpaid, payment_enums.PaymentStatusPending, payment_enums.PaymentStatusPaid, payment_enums.PaymentStatusPartiallyPaid, payment_enums.PaymentStatusRefunded, payment_enums.PaymentStatusPartialRefunded}, invalid: payment_enums.PaymentStatus("__invalid__")},
		{name: "paymentenum.RecoveryDecision", valid: []stringEnum{payment_enums.RecoveryDecisionPending, payment_enums.RecoveryDecisionApproved, payment_enums.RecoveryDecisionDeclined}, invalid: payment_enums.RecoveryDecision("__invalid__")},
		{name: "settlementenum.SettlementType", valid: []stringEnum{settlement_enums.SettlementTypeSettlement, settlement_enums.SettlementTypeEnquiry}, invalid: settlement_enums.SettlementType("__invalid__")},
		{name: "terminalenum.TerminalConnectionMode", valid: []stringEnum{terminal_enums.TerminalConnectionModeCloudSync, terminal_enums.TerminalConnectionModeCloudAsync, terminal_enums.TerminalConnectionModeLocal}, invalid: terminal_enums.TerminalConnectionMode("__invalid__")},
		{name: "terminalenum.TerminalProvider", valid: []stringEnum{terminal_enums.TerminalProviderMx51, terminal_enums.TerminalProviderStripe}, invalid: terminal_enums.TerminalProvider("__invalid__")},
		{name: "terminalenum.TerminalRefundType", valid: []stringEnum{terminal_enums.TerminalRefundTypeReferenced, terminal_enums.TerminalRefundTypeUnreferenced}, invalid: terminal_enums.TerminalRefundType("__invalid__")},
		{name: "terminalenum.TerminalStatus", valid: []stringEnum{terminal_enums.TerminalStatusRegistered, terminal_enums.TerminalStatusActive, terminal_enums.TerminalStatusDeregistered, terminal_enums.TerminalStatusExpired, terminal_enums.TerminalStatusError}, invalid: terminal_enums.TerminalStatus("__invalid__")},
		{name: "terminalenum.TerminalTxFinancialStatus", valid: []stringEnum{terminal_enums.TerminalTxFinancialStatusApproved, terminal_enums.TerminalTxFinancialStatusDeclined, terminal_enums.TerminalTxFinancialStatusCancelled, terminal_enums.TerminalTxFinancialStatusUnknown}, invalid: terminal_enums.TerminalTxFinancialStatus("__invalid__")},
		{name: "terminalenum.TerminalTxStatus", valid: []stringEnum{terminal_enums.TerminalTxStatusUnknown, terminal_enums.TerminalTxStatusPending, terminal_enums.TerminalTxStatusAwaitingAction, terminal_enums.TerminalTxStatusFinalised, terminal_enums.TerminalTxStatusOverridePending, terminal_enums.TerminalTxStatusOverrideResolved}, invalid: terminal_enums.TerminalTxStatus("__invalid__")},
		{name: "terminalenum.TerminalTxType", valid: []stringEnum{terminal_enums.TerminalTxTypePurchase, terminal_enums.TerminalTxTypeRefund, terminal_enums.TerminalTxTypeReversal, terminal_enums.TerminalTxTypeCashout, terminal_enums.TerminalTxTypePurchaseWithCashout, terminal_enums.TerminalTxTypeMOTO, terminal_enums.TerminalTxTypeSettlement, terminal_enums.TerminalTxTypeSettlementEnquiry}, invalid: terminal_enums.TerminalTxType("__invalid__")},
	})
}
