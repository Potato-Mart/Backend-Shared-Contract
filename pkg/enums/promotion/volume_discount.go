package promotionenum

// VolumeDiscountAppliesTo restricts which customer channel a
// carton-based volume discount tier applies to.
type VolumeDiscountAppliesTo string

const (
	VolumeDiscountAppliesToAll       VolumeDiscountAppliesTo = "ALL"
	VolumeDiscountAppliesToWholesale VolumeDiscountAppliesTo = "WHOLESALE"
	VolumeDiscountAppliesToRetail    VolumeDiscountAppliesTo = "RETAIL"
)

func (v VolumeDiscountAppliesTo) IsValid() bool {
	switch v {
	case VolumeDiscountAppliesToAll, VolumeDiscountAppliesToWholesale, VolumeDiscountAppliesToRetail:
		return true
	}
	return false
}

func (v VolumeDiscountAppliesTo) String() string { return string(v) }
