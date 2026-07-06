package warehouseenum

type PackingDiscrepancyKind string

const (
	PackingDiscrepancyKindShortage   PackingDiscrepancyKind = "shortage"
	PackingDiscrepancyKindOverweight PackingDiscrepancyKind = "overweight"
	PackingDiscrepancyKindDamaged    PackingDiscrepancyKind = "damaged"
	PackingDiscrepancyKindPending    PackingDiscrepancyKind = "pending"
)

// IsValid reports whether s is a known PackingDiscrepancyKind.
func (s PackingDiscrepancyKind) IsValid() bool {
	switch s {
	case PackingDiscrepancyKindShortage, PackingDiscrepancyKindOverweight, PackingDiscrepancyKindDamaged, PackingDiscrepancyKindPending:
		return true
	}
	return false
}

func (s PackingDiscrepancyKind) String() string { return string(s) }

// PackingSessionStatus is the workflow state for an order packing session.
type PackingSessionStatus string

const (
	PackingSessionStatusPending     PackingSessionStatus = "pending"
	PackingSessionStatusPacking     PackingSessionStatus = "packing"
	PackingSessionStatusPacked      PackingSessionStatus = "packed"
	PackingSessionStatusSyncPending PackingSessionStatus = "sync_pending"
	PackingSessionStatusResolved    PackingSessionStatus = "resolved"
)

// IsValid reports whether s is a known packing session state.
func (s PackingSessionStatus) IsValid() bool {
	switch s {
	case PackingSessionStatusPending, PackingSessionStatusPacking, PackingSessionStatusPacked,
		PackingSessionStatusSyncPending, PackingSessionStatusResolved:
		return true
	}
	return false
}

func (s PackingSessionStatus) String() string { return string(s) }

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
