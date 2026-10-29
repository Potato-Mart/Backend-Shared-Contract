package fulfilment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/orders/order/order_enums"
)

// PackingRecord describes physical packing for a specific order composition.
// Requested and packed quantities are represented in each line's package
// composition. The order owns its current fulfilment status and history.
type PackingRecord struct {
	ID                    string                        `json:"id"`
	OrderNumber           string                        `json:"order_number"`
	FulfillmentGeneration int64                         `json:"fulfillment_generation"`
	Revision              int64                         `json:"revision"`
	Status                order_enums.FulfillmentStatus `json:"status"`
	Operator              string                        `json:"operator,omitempty"`
	Lines                 []PackingLine                 `json:"lines"`
	Containers            []OutboundContainerPlan       `json:"containers,omitempty"`
	Damages               []PackingDamage               `json:"damages,omitempty"`
	Discrepancies         []PackingDiscrepancy          `json:"discrepancies,omitempty"`
	StartedAt             *time.Time                    `json:"started_at,omitempty"`
	PackedAt              *time.Time                    `json:"packed_at,omitempty"`
	ClosedAt              *time.Time                    `json:"closed_at,omitempty"`
	CreatedAt             time.Time                     `json:"created_at"`
	UpdatedAt             time.Time                     `json:"updated_at"`
}
