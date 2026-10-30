package possettlement

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// TenderInstrument identifies an original owner debit and its money valuation.
// A gift card binds both the Payments record and Pricing wallet transaction;
// points bind the original Pricing reservation and ledger entry.
type TenderInstrument struct {
	Kind                string      `json:"kind"`
	PaymentID           string      `json:"payment_id,omitempty"`
	ReservationID       string      `json:"reservation_id,omitempty"`
	GiftCardCode        string      `json:"gift_card_code,omitempty"`
	WalletTransactionID string      `json:"wallet_transaction_id,omitempty"`
	PointsLedgerEntryID string      `json:"points_ledger_entry_id,omitempty"`
	Amount              money.Money `json:"amount"`
}
