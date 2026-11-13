package warehouse_enums

// PackingShortageSettlementReason is why the packing station refunded units
// of a recorded packing shortage. NoStock means the goods are not there: the
// stock hold of the line is released and stock the books still show is
// written off. Overweight means the goods are there but were not shipped: the
// hold is released and the units are sellable again.
type PackingShortageSettlementReason string

const (
	PackingShortageSettlementReasonNoStock    PackingShortageSettlementReason = "no_stock"
	PackingShortageSettlementReasonOverweight PackingShortageSettlementReason = "overweight"
)

// IsValid reports whether r is a known PackingShortageSettlementReason.
func (r PackingShortageSettlementReason) IsValid() bool {
	switch r {
	case PackingShortageSettlementReasonNoStock, PackingShortageSettlementReasonOverweight:
		return true
	}
	return false
}

func (r PackingShortageSettlementReason) String() string { return string(r) }
