package inventory

// UnallocatedNoLotPhysicalEvidence identifies physical stock held without a
// lot or return-receiving identity. SourceBucketID is the actual original
// allocation bucket and must equal the allocation bucket ID. For a staged
// movement the owner verifies it against the staging source bucket. Quantity is the enclosing allocation or movement's physical
// composition. This evidence does not invent a receipt, unit or cost.
type UnallocatedNoLotPhysicalEvidence struct {
	SourceBucketID string `json:"source_bucket_id"`
}
