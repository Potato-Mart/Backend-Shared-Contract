package warehouse_enums

// InventoryDateMarkKind identifies the meaning of a lot's date mark.
type InventoryDateMarkKind string

const (
	InventoryDateMarkBestBefore InventoryDateMarkKind = "BEST_BEFORE"
	InventoryDateMarkExpiry     InventoryDateMarkKind = "EXPIRY"
)

func (k InventoryDateMarkKind) IsValid() bool {
	switch k {
	case InventoryDateMarkBestBefore, InventoryDateMarkExpiry:
		return true
	}
	return false
}

func (k InventoryDateMarkKind) String() string { return string(k) }
