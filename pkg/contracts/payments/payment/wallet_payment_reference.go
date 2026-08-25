package payment

// WalletPaymentReference links a completed gift-card order payment to the
// authoritative wallet ledger transaction returned by Pricing.
type WalletPaymentReference struct {
	GiftCardCode        string `json:"gift_card_code"`
	WalletTransactionID string `json:"wallet_transaction_id"`
}
