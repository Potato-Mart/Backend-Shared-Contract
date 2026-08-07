package identity_enums

// Portal identifies which front-door application a user, session, or device
// interaction belongs to. Admission policy is owned by each backend.
type Portal string

const (
	PortalControl   Portal = "control"
	PortalRetail    Portal = "retail"
	PortalWholesale Portal = "wholesale"
)

func (p Portal) IsValid() bool {
	switch p {
	case PortalControl, PortalRetail, PortalWholesale:
		return true
	default:
		return false
	}
}

func (p Portal) String() string { return string(p) }
