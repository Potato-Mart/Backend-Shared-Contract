package group_order

import "time"

// GroupOrderFulfilmentPlan is the consolidated parent-owned inventory and
// participant-share snapshot for one group order.
type GroupOrderFulfilmentPlan struct {
	ID                 string                       `json:"id"`
	GroupOrderCode     string                       `json:"group_order_code"`
	ParentOrderNumber  string                       `json:"parent_order_number"`
	ParentFulfilmentID string                       `json:"parent_fulfilment_id"`
	AggregateLines     []GroupOrderAggregateLine    `json:"aggregate_lines"`
	ParticipantShares  []GroupOrderParticipantShare `json:"participant_shares"`
	Revision           int64                        `json:"revision"`
	Timezone           string                       `json:"timezone"`
	CapturedAt         time.Time                    `json:"captured_at"`
}
