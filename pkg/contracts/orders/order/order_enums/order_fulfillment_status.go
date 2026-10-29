package order_enums

type FulfillmentStatus string

const (
	FulfillmentStatusUnfulfilled       FulfillmentStatus = "unfulfilled"
	FulfillmentStatusPickingInProgress FulfillmentStatus = "picking_in_progress"
	FulfillmentStatusPacking           FulfillmentStatus = "packing"
	FulfillmentStatusPacked            FulfillmentStatus = "packed"
	FulfillmentStatusFulfilled         FulfillmentStatus = "fulfilled"
	// FulfillmentStatusCancelled means no further fulfilment work is scheduled.
	// It does not erase recorded picking, packing, shipment, or item facts and is
	// distinct from FulfillmentStatusFulfilled.
	FulfillmentStatusCancelled FulfillmentStatus = "cancelled"
)

// IsValid reports whether p is a known FulfillmentStatus.
func (p FulfillmentStatus) IsValid() bool {
	switch p {
	case FulfillmentStatusUnfulfilled, FulfillmentStatusPickingInProgress, FulfillmentStatusPacking,
		FulfillmentStatusPacked, FulfillmentStatusFulfilled,
		FulfillmentStatusCancelled:
		return true
	}
	return false
}

func (p FulfillmentStatus) String() string { return string(p) }
