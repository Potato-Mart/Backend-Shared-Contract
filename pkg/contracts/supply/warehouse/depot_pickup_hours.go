package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal/temporal_enums"
)

// DepotPickupHours is one weekday of a depot's customer pickup schedule.
// Opens and Closes are "HH:MM" wall-clock times in the depot Timezone. A
// Closed day carries no times. Supply validates the schedule.
type DepotPickupHours struct {
	DayOfWeek temporal_enums.DayOfWeek `json:"day_of_week"`
	Opens     temporal.TimeOfDay       `json:"opens,omitempty"`
	Closes    temporal.TimeOfDay       `json:"closes,omitempty"`
	Closed    bool                     `json:"closed,omitempty"`
}
