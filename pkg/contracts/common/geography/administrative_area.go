package geography

// AdministrativeAreaType identifies the kind of first-level or equivalent
// subdivision represented by an administrative-area reference.
type AdministrativeAreaType string

const (
	AdministrativeAreaState      AdministrativeAreaType = "STATE"
	AdministrativeAreaTerritory  AdministrativeAreaType = "TERRITORY"
	AdministrativeAreaProvince   AdministrativeAreaType = "PROVINCE"
	AdministrativeAreaPrefecture AdministrativeAreaType = "PREFECTURE"
	AdministrativeAreaRegion     AdministrativeAreaType = "REGION"
	AdministrativeAreaDistrict   AdministrativeAreaType = "DISTRICT"
)

func (a AdministrativeAreaType) IsValid() bool {
	switch a {
	case AdministrativeAreaState, AdministrativeAreaTerritory,
		AdministrativeAreaProvince, AdministrativeAreaPrefecture,
		AdministrativeAreaRegion, AdministrativeAreaDistrict:
		return true
	default:
		return false
	}
}

func (a AdministrativeAreaType) String() string { return string(a) }
