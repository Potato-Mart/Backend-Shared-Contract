package ads_enums

// AdPostIneligibilityReason says why an existing page or Instagram post cannot
// be promoted as an ad. It is a code only: the explanation shown to staff is
// admin-owned, and the platform wording is never passed through.
type AdPostIneligibilityReason string

const (
	// AdPostIneligibilityReasonLicensedMusic is a post whose soundtrack is
	// licensed for organic use only.
	AdPostIneligibilityReasonLicensedMusic AdPostIneligibilityReason = "licensed_music"
	// AdPostIneligibilityReasonSharedContent is a post that shares someone
	// else's content rather than our own.
	AdPostIneligibilityReasonSharedContent AdPostIneligibilityReason = "shared_content"
	// AdPostIneligibilityReasonUnsupportedType is a post whose media type
	// cannot be promoted.
	AdPostIneligibilityReasonUnsupportedType AdPostIneligibilityReason = "unsupported_type"
)

// IsValid reports whether r is a known AdPostIneligibilityReason.
func (r AdPostIneligibilityReason) IsValid() bool {
	switch r {
	case AdPostIneligibilityReasonLicensedMusic, AdPostIneligibilityReasonSharedContent, AdPostIneligibilityReasonUnsupportedType:
		return true
	}
	return false
}

func (r AdPostIneligibilityReason) String() string { return string(r) }
