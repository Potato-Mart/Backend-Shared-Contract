package selection

// ProductSelectionItem is one product membership of a ProductSelection.
// Position is the zero-based curated order within the selection.
type ProductSelectionItem struct {
	SelectionCode string `json:"selection_code"`
	SKUCode       string `json:"sku_code"`
	Position      int    `json:"position"`
}
