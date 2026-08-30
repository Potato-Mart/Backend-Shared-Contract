package membership

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/wallet_enums"
)

// Reward is a catalog item that a member can redeem with membership points.
//
// Type selects the fulfilment family and Benefit carries that family's
// configuration, so a new benefit family extends RewardBenefit instead of
// widening this struct.
type Reward struct {
	ID                     string                              `json:"id"`
	Code                   string                              `json:"code"`
	Names                  []localization.LocalizedName        `json:"names"`
	Descriptions           []localization.LocalizedDescription `json:"descriptions,omitempty"`
	Type                   wallet_enums.RewardType             `json:"type"`
	PointsCost             int                                 `json:"points_cost"`
	Benefit                RewardBenefit                       `json:"benefit"`
	StartsAt               *time.Time                          `json:"starts_at,omitempty"`
	ExpiresAt              *time.Time                          `json:"expires_at,omitempty"`
	IsActive               bool                                `json:"is_active"`
	UsageLimit             int                                 `json:"usage_limit,omitempty"`
	UsedCount              int                                 `json:"used_count"`
	PerMemberLimit         int                                 `json:"per_member_limit,omitempty"`
	MinimumTierKey         string                              `json:"minimum_tier_key,omitempty"`
	TriggerTierKey         string                              `json:"trigger_tier_key,omitempty"`
	IssueOnTierAchievement bool                                `json:"issue_on_tier_achievement,omitempty"`
	Metadata               metadata.Metadata                   `json:"metadata,omitempty"`
	History                []security.HistoryEntry             `json:"history,omitempty"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	audit.AuditFields
}
