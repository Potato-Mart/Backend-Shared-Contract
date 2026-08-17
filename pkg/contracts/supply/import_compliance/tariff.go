package import_compliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/temporal"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

// RateValue preserves official source text because tariff schedules may use
// compound or non-percentage rates. BasisPoints is populated only when the
// source can be faithfully represented as a percentage.
type RateValue struct {
	Raw         string `json:"raw"`
	BasisPoints *int64 `json:"basis_points,omitempty"`
}

type TariffClassification struct {
	Jurisdiction import_compliance_enums.Jurisdiction `json:"jurisdiction"`
	Code         string                               `json:"code"`
	DutyRate     RateValue                            `json:"duty_rate"`
	GSTRate      *RateValue                           `json:"gst_rate,omitempty"`
	Catalogue    CatalogueReference                   `json:"catalogue"`
	Evidence     []EvidenceReference                  `json:"evidence,omitempty"`
}

// TariffLineSnapshot preserves every purchase-order/product field displayed by
// the tariff worksheet before classification begins.
type TariffLineSnapshot struct {
	ID                  string `json:"id"`
	PurchaseOrderLineID string `json:"purchase_order_line_id,omitempty"`
	SKUID               string `json:"sku_id,omitempty"`
	Barcode             string `json:"barcode,omitempty"`
	ProductName         string `json:"product_name,omitempty"`
	AlternateNames      string `json:"alternate_names,omitempty"`
	Brand               string `json:"brand,omitempty"`
	OrderedQuantity     *int64 `json:"ordered_quantity,omitempty"`
	ProductTaxed        *bool  `json:"product_taxed,omitempty"`
}

type TariffAssessmentRow struct {
	ID                       string               `json:"id"`
	Product                  TariffLineSnapshot   `json:"product"`
	Taiwan                   TariffClassification `json:"taiwan"`
	Australia                TariffClassification `json:"australia"`
	Trademark                string               `json:"trademark,omitempty"`
	TrademarkEvidenceIDs     []string             `json:"trademark_evidence_ids,omitempty"`
	Source                   string               `json:"source"`
	Notes                    string               `json:"notes,omitempty"`
	PublishedProfileRevision *int64               `json:"published_profile_revision,omitempty"`
}

// TariffAssessment is the revisioned, purchase-order-specific worksheet. An
// approved assessment may publish separate reusable TariffProfile revisions.
type TariffAssessment struct {
	ID            string                `json:"id"`
	Revision      RevisionMetadata      `json:"revision"`
	PurchaseOrder PurchaseOrderSnapshot `json:"purchase_order"`
	Rows          []TariffAssessmentRow `json:"rows"`
	Evidence      []EvidenceReference   `json:"evidence,omitempty"`
	Artifacts     []ArtifactReference   `json:"artifacts,omitempty"`

	audit.AuditFields
}

// TariffProfile is an approved reusable classification for one SKU and
// jurisdiction. Assessment records retain their own frozen classification.
type TariffProfile struct {
	ID                   string                               `json:"id"`
	Revision             RevisionMetadata                     `json:"revision"`
	SKUID                string                               `json:"sku_id"`
	Jurisdiction         import_compliance_enums.Jurisdiction `json:"jurisdiction"`
	Classification       TariffClassification                 `json:"classification"`
	EffectiveFrom        temporal.Date                        `json:"effective_from,omitempty"`
	EffectiveTo          temporal.Date                        `json:"effective_to,omitempty"`
	TrademarkEvidenceIDs []string                             `json:"trademark_evidence_ids,omitempty"`
	Evidence             []EvidenceReference                  `json:"evidence,omitempty"`

	audit.AuditFields
}

// TrademarkEvidence is a cited search result or manually verified record. It
// does not itself make a legal-clearance claim.
type TrademarkEvidence struct {
	ID                 string                               `json:"id"`
	Revision           RevisionMetadata                     `json:"revision"`
	SKUID              string                               `json:"sku_id,omitempty"`
	Jurisdiction       import_compliance_enums.Jurisdiction `json:"jurisdiction"`
	Mark               string                               `json:"mark"`
	Status             string                               `json:"status,omitempty"`
	Owner              string                               `json:"owner,omitempty"`
	RegistrationNumber string                               `json:"registration_number,omitempty"`
	Classes            []string                             `json:"classes,omitempty"`
	Source             string                               `json:"source"`
	CheckedAt          *time.Time                           `json:"checked_at,omitempty"`
	Evidence           []EvidenceReference                  `json:"evidence,omitempty"`

	audit.AuditFields
}
