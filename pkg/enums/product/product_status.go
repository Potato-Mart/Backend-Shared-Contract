package productenum

// ProductStatus is the admin-controlled lifecycle state of a product. It
// records merchandising intent only — it deliberately does NOT encode
// derived runtime state such as "new", "restocked", or "out of stock".
// Those are computed at read time from first_listed_at / restocked_at /
// current_stock (see product.Product.DisplayStatus); folding them into
// the stored status would require a background job and could drift.
type ProductStatus string

const (
	// ProductStatusDraft is a product being prepared that is not yet
	// publicly listed or sellable.
	ProductStatusDraft ProductStatus = "draft"
	// ProductStatusActive is a product that is publicly listed and on
	// sale. The first transition into this status stamps first_listed_at
	// and starts the 14-day NEW (新品) window.
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
