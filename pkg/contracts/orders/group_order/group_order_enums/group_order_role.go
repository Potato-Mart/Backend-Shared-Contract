package group_order_enums

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
