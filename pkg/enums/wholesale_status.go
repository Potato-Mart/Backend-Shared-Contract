package enums

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

// WholesaleMembershipStatus describes the lifecycle for a user's membership in
// a wholesale organisation.
type WholesaleMembershipStatus string

const (
	// WholesaleMembershipStatusPending means the membership is awaiting activation.
	WholesaleMembershipStatusPending WholesaleMembershipStatus = "pending"
	// WholesaleMembershipStatusActive means the membership can be used.
	WholesaleMembershipStatusActive WholesaleMembershipStatus = "active"
	// WholesaleMembershipStatusSuspended means the membership is temporarily blocked.
	WholesaleMembershipStatusSuspended WholesaleMembershipStatus = "suspended"
	// WholesaleMembershipStatusRevoked means the membership has been removed.
	WholesaleMembershipStatusRevoked WholesaleMembershipStatus = "revoked"
)

// IsValid reports whether s is a known WholesaleMembershipStatus value.
func (s WholesaleMembershipStatus) IsValid() bool {
	switch s {
	case WholesaleMembershipStatusPending, WholesaleMembershipStatusActive,
		WholesaleMembershipStatusSuspended, WholesaleMembershipStatusRevoked:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s WholesaleMembershipStatus) String() string { return string(s) }
