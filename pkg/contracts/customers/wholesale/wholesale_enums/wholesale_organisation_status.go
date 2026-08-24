package wholesale_enums

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
