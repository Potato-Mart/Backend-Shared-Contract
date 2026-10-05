package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/supply/catalogue/classification"

	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pricing/membership/membership_enums"
)

// MembershipTier defines qualification and benefits within a country's tier
// group. Each country has one group and currency, identified by CountryCode
// and MinQualifyingSpend.Currency. TierKey remains globally unique. Spending
// points does not affect tier qualification.
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
	// MarketCodes selects membership availability and staff geographic access.
	// Pricing must validate a nonempty, unique list of markets in CountryCode using
	// the group's currency, shared by every tier in the group. Empty or missing
	// lists do not imply global availability. Pricing owns validation, role
	// permissions and access enforcement.
	MarketCodes []string              `json:"market_codes"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	audit.AuditFields
}
