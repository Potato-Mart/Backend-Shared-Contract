package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/supply/catalogue/classification"

	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/pricing/membership/membership_enums"
)

// MembershipTier defines the qualification and benefit rules for the global
// membership programme. Spending points does not affect tier qualification.
type MembershipTier struct {
	TierKey string `json:"tier_key"`
	// Label is extensible locale data. The service requires en, zh-TW and zh-CN.
	Label               []localization.LocalizedText          `json:"label"`
	QualificationMetric membership_enums.MembershipTierMetric `json:"qualification_metric"`
	MinQualifyingSpend  money.Money                           `json:"min_qualifying_spend"`
	PointMultiplier     float64                               `json:"point_multiplier"`
	Metadata            []MembershipTierMetadataEntry         `json:"metadata,omitempty"`
	TierCard            *classification.ObjectMediaRef        `json:"tier_card,omitempty"`
	// Benefits is the typed, localized benefit list.
	Benefits []TierBenefit `json:"benefits,omitempty"`
	IsSystem bool          `json:"is_system"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	audit.AuditFields
}
