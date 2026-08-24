package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/customers/retail/retail_enums"
)

// BuyerContext is the shared, channel-independent description of who is
// buying. It is intentionally separate from the order channel
// (Order.Channel / Cart.Channel, a commerce_enums.OrderType): POS is a channel,
// never a buyer type. A wholesale organisation buying in the physical shop
// is Type=BuyerTypeWholesaleOrganisation on Channel=OrderTypePOS, whereas a
// walk-in customer on the same channel is Type=BuyerTypeGuestRetail.
//
// The optional reference IDs resolve the buyer to a concrete identity when
// one exists. WholesaleOrganisationCode / OrganisationAccessID carry the B2B
// linkage that the existing Cart.CustomerNumber / Order.Customer fields cannot.
type BuyerContext struct {
	Type                      retail_enums.BuyerType `json:"type,omitempty"`
	RetailCustomerNumber      string                 `json:"retail_customer_number,omitempty"`
	WholesaleOrganisationCode string                 `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string                 `json:"organisation_access_id,omitempty"`
}
