package courier_enums

// PostalCodeMatchMode makes broad matching explicit instead of treating an
// empty include list as permission to deliver anywhere in a country.
type PostalCodeMatchMode string

const (
	PostalCodeMatchModeIncludeOnly PostalCodeMatchMode = "include_only"
	PostalCodeMatchModeAllExcept   PostalCodeMatchMode = "all_except"
)

func (m PostalCodeMatchMode) IsValid() bool {
	switch m {
	case PostalCodeMatchModeIncludeOnly, PostalCodeMatchModeAllExcept:
		return true
	default:
		return false
	}
}

func (m PostalCodeMatchMode) String() string { return string(m) }
