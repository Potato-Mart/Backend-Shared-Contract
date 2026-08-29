package procurement

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/procurement/purchase_enums"
)

// PurchaseOrder is the independently mutable procurement order root.
type PurchaseOrder struct {
	ID          string                             `json:"id"`
	OrderNumber string                             `json:"order_number"`
	Status      purchase_enums.PurchaseOrderStatus `json:"status"`
	// DepotCode, MarketCode, and CountryCode are the denormalized receiving
	// site, market, and country, carried so a geographically scoped staff
	// query is a plain indexed match.
	DepotCode    string                  `json:"depot_code,omitempty"`
	MarketCode   string                  `json:"market_code,omitempty"`
	CountryCode  geography.CountryCode   `json:"country_code,omitempty"`
	Currency     money.CurrencyCode      `json:"currency"`
	Items        []PurchaseOrderItem     `json:"items"`
	Subtotal     money.Money             `json:"subtotal"`
	TaxAmount    money.Money             `json:"tax_amount"`
	ShippingCost *money.Money            `json:"shipping_cost,omitempty"`
	Total        money.Money             `json:"total"`
	Reference    string                  `json:"reference,omitempty"`
	SupplierCode string                  `json:"supplier_code"`
	SupplierName string                  `json:"supplier_name,omitempty"`
	ExpectedAt   *time.Time              `json:"expected_at,omitempty"`
	SubmittedAt  *time.Time              `json:"submitted_at,omitempty"`
	ConfirmedAt  *time.Time              `json:"confirmed_at,omitempty"`
	CancelledAt  *time.Time              `json:"cancelled_at,omitempty"`
	CompletedAt  *time.Time              `json:"completed_at,omitempty"`
	Note         string                  `json:"note,omitempty"`
	InternalNote string                  `json:"internal_note,omitempty"`
	History      []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}
