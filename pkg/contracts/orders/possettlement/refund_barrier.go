package possettlement

// RefundBarrier identifies a complete order-scoped prefix of Payments-owned
// refund admissions. Pricing acknowledges a prefix only after every admitted
// effect has been applied or authoritatively resolved. A count, timestamp or
// empty query does not establish this evidence. Sequence zero is the initial
// prefix, not proof that a legacy order has registered all historical writers.
type RefundBarrier struct {
	OrderNumber      string `json:"order_number"`
	Sequence         int64  `json:"sequence"`
	PrefixSHA256     string `json:"prefix_sha256"`
	CoverageRevision string `json:"coverage_revision"`
}
