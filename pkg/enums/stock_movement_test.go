package enums_test

import (
	"testing"

	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/warehouse"
)

func TestStockMovementTypeIsValid(t *testing.T) {
	validTypes := []warehouseenum.StockMovementType{
		warehouseenum.StockMovementTypePurchaseReceipt,
		warehouseenum.StockMovementTypeSaleReserve,
		warehouseenum.StockMovementTypeSaleCommit,
		warehouseenum.StockMovementTypeSaleRelease,
		warehouseenum.StockMovementTypeAdjustment,
		warehouseenum.StockMovementTypeDamage,
		warehouseenum.StockMovementTypeReturn,
		warehouseenum.StockMovementTypeTransferIn,
		warehouseenum.StockMovementTypeTransferOut,
		warehouseenum.StockMovementTypeStocktake,
	}

	for _, movementType := range validTypes {
		if !movementType.IsValid() {
			t.Fatalf("%q should be valid", movementType)
		}
		if movementType.String() != string(movementType) {
			t.Fatalf("String() = %q, want %q", movementType.String(), string(movementType))
		}
	}

	if warehouseenum.StockMovementType("UNKNOWN").IsValid() {
		t.Fatal("UNKNOWN should be invalid")
	}
}
