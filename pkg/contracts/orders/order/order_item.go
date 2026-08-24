package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/packaging"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/product"
)

// OrderItem is the frozen product, package, pricing, and fulfilment state for
// one purchased line.
type OrderItem struct {
	ID string `json:"id"`
	// SKUCode is the frozen SKU code captured when the line was priced.
	SKUCode              string                                   `json:"sku_code"`
	ProductName          string                                   `json:"product_name"`
	ProductImage         *security.ObjectMedia                    `json:"product_image,omitempty"`
	ProductPackageOption product.ProductPackageOption             `json:"product_package_option"`
	CapturedAt           time.Time                                `json:"captured_at"`
	VariantTitle         string                                   `json:"variant_title,omitempty"`
	Components           []PricedPackageComponent                 `json:"components"`
	TotalBaseUnits       int64                                    `json:"total_base_units"`
	Pricing              *PricingContext                          `json:"pricing,omitempty"`
	SubstitutionPolicy   LooseSubstitutionPolicySnapshot          `json:"substitution_policy"`
	RequestedComposition packaging.PackageCompositionSnapshot     `json:"requested_composition"`
	AllocatedComposition packaging.PackageCompositionSnapshot     `json:"allocated_composition"`
	PickedComposition    packaging.PackageCompositionSnapshot     `json:"picked_composition"`
	PackedComposition    packaging.PackageCompositionSnapshot     `json:"packed_composition"`
	ReturnedComposition  packaging.PackageCompositionSnapshot     `json:"returned_composition"`
	RefundedComposition  packaging.PackageCompositionSnapshot     `json:"refunded_composition"`
	Substitutions        []operations.PackageSubstitutionSnapshot `json:"substitutions,omitempty"`
	DiscountAmount       money.Money                              `json:"discount_amount"`
	Total                money.Money                              `json:"total"`
	Preorder             *PreorderItemState                       `json:"preorder,omitempty"`
}
