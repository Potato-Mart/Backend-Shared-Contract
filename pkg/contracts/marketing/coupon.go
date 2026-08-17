package marketing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/commerce/commerce_enums"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/marketing/marketing_enums"
)

// Coupon is the public coupon definition identified by CouponCode. It keeps
// customer-facing content and eligibility shape together without exposing any
// customer assignment, redemption, or provider data.
type Coupon struct {
	CouponCode       string                       `json:"coupon_code"`
	CouponName       []localization.LocalizedName `json:"coupon_name"`
	CouponCover      *security.ObjectMedia        `json:"coupon_cover,omitempty"`
	CouponDetail     CouponDetail                 `json:"coupon_detail"`
	CouponScope      CouponScope                  `json:"coupon_scope"`
	CouponStatus     CouponStatus                 `json:"coupon_status"`
	CouponPosition   CouponPosition               `json:"coupon_position"`
	CouponConditions CouponConditions             `json:"coupon_conditions"`

	audit.AuditFields
}

// CouponDetail contains public coupon presentation and benefit type.
type CouponDetail struct {
	Description []localization.LocalizedDescription `json:"description,omitempty"`
	CouponType  marketing_enums.CouponType          `json:"coupon_type"`
}

// CouponScope identifies the target family and detailed benefit applied to it.
type CouponScope struct {
	ScopeType   marketing_enums.CouponScopeType `json:"scope_type"`
	Targets     []ScopeTarget                   `json:"targets,omitempty"`
	ScopeDetail ScopeDetail                     `json:"scope_detail"`
}

// CouponStatus combines the coupon lifecycle state with its active window.
type CouponStatus struct {
	Status      marketing_enums.CouponStatus `json:"status"`
	Dismissible bool                         `json:"dismissible"`
	StartsAt    *time.Time                   `json:"starts_at,omitempty"`
	EndsAt      *time.Time                   `json:"ends_at,omitempty"`
}

// CouponPosition carries the commercial geography and authored schedule zone.
type CouponPosition struct {
	GeographicScope  geography.GeographicScope `json:"geographic_scope"`
	ScheduleTimezone string                    `json:"schedule_timezone"`
}

// CouponConditions carries public application controls. Nil UsageLimit means
// the aggregate does not declare a global usage cap.
type CouponConditions struct {
	UsageLimit *int                       `json:"usage_limit,omitempty"`
	Stackable  bool                       `json:"stackable"`
	Channels   []commerce_enums.OrderType `json:"channels,omitempty"`
}
