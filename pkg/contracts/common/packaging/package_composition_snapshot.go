package packaging

// PackageCompositionSnapshot is the package-aware representation of a product
// quantity at a point in a sales or fulfilment workflow.
type PackageCompositionSnapshot struct {
	TotalBaseUnits int64                      `json:"total_base_units"`
	Components     []PackageComponentSnapshot `json:"components"`
}
