package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/shipping/shipping_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
)

// DeliverySelection freezes the server-validated delivery choice accepted for
// an order and carried into fulfilment. Services preserve it across retries;
// it is not booking acceptance or evidence of physical dispatch. StartAt and
// EndAt are UTC instants; Date is the local delivery date in Timezone (IANA).
// ScheduleCode identifies the configured window or provider schedule; SlotID
// identifies its dated offer. Revisions refer to the accepted configuration.
// TimeBand is the accepted slot's band. TimeAcknowledgedAt is the UTC instant
// the customer confirmed a late-night or early-morning window; it is empty
// when no acknowledgement was required or recorded.
type DeliverySelection struct {
	DeliveryCompany    courier.DeliveryCompanyRef       `json:"delivery_company"`
	ScheduleCode       string                           `json:"schedule_code"`
	ScheduleRevision   int64                            `json:"schedule_revision"`
	SlotID             string                           `json:"slot_id"`
	Date               temporal.Date                    `json:"date"`
	StartAt            time.Time                        `json:"start_at"`
	EndAt              time.Time                        `json:"end_at"`
	Timezone           string                           `json:"timezone"`
	Source             courier_enums.DeliverySlotSource `json:"source"`
	TimeBand           shipping_enums.DeliveryTimeBand  `json:"time_band,omitempty"`
	TimeAcknowledgedAt *time.Time                       `json:"time_acknowledged_at,omitempty"`
}
