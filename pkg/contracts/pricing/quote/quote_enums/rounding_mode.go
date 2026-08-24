package quote_enums

type RoundingMode string

const (
	RoundingModeHalfUp           RoundingMode = "half_up"
	RoundingModeHalfCentUp       RoundingMode = "half_cent_up"
	RoundingModeCharmNineUp      RoundingMode = "charm_nine_up"
	RoundingModeCashIncrement    RoundingMode = "cash_increment"
	RoundingModeLargestRemainder RoundingMode = "largest_remainder"
)

func (m RoundingMode) IsValid() bool {
	switch m {
	case RoundingModeHalfUp, RoundingModeHalfCentUp, RoundingModeCharmNineUp, RoundingModeCashIncrement, RoundingModeLargestRemainder:
		return true
	}
	return false
}
func (m RoundingMode) String() string { return string(m) }
