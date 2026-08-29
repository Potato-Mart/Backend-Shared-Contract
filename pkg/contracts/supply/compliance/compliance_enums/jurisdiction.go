package compliance_enums

type Jurisdiction string

const (
	JurisdictionAustralia Jurisdiction = "AU"
	JurisdictionTaiwan    Jurisdiction = "TW"
)

func (j Jurisdiction) IsValid() bool {
	switch j {
	case JurisdictionAustralia, JurisdictionTaiwan:
		return true
	}
	return false
}

func (j Jurisdiction) String() string { return string(j) }
