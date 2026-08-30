package compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/measurement"
)

// LabelMaster is one product/package label revision. It remains independent of
// the product catalogue and retains only its compliance-owned source evidence.
type LabelMaster struct {
	ID       string           `json:"id"`
	Revision RevisionMetadata `json:"revision"`
	// MarketCode and CountryCode are the denormalized market and country the
	// record belongs to, carried so a geographically scoped staff query is
	// a plain indexed match.
	MarketCode             string                 `json:"market_code,omitempty"`
	CountryCode            geography.CountryCode  `json:"country_code,omitempty"`
	SourceProductEvidence  LabelProductEvidence   `json:"source_product_evidence"`
	SKUCode                string                 `json:"sku_code"`
	VariantCode            string                 `json:"variant_code"`
	Brand                  string                 `json:"brand,omitempty"`
	EnglishName            string                 `json:"english_name"`
	ChineseName            string                 `json:"chinese_name,omitempty"`
	Barcode                string                 `json:"barcode,omitempty"`
	NetWeightGrams         int64                  `json:"net_weight_grams"`
	PackageDimensions      measurement.Dimensions `json:"package_dimensions"`
	Ingredients            string                 `json:"ingredients"`
	Allergens              string                 `json:"allergens,omitempty"`
	ManufacturingProcess   string                 `json:"manufacturing_process,omitempty"`
	BestBefore             string                 `json:"best_before,omitempty"`
	ShelfLife              string                 `json:"shelf_life,omitempty"`
	Importer               LabelImporter          `json:"importer"`
	CountryOfOrigin        string                 `json:"country_of_origin,omitempty"`
	SecondNutritionEnabled bool                   `json:"second_nutrition_enabled"`
	NutritionPanels        []NutritionPanel       `json:"nutrition_panels"`
	PackagePhotoMediaCode  string                 `json:"package_photo_media_code,omitempty"`
	PackagePhotoName       string                 `json:"package_photo_name,omitempty"`
	Layout                 LabelLayout            `json:"layout"`
	Evidence               []EvidenceReference    `json:"evidence,omitempty"`
	Artifacts              []ArtifactReference    `json:"artifacts,omitempty"`

	audit.AuditFields
}
