package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/purchase/purchase_enums"
)

func TestPurchaseEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "purchaseenum.PurchaseOrderStatus", valid: []stringEnum{purchase_enums.PurchaseOrderStatusDraft, purchase_enums.PurchaseOrderStatusSubmitted, purchase_enums.PurchaseOrderStatusConfirmed, purchase_enums.PurchaseOrderStatusPartiallyReceived, purchase_enums.PurchaseOrderStatusReceived, purchase_enums.PurchaseOrderStatusCancelled, purchase_enums.PurchaseOrderStatusRefunded}, invalid: purchase_enums.PurchaseOrderStatus("__invalid__")},
	})
}
