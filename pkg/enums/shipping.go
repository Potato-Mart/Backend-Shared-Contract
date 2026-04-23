package enums

type ShippingRateName string

const (
	ShippingRateNameStandard ShippingRateName = "STANDARD"
	ShippingRateNameExpress  ShippingRateName = "EXPRESS"
	ShippingRateNamePickup   ShippingRateName = "PICKUP"
)

// IsValid reports whether p is a known ShippingRateName.
func (p ShippingRateName) IsValid() bool {
	switch p {
	case ShippingRateNameStandard, ShippingRateNameExpress, ShippingRateNamePickup:
		return true
	}
	return false
}

func (p ShippingRateName) String() string { return string(p) }
