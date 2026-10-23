package coupon_enums

type AutomaticPolicy string

const (
	AutomaticPolicyOnceWhenLive AutomaticPolicy = "once_when_live"
	AutomaticPolicyWhileActive  AutomaticPolicy = "while_active"
)

func (v AutomaticPolicy) IsValid() bool {
	return v == AutomaticPolicyOnceWhenLive || v == AutomaticPolicyWhileActive
}
func (v AutomaticPolicy) String() string { return string(v) }
