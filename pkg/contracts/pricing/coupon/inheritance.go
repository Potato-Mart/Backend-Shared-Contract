package coupon

import "github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/pricing/coupon/coupon_enums"

type Inheritance struct {
	Scope             coupon_enums.InheritanceSource `json:"scope"`
	AcquisitionWindow coupon_enums.InheritanceSource `json:"acquisition_window"`
	RedemptionWindow  coupon_enums.InheritanceSource `json:"redemption_window"`
}
