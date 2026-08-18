// Package pricebook holds the Pricing-owned authoritative price container:
// the price book, its per-SKU entries, and the assignments that bind a book to
// buyers in a market. Authoritative prices live here and never inside Product
// or SKU, and one country's price is never derived from another country's
// price through currency conversion.
package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/customers/wholesale/wholesale_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/product/product_enums"
)

// PriceBook is one pricing context inside exactly one market. It owns the
// currency and exponent its entries are stated in, the channel and audience it
// serves, whether its amounts include tax, and the price-ending policy applied
// when drafts are generated for it.
type PriceBook struct {
	ID               string                            `json:"id"`
	Code             string                            `json:"code"`
	Name             string                            `json:"name"`
	MarketCode       string                            `json:"market_code"`
	Currency         money.CurrencyCode                `json:"currency"`
	CurrencyExponent money.CurrencyExponent            `json:"currency_exponent"`
	Channel          commerce_enums.OrderType          `json:"channel"`
	Audience         product_enums.PriceAudience       `json:"audience"`
	TaxInclusion     pricebook_enums.PriceTaxInclusion `json:"tax_inclusion"`
	PriceEnding      pricebook_enums.PriceEndingPolicy `json:"price_ending"`
	Status           pricebook_enums.PriceBookStatus   `json:"status"`
	ValidFrom        time.Time                         `json:"valid_from"`
	ValidUntil       *time.Time                        `json:"valid_until,omitempty"`
	Revision         int64                             `json:"revision"`

	audit.AuditFields
}

// PriceEntry is one SKU's commercial amount inside one price book. Only an
// approved entry inside its validity window is resolvable. Changing an entry
// never changes an already captured transaction snapshot.
type PriceEntry struct {
	ID            string                           `json:"id"`
	PriceBookCode string                           `json:"price_book_code"`
	SKUCode       string                           `json:"sku_code"`
	Amount        money.Money                      `json:"amount"`
	Status        pricebook_enums.PriceEntryStatus `json:"status"`
	Derivation    pricebook_enums.PriceDerivation  `json:"derivation"`
	ValidFrom     time.Time                        `json:"valid_from"`
	ValidUntil    *time.Time                       `json:"valid_until,omitempty"`
	Approval      *audit.LifecycleAction           `json:"approval,omitempty"`
	Rejection     *audit.LifecycleAction           `json:"rejection,omitempty"`
	Withdrawal    *audit.LifecycleAction           `json:"withdrawal,omitempty"`
	// SourceBaseCostRevision is the base acquisition cost revision a
	// generated suggestion was derived from. It is absent on manually
	// entered and imported amounts.
	SourceBaseCostRevision *int64 `json:"source_base_cost_revision,omitempty"`
	Revision               int64  `json:"revision"`

	audit.AuditFields
}

// PriceBookAssignment binds one price book to buyers inside one market. Kind
// selects which resolution level the assignment participates in, and only the
// fields relevant to that level are populated.
type PriceBookAssignment struct {
	ID            string                                  `json:"id"`
	MarketCode    string                                  `json:"market_code"`
	PriceBookCode string                                  `json:"price_book_code"`
	Kind          pricebook_enums.PriceBookAssignmentKind `json:"kind"`
	Channel       commerce_enums.OrderType                `json:"channel,omitempty"`
	// OrganisationCategory is set only for organisation-category
	// assignments.
	OrganisationCategory wholesale_enums.WholesaleOrganisationCategory `json:"organisation_category,omitempty"`
	// OrganisationCode is set only for exact-organisation overrides.
	OrganisationCode string                          `json:"organisation_code,omitempty"`
	Status           pricebook_enums.PriceBookStatus `json:"status"`
	ValidFrom        time.Time                       `json:"valid_from"`
	ValidUntil       *time.Time                      `json:"valid_until,omitempty"`
	Revision         int64                           `json:"revision"`

	audit.AuditFields
}
