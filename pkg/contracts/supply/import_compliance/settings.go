package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
)

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

	audit.AuditFields
}

type IngredientDeclarationRule struct {
	ID          string `json:"id"`
	Ingredients string `json:"ingredients"`
	Statement   string `json:"statement"`
}

// AirCargoSettings preserves the current import-settings page fields while
// using integer grams/cubic centimetres and fixed-precision Money values.
type AirCargoSettings struct {
	ReferenceWeightGrams                         int64       `json:"reference_weight_grams"`
	TaiwanInspectionCost                         money.Money `json:"taiwan_inspection_cost"`
	StorageThresholdGrams                        int64       `json:"storage_threshold_grams"`
	StorageUnderThresholdCost                    money.Money `json:"storage_under_threshold_cost"`
	StorageAtOrOverThresholdCost                 money.Money `json:"storage_at_or_over_threshold_cost"`
	VolumetricDivisorCubicCentimetresPerKilogram int64       `json:"volumetric_divisor_cubic_centimetres_per_kilogram"`
	TaiwanFreightCost                            money.Money `json:"taiwan_freight_cost"`
	TaiwanHandlingCost                           money.Money `json:"taiwan_handling_cost"`
	TaiwanTransportCost                          money.Money `json:"taiwan_transport_cost"`
	TaiwanBrokerageCost                          money.Money `json:"taiwan_brokerage_cost"`
	TaiwanTruckCost                              money.Money `json:"taiwan_truck_cost"`
	AustraliaTerminalCost                        money.Money `json:"australia_terminal_cost"`
	AustraliaTruckCost                           money.Money `json:"australia_truck_cost"`
	AustraliaDocumentCost                        money.Money `json:"australia_document_cost"`
	AustraliaBrokerageCost                       money.Money `json:"australia_brokerage_cost"`
	AustraliaQuarantineCost                      money.Money `json:"australia_quarantine_cost"`
}

type AmbientSeaSettings struct {
	ReferenceVolumeCubicCentimetres int64       `json:"reference_volume_cubic_centimetres"`
	TaiwanConsolidationCost         money.Money `json:"taiwan_consolidation_cost"`
	TaiwanHandlingCost              money.Money `json:"taiwan_handling_cost"`
	TaiwanBrokerageCost             money.Money `json:"taiwan_brokerage_cost"`
	AustraliaTerminalCost           money.Money `json:"australia_terminal_cost"`
	AustraliaTruckCost              money.Money `json:"australia_truck_cost"`
	AustraliaFreightCost            money.Money `json:"australia_freight_cost"`
	AustraliaDocumentCost           money.Money `json:"australia_document_cost"`
	AustraliaBrokerageCost          money.Money `json:"australia_brokerage_cost"`
	AustraliaQuarantineCost         money.Money `json:"australia_quarantine_cost"`
}

type FrozenSeaSettings struct {
	ReferenceVolumeCubicCentimetres int64       `json:"reference_volume_cubic_centimetres"`
	TaiwanMiscellaneousCost         money.Money `json:"taiwan_miscellaneous_cost"`
	TaiwanFreightCost               money.Money `json:"taiwan_freight_cost"`
	TaiwanStuffingCost              money.Money `json:"taiwan_stuffing_cost"`
	AustraliaPortCost               money.Money `json:"australia_port_cost"`
	AustraliaBrokerageCost          money.Money `json:"australia_brokerage_cost"`
}
