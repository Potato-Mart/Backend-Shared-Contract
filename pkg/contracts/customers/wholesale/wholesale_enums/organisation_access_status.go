package wholesale_enums

// OrganisationAccessStatus describes the lifecycle for a user's access to a
// wholesale organisation.
type OrganisationAccessStatus string

const (
	// OrganisationAccessStatusPending means the access grant is awaiting activation.
	OrganisationAccessStatusPending OrganisationAccessStatus = "pending"
	// OrganisationAccessStatusActive means the access grant can be used.
	OrganisationAccessStatusActive OrganisationAccessStatus = "active"
	// OrganisationAccessStatusSuspended means the access grant is temporarily blocked.
	OrganisationAccessStatusSuspended OrganisationAccessStatus = "suspended"
	// OrganisationAccessStatusRevoked means the access grant has been removed.
	OrganisationAccessStatusRevoked OrganisationAccessStatus = "revoked"
)

// IsValid reports whether s is a known OrganisationAccessStatus value.
func (s OrganisationAccessStatus) IsValid() bool {
	switch s {
	case OrganisationAccessStatusPending, OrganisationAccessStatusActive,
		OrganisationAccessStatusSuspended, OrganisationAccessStatusRevoked:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s OrganisationAccessStatus) String() string { return string(s) }
