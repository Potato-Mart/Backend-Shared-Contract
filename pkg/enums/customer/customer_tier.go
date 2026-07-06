package customerenum

type CustomerTier string

const (
	CustomerTierStandard CustomerTier = "standard"
	CustomerTierSilver   CustomerTier = "silver"
	CustomerTierGold     CustomerTier = "gold"
	CustomerTierPlatinum CustomerTier = "platinum"
)

// IsValid reports whether p is a known CustomerTier.
func (p CustomerTier) IsValid() bool {
	switch p {
	case CustomerTierStandard, CustomerTierSilver, CustomerTierGold, CustomerTierPlatinum:
		return true
	}
	return false
}

func (p CustomerTier) String() string { return string(p) }
