package enums_test

import (
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	"testing"

	orderenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/orders/order"
)

func TestOrderEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "orderenum.LooseSubstitutionPolicySource", valid: []stringEnum{orderenum.LooseSubstitutionPolicySourceChannelDefault, orderenum.LooseSubstitutionPolicySourceBuyerSelected, orderenum.LooseSubstitutionPolicySourceGroupManager}, invalid: orderenum.LooseSubstitutionPolicySource("__invalid__")},
		{name: "orderenum.GroupOrderRole", valid: []stringEnum{orderenum.GroupOrderRoleConsolidatedParent, orderenum.GroupOrderRoleParticipant}, invalid: orderenum.GroupOrderRole("__invalid__")},
		{name: "orderenum.FulfillmentStatus", valid: []stringEnum{orderenum.FulfillmentStatusUnfulfilled, orderenum.FulfillmentStatusPickingPrinted, orderenum.FulfillmentStatusPacking, orderenum.FulfillmentStatusPacked, orderenum.FulfillmentStatusPartial, orderenum.FulfillmentStatusFulfilled}, invalid: orderenum.FulfillmentStatus("__invalid__")},
		{name: "orderenum.OrderSourceDeviceType", valid: []stringEnum{orderenum.OrderSourceDeviceTypeIOS, orderenum.OrderSourceDeviceTypeAndroid, orderenum.OrderSourceDeviceTypePC, orderenum.OrderSourceDeviceTypeMobileWeb, orderenum.OrderSourceDeviceTypeTablet, orderenum.OrderSourceDeviceTypePos, orderenum.OrderSourceDeviceTypeManual, orderenum.OrderSourceDeviceTypePhone, orderenum.OrderSourceDeviceTypeVR}, invalid: orderenum.OrderSourceDeviceType("__invalid__")},
		{name: "common.OrderType", valid: []stringEnum{common.OrderTypeOnline, common.OrderTypePOS, common.OrderTypeB2B, common.OrderTypeRelay, common.OrderTypeManual, common.OrderTypeImport}, invalid: common.OrderType("__invalid__")},
		{name: "orderenum.PreorderAllocationStatus", valid: []stringEnum{orderenum.PreorderAllocationStatusWaitingForStock, orderenum.PreorderAllocationStatusPartiallyAllocated, orderenum.PreorderAllocationStatusStockAllocated}, invalid: orderenum.PreorderAllocationStatus("__invalid__")},
		{name: "orderenum.FulfillmentReadiness", valid: []stringEnum{orderenum.FulfillmentReadinessReady, orderenum.FulfillmentReadinessWaitingForPreorderStock}, invalid: orderenum.FulfillmentReadiness("__invalid__")},
		{name: "orderenum.SalesOrderStatus", valid: []stringEnum{orderenum.SalesOrderStatusPending, orderenum.SalesOrderStatusConfirmed, orderenum.SalesOrderStatusPaid, orderenum.SalesOrderStatusProcessing, orderenum.SalesOrderStatusPicking, orderenum.SalesOrderStatusPacked, orderenum.SalesOrderStatusShipped, orderenum.SalesOrderStatusDelivered, orderenum.SalesOrderStatusCompleted, orderenum.SalesOrderStatusCancelled, orderenum.SalesOrderStatusRefunded}, invalid: orderenum.SalesOrderStatus("__invalid__")},
	})
}
