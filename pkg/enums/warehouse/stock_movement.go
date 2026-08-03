package warehouseenum

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
		StockMovementTypeScrap:
		return true
	}
	return false
}

func (t StockMovementType) String() string { return string(t) }
