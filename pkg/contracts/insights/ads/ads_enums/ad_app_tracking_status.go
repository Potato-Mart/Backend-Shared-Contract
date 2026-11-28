package ads_enums

// AdAppTrackingStatus is the iOS app tracking transparency answer recorded
// with an ad click captured in the iOS app. Only an authorized answer permits
// a click id or a hashed contact digest to leave our systems; every other
// answer withholds both.
type AdAppTrackingStatus string

const (
	// AdAppTrackingStatusAuthorized means the person allowed tracking.
	AdAppTrackingStatusAuthorized AdAppTrackingStatus = "authorized"
	// AdAppTrackingStatusDenied means the person refused tracking.
	AdAppTrackingStatusDenied AdAppTrackingStatus = "denied"
	// AdAppTrackingStatusRestricted means tracking is not permitted on the
	// device, for example under a device policy.
	AdAppTrackingStatusRestricted AdAppTrackingStatus = "restricted"
	// AdAppTrackingStatusNotDetermined means the person has not been asked
	// yet.
	AdAppTrackingStatusNotDetermined AdAppTrackingStatus = "not_determined"
)

// IsValid reports whether s is a known AdAppTrackingStatus.
func (s AdAppTrackingStatus) IsValid() bool {
	switch s {
	case AdAppTrackingStatusAuthorized, AdAppTrackingStatusDenied,
		AdAppTrackingStatusRestricted, AdAppTrackingStatusNotDetermined:
		return true
	}
	return false
}

func (s AdAppTrackingStatus) String() string { return string(s) }
