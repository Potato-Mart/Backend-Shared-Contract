package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/commerce/commerce_enums"
	group_order_enums "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/group_order/group_order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order/order_enums"
)

func TestOrderEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "orderenum.OfflineCashierAuthKind", valid: []stringEnum{order_enums.OfflineCashierAuthKindPIN, order_enums.OfflineCashierAuthKindPasswordSession}, invalid: order_enums.OfflineCashierAuthKind("__invalid__")},
		{name: "orderenum.LooseSubstitutionPolicySource", valid: []stringEnum{order_enums.LooseSubstitutionPolicySourceChannelDefault, order_enums.LooseSubstitutionPolicySourceBuyerSelected, order_enums.LooseSubstitutionPolicySourceGroupManager}, invalid: order_enums.LooseSubstitutionPolicySource("__invalid__")},
		{name: "orderenum.GroupOrderRole", valid: []stringEnum{group_order_enums.GroupOrderRoleConsolidatedParent, group_order_enums.GroupOrderRoleParticipant}, invalid: group_order_enums.GroupOrderRole("__invalid__")},
		{name: "orderenum.FulfillmentStatus", valid: []stringEnum{order_enums.FulfillmentStatusUnfulfilled, order_enums.FulfillmentStatusPickingInProgress, order_enums.FulfillmentStatusPacking, order_enums.FulfillmentStatusPacked, order_enums.FulfillmentStatusFulfilled, order_enums.FulfillmentStatusCancelled}, invalid: order_enums.FulfillmentStatus("__invalid__")},
		{name: "orderenum.OrderSourceDeviceType", valid: []stringEnum{order_enums.OrderSourceDeviceTypeIOS, order_enums.OrderSourceDeviceTypeAndroid, order_enums.OrderSourceDeviceTypePC, order_enums.OrderSourceDeviceTypeMobileWeb, order_enums.OrderSourceDeviceTypeTablet, order_enums.OrderSourceDeviceTypePos, order_enums.OrderSourceDeviceTypeManual, order_enums.OrderSourceDeviceTypePhone, order_enums.OrderSourceDeviceTypeVR}, invalid: order_enums.OrderSourceDeviceType("__invalid__")},
		{name: "commerce_enums.OrderType", valid: []stringEnum{commerce_enums.OrderTypeOnline, commerce_enums.OrderTypePOS, commerce_enums.OrderTypeB2B, commerce_enums.OrderTypeRelay, commerce_enums.OrderTypeManual, commerce_enums.OrderTypeImport}, invalid: commerce_enums.OrderType("__invalid__")},
		{name: "orderenum.PreorderAllocationStatus", valid: []stringEnum{order_enums.PreorderAllocationStatusWaitingForStock, order_enums.PreorderAllocationStatusPartiallyAllocated, order_enums.PreorderAllocationStatusStockAllocated}, invalid: order_enums.PreorderAllocationStatus("__invalid__")},
		{name: "orderenum.FulfillmentReadiness", valid: []stringEnum{order_enums.FulfillmentReadinessReady, order_enums.FulfillmentReadinessWaitingForPreorderStock}, invalid: order_enums.FulfillmentReadiness("__invalid__")},
		{name: "orderenum.SalesOrderStatus", valid: []stringEnum{order_enums.SalesOrderStatusPending, order_enums.SalesOrderStatusConfirmed, order_enums.SalesOrderStatusPaid, order_enums.SalesOrderStatusProcessing, order_enums.SalesOrderStatusPicking, order_enums.SalesOrderStatusPacked, order_enums.SalesOrderStatusShipped, order_enums.SalesOrderStatusDelivered, order_enums.SalesOrderStatusCompleted, order_enums.SalesOrderStatusCancelled, order_enums.SalesOrderStatusRefunded}, invalid: order_enums.SalesOrderStatus("__invalid__")},
		{name: "orderenum.StatusHistoryKind", valid: []stringEnum{order_enums.StatusHistoryKindOverride, order_enums.StatusHistoryKindPaymentCorrection}, invalid: order_enums.StatusHistoryKind("__invalid__")},
		{name: "orderenum.GiftCardIssueStatus", valid: []stringEnum{order_enums.GiftCardIssueStatusPending, order_enums.GiftCardIssueStatusIssued, order_enums.GiftCardIssueStatusFailed, order_enums.GiftCardIssueStatusReversalPending, order_enums.GiftCardIssueStatusReversed}, invalid: order_enums.GiftCardIssueStatus("__invalid__")},
		{name: "orderenum.TenderPartStatus", valid: []stringEnum{order_enums.TenderPartStatusPending, order_enums.TenderPartStatusInProgress, order_enums.TenderPartStatusCaptured, order_enums.TenderPartStatusDeclined, order_enums.TenderPartStatusCancelled, order_enums.TenderPartStatusRefunded}, invalid: order_enums.TenderPartStatus("__invalid__")},
	})
}

func TestTenderPartStatusWireValuesAreStable(t *testing.T) {
	want := map[order_enums.TenderPartStatus]string{
		order_enums.TenderPartStatusPending:    "pending",
		order_enums.TenderPartStatusInProgress: "in_progress",
		order_enums.TenderPartStatusCaptured:   "captured",
		order_enums.TenderPartStatusDeclined:   "declined",
		order_enums.TenderPartStatusCancelled:  "cancelled",
		order_enums.TenderPartStatusRefunded:   "refunded",
	}
	if len(want) != 6 {
		t.Fatalf("TenderPartStatus has %d distinct values, want 6", len(want))
	}
	for value, wire := range want {
		if got := value.String(); got != wire {
			t.Errorf("TenderPartStatus wire value = %q, want %q", got, wire)
		}
	}
	for _, unknown := range []order_enums.TenderPartStatus{"", "partially_captured", "CAPTURED"} {
		if unknown.IsValid() {
			t.Errorf("TenderPartStatus %q must not validate", unknown)
		}
	}
}

func TestGiftCardIssueStatusWireValuesAreStable(t *testing.T) {
	for value, want := range map[order_enums.GiftCardIssueStatus]string{
		order_enums.GiftCardIssueStatusPending:         "pending",
		order_enums.GiftCardIssueStatusIssued:          "issued",
		order_enums.GiftCardIssueStatusFailed:          "failed",
		order_enums.GiftCardIssueStatusReversalPending: "reversal_pending",
		order_enums.GiftCardIssueStatusReversed:        "reversed",
	} {
		if got := value.String(); got != want {
			t.Errorf("GiftCardIssueStatus wire value = %q, want %q", got, want)
		}
	}
}
