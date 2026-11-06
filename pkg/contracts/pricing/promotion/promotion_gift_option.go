package promotion

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
)

// PromotionGiftOption is one SKU a gift-type promotion relation can award.
// PackageOption is the authored package the gift is given in and is omitted
// when the promotion names none. MaxBaseUnits is the most base units of this
// SKU the relation awards. UnitPrice is the price charged per package for an
// add-on and is omitted for a free gift or when no price is established.
type PromotionGiftOption struct {
	SKUCode       string                      `json:"sku_code"`
	PackageOption *packaging.PackageOptionRef `json:"package_option,omitempty"`
	MaxBaseUnits  int64                       `json:"max_base_units"`
	UnitPrice     *money.Money                `json:"unit_price,omitempty"`
}
