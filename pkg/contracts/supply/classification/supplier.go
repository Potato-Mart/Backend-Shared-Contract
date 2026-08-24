package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/party"
)

// Supplier is the full supplier record. A supplier is an organisation, so it
// carries the complete organisation profile via party.OrganisationDetail
// (which embeds PartyRef for id / name / phone / email, plus registration,
// tax, addresses, branding and other organisation fields).
type Supplier struct {
	party.OrganisationDetail
	GeographicLocation   *geography.Address           `json:"geographic_location,omitempty"`
	AvailableMarketCodes []string                     `json:"available_market_codes,omitempty"`
	AvailableProducts    []SupplierAvailableProduct   `json:"available_products,omitempty"`
	AvailablePromotions  []SupplierAvailablePromotion `json:"available_promotions,omitempty"`
	audit.AuditFields
}

// SupplierAvailableProduct describes one supplier's terms and identity for a
// Potato Mart SKU. Unobserved commercial terms remain absent.
type SupplierAvailableProduct struct {
	SKUCode                 string                       `json:"sku_code"`
	SupplierSKUCode         string                       `json:"supplier_sku_code,omitempty"`
	ProductNames            []localization.LocalizedName `json:"product_names"`
	OfferedPrice            *money.Money                 `json:"offered_price,omitempty"`
	MinimumPurchaseQuantity int64                        `json:"minimum_purchase_quantity,omitempty"`
	Manufacturing           *ProductManufacturing        `json:"manufacturing,omitempty"`
}

// SupplierAvailablePromotion is a locale-aware supplier promotion without
// product-specific qualification or pricing policy.
type SupplierAvailablePromotion struct {
	Names        []localization.LocalizedName        `json:"names"`
	Descriptions []localization.LocalizedDescription `json:"descriptions,omitempty"`
}

// ProductSupplierRef is the code-only supplier relationship persisted with a
// product. Display names, contacts, addresses, legal identifiers, and
// operational fields are resolved from the supplier master.
type ProductSupplierRef struct {
	Code string `json:"code"`
}

// ProductManufacturing contains customer-safe product manufacturing details.
// Its fields are optional so partially known declarations remain representable.
type ProductManufacturing struct {
	CompanyName string          `json:"company_name,omitempty"`
	CountryRef  *CountryCodeRef `json:"country_ref,omitempty"`
}

// ProductSupply groups customer-safe supplier and manufacturing information.
// The two sections are independently optional because manufacturing details
// may be known even when no supplier reference is available, and vice versa.
type ProductSupply struct {
	Suppliers     []ProductSupplierRef  `json:"suppliers,omitempty"`
	Manufacturing *ProductManufacturing `json:"manufacturing,omitempty"`
}
