// Package import_compliance defines reusable import-compliance records and
// snapshots. Transport DTOs, validation, authorization, calculations, and
// lifecycle transitions belong to the owning backend.
package import_compliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/temporal"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/import_compliance/import_compliance_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/purchase/purchase_enums"
)

// RevisionMetadata identifies one immutable-or-editable revision and records
// the review actions already applied to it. The owning backend defines which
// transitions are permitted.
type RevisionMetadata struct {
	Number             int64                               `json:"number"`
	BaseRevisionNumber *int64                              `json:"base_revision_number,omitempty"`
	State              import_compliance_enums.ReviewState `json:"state"`
	Submitted          *audit.LifecycleAction              `json:"submitted,omitempty"`
	Approved           *audit.LifecycleAction              `json:"approved,omitempty"`
	Rejected           *audit.LifecycleAction              `json:"rejected,omitempty"`
	Archived           *audit.LifecycleAction              `json:"archived,omitempty"`
}

// EvidenceReference points at source material without embedding regulated or
// potentially sensitive file contents in the shared model.
type EvidenceReference struct {
	ID                   string                               `json:"id"`
	Kind                 import_compliance_enums.EvidenceKind `json:"kind"`
	MediaID              string                               `json:"media_id,omitempty"`
	SourceURL            string                               `json:"source_url,omitempty"`
	SourceTitle          string                               `json:"source_title,omitempty"`
	SourceVersion        string                               `json:"source_version,omitempty"`
	SourceChecksumSHA256 string                               `json:"source_checksum_sha256,omitempty"`
	CapturedAt           *time.Time                           `json:"captured_at,omitempty"`
	Note                 string                               `json:"note,omitempty"`
}

// CatalogueReference pins a classification to the exact official catalogue
// version and entry that a reviewer saw.
type CatalogueReference struct {
	ID                   string                               `json:"id"`
	Jurisdiction         import_compliance_enums.Jurisdiction `json:"jurisdiction"`
	Version              string                               `json:"version"`
	EntryID              string                               `json:"entry_id,omitempty"`
	SourceURL            string                               `json:"source_url"`
	SourceChecksumSHA256 string                               `json:"source_checksum_sha256"`
	EffectiveFrom        temporal.Date                        `json:"effective_from,omitempty"`
	RetrievedAt          time.Time                            `json:"retrieved_at"`
}

// ArtifactReference identifies a deterministic generated artifact stored by a
// backend-managed media service.
type ArtifactReference struct {
	ID             string                               `json:"id"`
	Kind           import_compliance_enums.ArtifactKind `json:"kind"`
	MediaID        string                               `json:"media_id"`
	Filename       string                               `json:"filename"`
	MIMEType       string                               `json:"mime_type"`
	ChecksumSHA256 string                               `json:"checksum_sha256"`
	RevisionNumber int64                                `json:"revision_number"`
	GeneratedAt    time.Time                            `json:"generated_at"`
}

// LabelProductEvidence freezes only the product facts used to author and
// substantiate a compliance label. It is compliance-owned evidence, not a
// second catalogue product model.
type LabelProductEvidence struct {
	SKUID          string                       `json:"sku_id"`
	Barcode        string                       `json:"barcode,omitempty"`
	EnglishName    string                       `json:"english_name,omitempty"`
	ChineseName    string                       `json:"chinese_name,omitempty"`
	AlternateNames []localization.LocalizedName `json:"alternate_names,omitempty"`
	Brand          string                       `json:"brand,omitempty"`
	// MarketID and TaxCategoryID replace the retired product-level Taxed
	// flag: taxability is a market listing fact in v27, so compliance
	// freezes the market and its Pricing-owned tax category instead.
	MarketID             string    `json:"market_id,omitempty"`
	TaxCategoryID        string    `json:"tax_category_id,omitempty"`
	CapturedAt           time.Time `json:"captured_at"`
	SourceChecksumSHA256 string    `json:"source_checksum_sha256,omitempty"`
}

// PurchaseOrderSnapshot freezes the purchase-order header used by a
// declaration or tariff assessment. Each aggregate owns its corresponding
// line snapshots.
type PurchaseOrderSnapshot struct {
	ID                   string                             `json:"id"`
	OrderNumber          string                             `json:"order_number"`
	Status               purchase_enums.PurchaseOrderStatus `json:"status"`
	SupplierCode         string                             `json:"supplier_code,omitempty"`
	SupplierName         string                             `json:"supplier_name,omitempty"`
	ExpectedArrival      temporal.Date                      `json:"expected_arrival,omitempty"`
	CapturedAt           time.Time                          `json:"captured_at"`
	SourceChecksumSHA256 string                             `json:"source_checksum_sha256,omitempty"`
}
