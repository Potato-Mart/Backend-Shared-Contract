package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

// ManufacturerDeclaration is a revisioned declaration backed by an immutable
// purchase-order snapshot. Multiple declarations may reference the same order.
type ManufacturerDeclaration struct {
	ID                     string                `json:"id"`
	Revision               RevisionMetadata      `json:"revision"`
	PurchaseOrder          PurchaseOrderSnapshot `json:"purchase_order"`
	SettingsRevisionNumber *int64                `json:"settings_revision_number,omitempty"`
	DeclarationReference   string                `json:"declaration_reference"`
	Shipment               DeclarationShipment   `json:"shipment"`
	Manufacturer           ManufacturerDetails   `json:"manufacturer"`
	Signatory              DeclarationSignatory  `json:"signatory"`
	Lines                  []DeclarationLine     `json:"lines"`
	Evidence               []EvidenceReference   `json:"evidence,omitempty"`
	Artifacts              []ArtifactReference   `json:"artifacts,omitempty"`

	audit.AuditFields
}

type DeclarationShipment struct {
	DeclarationDate       temporal.Date                      `json:"declaration_date,omitempty"`
	ImportMode            import_compliance_enums.ImportMode `json:"import_mode,omitempty"`
	ShippingMethod        string                             `json:"shipping_method,omitempty"`
	TransportReference    string                             `json:"transport_reference,omitempty"`
	ConsignmentIdentifier string                             `json:"consignment_identifier,omitempty"`
	PortOfLoading         string                             `json:"port_of_loading,omitempty"`
	DestinationPort       string                             `json:"destination_port,omitempty"`
}

type ManufacturerDetails struct {
	Name    string             `json:"name"`
	Address *geography.Address `json:"address,omitempty"`
	Phone   string             `json:"phone,omitempty"`
}

// DeclarationSignatory references a managed image rather than embedding a
// base64 payload in the declaration record.
type DeclarationSignatory struct {
	Name             string `json:"name"`
	Title            string `json:"title,omitempty"`
	SignatureMediaID string `json:"signature_media_id,omitempty"`
}

type DeclarationLine struct {
	ID                        string        `json:"id"`
	SourceLineID              string        `json:"source_line_id,omitempty"`
	SourceLabelID             string        `json:"source_label_id,omitempty"`
	SourceLabelRevisionNumber *int64        `json:"source_label_revision_number,omitempty"`
	ProductReference          string        `json:"product_reference,omitempty"`
	EnglishName               string        `json:"english_name"`
	ChineseName               string        `json:"chinese_name,omitempty"`
	OrderedQuantity           int64         `json:"ordered_quantity"`
	CartonCount               int64         `json:"carton_count"`
	SingleNetWeightGrams      int64         `json:"single_net_weight_grams"`
	TotalNetWeightGrams       int64         `json:"total_net_weight_grams"`
	TotalGrossWeightGrams     int64         `json:"total_gross_weight_grams"`
	ExpiryDate                temporal.Date `json:"expiry_date,omitempty"`
	Ingredients               string        `json:"ingredients,omitempty"`
	ManufacturingProcess      string        `json:"manufacturing_process,omitempty"`
	Note                      string        `json:"note,omitempty"`
}
