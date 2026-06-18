package enums

// PortalAccessStatus describes whether an account can enter a specific portal.
type PortalAccessStatus string

const (
	// PortalAccessStatusPending means access has been requested but not granted.
	PortalAccessStatusPending PortalAccessStatus = "pending"
	// PortalAccessStatusActive means the account can enter the portal.
	PortalAccessStatusActive PortalAccessStatus = "active"
	// PortalAccessStatusSuspended means portal access is temporarily blocked.
	PortalAccessStatusSuspended PortalAccessStatus = "suspended"
	// PortalAccessStatusRevoked means portal access has been removed.
	PortalAccessStatusRevoked PortalAccessStatus = "revoked"
)

// IsValid reports whether s is a known PortalAccessStatus value.
func (s PortalAccessStatus) IsValid() bool {
	switch s {
	case PortalAccessStatusPending, PortalAccessStatusActive,
		PortalAccessStatusSuspended, PortalAccessStatusRevoked:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s PortalAccessStatus) String() string { return string(s) }

// CanAccess reports whether s permits portal entry.
func (s PortalAccessStatus) CanAccess() bool {
	return s == PortalAccessStatusActive
}
