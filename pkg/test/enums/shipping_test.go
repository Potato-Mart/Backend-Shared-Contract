package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/orders/shipping/shipping_enums"
)

func TestShippingEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "shippingenum.DeliveryMethod", valid: []stringEnum{shipping_enums.DeliveryMethodDelivery, shipping_enums.DeliveryMethodPickup, shipping_enums.DeliveryMethodOutsourced}, invalid: shipping_enums.DeliveryMethod("__invalid__")},
		{name: "shippingenum.FulfilmentIntent", valid: []stringEnum{shipping_enums.FulfilmentIntentDelivery, shipping_enums.FulfilmentIntentPickup, shipping_enums.FulfilmentIntentInStoreCarry}, invalid: shipping_enums.FulfilmentIntent("__invalid__")},
		{name: "shippingenum.ShippingRateName", valid: []stringEnum{shipping_enums.ShippingRateNameStandard, shipping_enums.ShippingRateNameExpress, shipping_enums.ShippingRateNamePickup}, invalid: shipping_enums.ShippingRateName("__invalid__")},
	})
}
