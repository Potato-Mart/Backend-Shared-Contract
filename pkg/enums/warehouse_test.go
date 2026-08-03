package enums_test

import (
	"testing"

	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

func TestWarehouseEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "warehouseenum.PackingDamageHandling", valid: []stringEnum{warehouseenum.PackingDamageReplaceFromStock, warehouseenum.PackingDamageShortShipRefund}, invalid: warehouseenum.PackingDamageHandling("__invalid__")},
		{name: "warehouseenum.PackingSessionStatus", valid: []stringEnum{warehouseenum.PackingSessionStatusPending, warehouseenum.PackingSessionStatusPacking, warehouseenum.PackingSessionStatusPacked, warehouseenum.PackingSessionStatusResolved}, invalid: warehouseenum.PackingSessionStatus("__invalid__")},
		{name: "warehouseenum.CameraProjection", valid: []stringEnum{warehouseenum.CameraPerspective, warehouseenum.CameraOrthographic}, invalid: warehouseenum.CameraProjection("__invalid__")},
		{name: "warehouseenum.DamageStage", valid: []stringEnum{warehouseenum.DamageStageInbound, warehouseenum.DamageStagePicking, warehouseenum.DamageStagePacking, warehouseenum.DamageStageStorage}, invalid: warehouseenum.DamageStage("__invalid__")},
		{name: "warehouseenum.InboundReceiptStatus", valid: []stringEnum{warehouseenum.InboundReceiptStatusDraft, warehouseenum.InboundReceiptStatusConfirmed}, invalid: warehouseenum.InboundReceiptStatus("__invalid__")},
		{name: "warehouseenum.LayoutNodeType", valid: []stringEnum{warehouseenum.LayoutNodeZone, warehouseenum.LayoutNodeAisle, warehouseenum.LayoutNodeRack, warehouseenum.LayoutNodeShelf, warehouseenum.LayoutNodeBin}, invalid: warehouseenum.LayoutNodeType("__invalid__")},
		{name: "warehouseenum.ModelFormat", valid: []stringEnum{warehouseenum.ModelFormatGLB, warehouseenum.ModelFormatGLTF, warehouseenum.ModelFormatOBJ, warehouseenum.ModelFormatFBX, warehouseenum.ModelFormatUSDZ}, invalid: warehouseenum.ModelFormat("__invalid__")},
		{name: "warehouseenum.OutboundShipmentStatus", valid: []stringEnum{warehouseenum.OutboundShipmentStatusPacked, warehouseenum.OutboundShipmentStatusDispatched, warehouseenum.OutboundShipmentStatusDelivered}, invalid: warehouseenum.OutboundShipmentStatus("__invalid__")},
		{name: "warehouseenum.PackingDiscrepancyKind", valid: []stringEnum{warehouseenum.PackingDiscrepancyKindShortage, warehouseenum.PackingDiscrepancyKindOverweight, warehouseenum.PackingDiscrepancyKindDamaged, warehouseenum.PackingDiscrepancyKindPending}, invalid: warehouseenum.PackingDiscrepancyKind("__invalid__")},
		{name: "warehouseenum.PickingItemStatus", valid: []stringEnum{warehouseenum.PickingItemStatusPending, warehouseenum.PickingItemStatusPartial, warehouseenum.PickingItemStatusComplete, warehouseenum.PickingItemStatusSkipped}, invalid: warehouseenum.PickingItemStatus("__invalid__")},
		{name: "warehouseenum.PickingListStatus", valid: []stringEnum{warehouseenum.PickingListStatusPending, warehouseenum.PickingListStatusInProgress, warehouseenum.PickingListStatusComplete, warehouseenum.PickingListStatusCancelled}, invalid: warehouseenum.PickingListStatus("__invalid__")},
		{name: "warehouseenum.ShapeType", valid: []stringEnum{warehouseenum.ShapeBox, warehouseenum.ShapeCylinder, warehouseenum.ShapeSphere, warehouseenum.ShapePlane, warehouseenum.ShapeCustom}, invalid: warehouseenum.ShapeType("__invalid__")},
		{name: "warehouseenum.StockMovementType", valid: []stringEnum{warehouseenum.StockMovementTypePurchaseReceipt, warehouseenum.StockMovementTypeSaleReserve, warehouseenum.StockMovementTypeSaleCommit, warehouseenum.StockMovementTypeSaleRelease, warehouseenum.StockMovementTypeAdjustment, warehouseenum.StockMovementTypeDamage, warehouseenum.StockMovementTypeReturn, warehouseenum.StockMovementTypeTransferIn, warehouseenum.StockMovementTypeTransferOut, warehouseenum.StockMovementTypeStocktake}, invalid: warehouseenum.StockMovementType("__invalid__")},
		{name: "warehouseenum.StorageType", valid: []stringEnum{warehouseenum.StorageDry, warehouseenum.StorageChilled, warehouseenum.StorageFrozen}, invalid: warehouseenum.StorageType("__invalid__")},
		{name: "warehouseenum.WMSDraftStatus", valid: []stringEnum{warehouseenum.WMSDraftStatusDraft, warehouseenum.WMSDraftStatusSubmitted, warehouseenum.WMSDraftStatusCancelled}, invalid: warehouseenum.WMSDraftStatus("__invalid__")},
		{name: "warehouseenum.WMSDraftType", valid: []stringEnum{warehouseenum.WMSDraftTypeInbound, warehouseenum.WMSDraftTypeOutbound}, invalid: warehouseenum.WMSDraftType("__invalid__")},
	})
}
