package purchase

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/purchase/purchase_enums"
)

// SupplierTaxIdentity is the supplier's own registration evidence recorded on
// an invoice. Registration is an entity fact about the supplier and never a
// statement about a product's taxability.
type SupplierTaxIdentity struct {
	SupplierCode string `json:"supplier_code"`
	SupplierName string `json:"supplier_name"`
	// BusinessNumberScheme names the register the number belongs to, for
	// example "abn".
	BusinessNumberScheme string         `json:"business_number_scheme,omitempty"`
	BusinessNumber       string         `json:"business_number,omitempty"`
	TaxRegistered        bool           `json:"tax_registered"`
	TaxRegisteredFrom    *temporal.Date `json:"tax_registered_from,omitempty"`
}

// SupplierInvoiceDocument is the immutable reference to the received document.
// ContentSHA256 makes the stored artefact tamper-evident and, with
// DuplicateKey on the invoice, protects against re-recording the same invoice.
type SupplierInvoiceDocument struct {
	Reference     string    `json:"reference"`
	ContentSHA256 string    `json:"content_sha256"`
	MediaID       string    `json:"media_id,omitempty"`
	ReceivedAt    time.Time `json:"received_at"`
}

// SupplierInvoiceLine is one purchased line with its declared tax evidence.
// Freight and duty stay recorded here separately and are never blended into
// carrying cost.
type SupplierInvoiceLine struct {
	ID    string `json:"id"`
	SKUID string `json:"sku_id"`
	// SKUCode is the frozen SKU code captured when the invoice was recorded.
	SKUCode         string                               `json:"sku_code"`
	Description     string                               `json:"description,omitempty"`
	PackageOptionID string                               `json:"package_option_id,omitempty"`
	Composition     packaging.PackageCompositionSnapshot `json:"composition"`
	BaseUnits       int64                                `json:"base_units"`

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

// SupplierInvoice is the authoritative purchase tax record for one supplier
// document. An invoice carrying a line with unknown taxability or an unknown
// price basis cannot be confirmed.
type SupplierInvoice struct {
	ID            string              `json:"id"`
	InvoiceNumber string              `json:"invoice_number"`
	Supplier      SupplierTaxIdentity `json:"supplier"`
	IssueDate     temporal.Date       `json:"issue_date"`

	Currency         money.CurrencyCode     `json:"currency"`
	CurrencyExponent money.CurrencyExponent `json:"currency_exponent"`

	PurchaseOrderNumbers []string              `json:"purchase_order_numbers,omitempty"`
	ReceiptIDs           []string              `json:"receipt_ids,omitempty"`
	Lines                []SupplierInvoiceLine `json:"lines"`

	Subtotal money.Money `json:"subtotal"`
	// DeclaredTaxTotal is the supplier's own stated tax total, when given.
	DeclaredTaxTotal   *money.Money `json:"declared_tax_total,omitempty"`
	CalculatedTaxTotal money.Money  `json:"calculated_tax_total"`
	FreightTotal       *money.Money `json:"freight_total,omitempty"`
	DutyTotal          *money.Money `json:"duty_total,omitempty"`
	Total              money.Money  `json:"total"`

	TaxSource     purchase_enums.SupplierTaxSource   `json:"tax_source"`
	InputTaxClaim purchase_enums.InputTaxClaimStatus `json:"input_tax_claim"`
	// QualifyingTaxInvoice records whether the document meets the market's
	// tax-invoice evidence requirements for an input-tax claim.
	QualifyingTaxInvoice bool `json:"qualifying_tax_invoice"`

	Document SupplierInvoiceDocument `json:"document"`
	// DuplicateKey is the deterministic supplier plus invoice-number key a
	// unique index enforces so the same invoice cannot be recorded twice.
	DuplicateKey   string                               `json:"duplicate_key"`
	Status         purchase_enums.SupplierInvoiceStatus `json:"status"`
	Reconciliation *audit.LifecycleAction               `json:"reconciliation,omitempty"`
	Revision       int64                                `json:"revision"`

	audit.AuditFields
}
