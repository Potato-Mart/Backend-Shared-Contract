package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
)

// RefundRequestClosedEvent is emitted on the refund-events topic, with event
// type refund.request_closed v1, when a pending refund request is closed
// without any refund being executed. It never announces money movement.
//
// Status is rejected (staff turned the request down) or cancelled (the
// request was withdrawn). Station is true when the request was raised by the
// packing station for a recorded packing shortage; DiscrepancyID then names
// that shortage when known. ItemIDs are the order items the request covered.
// OrderNumber is opaque, including custom external order numbers; consumers
// must not classify it by prefix. ClosedBy is the actor who closed the request.
//
// MarketCode and CountryCode are the denormalized geography the event belongs
// to. Empty values provide no geographic evidence; a consumer that persists a
// geographically scoped record must fail closed rather than defaulting them.
type RefundRequestClosedEvent struct {
	RefundRequestID string                `json:"refund_request_id"`
	OrderNumber     string                `json:"order_number"`
	Status          string                `json:"status"`
	Station         bool                  `json:"station"`
	DiscrepancyID   string                `json:"discrepancy_id,omitempty"`
	ItemIDs         []string              `json:"item_ids"`
	ClosedBy        string                `json:"closed_by"`
	ClosedAt        time.Time             `json:"closed_at"`
	MarketCode      string                `json:"market_code,omitempty"`
	CountryCode     geography.CountryCode `json:"country_code,omitempty"`
}
