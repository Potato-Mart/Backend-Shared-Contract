package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/pricing/pricebook/pricebook_enums"
)

// MembershipTierPriceBookAssignment binds a membership tier to one retail
// price book in a market and channel. PublicOfferEnabled records whether the
// resolved tier offer may be projected publicly; its zero value is false.
type MembershipTierPriceBookAssignment struct {
	ID                 string                          `json:"id"`
	MarketCode         string                          `json:"market_code"`
	Channel            commerce_enums.OrderType        `json:"channel"`
	MembershipTierKey  string                          `json:"membership_tier_key"`
	PriceBookCode      string                          `json:"price_book_code"`
	PublicOfferEnabled bool                            `json:"public_offer_enabled"`
	Status             pricebook_enums.PriceBookStatus `json:"status"`
	ValidFrom          time.Time                       `json:"valid_from"`
	ValidUntil         *time.Time                      `json:"valid_until,omitempty"`
	Revision           int64                           `json:"revision"`

	audit.AuditFields
}
