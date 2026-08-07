// Package importcompliance defines reusable import-compliance records and
// snapshots. Transport DTOs, validation, authorization, calculations, and
// lifecycle transitions belong to the owning backend.
package importcompliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/temporal"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/supply/importcompliance/importcompliance_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/supply/purchase/purchase_enums"
)

// RevisionMetadata identifies one immutable-or-editable revision and records
// the review actions already applied to it. The owning backend defines which
// transitions are permitted.
type RevisionMetadata struct {
	Number             int64                              `json:"number"`
	BaseRevisionNumber *int64                             `json:"base_revision_number,omitempty"`
	State              importcompliance_enums.ReviewState `json:"state"`
	Submitted          *audit.LifecycleAction             `json:"submitted,omitempty"`
	Approved           *audit.LifecycleAction             `json:"approved,omitempty"`
	Rejected           *audit.LifecycleAction             `json:"rejected,omitempty"`
	Archived           *audit.LifecycleAction             `json:"archived,omitempty"`
}

// EvidenceReference points at source material without embedding regulated or
// potentially sensitive file contents in the shared model.
type EvidenceReference struct {
	ID                   string                              `json:"id"`
	Kind                 importcompliance_enums.EvidenceKind `json:"kind"`
	MediaID              string                              `json:"media_id,omitempty"`
	SourceURL            string                              `json:"source_url,omitempty"`
	SourceTitle          string                              `json:"source_title,omitempty"`
	SourceVersion        string                              `json:"source_version,omitempty"`
	SourceChecksumSHA256 string                              `json:"source_checksum_sha256,omitempty"`
	CapturedAt           *time.Time                          `json:"captured_at,omitempty"`
	Note                 string                              `json:"note,omitempty"`
}

// CatalogueReference pins a classification to the exact official catalogue
// version and entry that a reviewer saw.
type CatalogueReference struct {
	ID                   string                              `json:"id"`
	Jurisdiction         importcompliance_enums.Jurisdiction `json:"jurisdiction"`
	Version              string                              `json:"version"`
	EntryID              string                              `json:"entry_id,omitempty"`
	SourceURL            string                              `json:"source_url"`
	SourceChecksumSHA256 string                              `json:"source_checksum_sha256"`
	EffectiveFrom        temporal.Date                       `json:"effective_from,omitempty"`
	RetrievedAt          time.Time                           `json:"retrieved_at"`
}

// ArtifactReference identifies a deterministic generated artifact stored by a
// backend-managed media service.
type ArtifactReference struct {
	ID             string                              `json:"id"`
	Kind           importcompliance_enums.ArtifactKind `json:"kind"`
	MediaID        string                              `json:"media_id"`
	Filename       string                              `json:"filename"`
	MIMEType       string                              `json:"mime_type"`
	ChecksumSHA256 string                              `json:"checksum_sha256"`
	RevisionNumber int64                               `json:"revision_number"`
	GeneratedAt    time.Time                           `json:"generated_at"`
}

// ProductSnapshot captures the product identity and display fields used by an
// import-compliance record at the time it was authored.
type ProductSnapshot struct {
	ID                   string                       `json:"id,omitempty"`
	SKUCode              string                       `json:"sku_code"`
	SKU                  string                       `json:"sku,omitempty"`
	Barcode              string                       `json:"barcode,omitempty"`
	EnglishName          string                       `json:"english_name,omitempty"`
	ChineseName          string                       `json:"chinese_name,omitempty"`
	AlternateNames       []localization.LocalizedName `json:"alternate_names,omitempty"`
	Brand                string                       `json:"brand,omitempty"`
	Taxed                *bool                        `json:"taxed,omitempty"`
	CapturedAt           time.Time                    `json:"captured_at"`
	SourceChecksumSHA256 string                       `json:"source_checksum_sha256,omitempty"`
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
