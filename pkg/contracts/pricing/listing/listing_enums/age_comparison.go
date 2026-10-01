package listing_enums

// AgeComparison identifies the strict blocking side of a retail age threshold.
// Completed age equal to the threshold is allowed for both values. Calendar
// calculation and eligibility decisions belong to services, not this enum.
type AgeComparison string

const (
	// AgeComparisonBelow blocks completed age strictly below AgeYears.
	AgeComparisonBelow AgeComparison = "below"
	// AgeComparisonAbove blocks completed age strictly above AgeYears.
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
