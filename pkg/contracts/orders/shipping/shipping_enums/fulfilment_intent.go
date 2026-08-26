package shipping_enums

// FulfilmentIntent is the buyer's intended fulfilment method for a cart or
// order. It is optional context — the authoritative fulfilment/shipping
// state still lives on the order's shipping fields.
type FulfilmentIntent string

const (
	// FulfilmentIntentDelivery is shipped/delivered to the buyer.
	FulfilmentIntentDelivery FulfilmentIntent = "delivery"
	// FulfilmentIntentPickup is collected by the buyer (click & collect).
	FulfilmentIntentPickup FulfilmentIntent = "pickup"
	// FulfilmentIntentInStoreCarry is bought and carried out in store.
	FulfilmentIntentInStoreCarry FulfilmentIntent = "in_store_carry"
	// FulfilmentIntentDigital is fulfilled without a delivery address or depot.
	// Its frozen geographic context records the global or market rules applied
	// when the cart was priced.
	FulfilmentIntentDigital FulfilmentIntent = "digital"
)

// IsValid reports whether f is a known FulfilmentIntent.
func (f FulfilmentIntent) IsValid() bool {
	switch f {
	case FulfilmentIntentDelivery, FulfilmentIntentPickup, FulfilmentIntentInStoreCarry, FulfilmentIntentDigital:
		return true
	}
	return false
}

func (f FulfilmentIntent) String() string { return string(f) }
