package warehouse_enums

// StockLocationRole marks a stock location with a special customer-facing
// role. The empty value is a normal location. The role does not change how
// stock at the location is counted, reserved or sold.
type StockLocationRole string

const (
	// StockLocationRoleDisplay is a customer-facing display location, for
	// example a frozen display A7FD-xx.
	StockLocationRoleDisplay StockLocationRole = "DISPLAY"
)

// IsValid reports whether r is a known StockLocationRole. The empty value
// (a normal location) is not a role and is not valid.
func (r StockLocationRole) IsValid() bool {
	switch r {
	case StockLocationRoleDisplay:
		return true
	default:
		return false
	}
}

func (r StockLocationRole) String() string { return string(r) }
