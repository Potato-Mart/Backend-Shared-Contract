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

// GroupOrderRole identifies whether an order owns consolidated fulfilment or
// references that fulfilment as a participant.
type GroupOrderRole string

const (
	GroupOrderRoleConsolidatedParent GroupOrderRole = "CONSOLIDATED_PARENT"
	GroupOrderRoleParticipant        GroupOrderRole = "PARTICIPANT"
)

// IsValid reports whether r is a known GroupOrderRole value.
func (r GroupOrderRole) IsValid() bool {
	switch r {
	case GroupOrderRoleConsolidatedParent, GroupOrderRoleParticipant:
		return true
	}
	return false
}

func (r GroupOrderRole) String() string { return string(r) }
