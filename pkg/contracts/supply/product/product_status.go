package product

// ProductStatus is the admin-controlled lifecycle state of a product. Live
// availability is represented separately by ProductStockSummary.
type ProductStatus string

const (
	// ProductStatusDraft is a product being prepared that is not yet
	// publicly listed or sellable.
	ProductStatusDraft ProductStatus = "draft"
	// ProductStatusActive is a product that is publicly listed and on sale.
	ProductStatusActive ProductStatus = "active"
	// ProductStatusArchived is a product delisted from sale but retained
	// for history/reporting (the soft-removed state; replaces the former
	// "dismiss").
	ProductStatusArchived ProductStatus = "archived"
	// ProductStatusDiscontinued is a product permanently retired and not
	// expected to return.
	ProductStatusDiscontinued ProductStatus = "discontinued"
)

// IsValid reports whether p is a known ProductStatus.
func (p ProductStatus) IsValid() bool {
	switch p {
	case ProductStatusDraft, ProductStatusActive, ProductStatusArchived, ProductStatusDiscontinued:
		return true
	}
	return false
}

func (p ProductStatus) String() string { return string(p) }
