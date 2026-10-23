package coupon_enums

// Visibility controls public discovery independently of issuance policy.
type Visibility string

const (
	VisibilityPublic       Visibility = "public"
	VisibilityCampaignOnly Visibility = "campaign_only"
	VisibilityUnlisted     Visibility = "unlisted"
)

func (v Visibility) IsValid() bool {
	return v == VisibilityPublic || v == VisibilityCampaignOnly || v == VisibilityUnlisted
}
func (v Visibility) String() string { return string(v) }
