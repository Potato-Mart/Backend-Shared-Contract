package listing

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/listing/listing_enums"
)

// SaleRestriction is one market-specific restriction recorded against a
// listing. Channels narrows the restriction to specific order channels when it
// does not apply everywhere.
type SaleRestriction struct {
	Kind     listing_enums.SaleRestrictionKind `json:"kind"`
	Channels []commerce_enums.OrderType        `json:"channels,omitempty"`
	Value    int64                             `json:"value,omitempty"`
	Note     string                            `json:"note,omitempty"`
}
