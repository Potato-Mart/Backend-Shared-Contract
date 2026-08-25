package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"
)

// ProductOrigin is the customer-facing origin display block for a product.
type ProductOrigin struct {
	CountryCode geography.CountryCode        `json:"country_code,omitempty"`
	Label       []localization.LocalizedText `json:"label,omitempty"`
	Statement   []localization.LocalizedText `json:"statement,omitempty"`
}
