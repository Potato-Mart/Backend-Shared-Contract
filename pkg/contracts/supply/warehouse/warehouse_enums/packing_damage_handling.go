package warehouse_enums

// PackingDamageHandling captures what the operator did for a damaged packed unit.
type PackingDamageHandling string

const (
	PackingDamageReplaceFromStock PackingDamageHandling = "replace_from_stock"
	PackingDamageShortShipRefund  PackingDamageHandling = "short_ship_refund"
)

// IsValid reports whether h is a known customer handling mode.
func (h PackingDamageHandling) IsValid() bool {
	switch h {
	case PackingDamageReplaceFromStock, PackingDamageShortShipRefund:
		return true
	}
	return false
}

func (h PackingDamageHandling) String() string { return string(h) }
