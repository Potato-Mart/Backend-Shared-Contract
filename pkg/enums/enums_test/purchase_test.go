package enums_test

import (
	"testing"

	purchaseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/purchase"
)

func TestPurchaseEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "purchaseenum.PurchaseOrderStatus", valid: []stringEnum{purchaseenum.PurchaseOrderStatusDraft, purchaseenum.PurchaseOrderStatusSubmitted, purchaseenum.PurchaseOrderStatusConfirmed, purchaseenum.PurchaseOrderStatusPartiallyReceived, purchaseenum.PurchaseOrderStatusReceived, purchaseenum.PurchaseOrderStatusCancelled, purchaseenum.PurchaseOrderStatusRefunded}, invalid: purchaseenum.PurchaseOrderStatus("__invalid__")},
	})
}
