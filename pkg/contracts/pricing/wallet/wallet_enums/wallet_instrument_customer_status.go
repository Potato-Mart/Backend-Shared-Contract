package wallet_enums

// WalletInstrumentCustomerStatus is Pricing's customer-facing usage projection,
// independent of the instrument issuance/redemption lifecycle. It is not checkout
// authorization; Pricing still evaluates current eligibility at use.
type WalletInstrumentCustomerStatus string

const (
	WalletInstrumentCustomerStatusAvailable WalletInstrumentCustomerStatus = "available"
	WalletInstrumentCustomerStatusUsed      WalletInstrumentCustomerStatus = "used"
	WalletInstrumentCustomerStatusExpired   WalletInstrumentCustomerStatus = "expired"
)

func (s WalletInstrumentCustomerStatus) IsValid() bool {
	switch s {
	case WalletInstrumentCustomerStatusAvailable, WalletInstrumentCustomerStatusUsed, WalletInstrumentCustomerStatusExpired:
		return true
	}
	return false
}
func (s WalletInstrumentCustomerStatus) String() string { return string(s) }
