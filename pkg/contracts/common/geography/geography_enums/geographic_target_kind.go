package geography_enums

// GeographicTargetKind identifies the hierarchy level named by one target.
type GeographicTargetKind string

const (
	GeographicTargetCountry     GeographicTargetKind = "COUNTRY"
	GeographicTargetSubdivision GeographicTargetKind = "SUBDIVISION"
	GeographicTargetDepotRegion GeographicTargetKind = "DEPOT_REGION"
	GeographicTargetDepot       GeographicTargetKind = "DEPOT"
)

func (k GeographicTargetKind) IsValid() bool {
	switch k {
	case GeographicTargetCountry, GeographicTargetSubdivision,
		GeographicTargetDepotRegion, GeographicTargetDepot:
		return true
	default:
		return false
	}
}

func (k GeographicTargetKind) String() string { return string(k) }
