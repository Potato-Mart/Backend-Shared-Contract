package courier

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal"

// CarrierBookingTimeWindow is a wall-clock delivery window edit sent to a
// carrier. FromTime and ToTime are interpreted on the booking's delivery date
// in the delivery (recipient's) timezone. An edit may be one-sided: a nil side
// was not edited and keeps the value originally sent to the carrier. When both
// sides are present, FromTime is before ToTime.
type CarrierBookingTimeWindow struct {
	FromTime *temporal.TimeOfDay `json:"from_time,omitempty"`
	ToTime   *temporal.TimeOfDay `json:"to_time,omitempty"`
}
