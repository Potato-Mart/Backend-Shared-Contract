package wholesaleenum

// WholesaleOrganisationStatus describes the approval lifecycle for a
// wholesale organisation.
type WholesaleOrganisationStatus string

const (
	// WholesaleOrganisationStatusPending means the organisation is awaiting review.
	WholesaleOrganisationStatusPending WholesaleOrganisationStatus = "pending"
	// WholesaleOrganisationStatusApproved means the organisation is approved.
	WholesaleOrganisationStatusApproved WholesaleOrganisationStatus = "approved"
	// WholesaleOrganisationStatusSuspended means the organisation is temporarily blocked.
	WholesaleOrganisationStatusSuspended WholesaleOrganisationStatus = "suspended"
	// WholesaleOrganisationStatusRejected means the organisation was not approved.
	WholesaleOrganisationStatusRejected WholesaleOrganisationStatus = "rejected"
	// WholesaleOrganisationStatusClosed means the organisation is closed.
	WholesaleOrganisationStatusClosed WholesaleOrganisationStatus = "closed"
)

// IsValid reports whether s is a known WholesaleOrganisationStatus value.
func (s WholesaleOrganisationStatus) IsValid() bool {
	switch s {
	case WholesaleOrganisationStatusPending, WholesaleOrganisationStatusApproved,
		WholesaleOrganisationStatusSuspended, WholesaleOrganisationStatusRejected,
		WholesaleOrganisationStatusClosed:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s WholesaleOrganisationStatus) String() string { return string(s) }

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
