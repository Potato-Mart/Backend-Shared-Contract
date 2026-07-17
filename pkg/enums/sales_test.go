package enums_test

import (
	"testing"

	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/sales"
)

func TestSalesEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "salesenum.FulfillmentStatus", valid: []stringEnum{salesenum.FulfillmentStatusUnfulfilled, salesenum.FulfillmentStatusPickingPrinted, salesenum.FulfillmentStatusPacking, salesenum.FulfillmentStatusPacked, salesenum.FulfillmentStatusPartial, salesenum.FulfillmentStatusFulfilled}, invalid: salesenum.FulfillmentStatus("__invalid__")},
		{name: "salesenum.OrderSourceDeviceType", valid: []stringEnum{salesenum.OrderSourceDeviceTypeIOS, salesenum.OrderSourceDeviceTypeAndroid, salesenum.OrderSourceDeviceTypePC, salesenum.OrderSourceDeviceTypeMobileWeb, salesenum.OrderSourceDeviceTypeTablet, salesenum.OrderSourceDeviceTypePos, salesenum.OrderSourceDeviceTypeManual, salesenum.OrderSourceDeviceTypePhone, salesenum.OrderSourceDeviceTypeVR}, invalid: salesenum.OrderSourceDeviceType("__invalid__")},
		{name: "salesenum.OrderType", valid: []stringEnum{salesenum.OrderTypeOnline, salesenum.OrderTypePOS, salesenum.OrderTypeB2B, salesenum.OrderTypeRelay, salesenum.OrderTypeManual, salesenum.OrderTypeImport}, invalid: salesenum.OrderType("__invalid__")},
		{name: "salesenum.PreorderAllocationStatus", valid: []stringEnum{salesenum.PreorderAllocationStatusWaitingForStock, salesenum.PreorderAllocationStatusPartiallyAllocated, salesenum.PreorderAllocationStatusStockAllocated}, invalid: salesenum.PreorderAllocationStatus("__invalid__")},
		{name: "salesenum.FulfillmentReadiness", valid: []stringEnum{salesenum.FulfillmentReadinessReady, salesenum.FulfillmentReadinessWaitingForPreorderStock}, invalid: salesenum.FulfillmentReadiness("__invalid__")},
		{name: "salesenum.SalesOrderStatus", valid: []stringEnum{salesenum.SalesOrderStatusPending, salesenum.SalesOrderStatusConfirmed, salesenum.SalesOrderStatusPaid, salesenum.SalesOrderStatusProcessing, salesenum.SalesOrderStatusPicking, salesenum.SalesOrderStatusPacked, salesenum.SalesOrderStatusShipped, salesenum.SalesOrderStatusDelivered, salesenum.SalesOrderStatusCompleted, salesenum.SalesOrderStatusCancelled, salesenum.SalesOrderStatusRefunded}, invalid: salesenum.SalesOrderStatus("__invalid__")},
	})
}
