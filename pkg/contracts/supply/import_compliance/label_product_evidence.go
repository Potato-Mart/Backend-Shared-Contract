package import_compliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"
)

// LabelProductEvidence freezes only the product facts used to author and
// substantiate a compliance label. It is compliance-owned evidence, not a
// second catalogue product model.
type LabelProductEvidence struct {
	SKUCode        string                       `json:"sku_code"`
	Barcode        string                       `json:"barcode,omitempty"`
	EnglishName    string                       `json:"english_name,omitempty"`
	ChineseName    string                       `json:"chinese_name,omitempty"`
	AlternateNames []localization.LocalizedName `json:"alternate_names,omitempty"`
	Brand          string                       `json:"brand,omitempty"`
	// MarketCode and TaxCategoryCode replace the retired product-level Taxed
	// flag: taxability is a market listing fact in v27, so compliance
	// freezes the market and its Pricing-owned tax category instead.
	MarketCode           string    `json:"market_code,omitempty"`
	TaxCategoryCode      string    `json:"tax_category_code,omitempty"`
	CapturedAt           time.Time `json:"captured_at"`
	SourceChecksumSHA256 string    `json:"source_checksum_sha256,omitempty"`
}
