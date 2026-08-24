package order_enums

// LooseSubstitutionPolicySource identifies who selected the case-to-each
// substitution setting captured on a cart or order line.
type LooseSubstitutionPolicySource string

const (
	LooseSubstitutionPolicySourceChannelDefault LooseSubstitutionPolicySource = "CHANNEL_DEFAULT"
	LooseSubstitutionPolicySourceBuyerSelected  LooseSubstitutionPolicySource = "BUYER_SELECTED"
	LooseSubstitutionPolicySourceGroupManager   LooseSubstitutionPolicySource = "GROUP_MANAGER_SELECTED"
)

// IsValid reports whether s is a known LooseSubstitutionPolicySource value.
func (s LooseSubstitutionPolicySource) IsValid() bool {
	switch s {
	case LooseSubstitutionPolicySourceChannelDefault,
		LooseSubstitutionPolicySourceBuyerSelected,
		LooseSubstitutionPolicySourceGroupManager:
		return true
	}
	return false
}

func (s LooseSubstitutionPolicySource) String() string { return string(s) }
