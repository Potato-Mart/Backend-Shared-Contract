package enums

import "testing"

func TestStockMovementTypeIsValid(t *testing.T) {
	validTypes := []StockMovementType{
		StockMovementTypePurchaseReceipt,
		StockMovementTypeSaleReserve,
		StockMovementTypeSaleCommit,
		StockMovementTypeSaleRelease,
		StockMovementTypeAdjustment,
		StockMovementTypeDamage,
		StockMovementTypeReturn,
		StockMovementTypeTransferIn,
		StockMovementTypeTransferOut,
		StockMovementTypeStocktake,
	}

	for _, movementType := range validTypes {
		if !movementType.IsValid() {
			t.Fatalf("%q should be valid", movementType)
		}
		if movementType.String() != string(movementType) {
			t.Fatalf("String() = %q, want %q", movementType.String(), string(movementType))
		}
	}

	if StockMovementType("UNKNOWN").IsValid() {
		t.Fatal("UNKNOWN should be invalid")
	}
}
