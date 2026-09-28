package courier

// DeliveryAPIRequestConfigurations groups the independently configured
// connection, shipping-area, and time-slot requests. Nil requests and omitted
// field groups mean unconfigured; the shared contract supplies no defaults.
type DeliveryAPIRequestConfigurations struct {
	Connection    *DeliveryAPIRequestConfiguration `json:"connection,omitempty"`
	ShippingAreas *DeliveryAPIRequestConfiguration `json:"shipping_areas,omitempty"`
	TimeSlots     *DeliveryAPIRequestConfiguration `json:"time_slots,omitempty"`
}
