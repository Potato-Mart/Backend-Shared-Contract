package membership

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
)

// SubscriptionPlan defines a recurring purchase option available through the
// membership domain. It remains separate from points accounting.
type SubscriptionPlan struct {
	ID              string      `json:"id"`
	SKUCode         string      `json:"sku_code"`
	UnitPrice       money.Money `json:"unit_price"`
	FrequencyDays   int         `json:"frequency_days"`
	FrequencyLabel  string      `json:"frequency_label"`
	DiscountPercent float64     `json:"discount_percent"`
	MinCycles       int         `json:"min_cycles"`
	IsActive        bool        `json:"is_active"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	audit.AuditFields
}
