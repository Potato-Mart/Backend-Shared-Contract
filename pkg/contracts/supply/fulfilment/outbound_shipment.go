package fulfilment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/party"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/shipping"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// OutboundShipment is one Supply dispatch unit for an order.
//
// Recipient is the delivery contact handed to the carrier; CustomerName keeps
// its meaning. Recipient and Address are personal data: a shipment carrying
// either must set ContainsPII in its DataProtectionFields.
type OutboundShipment struct {
	ID        string `json:"id"`
	DepotCode string `json:"depot_code"`
	// MarketCode and CountryCode are the denormalized market and country of
	// the dispatching depot, carried so a geographically scoped staff query
	// is a plain indexed match.
	MarketCode   string                                 `json:"market_code,omitempty"`
	CountryCode  geography.CountryCode                  `json:"country_code,omitempty"`
	OrderNumber  string                                 `json:"order_number"`
	CustomerName string                                 `json:"customer_name,omitempty"`
	Address      *geography.Address                     `json:"address,omitempty"`
	Recipient    *party.Recipient                       `json:"recipient,omitempty"`
	Operator     string                                 `json:"operator"`
	Status       warehouse_enums.OutboundShipmentStatus `json:"status"`
	Containers   []OutboundContainerPlan                `json:"containers,omitempty"`
	// Carrier is the optional delivery-company code recorded by Supply
	// (for example detrack, bcrc, aupost).
	Carrier string `json:"carrier,omitempty"`
	// DeliverySelection is copied from the accepted order, not re-routed.
	DeliverySelection *shipping.DeliverySelection `json:"delivery_selection,omitempty"`
	TrackingNumber    string                      `json:"tracking_number,omitempty"`
	Note              string                      `json:"note,omitempty"`
	DispatchedAt      *time.Time                  `json:"dispatched_at,omitempty"`
	DeliveredAt       *time.Time                  `json:"delivered_at,omitempty"`
	History           []security.HistoryEntry     `json:"history,omitempty"`
	// LatestCarrierBooking summarizes the order's latest
	// CarrierBookingAttempt. One attempt covers all of the order's shipments,
	// so every shipment it lists carries the same summary.
	LatestCarrierBooking *courier.CarrierBookingSummary `json:"latest_carrier_booking,omitempty"`
	// Warnings lists non-blocking conditions staff must see, such as a
	// carrier booking that still needs manual cancellation or an unpaid
	// order. Supply fills them on read; they never block dispatch.
	Warnings []ShipmentWarning `json:"warnings,omitempty"`
	audit.AuditFields
	security.DataProtectionFields
}
