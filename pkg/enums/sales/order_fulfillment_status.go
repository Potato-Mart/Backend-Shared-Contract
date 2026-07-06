package salesenum

type FulfillmentStatus string

const (
	FulfillmentStatusUnfulfilled    FulfillmentStatus = "unfulfilled"
	FulfillmentStatusPickingPrinted FulfillmentStatus = "picking_printed"
	FulfillmentStatusPacking        FulfillmentStatus = "packing"
	FulfillmentStatusPacked         FulfillmentStatus = "packed"
	FulfillmentStatusPartial        FulfillmentStatus = "partial"
	FulfillmentStatusFulfilled      FulfillmentStatus = "fulfilled"
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
