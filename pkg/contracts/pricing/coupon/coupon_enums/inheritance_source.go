package coupon_enums

type InheritanceSource string

const (
	InheritanceIndependent InheritanceSource = "independent"
	InheritanceCampaign    InheritanceSource = "campaign"
)

func (v InheritanceSource) IsValid() bool {
	return v == InheritanceIndependent || v == InheritanceCampaign
}
func (v InheritanceSource) String() string { return string(v) }
