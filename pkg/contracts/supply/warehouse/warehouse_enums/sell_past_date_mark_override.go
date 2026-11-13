package warehouse_enums

// SellPastDateMarkOverride is one product's own setting. Inherit follows the
// market default; the other values replace it for that product in every
// market.
type SellPastDateMarkOverride string

const (
	SellPastDateMarkOverrideInherit             SellPastDateMarkOverride = "inherit"
	SellPastDateMarkOverrideNone                SellPastDateMarkOverride = "none"
	SellPastDateMarkOverrideBestBefore          SellPastDateMarkOverride = "best_before"
	SellPastDateMarkOverrideBestBeforeAndExpiry SellPastDateMarkOverride = "best_before_and_expiry"
)

func (o SellPastDateMarkOverride) IsValid() bool {
	switch o {
	case SellPastDateMarkOverrideInherit, SellPastDateMarkOverrideNone, SellPastDateMarkOverrideBestBefore, SellPastDateMarkOverrideBestBeforeAndExpiry:
		return true
	}
	return false
}

func (o SellPastDateMarkOverride) String() string { return string(o) }
