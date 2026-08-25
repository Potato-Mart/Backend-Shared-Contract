package operations

// InventoryCauseRef identifies the contract record that caused an inventory
// change.
type InventoryCauseRef struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
