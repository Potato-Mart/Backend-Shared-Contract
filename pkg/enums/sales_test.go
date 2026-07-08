package enums_test

import (
	"testing"

	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/enums/sales"
)

func TestSalesEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "salesenum.FulfillmentStatus", valid: []stringEnum{salesenum.FulfillmentStatusUnfulfilled, salesenum.FulfillmentStatusPickingPrinted, salesenum.FulfillmentStatusPacking, salesenum.FulfillmentStatusPacked, salesenum.FulfillmentStatusPartial, salesenum.FulfillmentStatusFulfilled}, invalid: salesenum.FulfillmentStatus("__invalid__")},
		{name: "salesenum.OrderSourceDeviceType", valid: []stringEnum{salesenum.OrderSourceDeviceTypeIOS, salesenum.OrderSourceDeviceTypeAndroid, salesenum.OrderSourceDeviceTypePC, salesenum.OrderSourceDeviceTypeMobileWeb, salesenum.OrderSourceDeviceTypeTablet, salesenum.OrderSourceDeviceTypePos, salesenum.OrderSourceDeviceTypeManual, salesenum.OrderSourceDeviceTypePhone, salesenum.OrderSourceDeviceTypeVR}, invalid: salesenum.OrderSourceDeviceType("__invalid__")},
		{name: "salesenum.OrderType", valid: []stringEnum{salesenum.OrderTypeOnline, salesenum.OrderTypePOS, salesenum.OrderTypeB2B, salesenum.OrderTypeRelay, salesenum.OrderTypeManual, salesenum.OrderTypeImport}, invalid: salesenum.OrderType("__invalid__")},
		{name: "salesenum.PreorderStatus", valid: []stringEnum{salesenum.PreorderStatusRequested, salesenum.PreorderStatusAccepted, salesenum.PreorderStatusRejected, salesenum.PreorderStatusCancelled, salesenum.PreorderStatusConverted, salesenum.PreorderStatusFulfilled, salesenum.PreorderStatusExpired}, invalid: salesenum.PreorderStatus("__invalid__")},
		{name: "salesenum.SalesOrderStatus", valid: []stringEnum{salesenum.SalesOrderStatusPending, salesenum.SalesOrderStatusConfirmed, salesenum.SalesOrderStatusPaid, salesenum.SalesOrderStatusProcessing, salesenum.SalesOrderStatusPicking, salesenum.SalesOrderStatusPacked, salesenum.SalesOrderStatusShipped, salesenum.SalesOrderStatusDelivered, salesenum.SalesOrderStatusCompleted, salesenum.SalesOrderStatusCancelled, salesenum.SalesOrderStatusRefunded}, invalid: salesenum.SalesOrderStatus("__invalid__")},
	})
}

func TestSalesOrderStatusTransitions(t *testing.T) {
	allowed := []struct {
		from salesenum.SalesOrderStatus
		to   salesenum.SalesOrderStatus
	}{
		{salesenum.SalesOrderStatusPending, salesenum.SalesOrderStatusConfirmed},
		{salesenum.SalesOrderStatusConfirmed, salesenum.SalesOrderStatusPaid},
		{salesenum.SalesOrderStatusPaid, salesenum.SalesOrderStatusProcessing},
		{salesenum.SalesOrderStatusProcessing, salesenum.SalesOrderStatusPicking},
		{salesenum.SalesOrderStatusPicking, salesenum.SalesOrderStatusPacked},
		{salesenum.SalesOrderStatusPacked, salesenum.SalesOrderStatusShipped},
		{salesenum.SalesOrderStatusShipped, salesenum.SalesOrderStatusDelivered},
		{salesenum.SalesOrderStatusDelivered, salesenum.SalesOrderStatusCompleted},
		{salesenum.SalesOrderStatusCompleted, salesenum.SalesOrderStatusRefunded},
	}

	for _, transition := range allowed {
		if !transition.from.CanTransitionTo(transition.to) {
			t.Fatalf("%s should transition to %s", transition.from, transition.to)
		}
	}

	rejected := []struct {
		from salesenum.SalesOrderStatus
		to   salesenum.SalesOrderStatus
	}{
		{salesenum.SalesOrderStatusPending, salesenum.SalesOrderStatusPacked},
		{salesenum.SalesOrderStatusProcessing, salesenum.SalesOrderStatusCancelled},
		{salesenum.SalesOrderStatusCancelled, salesenum.SalesOrderStatusPending},
		{salesenum.SalesOrderStatusRefunded, salesenum.SalesOrderStatusPaid},
	}

	for _, transition := range rejected {
		if transition.from.CanTransitionTo(transition.to) {
			t.Fatalf("%s should not transition to %s", transition.from, transition.to)
		}
	}

	if !salesenum.SalesOrderStatusCancelled.IsTerminal() || !salesenum.SalesOrderStatusRefunded.IsTerminal() {
		t.Fatal("cancelled and refunded sales orders should be terminal")
	}
	if salesenum.SalesOrderStatusCompleted.IsTerminal() {
		t.Fatal("completed sales order should remain refundable, not terminal")
	}
}
