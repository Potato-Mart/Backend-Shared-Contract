package enums

type CustomerTier string

const (
	CustomerTierStandard CustomerTier = "STANDARD"
	CustomerTierSilver   CustomerTier = "SILVER"
	CustomerTierGold     CustomerTier = "GOLD"
	CustomerTierPlatinum CustomerTier = "PLATINUM"
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
