package ads

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdPerformanceRow is one platform object at one level, with the spend the
// platform reports and the paid orders we attributed to it.
//
// Revision is the opaque object revision a delivery change must quote as its
// expected revision; it is carried exactly, never parsed or compared as a
// number. Reach is a Meta figure only and is absent for Google. Objective is
// the platform's own objective code, kept open because platforms add to it.
// PlatformConversions and PlatformConversionValue are the platform's own
// counts, while Orders, NewCustomers and Revenue are ours; the two are
// measured differently and are not expected to match. Buckets are present only
// when a daily breakdown was asked for. Money is a list per currency so that
// accounts in different currencies are never summed.
type AdPerformanceRow struct {
	Platform                ads_enums.AdPlatform       `json:"platform,omitempty"`
	Level                   ads_enums.AdLevel          `json:"level,omitempty"`
	ID                      string                     `json:"id,omitempty"`
	Name                    string                     `json:"name,omitempty"`
	ParentID                string                     `json:"parent_id,omitempty"`
	Status                  ads_enums.AdEntityStatus   `json:"status,omitempty"`
	ReviewStatus            ads_enums.AdReviewStatus   `json:"review_status,omitempty"`
	Objective               string                     `json:"objective,omitempty"`
	Revision                string                     `json:"revision,omitempty"`
	SourcePlatform          ads_enums.AdSourcePlatform `json:"source_platform,omitempty"`
	Budget                  *AdBudget                  `json:"budget,omitempty"`
	Bidding                 *AdBidding                 `json:"bid_strategy,omitempty"`
	Spend                   []money.Money              `json:"spend,omitempty"`
	Impressions             int64                      `json:"impressions,omitempty"`
	Reach                   int64                      `json:"reach,omitempty"`
	Clicks                  int64                      `json:"clicks,omitempty"`
	Installs                int64                      `json:"installs,omitempty"`
	PlatformConversions     int64                      `json:"platform_conversions,omitempty"`
	PlatformConversionValue []money.Money              `json:"platform_conversion_value,omitempty"`
	Orders                  int64                      `json:"orders,omitempty"`
	NewCustomers            int64                      `json:"new_customers,omitempty"`
	Revenue                 []money.Money              `json:"revenue,omitempty"`
	Buckets                 []AdPerformanceBucket      `json:"buckets,omitempty"`
}
