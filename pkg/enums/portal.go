package enums

// Portal identifies which front-door application a user, session, or
// device interaction belongs to.
type Portal string

const (
	PortalControl Portal = "control"
	PortalStore   Portal = "store"
	PortalPartner Portal = "partner"
)

// IsValid reports whether p is a known Portal.
func (p Portal) IsValid() bool {
	switch p {
	case PortalControl, PortalStore, PortalPartner:
		return true
	}
	return false
}

func (p Portal) String() string { return string(p) }
