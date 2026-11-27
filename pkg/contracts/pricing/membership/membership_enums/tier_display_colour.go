package membership_enums

// TierDisplayColour is the colour token clients use to draw a membership tier
// chip. It is a token, not a colour value: each client maps it to its own
// palette, and the contract carries no hex or RGB code.
type TierDisplayColour string

const (
	TierDisplayColourBronze   TierDisplayColour = "bronze"
	TierDisplayColourSilver   TierDisplayColour = "silver"
	TierDisplayColourGold     TierDisplayColour = "gold"
	TierDisplayColourPlatinum TierDisplayColour = "platinum"
	TierDisplayColourBlack    TierDisplayColour = "black"
	TierDisplayColourDiamond  TierDisplayColour = "diamond"
)

// IsValid reports whether c is a known TierDisplayColour.
func (c TierDisplayColour) IsValid() bool {
	switch c {
	case TierDisplayColourBronze, TierDisplayColourSilver, TierDisplayColourGold, TierDisplayColourPlatinum, TierDisplayColourBlack, TierDisplayColourDiamond:
		return true
	}
	return false
}

func (c TierDisplayColour) String() string { return string(c) }
