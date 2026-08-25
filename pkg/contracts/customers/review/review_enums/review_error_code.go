package review_enums

// ReviewErrorCode is a stable domain error code shared by review clients.
type ReviewErrorCode string

const (
	ReviewErrorCodeNotFound              ReviewErrorCode = "REVIEW_NOT_FOUND"
	ReviewErrorCodeQualificationRequired ReviewErrorCode = "REVIEW_QUALIFICATION_REQUIRED"
)

func (c ReviewErrorCode) String() string { return string(c) }

func (c ReviewErrorCode) IsValid() bool {
	switch c {
	case ReviewErrorCodeNotFound, ReviewErrorCodeQualificationRequired:
		return true
	default:
		return false
	}
}
