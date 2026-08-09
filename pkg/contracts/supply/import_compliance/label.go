package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

// LabelMaster is one product/package label revision. It remains independent of
// the product catalogue and retains only its compliance-owned source evidence.
type LabelMaster struct {
	ID                     string                 `json:"id"`
	Revision               RevisionMetadata       `json:"revision"`
	SourceProductEvidence  LabelProductEvidence   `json:"source_product_evidence"`
	ProductSKUCode         string                 `json:"product_sku_code"`
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
	PackagePhotoMediaID    string                 `json:"package_photo_media_id,omitempty"`
	PackagePhotoName       string                 `json:"package_photo_name,omitempty"`
	Layout                 LabelLayout            `json:"layout"`
	Evidence               []EvidenceReference    `json:"evidence,omitempty"`
	Artifacts              []ArtifactReference    `json:"artifacts,omitempty"`

	audit.AuditFields
}

type LabelImporter struct {
	Name    string             `json:"name,omitempty"`
	Address *geography.Address `json:"address,omitempty"`
	Phone   string             `json:"phone,omitempty"`
}

// LabelLayout stores the local preview controls. FontScaleBasisPoints uses
// 10,000 as a scale of 1.0.
type LabelLayout struct {
	Size                 import_compliance_enums.LabelSize        `json:"size"`
	Orientation          import_compliance_enums.LabelOrientation `json:"orientation"`
	FontScaleBasisPoints int64                                    `json:"font_scale_basis_points"`
	IncludeBarcode       bool                                     `json:"include_barcode"`
}

// NutritionPanel intentionally preserves author-entered display text and
// units. It is regulatory label copy, not a value used in shared arithmetic.
type NutritionPanel struct {
	Title                   string `json:"title,omitempty"`
	ServingsPerPack         string `json:"servings_per_pack,omitempty"`
	ServingSize             string `json:"serving_size,omitempty"`
	EnergyPerServe          string `json:"energy_per_serve,omitempty"`
	EnergyPer100Grams       string `json:"energy_per_100_grams,omitempty"`
	ProteinPerServe         string `json:"protein_per_serve,omitempty"`
	ProteinPer100Grams      string `json:"protein_per_100_grams,omitempty"`
	FatPerServe             string `json:"fat_per_serve,omitempty"`
	FatPer100Grams          string `json:"fat_per_100_grams,omitempty"`
	SaturatedFatPerServe    string `json:"saturated_fat_per_serve,omitempty"`
	SaturatedFatPer100Grams string `json:"saturated_fat_per_100_grams,omitempty"`
	CarbohydratePerServe    string `json:"carbohydrate_per_serve,omitempty"`
	CarbohydratePer100Grams string `json:"carbohydrate_per_100_grams,omitempty"`
	SugarsPerServe          string `json:"sugars_per_serve,omitempty"`
	SugarsPer100Grams       string `json:"sugars_per_100_grams,omitempty"`
	SodiumPerServe          string `json:"sodium_per_serve,omitempty"`
	SodiumPer100Grams       string `json:"sodium_per_100_grams,omitempty"`
}
