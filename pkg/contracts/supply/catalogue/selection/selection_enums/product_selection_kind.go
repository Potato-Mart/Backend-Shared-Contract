package selection_enums

// ProductSelectionKind is the presentation role of a product selection.
type ProductSelectionKind string

const (
	// ProductSelectionKindBadge is a selection presented as a badge tile
	// with its own icon.
	ProductSelectionKindBadge ProductSelectionKind = "badge"
	// ProductSelectionKindSection is a selection presented as a product
	// section rail.
	ProductSelectionKindSection ProductSelectionKind = "section"
)

// IsValid reports whether k is a known ProductSelectionKind.
func (k ProductSelectionKind) IsValid() bool {
	switch k {
	case ProductSelectionKindBadge, ProductSelectionKindSection:
		return true
	}
	return false
}

func (k ProductSelectionKind) String() string { return string(k) }
