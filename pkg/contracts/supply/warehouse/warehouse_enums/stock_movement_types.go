package warehouse_enums

// StockMovementType identifies the physical inventory change represented by a
// stock movement.
type StockMovementType string

const (
	StockMovementTypeReceipt           StockMovementType = "RECEIPT"
	StockMovementTypeTransfer          StockMovementType = "TRANSFER"
	StockMovementTypeStage             StockMovementType = "STAGE"
	StockMovementTypeUnstage           StockMovementType = "UNSTAGE"
	StockMovementTypePackageConversion StockMovementType = "PACKAGE_CONVERSION"
	StockMovementTypeSaleCommit        StockMovementType = "SALE_COMMIT"
	StockMovementTypeReturn            StockMovementType = "RETURN"
	StockMovementTypeQualityHold       StockMovementType = "QUALITY_HOLD"
	StockMovementTypeQualityRelease    StockMovementType = "QUALITY_RELEASE"
	StockMovementTypeQualityReject     StockMovementType = "QUALITY_REJECT"
	StockMovementTypeAdjustment        StockMovementType = "ADJUSTMENT"
	StockMovementTypeReturnToVendor    StockMovementType = "RETURN_TO_VENDOR"
	StockMovementTypeScrap             StockMovementType = "SCRAP"
	// StockMovementTypeLotReassignment moves base units from one lot's
	// bucket to another lot's bucket at the same location and package grain
	// in a single movement. It is value-neutral: the source lot's unit cost
	// carries across to the destination lot. On a StockMovement, lot_id is
	// the source lot; the destination lot is identified through
	// destination_bucket_id. No additional movement fields are used.
	StockMovementTypeLotReassignment StockMovementType = "LOT_REASSIGNMENT"
	// StockMovementTypeReceiptReversal removes base units booked by a
	// confirmed purchase receipt that is being reversed. On a StockMovement,
	// source_bucket_id is the bucket the units leave, purchase_receipt_id is
	// the reversed receipt, and there is no destination.
	StockMovementTypeReceiptReversal StockMovementType = "RECEIPT_REVERSAL"
)

// IsValid reports whether t is a known StockMovementType.
func (t StockMovementType) IsValid() bool {
	switch t {
	case StockMovementTypeReceipt, StockMovementTypeTransfer,
		StockMovementTypeStage, StockMovementTypeUnstage,
		StockMovementTypePackageConversion, StockMovementTypeSaleCommit,
		StockMovementTypeReturn, StockMovementTypeQualityHold,
		StockMovementTypeQualityRelease, StockMovementTypeQualityReject,
		StockMovementTypeAdjustment, StockMovementTypeReturnToVendor,
		StockMovementTypeScrap, StockMovementTypeLotReassignment,
		StockMovementTypeReceiptReversal:
		return true
	}
	return false
}

func (t StockMovementType) String() string { return string(t) }
