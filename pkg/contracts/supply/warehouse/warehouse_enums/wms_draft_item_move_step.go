package warehouse_enums

// WMSDraftItemMoveStep is the progress of the stock move a WMS draft item
// submits.
type WMSDraftItemMoveStep string

const (
	WMSDraftItemMoveStepPending WMSDraftItemMoveStep = "PENDING"
	WMSDraftItemMoveStepMoved   WMSDraftItemMoveStep = "MOVED"
	WMSDraftItemMoveStepDone    WMSDraftItemMoveStep = "DONE"
)

func (w WMSDraftItemMoveStep) IsValid() bool {
	switch w {
	case WMSDraftItemMoveStepPending, WMSDraftItemMoveStepMoved, WMSDraftItemMoveStepDone:
		return true
	}
	return false
}

func (w WMSDraftItemMoveStep) String() string { return string(w) }
