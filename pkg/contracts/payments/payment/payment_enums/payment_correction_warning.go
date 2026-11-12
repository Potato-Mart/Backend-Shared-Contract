package payment_enums

// PaymentCorrectionWarning names a side effect a staff payment correction did
// not perform. A correction changes only the Payments record: the gateway
// charge, the terminal transaction and the wallet ledger keep their original
// amounts, and staff must reconcile them separately.
type PaymentCorrectionWarning string

const (
	PaymentCorrectionWarningGatewayAmountUnchanged  PaymentCorrectionWarning = "gateway_amount_unchanged"
	PaymentCorrectionWarningTerminalAmountUnchanged PaymentCorrectionWarning = "terminal_amount_unchanged"
	PaymentCorrectionWarningWalletLedgerUnchanged   PaymentCorrectionWarning = "wallet_ledger_unchanged"
)

// IsValid reports whether w is a known PaymentCorrectionWarning.
func (w PaymentCorrectionWarning) IsValid() bool {
	switch w {
	case PaymentCorrectionWarningGatewayAmountUnchanged, PaymentCorrectionWarningTerminalAmountUnchanged,
		PaymentCorrectionWarningWalletLedgerUnchanged:
		return true
	}
	return false
}

func (w PaymentCorrectionWarning) String() string { return string(w) }
