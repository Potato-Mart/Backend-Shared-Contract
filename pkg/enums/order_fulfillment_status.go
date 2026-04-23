package enums

type FulfillmentStatus string

const (
	FulfillmentStatusUnfulfilled    FulfillmentStatus = "UNFULFILLED"
	FulfillmentStatusPickingPrinted FulfillmentStatus = "PICKING_PRINTED"
	FulfillmentStatusPacking        FulfillmentStatus = "PACKING"
	FulfillmentStatusPacked         FulfillmentStatus = "PACKED"
	FulfillmentStatusPartial        FulfillmentStatus = "PARTIAL"
	FulfillmentStatusFulfilled      FulfillmentStatus = "FULFILLED"
)

// IsValid reports whether p is a known FulfillmentStatus.
func (p FulfillmentStatus) IsValid() bool {
	switch p {
	case FulfillmentStatusUnfulfilled, FulfillmentStatusPickingPrinted, FulfillmentStatusPacking,
		FulfillmentStatusPacked, FulfillmentStatusPartial, FulfillmentStatusFulfilled:
		return true
	}
	return false
}

func (p FulfillmentStatus) String() string { return string(p) }
