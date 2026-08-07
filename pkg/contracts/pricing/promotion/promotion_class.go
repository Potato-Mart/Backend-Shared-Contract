package promotion

// PromotionClass separates standing price promotions (常態特價) from
// time-boxed special campaigns (特殊活動). It is orthogonal to
// PromotionType, which classifies the engine mechanics (auto_discount,
// bogo, …): a special campaign can still be an auto_discount under the
// hood.
//
//	normal_promotion -> 常態特價: recurring/standing promotion
//	special_campaign -> 特殊活動: event promotion for a specific period;
//	                    always overrides a normal_promotion that targets
//	                    the same product or category tag in the same window
type PromotionClass string

const (
	PromotionClassNormal          PromotionClass = "normal_promotion"
	PromotionClassSpecialCampaign PromotionClass = "special_campaign"
)

func (p PromotionClass) IsValid() bool {
	switch p {
	case PromotionClassNormal, PromotionClassSpecialCampaign:
		return true
	}
	return false
}

func (p PromotionClass) String() string { return string(p) }
