package warehouse_enums

// DamageSaleTier is the approved reduced-sale condition tier recorded against
// damaged inventory. The tier names the assessed damage band; the commercial
// multiplier that Pricing applies to it is a fixed tier mapping owned by
// Pricing and is deliberately not one minus the damage percentage.
type DamageSaleTier string

const (
	// DamageSaleTier30 is the 30% damage band.
	DamageSaleTier30 DamageSaleTier = "tier_30"
	// DamageSaleTier50 is the 50% damage band.
	DamageSaleTier50 DamageSaleTier = "tier_50"
	// DamageSaleTier80 is the 80% damage band.
	DamageSaleTier80 DamageSaleTier = "tier_80"
)

// IsValid reports whether t is a known DamageSaleTier.
func (t DamageSaleTier) IsValid() bool {
	switch t {
	case DamageSaleTier30, DamageSaleTier50, DamageSaleTier80:
		return true
	}
	return false
}

func (t DamageSaleTier) String() string { return string(t) }
