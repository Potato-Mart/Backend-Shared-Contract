package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/orders/order/order_enums"
)

func TestOrderEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "orderenum.LooseSubstitutionPolicySource", valid: []stringEnum{order_enums.LooseSubstitutionPolicySourceChannelDefault, order_enums.LooseSubstitutionPolicySourceBuyerSelected, order_enums.LooseSubstitutionPolicySourceGroupManager}, invalid: order_enums.LooseSubstitutionPolicySource("__invalid__")},
		{name: "orderenum.GroupOrderRole", valid: []stringEnum{order_enums.GroupOrderRoleConsolidatedParent, order_enums.GroupOrderRoleParticipant}, invalid: order_enums.GroupOrderRole("__invalid__")},
		{name: "orderenum.FulfillmentStatus", valid: []stringEnum{order_enums.FulfillmentStatusUnfulfilled, order_enums.FulfillmentStatusPickingPrinted, order_enums.FulfillmentStatusPacking, order_enums.FulfillmentStatusPacked, order_enums.FulfillmentStatusPartial, order_enums.FulfillmentStatusFulfilled}, invalid: order_enums.FulfillmentStatus("__invalid__")},
		{name: "orderenum.OrderSourceDeviceType", valid: []stringEnum{order_enums.OrderSourceDeviceTypeIOS, order_enums.OrderSourceDeviceTypeAndroid, order_enums.OrderSourceDeviceTypePC, order_enums.OrderSourceDeviceTypeMobileWeb, order_enums.OrderSourceDeviceTypeTablet, order_enums.OrderSourceDeviceTypePos, order_enums.OrderSourceDeviceTypeManual, order_enums.OrderSourceDeviceTypePhone, order_enums.OrderSourceDeviceTypeVR}, invalid: order_enums.OrderSourceDeviceType("__invalid__")},
		{name: "commerce_enums.OrderType", valid: []stringEnum{commerce_enums.OrderTypeOnline, commerce_enums.OrderTypePOS, commerce_enums.OrderTypeB2B, commerce_enums.OrderTypeRelay, commerce_enums.OrderTypeManual, commerce_enums.OrderTypeImport}, invalid: commerce_enums.OrderType("__invalid__")},
		{name: "orderenum.PreorderAllocationStatus", valid: []stringEnum{order_enums.PreorderAllocationStatusWaitingForStock, order_enums.PreorderAllocationStatusPartiallyAllocated, order_enums.PreorderAllocationStatusStockAllocated}, invalid: order_enums.PreorderAllocationStatus("__invalid__")},
		{name: "orderenum.FulfillmentReadiness", valid: []stringEnum{order_enums.FulfillmentReadinessReady, order_enums.FulfillmentReadinessWaitingForPreorderStock}, invalid: order_enums.FulfillmentReadiness("__invalid__")},
		{name: "orderenum.SalesOrderStatus", valid: []stringEnum{order_enums.SalesOrderStatusPending, order_enums.SalesOrderStatusConfirmed, order_enums.SalesOrderStatusPaid, order_enums.SalesOrderStatusProcessing, order_enums.SalesOrderStatusPicking, order_enums.SalesOrderStatusPacked, order_enums.SalesOrderStatusShipped, order_enums.SalesOrderStatusDelivered, order_enums.SalesOrderStatusCompleted, order_enums.SalesOrderStatusCancelled, order_enums.SalesOrderStatusRefunded}, invalid: order_enums.SalesOrderStatus("__invalid__")},
	})
}
