package importcompliance

import "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"

// ImportSettings is one revision of the shared import-cost and declaration
// configuration. TWDPerAUDMicros stores the TWD-per-AUD exchange rate scaled by
// 1,000,000. Margin and tax values use basis points (100 = one percent).
type ImportSettings struct {
	ID                      string                      `json:"id"`
	Revision                RevisionMetadata            `json:"revision"`
	TWDPerAUDMicros         int64                       `json:"twd_per_aud_micros"`
	NextInvoiceNumber       int64                       `json:"next_invoice_number"`
	ExportMarginBasisPoints int64                       `json:"export_margin_basis_points"`
	DefaultSignatoryName    string                      `json:"default_signatory_name,omitempty"`
	TaiwanTaxBasisPoints    int64                       `json:"taiwan_tax_basis_points"`
	AirCargo                AirCargoSettings            `json:"air_cargo"`
	AmbientSea              AmbientSeaSettings          `json:"ambient_sea"`
	FrozenSea               FrozenSeaSettings           `json:"frozen_sea"`
	IngredientRules         []IngredientDeclarationRule `json:"ingredient_rules,omitempty"`

	common.AuditFields
}

type IngredientDeclarationRule struct {
	ID          string `json:"id"`
	Ingredients string `json:"ingredients"`
	Statement   string `json:"statement"`
}

// AirCargoSettings preserves the current import-settings page fields while
// using integer grams/cubic centimetres and fixed-precision Money values.
type AirCargoSettings struct {
	ReferenceWeightGrams                         int64        `json:"reference_weight_grams"`
	TaiwanInspectionCost                         common.Money `json:"taiwan_inspection_cost"`
	StorageThresholdGrams                        int64        `json:"storage_threshold_grams"`
	StorageUnderThresholdCost                    common.Money `json:"storage_under_threshold_cost"`
	StorageAtOrOverThresholdCost                 common.Money `json:"storage_at_or_over_threshold_cost"`
	VolumetricDivisorCubicCentimetresPerKilogram int64        `json:"volumetric_divisor_cubic_centimetres_per_kilogram"`
	TaiwanFreightCost                            common.Money `json:"taiwan_freight_cost"`
	TaiwanHandlingCost                           common.Money `json:"taiwan_handling_cost"`
	TaiwanTransportCost                          common.Money `json:"taiwan_transport_cost"`
	TaiwanBrokerageCost                          common.Money `json:"taiwan_brokerage_cost"`
	TaiwanTruckCost                              common.Money `json:"taiwan_truck_cost"`
	AustraliaTerminalCost                        common.Money `json:"australia_terminal_cost"`
	AustraliaTruckCost                           common.Money `json:"australia_truck_cost"`
	AustraliaDocumentCost                        common.Money `json:"australia_document_cost"`
	AustraliaBrokerageCost                       common.Money `json:"australia_brokerage_cost"`
	AustraliaQuarantineCost                      common.Money `json:"australia_quarantine_cost"`
}

type AmbientSeaSettings struct {
	ReferenceVolumeCubicCentimetres int64        `json:"reference_volume_cubic_centimetres"`
	TaiwanConsolidationCost         common.Money `json:"taiwan_consolidation_cost"`
	TaiwanHandlingCost              common.Money `json:"taiwan_handling_cost"`
	TaiwanBrokerageCost             common.Money `json:"taiwan_brokerage_cost"`
	AustraliaTerminalCost           common.Money `json:"australia_terminal_cost"`
	AustraliaTruckCost              common.Money `json:"australia_truck_cost"`
	AustraliaFreightCost            common.Money `json:"australia_freight_cost"`
	AustraliaDocumentCost           common.Money `json:"australia_document_cost"`
	AustraliaBrokerageCost          common.Money `json:"australia_brokerage_cost"`
	AustraliaQuarantineCost         common.Money `json:"australia_quarantine_cost"`
}

type FrozenSeaSettings struct {
	ReferenceVolumeCubicCentimetres int64        `json:"reference_volume_cubic_centimetres"`
	TaiwanMiscellaneousCost         common.Money `json:"taiwan_miscellaneous_cost"`
	TaiwanFreightCost               common.Money `json:"taiwan_freight_cost"`
	TaiwanStuffingCost              common.Money `json:"taiwan_stuffing_cost"`
	AustraliaPortCost               common.Money `json:"australia_port_cost"`
	AustraliaBrokerageCost          common.Money `json:"australia_brokerage_cost"`
}
