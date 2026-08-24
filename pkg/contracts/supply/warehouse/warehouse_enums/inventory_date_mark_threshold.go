package warehouse_enums

// InventoryDateMarkThreshold identifies a date-mark threshold crossing.
type InventoryDateMarkThreshold string

const (
	InventoryDateMarkThresholdApproaching InventoryDateMarkThreshold = "APPROACHING"
	InventoryDateMarkThresholdReached     InventoryDateMarkThreshold = "REACHED"
	InventoryDateMarkThresholdPassed      InventoryDateMarkThreshold = "PASSED"
)

func (t InventoryDateMarkThreshold) IsValid() bool {
	switch t {
	case InventoryDateMarkThresholdApproaching, InventoryDateMarkThresholdReached,
		InventoryDateMarkThresholdPassed:
		return true
	}
	return false
}

func (t InventoryDateMarkThreshold) String() string { return string(t) }
