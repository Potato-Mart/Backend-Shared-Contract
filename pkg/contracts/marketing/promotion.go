package marketing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/commerce/commerce_enums"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/marketing/marketing_enums"
)

// Promotion is the public, code-addressed marketing definition for an
// automatic pricing benefit. Pricing services remain authoritative for
// evaluation and reservation decisions.
type Promotion struct {
	PromotionCode       string                       `json:"promotion_code"`
	PromotionName       []localization.LocalizedName `json:"promotion_name"`
	PromotionCover      *security.ObjectMedia        `json:"promotion_cover,omitempty"`
	PromotionDetail     PromotionDetail              `json:"promotion_detail"`
	PromotionScope      PromotionScope               `json:"promotion_scope"`
	PromotionStatus     PromotionStatus              `json:"promotion_status"`
	PromotionPosition   PromotionPosition            `json:"promotion_position"`
	PromotionConditions PromotionConditions          `json:"promotion_conditions"`
	ScopeRelations      ScopeRelations               `json:"scope_relations"`

	audit.AuditFields
}

// PromotionDetail contains public promotion presentation and mechanic type.
type PromotionDetail struct {
	Description   []localization.LocalizedDescription `json:"description,omitempty"`
	PromotionType marketing_enums.PromotionType       `json:"promotion_type"`
}

// PromotionScope identifies the target family and detailed benefit applied to
// that scope.
type PromotionScope struct {
	ScopeType   marketing_enums.PromotionScopeType `json:"scope_type"`
	Targets     []ScopeTarget                      `json:"targets,omitempty"`
	ScopeDetail ScopeDetail                        `json:"scope_detail"`
}

// PromotionStatus combines the promotion lifecycle state with its active
// window.
type PromotionStatus struct {
	Status   marketing_enums.PromotionStatus `json:"status"`
	StartsAt *time.Time                      `json:"starts_at,omitempty"`
	EndsAt   *time.Time                      `json:"ends_at,omitempty"`
}

// PromotionPosition carries the commercial geography and authored schedule
// zone for a promotion.
type PromotionPosition struct {
	GeographicScope  geography.GeographicScope `json:"geographic_scope"`
	ScheduleTimezone string                    `json:"schedule_timezone"`
}

// PromotionConditions carries public application controls. Nil caps mean the
// aggregate does not declare that cap.
type PromotionConditions struct {
	UsageLimit       *int                       `json:"usage_limit,omitempty"`
	PerCustomerLimit *int                       `json:"per_customer_limit,omitempty"`
	Stackable        bool                       `json:"stackable"`
	Channels         []commerce_enums.OrderType `json:"channels,omitempty"`
}

// ScopeRelations captures the public relation inputs required by the finite
// promotion mechanics. It intentionally contains no resolved cart, customer,
// or provider state.
type ScopeRelations struct {
	Targets                     []ScopeTarget `json:"targets,omitempty"`
	RequiredProducts            []ScopeTarget `json:"required_products,omitempty"`
	Tiers                       []ScopeDetail `json:"tiers,omitempty"`
	MixTargets                  []ScopeTarget `json:"mix_targets,omitempty"`
	BOGOGetQuantity             *int          `json:"bogo_get_quantity,omitempty"`
	AddOnMaxQuantity            *int          `json:"add_on_max_quantity,omitempty"`
	PointsMultiplierBasisPoints *int64        `json:"points_multiplier_basis_points,omitempty"`
}
