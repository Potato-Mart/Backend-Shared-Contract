package enums

// DeliveryMethod is how an order physically reaches the buyer, managed by
// staff on the order. It is authoritative — unlike FulfilmentIntent, which
// is optional buyer context — and independent of the shipping rate that
// priced the order.
type DeliveryMethod string

const (
	// DeliveryMethodDelivery is delivered by our own or partner fleet.
	DeliveryMethodDelivery DeliveryMethod = "delivery"
	// DeliveryMethodPickup is collected by the buyer.
	DeliveryMethodPickup DeliveryMethod = "pickup"
	// DeliveryMethodOutsourced is delivered by a third-party company other
	// than the partner fleet; the company name lives on the order's
	// outsourced_carrier field.
	DeliveryMethodOutsourced DeliveryMethod = "outsourced"
)

// IsValid reports whether d is a known DeliveryMethod.
func (d DeliveryMethod) IsValid() bool {
	switch d {
	case DeliveryMethodDelivery, DeliveryMethodPickup, DeliveryMethodOutsourced:
		return true
	}
	return false
}

func (d DeliveryMethod) String() string { return string(d) }
