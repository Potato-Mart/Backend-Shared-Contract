package selection_enums

// ProductSelectionStatus is the admin-controlled lifecycle state of a product
// selection. Soft deletion is recorded separately by the audit fields.
type ProductSelectionStatus string

const (
	// ProductSelectionStatusDraft is a selection being prepared that is not
	// publicly visible.
	ProductSelectionStatusDraft ProductSelectionStatus = "draft"
	// ProductSelectionStatusPublished is a selection that is publicly
	// visible.
	ProductSelectionStatusPublished ProductSelectionStatus = "published"
	// ProductSelectionStatusArchived is a selection withdrawn from public
	// view but retained.
	ProductSelectionStatusArchived ProductSelectionStatus = "archived"
)

// IsValid reports whether s is a known ProductSelectionStatus.
func (s ProductSelectionStatus) IsValid() bool {
	switch s {
	case ProductSelectionStatusDraft, ProductSelectionStatusPublished, ProductSelectionStatusArchived:
		return true
	}
	return false
}

func (s ProductSelectionStatus) String() string { return string(s) }
