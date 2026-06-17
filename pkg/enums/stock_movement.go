package enums

// StockMovementType identifies why a product's stock balance changed.
type StockMovementType string

const (
	StockMovementTypePurchaseReceipt StockMovementType = "PURCHASE_RECEIPT"
	StockMovementTypeSaleReserve     StockMovementType = "SALE_RESERVE"
	StockMovementTypeSaleCommit      StockMovementType = "SALE_COMMIT"
	StockMovementTypeSaleRelease     StockMovementType = "SALE_RELEASE"
	StockMovementTypeAdjustment      StockMovementType = "ADJUSTMENT"
	StockMovementTypeDamage          StockMovementType = "DAMAGE"
	StockMovementTypeReturn          StockMovementType = "RETURN"
	StockMovementTypeTransferIn      StockMovementType = "TRANSFER_IN"
	StockMovementTypeTransferOut     StockMovementType = "TRANSFER_OUT"
	StockMovementTypeStocktake       StockMovementType = "STOCKTAKE"
)

// IsValid reports whether t is a known StockMovementType.
func (t StockMovementType) IsValid() bool {
	switch t {
	case StockMovementTypePurchaseReceipt, StockMovementTypeSaleReserve,
		StockMovementTypeSaleCommit, StockMovementTypeSaleRelease,
		StockMovementTypeAdjustment, StockMovementTypeDamage,
		StockMovementTypeReturn, StockMovementTypeTransferIn,
		StockMovementTypeTransferOut, StockMovementTypeStocktake:
		return true
	}
	return false
}

func (t StockMovementType) String() string { return string(t) }
