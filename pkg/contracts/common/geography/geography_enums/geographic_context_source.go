package geography_enums

// GeographicContextSource records the authoritative input used to resolve a
// geographic context, or that only globally scoped rules were eligible.
type GeographicContextSource string

const (
	GeographicContextSourceDeliveryAddress              GeographicContextSource = "DELIVERY_ADDRESS"
	GeographicContextSourceFulfilmentDepot              GeographicContextSource = "FULFILMENT_DEPOT"
	GeographicContextSourceWholesaleOrganisationProfile GeographicContextSource = "WHOLESALE_ORGANISATION_PROFILE"
	GeographicContextSourceGlobalFallback               GeographicContextSource = "GLOBAL_FALLBACK"
)

func (s GeographicContextSource) IsValid() bool {
	switch s {
	case GeographicContextSourceDeliveryAddress,
		GeographicContextSourceFulfilmentDepot,
		GeographicContextSourceWholesaleOrganisationProfile,
		GeographicContextSourceGlobalFallback:
		return true
	default:
		return false
	}
}

func (s GeographicContextSource) String() string { return string(s) }
