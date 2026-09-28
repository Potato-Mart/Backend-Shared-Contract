package courier_enums

// DeliveryAPIRequestPurpose identifies one manually configured courier API
// request supported by the Supply configuration API.
type DeliveryAPIRequestPurpose string

const (
	DeliveryAPIRequestPurposeConnection    DeliveryAPIRequestPurpose = "connection"
	DeliveryAPIRequestPurposeShippingAreas DeliveryAPIRequestPurpose = "shipping_areas"
	DeliveryAPIRequestPurposeTimeSlots     DeliveryAPIRequestPurpose = "time_slots"
)

func (p DeliveryAPIRequestPurpose) IsValid() bool {
	switch p {
	case DeliveryAPIRequestPurposeConnection,
		DeliveryAPIRequestPurposeShippingAreas,
		DeliveryAPIRequestPurposeTimeSlots:
		return true
	default:
		return false
	}
}

func (p DeliveryAPIRequestPurpose) String() string { return string(p) }
