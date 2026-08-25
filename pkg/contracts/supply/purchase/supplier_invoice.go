package purchase

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/purchase/purchase_enums"
)

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
