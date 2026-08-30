package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/wholesale/wholesale_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/pricebook/pricebook_enums"
)

// PriceBookAssignment binds a price book to buyers within one market.
type PriceBookAssignment struct {
	ID                   string                                        `json:"id"`
	MarketCode           string                                        `json:"market_code"`
	PriceBookCode        string                                        `json:"price_book_code"`
	Kind                 pricebook_enums.PriceBookAssignmentKind       `json:"kind"`
	Channel              commerce_enums.OrderType                      `json:"channel,omitempty"`
	OrganisationCategory wholesale_enums.WholesaleOrganisationCategory `json:"organisation_category,omitempty"`
	OrganisationCode     string                                        `json:"organisation_code,omitempty"`
	Status               pricebook_enums.PriceBookStatus               `json:"status"`
	ValidFrom            time.Time                                     `json:"valid_from"`
	ValidUntil           *time.Time                                    `json:"valid_until,omitempty"`
	Revision             int64                                         `json:"revision"`

	audit.AuditFields
}
