package warehouse_enums

// InventoryDisposition describes how inventory may be used or sold.
type InventoryDisposition string

const (
	InventoryDispositionStandardSellable InventoryDisposition = "STANDARD_SELLABLE"
	InventoryDispositionReducedSellable  InventoryDisposition = "REDUCED_SELLABLE"
	InventoryDispositionQualityHold      InventoryDisposition = "QUALITY_HOLD"
	InventoryDispositionQuarantined      InventoryDisposition = "QUARANTINED"
	InventoryDispositionBlocked          InventoryDisposition = "BLOCKED"
	InventoryDispositionRecalled         InventoryDisposition = "RECALLED"
	InventoryDispositionReturnToVendor   InventoryDisposition = "RETURN_TO_VENDOR"
	InventoryDispositionScrap            InventoryDisposition = "SCRAP"
)

func (d InventoryDisposition) IsValid() bool {
	switch d {
	case InventoryDispositionStandardSellable, InventoryDispositionReducedSellable,
		InventoryDispositionQualityHold, InventoryDispositionQuarantined,
		InventoryDispositionBlocked, InventoryDispositionRecalled,
		InventoryDispositionReturnToVendor, InventoryDispositionScrap:
		return true
	}
	return false
}

func (d InventoryDisposition) String() string { return string(d) }
