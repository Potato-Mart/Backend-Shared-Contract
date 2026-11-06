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
	// ProductSelectionKindShowcase is a selection presented as one of the
	// storefront Home theme boxes, ordered by display order.
	ProductSelectionKindShowcase ProductSelectionKind = "showcase"
)

// IsValid reports whether k is a known ProductSelectionKind.
func (k ProductSelectionKind) IsValid() bool {
	switch k {
	case ProductSelectionKindBadge, ProductSelectionKindSection, ProductSelectionKindShowcase:
		return true
	}
	return false
}

func (k ProductSelectionKind) String() string { return string(k) }
