package listing_enums

// AgeComparison identifies the allowed group for a retail age threshold.
// Completed age equal to the threshold is allowed only for Above. Calendar
// calculation and eligibility decisions belong to services, not this enum.
type AgeComparison string

const (
	// AgeComparisonBelow permits completed age strictly below AgeYears.
	AgeComparisonBelow AgeComparison = "below"
	// AgeComparisonAbove permits completed age greater than or equal to AgeYears.
	AgeComparisonAbove AgeComparison = "above"
)

func (c AgeComparison) IsValid() bool {
	switch c {
	case AgeComparisonBelow, AgeComparisonAbove:
		return true
	}
	return false
}

func (c AgeComparison) String() string { return string(c) }
