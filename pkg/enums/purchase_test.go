package enums_test

import (
	"testing"

	purchaseenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/purchase"
)

func TestPurchaseEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "purchaseenum.PurchaseOrderStatus", valid: []stringEnum{purchaseenum.PurchaseOrderStatusDraft, purchaseenum.PurchaseOrderStatusSubmitted, purchaseenum.PurchaseOrderStatusConfirmed, purchaseenum.PurchaseOrderStatusPartiallyReceived, purchaseenum.PurchaseOrderStatusReceived, purchaseenum.PurchaseOrderStatusCancelled, purchaseenum.PurchaseOrderStatusRefunded}, invalid: purchaseenum.PurchaseOrderStatus("__invalid__")},
	})
}

func TestPurchaseOrderStatusTransitions(t *testing.T) {
	allowed := []struct {
		from purchaseenum.PurchaseOrderStatus
		to   purchaseenum.PurchaseOrderStatus
	}{
		{purchaseenum.PurchaseOrderStatusDraft, purchaseenum.PurchaseOrderStatusSubmitted},
		{purchaseenum.PurchaseOrderStatusSubmitted, purchaseenum.PurchaseOrderStatusConfirmed},
		{purchaseenum.PurchaseOrderStatusConfirmed, purchaseenum.PurchaseOrderStatusPartiallyReceived},
		{purchaseenum.PurchaseOrderStatusConfirmed, purchaseenum.PurchaseOrderStatusReceived},
		{purchaseenum.PurchaseOrderStatusPartiallyReceived, purchaseenum.PurchaseOrderStatusReceived},
		{purchaseenum.PurchaseOrderStatusReceived, purchaseenum.PurchaseOrderStatusRefunded},
	}

	for _, transition := range allowed {
		if !transition.from.CanTransitionTo(transition.to) {
			t.Fatalf("%s should transition to %s", transition.from, transition.to)
		}
	}

	rejected := []struct {
		from purchaseenum.PurchaseOrderStatus
		to   purchaseenum.PurchaseOrderStatus
	}{
		{purchaseenum.PurchaseOrderStatusDraft, purchaseenum.PurchaseOrderStatusReceived},
		{purchaseenum.PurchaseOrderStatusSubmitted, purchaseenum.PurchaseOrderStatusRefunded},
		{purchaseenum.PurchaseOrderStatusCancelled, purchaseenum.PurchaseOrderStatusSubmitted},
		{purchaseenum.PurchaseOrderStatusRefunded, purchaseenum.PurchaseOrderStatusConfirmed},
	}

	for _, transition := range rejected {
		if transition.from.CanTransitionTo(transition.to) {
			t.Fatalf("%s should not transition to %s", transition.from, transition.to)
		}
	}

	if !purchaseenum.PurchaseOrderStatusCancelled.IsTerminal() || !purchaseenum.PurchaseOrderStatusRefunded.IsTerminal() {
		t.Fatal("cancelled and refunded purchase orders should be terminal")
	}
	if purchaseenum.PurchaseOrderStatusReceived.IsTerminal() {
		t.Fatal("received purchase order should remain refundable, not terminal")
	}
}
