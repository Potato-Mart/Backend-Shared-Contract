package payment_enums

// BankTransferConfirmationStatus records how much of a bank-transfer payment
// staff confirmed against the payer's transfer evidence when it was recorded.
//
// Both values complete the payment for the full amount due. ConfirmedPartial
// flags a transfer whose evidence did not show the full amount, so finance can
// follow it up; it never leaves a balance outstanding on the order.
type BankTransferConfirmationStatus string

const (
	BankTransferConfirmationStatusConfirmedFull    BankTransferConfirmationStatus = "confirmed_full"
	BankTransferConfirmationStatusConfirmedPartial BankTransferConfirmationStatus = "confirmed_partial"
)

func (s BankTransferConfirmationStatus) IsValid() bool {
	switch s {
	case BankTransferConfirmationStatusConfirmedFull, BankTransferConfirmationStatusConfirmedPartial:
		return true
	default:
		return false
	}
}

func (s BankTransferConfirmationStatus) String() string { return string(s) }
