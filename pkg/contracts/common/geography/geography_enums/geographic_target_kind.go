package geography_enums

// GeographicTargetKind identifies the hierarchy level named by one target.
type GeographicTargetKind string

const (
	GeographicTargetCountry GeographicTargetKind = "COUNTRY"
	GeographicTargetMarket  GeographicTargetKind = "MARKET"
	GeographicTargetState   GeographicTargetKind = "STATE"
	GeographicTargetDepot   GeographicTargetKind = "DEPOT"
)

func (k GeographicTargetKind) IsValid() bool {
	switch k {
	case GeographicTargetCountry, GeographicTargetMarket,
		GeographicTargetState, GeographicTargetDepot:
		return true
	default:
		return false
	}
}

func (k GeographicTargetKind) String() string { return string(k) }
