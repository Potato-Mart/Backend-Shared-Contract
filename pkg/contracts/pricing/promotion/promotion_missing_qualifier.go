package promotion

// PromotionMissingQualifier names the SKUs that can close one quantity
// shortfall and how many more units of them the cart needs.
type PromotionMissingQualifier struct {
	SKUCodes     []string `json:"sku_codes"`
	UnitsMissing int64    `json:"units_missing"`
}
