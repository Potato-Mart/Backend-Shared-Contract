package enums

// DeliveryRegion classifies where an order ships relative to the Melbourne
// warehouse. It is derived from the shipping address when the order is
// created (local when the address matches a shipping zone flagged is_local,
// interstate when the state is outside VIC, regional_vic otherwise).
type DeliveryRegion string

const (
	// DeliveryRegionLocalMelbourne is the Melbourne metro area.
	DeliveryRegionLocalMelbourne DeliveryRegion = "local_melbourne"
	// DeliveryRegionRegionalVIC is Victoria outside the metro area.
	DeliveryRegionRegionalVIC DeliveryRegion = "regional_vic"
	// DeliveryRegionInterstate is any state or territory other than VIC.
	DeliveryRegionInterstate DeliveryRegion = "interstate"
)

// IsValid reports whether r is a known DeliveryRegion.
func (r DeliveryRegion) IsValid() bool {
	switch r {
	case DeliveryRegionLocalMelbourne, DeliveryRegionRegionalVIC, DeliveryRegionInterstate:
		return true
	}
	return false
}

func (r DeliveryRegion) String() string { return string(r) }
