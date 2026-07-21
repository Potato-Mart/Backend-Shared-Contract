package salesenum

// CustomerOrderBucket groups sales-order statuses into the customer-facing
// order-history buckets. The status-set classification behind each bucket is
// owned by the sales-order service; unknown statuses classify as current.
type CustomerOrderBucket string

const (
	CustomerOrderBucketCurrent   CustomerOrderBucket = "current"
	CustomerOrderBucketCompleted CustomerOrderBucket = "completed"
	CustomerOrderBucketCancelled CustomerOrderBucket = "cancelled"
	CustomerOrderBucketRefunded  CustomerOrderBucket = "refunded"
)

func (b CustomerOrderBucket) IsValid() bool {
	switch b {
	case CustomerOrderBucketCurrent, CustomerOrderBucketCompleted,
		CustomerOrderBucketCancelled, CustomerOrderBucketRefunded:
		return true
	default:
		return false
	}
}

func (b CustomerOrderBucket) String() string { return string(b) }
