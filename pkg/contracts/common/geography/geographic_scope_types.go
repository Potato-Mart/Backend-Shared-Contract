package geography

// GeographicScopeMode distinguishes rules that apply everywhere from rules
// that apply only to their inclusive geographic targets.
type GeographicScopeMode string

const (
	GeographicScopeModeGlobal   GeographicScopeMode = "GLOBAL"
	GeographicScopeModeTargeted GeographicScopeMode = "TARGETED"
)

func (m GeographicScopeMode) IsValid() bool {
	switch m {
	case GeographicScopeModeGlobal, GeographicScopeModeTargeted:
		return true
	default:
		return false
	}
}

func (m GeographicScopeMode) String() string { return string(m) }

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

// GeographicContextSource records which buyer profile supplied a resolved
// geographic context, or that only globally scoped rules were eligible.
type GeographicContextSource string

const (
	GeographicContextSourceRetailCustomerProfile        GeographicContextSource = "RETAIL_CUSTOMER_PROFILE"
	GeographicContextSourceWholesaleOrganisationProfile GeographicContextSource = "WHOLESALE_ORGANISATION_PROFILE"
	GeographicContextSourceGlobalFallback               GeographicContextSource = "GLOBAL_FALLBACK"
)

func (s GeographicContextSource) IsValid() bool {
	switch s {
	case GeographicContextSourceRetailCustomerProfile,
		GeographicContextSourceWholesaleOrganisationProfile,
		GeographicContextSourceGlobalFallback:
		return true
	default:
		return false
	}
}

func (s GeographicContextSource) String() string { return string(s) }
