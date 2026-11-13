package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestWarehouseEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "warehouseenum.DamageSaleTier", valid: []stringEnum{warehouse_enums.DamageSaleTier30, warehouse_enums.DamageSaleTier50, warehouse_enums.DamageSaleTier80}, invalid: warehouse_enums.DamageSaleTier("__invalid__")},
		{name: "warehouseenum.PackingDamageHandling", valid: []stringEnum{warehouse_enums.PackingDamageReplaceFromStock, warehouse_enums.PackingDamageShortShipRefund}, invalid: warehouse_enums.PackingDamageHandling("__invalid__")},
		{name: "warehouseenum.CameraProjection", valid: []stringEnum{warehouse_enums.CameraPerspective, warehouse_enums.CameraOrthographic}, invalid: warehouse_enums.CameraProjection("__invalid__")},
		{name: "warehouseenum.DamageStage", valid: []stringEnum{warehouse_enums.DamageStageInbound, warehouse_enums.DamageStagePicking, warehouse_enums.DamageStagePacking, warehouse_enums.DamageStageStorage}, invalid: warehouse_enums.DamageStage("__invalid__")},
		{name: "warehouseenum.InboundReceiptStatus", valid: []stringEnum{warehouse_enums.InboundReceiptStatusDraft, warehouse_enums.InboundReceiptStatusConfirmed, warehouse_enums.InboundReceiptStatusReversed}, invalid: warehouse_enums.InboundReceiptStatus("__invalid__")},
		{name: "warehouseenum.LayoutNodeType", valid: []stringEnum{warehouse_enums.LayoutNodeZone, warehouse_enums.LayoutNodeAisle, warehouse_enums.LayoutNodeRack, warehouse_enums.LayoutNodeShelf, warehouse_enums.LayoutNodeBin}, invalid: warehouse_enums.LayoutNodeType("__invalid__")},
		{name: "warehouseenum.ModelFormat", valid: []stringEnum{warehouse_enums.ModelFormatGLB, warehouse_enums.ModelFormatGLTF, warehouse_enums.ModelFormatOBJ, warehouse_enums.ModelFormatFBX, warehouse_enums.ModelFormatUSDZ}, invalid: warehouse_enums.ModelFormat("__invalid__")},
		{name: "warehouseenum.OutboundShipmentStatus", valid: []stringEnum{warehouse_enums.OutboundShipmentStatusPacked, warehouse_enums.OutboundShipmentStatusDispatched, warehouse_enums.OutboundShipmentStatusDelivered}, invalid: warehouse_enums.OutboundShipmentStatus("__invalid__")},
		{name: "warehouseenum.ShipmentWarningCode", valid: []stringEnum{warehouse_enums.ShipmentWarningCodeCarrierCancellationRequired, warehouse_enums.ShipmentWarningCodeOrderUnpaid}, invalid: warehouse_enums.ShipmentWarningCode("__invalid__")},
		{name: "warehouseenum.PackingDiscrepancyKind", valid: []stringEnum{warehouse_enums.PackingDiscrepancyKindShortage, warehouse_enums.PackingDiscrepancyKindOverweight, warehouse_enums.PackingDiscrepancyKindDamaged, warehouse_enums.PackingDiscrepancyKindPending}, invalid: warehouse_enums.PackingDiscrepancyKind("__invalid__")},
		{name: "warehouseenum.PackingDiscrepancyWithdrawalMethod", valid: []stringEnum{warehouse_enums.PackingDiscrepancyWithdrawalMethodCommand, warehouse_enums.PackingDiscrepancyWithdrawalMethodScan, warehouse_enums.PackingDiscrepancyWithdrawalMethodReinstate}, invalid: warehouse_enums.PackingDiscrepancyWithdrawalMethod("__invalid__")},
		{name: "warehouseenum.PackingShortageSettlementReason", valid: []stringEnum{warehouse_enums.PackingShortageSettlementReasonNoStock, warehouse_enums.PackingShortageSettlementReasonOverweight}, invalid: warehouse_enums.PackingShortageSettlementReason("__invalid__")},
		{name: "warehouseenum.PackingShortageSettlementStatus", valid: []stringEnum{warehouse_enums.PackingShortageSettlementStatusApplied, warehouse_enums.PackingShortageSettlementStatusDeferred, warehouse_enums.PackingShortageSettlementStatusReversed}, invalid: warehouse_enums.PackingShortageSettlementStatus("__invalid__")},
		{name: "warehouseenum.PackingShortageSettlementReversalCause", valid: []stringEnum{warehouse_enums.PackingShortageSettlementReversalCauseWithdrawal, warehouse_enums.PackingShortageSettlementReversalCauseScan, warehouse_enums.PackingShortageSettlementReversalCauseProgress, warehouse_enums.PackingShortageSettlementReversalCauseSuperseded}, invalid: warehouse_enums.PackingShortageSettlementReversalCause("__invalid__")},
		{name: "warehouseenum.ShapeType", valid: []stringEnum{warehouse_enums.ShapeBox, warehouse_enums.ShapeCylinder, warehouse_enums.ShapeSphere, warehouse_enums.ShapePlane, warehouse_enums.ShapeCustom}, invalid: warehouse_enums.ShapeType("__invalid__")},
		{name: "warehouseenum.StockMovementType", valid: []stringEnum{warehouse_enums.StockMovementTypeReceipt, warehouse_enums.StockMovementTypeTransfer, warehouse_enums.StockMovementTypeStage, warehouse_enums.StockMovementTypeUnstage, warehouse_enums.StockMovementTypePackageConversion, warehouse_enums.StockMovementTypeSaleCommit, warehouse_enums.StockMovementTypeReturn, warehouse_enums.StockMovementTypeQualityHold, warehouse_enums.StockMovementTypeQualityRelease, warehouse_enums.StockMovementTypeQualityReject, warehouse_enums.StockMovementTypeAdjustment, warehouse_enums.StockMovementTypeReturnToVendor, warehouse_enums.StockMovementTypeScrap, warehouse_enums.StockMovementTypeLotReassignment, warehouse_enums.StockMovementTypeReceiptReversal}, invalid: warehouse_enums.StockMovementType("__invalid__")},
		{name: "warehouseenum.InventoryDateMarkKind", valid: []stringEnum{warehouse_enums.InventoryDateMarkBestBefore, warehouse_enums.InventoryDateMarkExpiry}, invalid: warehouse_enums.InventoryDateMarkKind("__invalid__")},
		{name: "warehouseenum.InventoryCondition", valid: []stringEnum{warehouse_enums.InventoryConditionGood, warehouse_enums.InventoryConditionPackagingDamagedMinor, warehouse_enums.InventoryConditionPackagingDamagedMajor, warehouse_enums.InventoryConditionProductDamaged, warehouse_enums.InventoryConditionOpened, warehouse_enums.InventoryConditionContaminated, warehouse_enums.InventoryConditionSpoiled}, invalid: warehouse_enums.InventoryCondition("__invalid__")},
		{name: "warehouseenum.InventoryDisposition", valid: []stringEnum{warehouse_enums.InventoryDispositionStandardSellable, warehouse_enums.InventoryDispositionReducedSellable, warehouse_enums.InventoryDispositionQualityHold, warehouse_enums.InventoryDispositionQuarantined, warehouse_enums.InventoryDispositionBlocked, warehouse_enums.InventoryDispositionRecalled, warehouse_enums.InventoryDispositionReturnToVendor, warehouse_enums.InventoryDispositionScrap}, invalid: warehouse_enums.InventoryDisposition("__invalid__")},
		{name: "warehouseenum.StockReservationStatus", valid: []stringEnum{warehouse_enums.StockReservationStatusReserved, warehouse_enums.StockReservationStatusPartiallyStaged, warehouse_enums.StockReservationStatusStaged, warehouse_enums.StockReservationStatusCommitted, warehouse_enums.StockReservationStatusReleased, warehouse_enums.StockReservationStatusExpired, warehouse_enums.StockReservationStatusCancelled}, invalid: warehouse_enums.StockReservationStatus("__invalid__")},
		{name: "warehouseenum.StockAvailabilityDirection", valid: []stringEnum{warehouse_enums.StockAvailabilityOutOfStock, warehouse_enums.StockAvailabilityRestocked}, invalid: warehouse_enums.StockAvailabilityDirection("__invalid__")},
		{name: "warehouseenum.InventoryDateMarkThreshold", valid: []stringEnum{warehouse_enums.InventoryDateMarkThresholdApproaching, warehouse_enums.InventoryDateMarkThresholdReached, warehouse_enums.InventoryDateMarkThresholdPassed}, invalid: warehouse_enums.InventoryDateMarkThreshold("__invalid__")},
		{name: "warehouseenum.StockLocationPurpose", valid: []stringEnum{warehouse_enums.StockLocationPurposeExpiryHold, warehouse_enums.StockLocationPurposeStandard, warehouse_enums.StockLocationPurposeQualityHold, warehouse_enums.StockLocationPurposeOnlineOrderStaging}, invalid: warehouse_enums.StockLocationPurpose("__invalid__")},
		{name: "warehouseenum.StockLocationHandlingMode", valid: []stringEnum{warehouse_enums.StockLocationHandlingEach, warehouse_enums.StockLocationHandlingCase, warehouse_enums.StockLocationHandlingMixed}, invalid: warehouse_enums.StockLocationHandlingMode("__invalid__")},
		{name: "warehouseenum.StockLocationAccess", valid: []stringEnum{warehouse_enums.StockLocationAccessCustomerAccessible, warehouse_enums.StockLocationAccessStaffOnly}, invalid: warehouse_enums.StockLocationAccess("__invalid__")},
		{name: "warehouseenum.StockLocationCollectionMode", valid: []stringEnum{warehouse_enums.StockLocationCollectionAllowList, warehouse_enums.StockLocationCollectionUnrestricted}, invalid: warehouse_enums.StockLocationCollectionMode("__invalid__")},
		{name: "warehouseenum.StockLocationCollectionRole", valid: []stringEnum{warehouse_enums.StockLocationCollectionPrimary, warehouse_enums.StockLocationCollectionOverflow}, invalid: warehouse_enums.StockLocationCollectionRole("__invalid__")},
		{name: "warehouseenum.StorageType", valid: []stringEnum{classification_enums.StorageAmbient, classification_enums.StorageChilled, classification_enums.StorageFrozen}, invalid: classification_enums.StorageType("DRY")},
		{name: "warehouseenum.LocationStorageType", valid: []stringEnum{warehouse_enums.LocationStorageAmbient, warehouse_enums.LocationStorageChilled, warehouse_enums.LocationStorageFrozen, warehouse_enums.LocationStorageMixed}, invalid: warehouse_enums.LocationStorageType("DRY")},
		{name: "warehouseenum.WMSDraftStatus", valid: []stringEnum{warehouse_enums.WMSDraftStatusDraft, warehouse_enums.WMSDraftStatusSubmitted, warehouse_enums.WMSDraftStatusCancelled}, invalid: warehouse_enums.WMSDraftStatus("__invalid__")},
		{name: "warehouseenum.WMSDraftItemMoveStep", valid: []stringEnum{warehouse_enums.WMSDraftItemMoveStepPending, warehouse_enums.WMSDraftItemMoveStepMoved, warehouse_enums.WMSDraftItemMoveStepDone}, invalid: warehouse_enums.WMSDraftItemMoveStep("__invalid__")},
		{name: "warehouseenum.WMSDraftType", valid: []stringEnum{warehouse_enums.WMSDraftTypeInbound, warehouse_enums.WMSDraftTypeOutbound}, invalid: warehouse_enums.WMSDraftType("__invalid__")},
	})
}

