package order

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/customers/retail"
	shippingenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/orders/shipping"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/product"
)

// BuyerContext is the shared, channel-independent description of who is
// buying. It is intentionally separate from the order channel
// (Order.Channel / Cart.Channel, a common.OrderType): POS is a channel,
// never a buyer type. A wholesale organisation buying in the physical shop
// is Type=BuyerTypeWholesaleOrganisation on Channel=OrderTypePOS, whereas a
// walk-in customer on the same channel is Type=BuyerTypeGuestRetail.
//
// The optional reference IDs resolve the buyer to a concrete identity when
// one exists. WholesaleOrganisationCode / OrganisationAccessID carry the B2B
// linkage that the existing Cart.CustomerNumber / Order.Customer fields cannot.
type BuyerContext struct {
	Type                      customerenum.BuyerType        `json:"type,omitempty"`
	RetailCustomerNumber      string                        `json:"retail_customer_number,omitempty"`
	WholesaleOrganisationCode string                        `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string                        `json:"organisation_access_id,omitempty"`
	FulfilmentIntent          shippingenum.FulfilmentIntent `json:"fulfilment_intent,omitempty"`
}

// PricingContext is the shared commercial pricing context under which a line
// price was determined. It is descriptive only: the contract records the
// audience and visibility that applied, never the resolver logic that picked
// the price.
type PricingContext struct {
	Audience          productenum.PriceAudience   `json:"audience,omitempty"`
	Visibility        productenum.PriceVisibility `json:"visibility,omitempty"`
	GeographicContext geography.GeographicContext `json:"geographic_context"`
}
