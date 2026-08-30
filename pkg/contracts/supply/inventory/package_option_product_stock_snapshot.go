package inventory

// PackageOptionProductStockSnapshot qualifies product stock by package option.
type PackageOptionProductStockSnapshot struct {
	PackageOptionCode string                       `json:"package_option_code"`
	Quantities        ProductStockQuantitySnapshot `json:"quantities"`
}
