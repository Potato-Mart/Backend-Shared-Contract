package coupon_enums

// Visibility controls public discovery independently of issuance policy.
type Visibility string

const (
	VisibilityPublic   Visibility = "PUBLIC"
	VisibilityUnlisted Visibility = "UNLISTED"
)

func (v Visibility) IsValid() bool  { return v == VisibilityPublic || v == VisibilityUnlisted }
func (v Visibility) String() string { return string(v) }
