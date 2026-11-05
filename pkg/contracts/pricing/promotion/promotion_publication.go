package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/marketing/audience"
)

// PromotionPublication is Pricing's canonical, customer-safe publication state
// for one promotion, observed at ObservedAt. Consumers such as Notification
// read it by promotion ID to confirm a PromotionChangedEvent is still current
// before fanning out.
//
// Revision equals Promotion.Revision and the revision carried by the
// promotion.changed event for the same publication. Published is the
// promotion's lifecycle publication state after campaign applicability is
// applied, so a promotion whose linked campaign is not applicable reads as
// unpublished. StartsAt and EndsAt are the effective UTC window, start
// inclusive and end exclusive ([StartsAt, EndsAt)); nil means unbounded on that
// side. Timezone is the promotion's scheduling timezone for display only.
//
// The record never carries promotion rules, terms, usage counters, or any
// customer data.
type PromotionPublication struct {
	PromotionID  string             `json:"promotion_id"`
	Published    bool               `json:"published"`
	Revision     int64              `json:"revision"`
	MarketCode   string             `json:"market_code,omitempty"`
	CampaignCode string             `json:"campaign_code,omitempty"`
	Audience     *audience.Audience `json:"audience,omitempty"`
	StartsAt     *time.Time         `json:"starts_at,omitempty"`
	EndsAt       *time.Time         `json:"ends_at,omitempty"`
	Timezone     string             `json:"timezone,omitempty"`
	ObservedAt   time.Time          `json:"observed_at"`
}
