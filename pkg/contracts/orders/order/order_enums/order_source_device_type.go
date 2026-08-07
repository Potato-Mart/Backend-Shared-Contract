package order_enums

// OrderSourceDeviceType identifies the customer-facing device/app surface where a
// sales order originated.
type OrderSourceDeviceType string

const (
	OrderSourceDeviceTypeIOS       OrderSourceDeviceType = "ios"
	OrderSourceDeviceTypeAndroid   OrderSourceDeviceType = "android"
	OrderSourceDeviceTypePC        OrderSourceDeviceType = "pc"
	OrderSourceDeviceTypeMobileWeb OrderSourceDeviceType = "mobile_web"
	OrderSourceDeviceTypeTablet    OrderSourceDeviceType = "tablet"
	OrderSourceDeviceTypePos       OrderSourceDeviceType = "pos"
	OrderSourceDeviceTypeManual    OrderSourceDeviceType = "manual"
	OrderSourceDeviceTypePhone     OrderSourceDeviceType = "phone"
	OrderSourceDeviceTypeVR        OrderSourceDeviceType = "vr"
)

// IsValid reports whether d is a known OrderSourceDeviceType.
func (d OrderSourceDeviceType) IsValid() bool {
	switch d {
	case OrderSourceDeviceTypeIOS,
		OrderSourceDeviceTypeAndroid,
		OrderSourceDeviceTypePC,
		OrderSourceDeviceTypeMobileWeb,
		OrderSourceDeviceTypeTablet,
		OrderSourceDeviceTypePos,
		OrderSourceDeviceTypeManual,
		OrderSourceDeviceTypePhone,
		OrderSourceDeviceTypeVR:
		return true
	}
	return false
}

func (d OrderSourceDeviceType) String() string { return string(d) }
