package ads_enums

// AdBidStrategy names how a platform is asked to bid, normalised across Meta
// and Google. Which strategies a level accepts, and which target a strategy
// requires, stay platform rules enforced by the ads module.
type AdBidStrategy string

const (
	// AdBidStrategyHighestVolume is the Meta highest-volume goal: spend the
	// budget on as many results as possible, with no cost target.
	AdBidStrategyHighestVolume AdBidStrategy = "highest_volume"
	// AdBidStrategyCostPerActionTarget is the Meta cost-per-result goal:
	// aim for an average cost per purchase.
	AdBidStrategyCostPerActionTarget AdBidStrategy = "cost_per_action_target"
	// AdBidStrategyBidCap is the Meta manual maximum bid per auction.
	AdBidStrategyBidCap AdBidStrategy = "bid_cap"
	// AdBidStrategyMinimumROAS is the Meta minimum return-on-ad-spend goal.
	AdBidStrategyMinimumROAS AdBidStrategy = "minimum_roas"
	// AdBidStrategyMaximiseConversions is the Google maximise-conversions
	// strategy, optionally with a target cost per action.
	AdBidStrategyMaximiseConversions AdBidStrategy = "maximise_conversions"
	// AdBidStrategyMaximiseConversionValue is the Google
	// maximise-conversion-value strategy, optionally with a target return on
	// ad spend.
	AdBidStrategyMaximiseConversionValue AdBidStrategy = "maximise_conversion_value"
)

// IsValid reports whether s is a known AdBidStrategy.
func (s AdBidStrategy) IsValid() bool {
	switch s {
	case AdBidStrategyHighestVolume, AdBidStrategyCostPerActionTarget,
		AdBidStrategyBidCap, AdBidStrategyMinimumROAS,
		AdBidStrategyMaximiseConversions, AdBidStrategyMaximiseConversionValue:
		return true
	}
	return false
}

func (s AdBidStrategy) String() string { return string(s) }
