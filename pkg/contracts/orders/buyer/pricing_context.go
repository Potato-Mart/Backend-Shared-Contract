package buyer

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/market/market_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/product/product_enums"
)

// PricingContext is the shared commercial pricing context under which a line
// price was determined. It is descriptive only: the contract records the
// audience and visibility that applied, never the resolver logic that picked
// the price.
type PricingContext struct {
	Audience          market_enums.PriceAudience    `json:"audience,omitempty"`
	Visibility        product_enums.PriceVisibility `json:"visibility,omitempty"`
	GeographicContext geography.GeographicContext   `json:"geographic_context"`
}
