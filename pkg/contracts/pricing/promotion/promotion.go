// Package promotion defines reusable promotion terms, scopes, relations, and
// frozen applications. It contains contract models only; pricing services own
// validation and evaluation.
package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/promotion/promotion_enums"
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

	Content   PromotionContent        `json:"content"`
	Period    PromotionPeriod         `json:"period"`
	Scope     PromotionScope          `json:"scope"`
	Relations []PromotionRelation     `json:"relations,omitempty"`
	Controls  PromotionControls       `json:"controls"`
	Source    *PromotionSource        `json:"source,omitempty"`
	History   []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}

// PromotionContent contains only approved localized customer-facing copy.
// Names and descriptions support authoring/display, while display and receipt
// messages are the strings safe to show during checkout and on receipts.
type PromotionContent struct {
	Names           []localization.LocalizedName        `json:"names,omitempty"`
	Descriptions    []localization.LocalizedDescription `json:"descriptions,omitempty"`
	DisplayMessages []localization.LocalizedText        `json:"display_messages,omitempty"`
	ReceiptMessages []localization.LocalizedText        `json:"receipt_messages,omitempty"`
}

// PromotionPeriod is the scheduled period for a promotion. A nil EndsAt means
// the promotion has no scheduled end; lifecycle status independently controls
// whether it is active.
type PromotionPeriod struct {
	StartsAt *time.Time `json:"starts_at,omitempty"`
	EndsAt   *time.Time `json:"ends_at,omitempty"`
	Timezone string     `json:"timezone"`
}

// PromotionSource records an optional upstream source using open identifiers.
// It intentionally does not constrain integration-specific source kinds.
type PromotionSource struct {
	Kind string `json:"kind,omitempty"`
	Ref  string `json:"ref,omitempty"`
}
