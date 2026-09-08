package quote_enums

// CustomPriceReason categorizes an audited custom price. It does not confer
// eligibility to sell stock or replace the required reason notes.
type CustomPriceReason string

const (
	CustomPriceReasonQuickSale  CustomPriceReason = "quick_sale"
	CustomPriceReasonSoonExpiry CustomPriceReason = "soon_expiry"
	CustomPriceReasonDamaged    CustomPriceReason = "damaged"
)

func (r CustomPriceReason) IsValid() bool {
	switch r {
	case CustomPriceReasonQuickSale, CustomPriceReasonSoonExpiry, CustomPriceReasonDamaged:
		return true
	}
	return false
}

func (r CustomPriceReason) String() string { return string(r) }
