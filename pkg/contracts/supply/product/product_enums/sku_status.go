package product_enums

// SKUStatus is the admin-controlled lifecycle state of one sellable SKU.
// Whether that SKU may actually be sold in a given market is a separate
// MarketListing fact, and live availability is separate stock evidence.
type SKUStatus string

const (
	// SKUStatusDraft is a SKU being prepared that may not be listed.
	SKUStatusDraft SKUStatus = "draft"
	// SKUStatusActive is a SKU available to be listed and stocked.
	SKUStatusActive SKUStatus = "active"
	// SKUStatusArchived is a SKU withdrawn from new listings but retained
	// for history and reporting.
	SKUStatusArchived SKUStatus = "archived"
	// SKUStatusDiscontinued is a SKU permanently retired.
	SKUStatusDiscontinued SKUStatus = "discontinued"
)

// IsValid reports whether s is a known SKUStatus.
func (s SKUStatus) IsValid() bool {
	switch s {
	case SKUStatusDraft, SKUStatusActive, SKUStatusArchived, SKUStatusDiscontinued:
		return true
	}
	return false
}

func (s SKUStatus) String() string { return string(s) }
