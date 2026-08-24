package deletion_enums

// ErasureDisposition is the record-free summary of what an
// erase_or_deidentify command achieved. Mixed means the service both erased
// data and deidentified policy-retained evidence; it never indicates which
// collection or source record was involved.
type ErasureDisposition string

const (
	ErasureDispositionErased           ErasureDisposition = "erased"
	ErasureDispositionDeidentified     ErasureDisposition = "deidentified"
	ErasureDispositionMixed            ErasureDisposition = "mixed"
	ErasureDispositionNoSubjectData    ErasureDisposition = "no_subject_data"
	ErasureDispositionRetainedByPolicy ErasureDisposition = "retained_by_policy"
)

// IsValid reports whether d is a supported erasure disposition.
func (d ErasureDisposition) IsValid() bool {
	switch d {
	case ErasureDispositionErased, ErasureDispositionDeidentified,
		ErasureDispositionMixed, ErasureDispositionNoSubjectData,
		ErasureDispositionRetainedByPolicy:
		return true
	default:
		return false
	}
}

// String returns the wire value for d.
func (d ErasureDisposition) String() string { return string(d) }
