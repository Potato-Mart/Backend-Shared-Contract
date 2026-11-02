package payment_enums

import "testing"

func TestBankTransferConfirmationStatus(t *testing.T) {
	for _, status := range []BankTransferConfirmationStatus{BankTransferConfirmationStatusConfirmedFull, BankTransferConfirmationStatusConfirmedPartial} {
		if !status.IsValid() || status.String() != string(status) {
			t.Fatal(status)
		}
	}
	for _, status := range []BankTransferConfirmationStatus{"", "pending", "unconfirmed", "CONFIRMED_FULL", "partial"} {
		if status.IsValid() {
			t.Fatalf("unknown bank-transfer confirmation status accepted: %q", status)
		}
	}
}
