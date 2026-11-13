package warehouse_enums

// SellPastDateMarkPolicy says which lots may still be sold once they are at
// or past their date mark. It is the market-level default and the resolved,
// effective value for one product in one market.
type SellPastDateMarkPolicy string

const (
	// SellPastDateMarkNone sells no stock at or past any date mark.
	SellPastDateMarkNone SellPastDateMarkPolicy = "none"
	// SellPastDateMarkBestBefore sells stock past a BEST_BEFORE mark only.
	SellPastDateMarkBestBefore SellPastDateMarkPolicy = "best_before"
	// SellPastDateMarkBestBeforeAndExpiry sells stock past a BEST_BEFORE or
	// an EXPIRY mark.
	SellPastDateMarkBestBeforeAndExpiry SellPastDateMarkPolicy = "best_before_and_expiry"
)

func (p SellPastDateMarkPolicy) IsValid() bool {
	switch p {
	case SellPastDateMarkNone, SellPastDateMarkBestBefore, SellPastDateMarkBestBeforeAndExpiry:
		return true
	}
	return false
}

func (p SellPastDateMarkPolicy) String() string { return string(p) }
