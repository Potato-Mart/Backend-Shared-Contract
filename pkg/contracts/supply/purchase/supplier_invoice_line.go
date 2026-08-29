package purchase

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/purchase/purchase_enums"
)

// SupplierInvoiceLine is one purchased line with its declared tax evidence.
// Freight and duty stay recorded here separately and are never blended into
// carrying cost.
type SupplierInvoiceLine struct {
	ID string `json:"id"`
	// SKUCode is the frozen SKU code captured when the invoice was recorded.
	SKUCode           string                               `json:"sku_code"`
	Description       string                               `json:"description,omitempty"`
	PackageOptionCode string                               `json:"package_option_code,omitempty"`
	Composition       packaging.PackageCompositionSnapshot `json:"composition"`
	BaseUnits         int64                                `json:"base_units"`

	UnitPrice  money.Money `json:"unit_price"`
	LineAmount money.Money `json:"line_amount"`

	TaxTreatment purchase_enums.LineTaxTreatment `json:"tax_treatment"`
	PriceBasis   purchase_enums.TaxPriceBasis    `json:"price_basis"`
	// DeclaredTax is the amount the supplier stated for this line.
	DeclaredTax *money.Money `json:"declared_tax,omitempty"`
	// CalculatedTax is present only when the line is explicitly taxable and
	// its price basis is known.
	CalculatedTax *money.Money                     `json:"calculated_tax,omitempty"`
	TaxSource     purchase_enums.SupplierTaxSource `json:"tax_source"`
	// RecoverableTax is the portion excluded from carrying cost.
	RecoverableTax *money.Money                       `json:"recoverable_tax,omitempty"`
	InputTaxClaim  purchase_enums.InputTaxClaimStatus `json:"input_tax_claim"`

	FreightAmount *money.Money `json:"freight_amount,omitempty"`
	DutyAmount    *money.Money `json:"duty_amount,omitempty"`

	PurchaseOrderNumber string `json:"purchase_order_number,omitempty"`
	ReceiptID           string `json:"receipt_id,omitempty"`
	Note                string `json:"note,omitempty"`
}
