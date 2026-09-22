package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/courier"
)

// DeliverySchedule is a cart-free, revisioned view of delivery windows for an
// area. DeliveryCompany applies to every slot when set; a mixed-company view
// leaves it absent and sets each slot's company. ExpiresAt bounds new selection,
// not fulfilment of an already accepted order. Carrier retains its legacy
// display-name meaning.
type DeliverySchedule struct {
	Availability      string                      `json:"availability"`
	UnavailableReason string                      `json:"unavailable_reason,omitempty"`
	Revision          int64                       `json:"revision"`
	Timezone          string                      `json:"timezone"`
	Carrier           string                      `json:"carrier,omitempty"`
	AreaRate          *DeliveryAreaRate           `json:"area_rate,omitempty"`
	DateGroups        []DeliveryDateGroup         `json:"date_groups"`
	DeliveryCompany   *courier.DeliveryCompanyRef `json:"delivery_company,omitempty"`
	ExpiresAt         *time.Time                  `json:"expires_at,omitempty"`
}
