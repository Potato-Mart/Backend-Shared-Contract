package supply_enums

// InventorySaleCostStatus reports whether a physical sale's cost is verified.
// Unresolved includes absent, provisional or otherwise unverified cost.
type InventorySaleCostStatus string

const (
	InventorySaleCostStatusValued     InventorySaleCostStatus = "valued"
	InventorySaleCostStatusUnresolved InventorySaleCostStatus = "unresolved"
)

func (status InventorySaleCostStatus) IsValid() bool {
	return status == InventorySaleCostStatusValued || status == InventorySaleCostStatusUnresolved
}

func (status InventorySaleCostStatus) String() string { return string(status) }
