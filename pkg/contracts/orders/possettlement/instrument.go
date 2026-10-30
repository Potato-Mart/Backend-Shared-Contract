package possettlement

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// Instrument references the original authoritative debit, never a balance.
// Kind is external, cash, gift_card or points. Coupons are discounts, not tender.
// Optional original-debit locators are absent in historical draft representations.
type Instrument struct {
	Kind                 string      `json:"kind"`
	Authority            string      `json:"authority"`
	ID                   string      `json:"id"`
	Amount               money.Money `json:"amount"`
	BenefitReservationID string      `json:"benefit_reservation_id,omitempty"`
	WalletTransactionID  string      `json:"wallet_transaction_id,omitempty"`
	PointReservationID   string      `json:"point_reservation_id,omitempty"`
	GiftCardCode         string      `json:"gift_card_code,omitempty"`
	PointsLedgerEntryID  string      `json:"points_ledger_entry_id,omitempty"`
}
