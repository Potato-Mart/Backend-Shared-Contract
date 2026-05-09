package payments

// Action codes a terminal transaction or settlement may instruct the
// POS to invoke. They cover both auto-actions (executed without
// merchant input as soon as the response is received) and button
// actions (executed when the merchant taps a button in the rendered
// Action Form).
//
// The full Action Form payload (layout, properties, details) is held
// opaquely on TerminalTransaction.POSInstructions and Settlement.POSInstructions
// because providers reserve the right to change its inner shape at any
// time. The constants below are the stable, documented action code
// vocabulary - safe to switch on.
const (
	ActionPrintMerchantReceipt = "PRINT_MERCHANT_RECEIPT"
	ActionPrintCustomerReceipt = "PRINT_CUSTOMER_RECEIPT"
	ActionTransactionComplete  = "TRANSACTION_COMPLETE"
	ActionRetryTransaction     = "RETRY_TRANSACTION"
	ActionSettlementComplete   = "SETTLEMENT_COMPLETE"
	ActionRetrySettlement      = "RETRY_SETTLEMENT"
	ActionTestAction           = "TEST_ACTION"

	// Submit-URL actions - sent back to the provider when the merchant
	// taps a button whose submit_url is non-null (e.g. signature
	// approve / decline). The exact string is the one the provider
	// expects in the action query parameter.
	ActionApproveSignature = "APPROVE_SIGNATURE"
	ActionDeclineSignature = "DECLINE_SIGNATURE"
)

// SettlementReceiptCode is one of the three-letter category codes used
// on settlement and settlement-enquiry receipts.
type SettlementReceiptCode string

const (
	SettlementReceiptCodePurchases  SettlementReceiptCode = "PUR"
	SettlementReceiptCodeTips       SettlementReceiptCode = "TIP"
	SettlementReceiptCodeSurcharges SettlementReceiptCode = "SUR"
	SettlementReceiptCodeRefunds    SettlementReceiptCode = "REF"
	SettlementReceiptCodeCashouts   SettlementReceiptCode = "CAS"
	SettlementReceiptCodeTotal      SettlementReceiptCode = "TOT"
)