// Supply compares a location's storage type with a product's by wire value,
// so the shared values must stay identical. MIXED is location-only.
func TestLocationStorageTypeSharesProductStorageWireValues(t *testing.T) {
	for location, product := range map[warehouse_enums.LocationStorageType]classification_enums.StorageType{
		warehouse_enums.LocationStorageAmbient: classification_enums.StorageAmbient,
		warehouse_enums.LocationStorageChilled: classification_enums.StorageChilled,
		warehouse_enums.LocationStorageFrozen:  classification_enums.StorageFrozen,
	} {
		if string(location) != string(product) {
			t.Errorf("location storage %q differs from product storage %q", location, product)
		}
	}
	if classification_enums.StorageType(warehouse_enums.LocationStorageMixed).IsValid() {
		t.Fatal("MIXED must not be a valid product storage type")
	}
}

func TestPackingDiscrepancyWithdrawalMethodWireValuesAreStable(t *testing.T) {
	for value, want := range map[warehouse_enums.PackingDiscrepancyWithdrawalMethod]string{
		warehouse_enums.PackingDiscrepancyWithdrawalMethodCommand:   "command",
		warehouse_enums.PackingDiscrepancyWithdrawalMethodScan:      "scan",
		warehouse_enums.PackingDiscrepancyWithdrawalMethodReinstate: "reinstate",
	} {
		if got := value.String(); got != want {
			t.Errorf("PackingDiscrepancyWithdrawalMethod wire value = %q, want %q", got, want)
		}
	}
}

func TestPackingShortageSettlementWireValuesAreStable(t *testing.T) {
	for value, want := range map[string]string{
		warehouse_enums.PackingShortageSettlementReasonNoStock.String():           "no_stock",
		warehouse_enums.PackingShortageSettlementReasonOverweight.String():        "overweight",
		warehouse_enums.PackingShortageSettlementStatusApplied.String():           "applied",
		warehouse_enums.PackingShortageSettlementStatusDeferred.String():          "deferred",
		warehouse_enums.PackingShortageSettlementStatusReversed.String():          "reversed",
		warehouse_enums.PackingShortageSettlementReversalCauseWithdrawal.String(): "withdrawal",
		warehouse_enums.PackingShortageSettlementReversalCauseScan.String():       "scan",
		warehouse_enums.PackingShortageSettlementReversalCauseProgress.String():   "progress",
		warehouse_enums.PackingShortageSettlementReversalCauseSuperseded.String(): "superseded",
	} {
		if value != want {
			t.Errorf("packing shortage settlement wire value = %q, want %q", value, want)
		}
	}
}
