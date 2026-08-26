package shipping_enums

import "testing"

func TestFulfilmentIntentDigitalIsKnown(t *testing.T) {
	for _, intent := range []FulfilmentIntent{
		FulfilmentIntentDelivery,
		FulfilmentIntentPickup,
		FulfilmentIntentInStoreCarry,
		FulfilmentIntentDigital,
	} {
		if !intent.IsValid() {
			t.Fatalf("known fulfilment intent %q is invalid", intent)
		}
	}
	if FulfilmentIntent("courier").IsValid() {
		t.Fatal("unknown fulfilment intent was classified as valid")
	}
}
