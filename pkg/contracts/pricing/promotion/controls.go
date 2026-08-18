package promotion

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/commerce/commerce_enums"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"
)

// PromotionControls contains reusable application limits and eligibility
// controls. Nil caps mean unlimited. GeographicScope and Channels constrain
// where the promotion is eligible without defining its mechanics.
type PromotionControls struct {
	Priority                       int                        `json:"priority,omitempty"`
	Stackable                      bool                       `json:"stackable,omitempty"`
	MaximumApplications            *int64                     `json:"maximum_applications,omitempty"`
	MaximumApplicationsPerCustomer *int64                     `json:"maximum_applications_per_customer,omitempty"`
	Channels                       []commerce_enums.OrderType `json:"channels,omitempty"`
	GeographicScope                geography.GeographicScope  `json:"geographic_scope"`
}
