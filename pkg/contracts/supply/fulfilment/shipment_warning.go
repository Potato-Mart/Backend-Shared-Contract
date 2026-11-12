package fulfilment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// ShipmentWarning is one non-blocking condition shown on an outbound
// shipment. Message is staff-facing text. ManualActionID, CarrierCode and
// CarrierReference identify the open courier.CarrierManualAction and the
// carrier booking it concerns, for a carrier_cancellation_required warning.
// CreatedAt is when the condition was first recorded.
type ShipmentWarning struct {
	Code             warehouse_enums.ShipmentWarningCode `json:"code"`
	Message          string                              `json:"message,omitempty"`
	ManualActionID   string                              `json:"manual_action_id,omitempty"`
	CarrierCode      string                              `json:"carrier_code,omitempty"`
	CarrierReference string                              `json:"carrier_reference,omitempty"`
	CreatedAt        *time.Time                          `json:"created_at,omitempty"`
}
