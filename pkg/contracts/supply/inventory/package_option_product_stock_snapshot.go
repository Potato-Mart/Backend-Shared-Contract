package inventory

import "github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/common/packaging"

// PackageOptionProductStockSnapshot qualifies product stock by package option.
type PackageOptionProductStockSnapshot struct {
	PackageOption packaging.PackageOptionRef   `json:"package_option"`
	Quantities    ProductStockQuantitySnapshot `json:"quantities"`
}
