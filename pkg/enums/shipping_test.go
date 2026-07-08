package enums_test

import (
	"testing"

	shippingenum "github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/enums/shipping"
)

func TestShippingEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "shippingenum.DeliveryMethod", valid: []stringEnum{shippingenum.DeliveryMethodDelivery, shippingenum.DeliveryMethodPickup, shippingenum.DeliveryMethodOutsourced}, invalid: shippingenum.DeliveryMethod("__invalid__")},
		{name: "shippingenum.DeliveryRegion", valid: []stringEnum{shippingenum.DeliveryRegionLocalMelbourne, shippingenum.DeliveryRegionRegionalVIC, shippingenum.DeliveryRegionInterstate}, invalid: shippingenum.DeliveryRegion("__invalid__")},
		{name: "shippingenum.FulfilmentIntent", valid: []stringEnum{shippingenum.FulfilmentIntentDelivery, shippingenum.FulfilmentIntentPickup, shippingenum.FulfilmentIntentInStoreCarry}, invalid: shippingenum.FulfilmentIntent("__invalid__")},
		{name: "shippingenum.ShippingRateName", valid: []stringEnum{shippingenum.ShippingRateNameStandard, shippingenum.ShippingRateNameExpress, shippingenum.ShippingRateNamePickup}, invalid: shippingenum.ShippingRateName("__invalid__")},
	})
}
