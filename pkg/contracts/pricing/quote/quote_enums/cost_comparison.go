package quote_enums

type CostComparison string

const (
	CostComparisonBelowCost     CostComparison = "below_cost"
	CostComparisonAtOrAboveCost CostComparison = "at_or_above_cost"
	CostComparisonUnavailable   CostComparison = "unavailable"
)

func (r CostComparison) IsValid() bool {
	switch r {
	case CostComparisonBelowCost, CostComparisonAtOrAboveCost, CostComparisonUnavailable:
		return true
	}
	return false
}
func (r CostComparison) String() string { return string(r) }
