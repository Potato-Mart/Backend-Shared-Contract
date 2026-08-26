// Package promotion defines reusable promotion terms, scopes, relations, and
// frozen applications. It contains contract models only; pricing services own
// validation and evaluation.
package promotion

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/promotion/promotion_enums"
)

// Promotion is a revisioned, open-ended promotion definition. Kind is an open
// string so new mechanics can be introduced without a shared-contract enum
// release. Relations express qualifier-to-target mechanics within one sale
// order, including simple discounts.
type Promotion struct {
	ID        string                          `json:"id"`
	SeriesKey string                          `json:"series_key"`
	Kind      string                          `json:"kind"`
	Status    promotion_enums.PromotionStatus `json:"status"`
	Revision  int64                           `json:"revision"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	Content   PromotionContent        `json:"content"`
	Period    PromotionPeriod         `json:"period"`
	Scope     PromotionScope          `json:"scope"`
	Relations []PromotionRelation     `json:"relations,omitempty"`
	Controls  PromotionControls       `json:"controls"`
	Source    *PromotionSource        `json:"source,omitempty"`
	History   []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}
