package enums

import (
	"reflect"
	"testing"
)

type stringEnum interface {
	IsValid() bool
	String() string
}

func TestExportedEnumsValidateKnownValues(t *testing.T) {
	tests := []struct {
		name    string
		valid   []stringEnum
		invalid stringEnum
	}{
		{name: "AuthAssuranceLevel", valid: []stringEnum{AuthAssuranceLevel1, AuthAssuranceLevel2, AuthAssuranceLevel3}, invalid: AuthAssuranceLevel("__invalid__")},
		{name: "AuthMethod", valid: []stringEnum{AuthMethodPassword, AuthMethodMFA, AuthMethodPasskey, AuthMethodSSO, AuthMethodRefreshToken, AuthMethodAPIKey}, invalid: AuthMethod("__invalid__")},
		{name: "CameraProjection", valid: []stringEnum{CameraPerspective, CameraOrthographic}, invalid: CameraProjection("__invalid__")},
		{name: "ChurnRisk", valid: []stringEnum{ChurnRiskLow, ChurnRiskMedium, ChurnRiskHigh}, invalid: ChurnRisk("__invalid__")},
		{name: "CouponAppliesTo", valid: []stringEnum{CouponAppliesToAll, CouponAppliesToSpecificProducts, CouponAppliesToSpecificCategories}, invalid: CouponAppliesTo("__invalid__")},
		{name: "CouponSource", valid: []stringEnum{CouponSourceManual, CouponSourceRFMComeback, CouponSourceBirthday, CouponSourceReferral, CouponSourceSignupBonus, CouponSourceCampaign}, invalid: CouponSource("__invalid__")},
		{name: "CustomerActivityType", valid: []stringEnum{CustomerActivityTypeNote, CustomerActivityTypeCall, CustomerActivityTypeEmail, CustomerActivityTypeSMS, CustomerActivityTypeLine, CustomerActivityTypeOrder, CustomerActivityTypeComplaint, CustomerActivityTypeReturn, CustomerActivityTypeRefund, CustomerActivityTypePointsAdjust, CustomerActivityTypeTierChange, CustomerActivityTypeStatusChange, CustomerActivityTypeReferral, CustomerActivityTypeCampaign}, invalid: CustomerActivityType("__invalid__")},
		{name: "CustomerIdentityKind", valid: []stringEnum{CustomerIdentityKindPhone, CustomerIdentityKindEmail, CustomerIdentityKindLine, CustomerIdentityKindMemberCard, CustomerIdentityKindPOSID, CustomerIdentityKindExternal}, invalid: CustomerIdentityKind("__invalid__")},
		{name: "CustomerProfileStatus", valid: []stringEnum{CustomerProfileStatusActive, CustomerProfileStatusInactive, CustomerProfileStatusBlocked}, invalid: CustomerProfileStatus("__invalid__")},
		{name: "CustomerSegment", valid: []stringEnum{CustomerSegmentWholesale, CustomerSegmentRetail}, invalid: CustomerSegment("__invalid__")},
		{name: "CustomerTier", valid: []stringEnum{CustomerTierStandard, CustomerTierSilver, CustomerTierGold, CustomerTierPlatinum}, invalid: CustomerTier("__invalid__")},
		{name: "DamageStage", valid: []stringEnum{DamageStageInbound, DamageStagePicking, DamageStagePacking, DamageStageStorage}, invalid: DamageStage("__invalid__")},
		{name: "DataClassification", valid: []stringEnum{DataClassificationPublic, DataClassificationInternal, DataClassificationConfidential, DataClassificationRestricted}, invalid: DataClassification("__invalid__")},
		{name: "DataProtectionBasis", valid: []stringEnum{DataProtectionBasisNotApplicable, DataProtectionBasisConsent, DataProtectionBasisContract, DataProtectionBasisLegalObligation, DataProtectionBasisLegitimateInterest}, invalid: DataProtectionBasis("__invalid__")},
		{name: "DeviceType", valid: []stringEnum{DeviceTypeDesktop, DeviceTypeMobile, DeviceTypeTablet, DeviceTypeAPI}, invalid: DeviceType("__invalid__")},
		{name: "DiscountScope", valid: []stringEnum{DiscountScopeAll, DiscountScopeCategory, DiscountScopeProduct}, invalid: DiscountScope("__invalid__")},
		{name: "DiscountType", valid: []stringEnum{DiscountTypePercentage, DiscountTypeFixedAmount, DiscountTypeFreeShipping, DiscountTypeFixedPrice}, invalid: DiscountType("__invalid__")},
		{name: "FulfillmentStatus", valid: []stringEnum{FulfillmentStatusUnfulfilled, FulfillmentStatusPickingPrinted, FulfillmentStatusPacking, FulfillmentStatusPacked, FulfillmentStatusPartial, FulfillmentStatusFulfilled}, invalid: FulfillmentStatus("__invalid__")},
		{name: "InboundReceiptStatus", valid: []stringEnum{InboundReceiptStatusDraft, InboundReceiptStatusConfirmed}, invalid: InboundReceiptStatus("__invalid__")},
		{name: "LayoutNodeType", valid: []stringEnum{LayoutNodeZone, LayoutNodeAisle, LayoutNodeRack, LayoutNodeShelf, LayoutNodeBin}, invalid: LayoutNodeType("__invalid__")},
		{name: "LoyaltyLedgerReason", valid: []stringEnum{LoyaltyLedgerReasonOrder, LoyaltyLedgerReasonBirthday, LoyaltyLedgerReasonRedeem, LoyaltyLedgerReasonAdminAdjust, LoyaltyLedgerReasonExpired, LoyaltyLedgerReasonReferral, LoyaltyLedgerReasonSignupBonus, LoyaltyLedgerReasonTierUpgrade, LoyaltyLedgerReasonManual}, invalid: LoyaltyLedgerReason("__invalid__")},
		{name: "LoyaltyPromotionTarget", valid: []stringEnum{LoyaltyPromotionTargetAll, LoyaltyPromotionTargetWholesale, LoyaltyPromotionTargetRetail, LoyaltyPromotionTargetTierSpecific}, invalid: LoyaltyPromotionTarget("__invalid__")},
		{name: "MarketingCampaignStatus", valid: []stringEnum{MarketingCampaignStatusDraft, MarketingCampaignStatusSending, MarketingCampaignStatusSent, MarketingCampaignStatusPartial, MarketingCampaignStatusFailed, MarketingCampaignStatusCancelled, MarketingCampaignStatusExported}, invalid: MarketingCampaignStatus("__invalid__")},
		{name: "MarketingChannel", valid: []stringEnum{MarketingChannelEmail, MarketingChannelSMS, MarketingChannelLine, MarketingChannelExport}, invalid: MarketingChannel("__invalid__")},
		{name: "MarketingRecipientStatus", valid: []stringEnum{MarketingRecipientStatusPending, MarketingRecipientStatusSent, MarketingRecipientStatusDelivered, MarketingRecipientStatusBounced, MarketingRecipientStatusOpened, MarketingRecipientStatusClicked, MarketingRecipientStatusFailed, MarketingRecipientStatusUnsubscribed}, invalid: MarketingRecipientStatus("__invalid__")},
		{name: "ModelFormat", valid: []stringEnum{ModelFormatGLB, ModelFormatGLTF, ModelFormatOBJ, ModelFormatFBX, ModelFormatUSDZ}, invalid: ModelFormat("__invalid__")},
		{name: "OrderSourceDeviceType", valid: []stringEnum{OrderSourceDeviceTypeIOS, OrderSourceDeviceTypeAndroid, OrderSourceDeviceTypePC, OrderSourceDeviceTypeMobileWeb, OrderSourceDeviceTypeTablet, OrderSourceDeviceTypePos, OrderSourceDeviceTypeManual, OrderSourceDeviceTypePhone, OrderSourceDeviceTypeVR}, invalid: OrderSourceDeviceType("__invalid__")},
		{name: "OrderType", valid: []stringEnum{OrderTypeOnline, OrderTypePOS, OrderTypeB2B, OrderTypeRelay, OrderTypeManual, OrderTypeImport}, invalid: OrderType("__invalid__")},
		{name: "OutboundShipmentStatus", valid: []stringEnum{OutboundShipmentStatusPacked, OutboundShipmentStatusDispatched}, invalid: OutboundShipmentStatus("__invalid__")},
		{name: "PackingDiscrepancyKind", valid: []stringEnum{PackingDiscrepancyKindShortage, PackingDiscrepancyKindOverweight, PackingDiscrepancyKindDamaged}, invalid: PackingDiscrepancyKind("__invalid__")},
		{name: "PaymentMethod", valid: []stringEnum{PaymentMethodCard, PaymentMethodCash, PaymentMethodQR, PaymentMethodBankTransfer, PaymentMethodLinePay, PaymentMethodApplePay, PaymentMethodGooglePay, PaymentMethodECPay, PaymentMethodManual, PaymentMethodEFTPOS, PaymentMethodMOTO, PaymentMethodCashout}, invalid: PaymentMethod("__invalid__")},
		{name: "PaymentRecordStatus", valid: []stringEnum{PaymentRecordStatusPending, PaymentRecordStatusProcessing, PaymentRecordStatusCompleted, PaymentRecordStatusFailed, PaymentRecordStatusCancelled, PaymentRecordStatusRefunded, PaymentRecordStatusAwaitingAction, PaymentRecordStatusUnknown}, invalid: PaymentRecordStatus("__invalid__")},
		{name: "PaymentStatus", valid: []stringEnum{PaymentStatusUnknown, PaymentStatusUnpaid, PaymentStatusPending, PaymentStatusPaid, PaymentStatusPartiallyPaid, PaymentStatusRefunded, PaymentStatusPartialRefunded}, invalid: PaymentStatus("__invalid__")},
		{name: "PickingItemStatus", valid: []stringEnum{PickingItemStatusPending, PickingItemStatusPartial, PickingItemStatusComplete, PickingItemStatusSkipped}, invalid: PickingItemStatus("__invalid__")},
		{name: "PickingListStatus", valid: []stringEnum{PickingListStatusPending, PickingListStatusInProgress, PickingListStatusComplete, PickingListStatusCancelled}, invalid: PickingListStatus("__invalid__")},
		{name: "Portal", valid: []stringEnum{PortalControl, PortalStore, PortalPartner}, invalid: Portal("__invalid__")},
		{name: "PromotionAddonTrigger", valid: []stringEnum{PromotionAddonTriggerAmount, PromotionAddonTriggerRequiredProducts}, invalid: PromotionAddonTrigger("__invalid__")},
		{name: "PromotionClass", valid: []stringEnum{PromotionClassNormal, PromotionClassSpecialCampaign}, invalid: PromotionClass("__invalid__")},
		{name: "PromotionDiscountTarget", valid: []stringEnum{PromotionDiscountTargetCart, PromotionDiscountTargetRequiredItems}, invalid: PromotionDiscountTarget("__invalid__")},
		{name: "PromotionQtyMode", valid: []stringEnum{PromotionQtyModePerProduct, PromotionQtyModeCombined}, invalid: PromotionQtyMode("__invalid__")},
		{name: "PromotionType", valid: []stringEnum{PromotionTypeAutoDiscount, PromotionTypeSpendGift, PromotionTypeAddonPurchase, PromotionTypeBOGO, PromotionTypeBundle, PromotionTypeTieredPricing}, invalid: PromotionType("__invalid__")},
		{name: "PurchaseOrderStatus", valid: []stringEnum{PurchaseOrderStatusDraft, PurchaseOrderStatusSubmitted, PurchaseOrderStatusConfirmed, PurchaseOrderStatusPartiallyReceived, PurchaseOrderStatusReceived, PurchaseOrderStatusCancelled, PurchaseOrderStatusRefunded}, invalid: PurchaseOrderStatus("__invalid__")},
		{name: "RecoveryDecision", valid: []stringEnum{RecoveryDecisionPending, RecoveryDecisionApproved, RecoveryDecisionDeclined}, invalid: RecoveryDecision("__invalid__")},
		{name: "SalesOrderStatus", valid: []stringEnum{SalesOrderStatusPending, SalesOrderStatusConfirmed, SalesOrderStatusPaid, SalesOrderStatusProcessing, SalesOrderStatusPicking, SalesOrderStatusPacked, SalesOrderStatusShipped, SalesOrderStatusDelivered, SalesOrderStatusCompleted, SalesOrderStatusCancelled, SalesOrderStatusRefunded}, invalid: SalesOrderStatus("__invalid__")},
		{name: "SecurityEventSeverity", valid: []stringEnum{SecurityEventSeverityInfo, SecurityEventSeverityLow, SecurityEventSeverityMedium, SecurityEventSeverityHigh, SecurityEventSeverityCritical}, invalid: SecurityEventSeverity("__invalid__")},
		{name: "SecurityEventStatus", valid: []stringEnum{SecurityEventStatusDetected, SecurityEventStatusTriaged, SecurityEventStatusInvestigating, SecurityEventStatusContained, SecurityEventStatusResolved, SecurityEventStatusFalsePositive}, invalid: SecurityEventStatus("__invalid__")},
		{name: "SecurityRiskLevel", valid: []stringEnum{SecurityRiskLevelLow, SecurityRiskLevelMedium, SecurityRiskLevelHigh, SecurityRiskLevelCritical}, invalid: SecurityRiskLevel("__invalid__")},
		{name: "SettlementType", valid: []stringEnum{SettlementTypeSettlement, SettlementTypeEnquiry}, invalid: SettlementType("__invalid__")},
		{name: "ShapeType", valid: []stringEnum{ShapeBox, ShapeCylinder, ShapeSphere, ShapePlane, ShapeCustom}, invalid: ShapeType("__invalid__")},
		{name: "ShippingRateName", valid: []stringEnum{ShippingRateNameStandard, ShippingRateNameExpress, ShippingRateNamePickup}, invalid: ShippingRateName("__invalid__")},
		{name: "StockMovementType", valid: []stringEnum{StockMovementTypePurchaseReceipt, StockMovementTypeSaleReserve, StockMovementTypeSaleCommit, StockMovementTypeSaleRelease, StockMovementTypeAdjustment, StockMovementTypeDamage, StockMovementTypeReturn, StockMovementTypeTransferIn, StockMovementTypeTransferOut, StockMovementTypeStocktake}, invalid: StockMovementType("__invalid__")},
		{name: "StorageType", valid: []stringEnum{StorageDry, StorageChilled, StorageFrozen}, invalid: StorageType("__invalid__")},
		{name: "SubscriptionStatus", valid: []stringEnum{SubscriptionStatusActive, SubscriptionStatusPaused, SubscriptionStatusCancelled}, invalid: SubscriptionStatus("__invalid__")},
		{name: "TerminalConnectionMode", valid: []stringEnum{TerminalConnectionModeCloudSync, TerminalConnectionModeCloudAsync, TerminalConnectionModeLocal}, invalid: TerminalConnectionMode("__invalid__")},
		{name: "TerminalProvider", valid: []stringEnum{TerminalProviderMx51}, invalid: TerminalProvider("__invalid__")},
		{name: "TerminalRefundType", valid: []stringEnum{TerminalRefundTypeReferenced, TerminalRefundTypeUnreferenced}, invalid: TerminalRefundType("__invalid__")},
		{name: "TerminalStatus", valid: []stringEnum{TerminalStatusRegistered, TerminalStatusActive, TerminalStatusDeregistered, TerminalStatusExpired, TerminalStatusError}, invalid: TerminalStatus("__invalid__")},
		{name: "TerminalTxFinancialStatus", valid: []stringEnum{TerminalTxFinancialStatusApproved, TerminalTxFinancialStatusDeclined, TerminalTxFinancialStatusCancelled, TerminalTxFinancialStatusUnknown}, invalid: TerminalTxFinancialStatus("__invalid__")},
		{name: "TerminalTxStatus", valid: []stringEnum{TerminalTxStatusUnknown, TerminalTxStatusPending, TerminalTxStatusAwaitingAction, TerminalTxStatusFinalised, TerminalTxStatusOverridePending, TerminalTxStatusOverrideResolved}, invalid: TerminalTxStatus("__invalid__")},
		{name: "TerminalTxType", valid: []stringEnum{TerminalTxTypePurchase, TerminalTxTypeRefund, TerminalTxTypeReversal, TerminalTxTypeCashout, TerminalTxTypePurchaseWithCashout, TerminalTxTypeMOTO, TerminalTxTypeSettlement, TerminalTxTypeSettlementEnquiry}, invalid: TerminalTxType("__invalid__")},
		{name: "UserPreferredLanguage", valid: []stringEnum{PreferredLanguageEnglish, PreferredLanguageTraditionalChinese, PreferredLanguageSimplifiedChinese}, invalid: UserPreferredLanguage("__invalid__")},
		{name: "UserRole", valid: []stringEnum{UserRoleSuperAdmin, UserRoleAdmin, UserRoleSales, UserRoleWarehouse, UserRoleWarehouseOperator, UserRoleMarketing, UserRoleCustomer, UserRoleClient}, invalid: UserRole("__invalid__")},
		{name: "VolumeDiscountAppliesTo", valid: []stringEnum{VolumeDiscountAppliesToAll, VolumeDiscountAppliesToWholesale, VolumeDiscountAppliesToRetail}, invalid: VolumeDiscountAppliesTo("__invalid__")},
		{name: "WMSDraftStatus", valid: []stringEnum{WMSDraftStatusDraft, WMSDraftStatusSubmitted, WMSDraftStatusCancelled}, invalid: WMSDraftStatus("__invalid__")},
		{name: "WMSDraftType", valid: []stringEnum{WMSDraftTypeInbound, WMSDraftTypeOutbound}, invalid: WMSDraftType("__invalid__")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, value := range tt.valid {
				if !value.IsValid() {
					t.Fatalf("%T(%q) should be valid", value, reflect.ValueOf(value).String())
				}
				if got, want := value.String(), reflect.ValueOf(value).String(); got != want {
					t.Fatalf("%T.String() = %q, want %q", value, got, want)
				}
			}

			if tt.invalid.IsValid() {
				t.Fatalf("%T(%q) should be invalid", tt.invalid, reflect.ValueOf(tt.invalid).String())
			}
		})
	}
}

func TestSalesOrderStatusTransitions(t *testing.T) {
	allowed := []struct {
		from SalesOrderStatus
		to   SalesOrderStatus
	}{
		{SalesOrderStatusPending, SalesOrderStatusConfirmed},
		{SalesOrderStatusConfirmed, SalesOrderStatusPaid},
		{SalesOrderStatusPaid, SalesOrderStatusProcessing},
		{SalesOrderStatusProcessing, SalesOrderStatusPicking},
		{SalesOrderStatusPicking, SalesOrderStatusPacked},
		{SalesOrderStatusPacked, SalesOrderStatusShipped},
		{SalesOrderStatusShipped, SalesOrderStatusDelivered},
		{SalesOrderStatusDelivered, SalesOrderStatusCompleted},
		{SalesOrderStatusCompleted, SalesOrderStatusRefunded},
	}

	for _, transition := range allowed {
		if !transition.from.CanTransitionTo(transition.to) {
			t.Fatalf("%s should transition to %s", transition.from, transition.to)
		}
	}

	rejected := []struct {
		from SalesOrderStatus
		to   SalesOrderStatus
	}{
		{SalesOrderStatusPending, SalesOrderStatusPacked},
		{SalesOrderStatusProcessing, SalesOrderStatusCancelled},
		{SalesOrderStatusCancelled, SalesOrderStatusPending},
		{SalesOrderStatusRefunded, SalesOrderStatusPaid},
	}

	for _, transition := range rejected {
		if transition.from.CanTransitionTo(transition.to) {
			t.Fatalf("%s should not transition to %s", transition.from, transition.to)
		}
	}

	if !SalesOrderStatusCancelled.IsTerminal() || !SalesOrderStatusRefunded.IsTerminal() {
		t.Fatal("cancelled and refunded sales orders should be terminal")
	}
	if SalesOrderStatusCompleted.IsTerminal() {
		t.Fatal("completed sales order should remain refundable, not terminal")
	}
}

func TestPurchaseOrderStatusTransitions(t *testing.T) {
	allowed := []struct {
		from PurchaseOrderStatus
		to   PurchaseOrderStatus
	}{
		{PurchaseOrderStatusDraft, PurchaseOrderStatusSubmitted},
		{PurchaseOrderStatusSubmitted, PurchaseOrderStatusConfirmed},
		{PurchaseOrderStatusConfirmed, PurchaseOrderStatusPartiallyReceived},
		{PurchaseOrderStatusConfirmed, PurchaseOrderStatusReceived},
		{PurchaseOrderStatusPartiallyReceived, PurchaseOrderStatusReceived},
		{PurchaseOrderStatusReceived, PurchaseOrderStatusRefunded},
	}

	for _, transition := range allowed {
		if !transition.from.CanTransitionTo(transition.to) {
			t.Fatalf("%s should transition to %s", transition.from, transition.to)
		}
	}

	rejected := []struct {
		from PurchaseOrderStatus
		to   PurchaseOrderStatus
	}{
		{PurchaseOrderStatusDraft, PurchaseOrderStatusReceived},
		{PurchaseOrderStatusSubmitted, PurchaseOrderStatusRefunded},
		{PurchaseOrderStatusCancelled, PurchaseOrderStatusSubmitted},
		{PurchaseOrderStatusRefunded, PurchaseOrderStatusConfirmed},
	}

	for _, transition := range rejected {
		if transition.from.CanTransitionTo(transition.to) {
			t.Fatalf("%s should not transition to %s", transition.from, transition.to)
		}
	}

	if !PurchaseOrderStatusCancelled.IsTerminal() || !PurchaseOrderStatusRefunded.IsTerminal() {
		t.Fatal("cancelled and refunded purchase orders should be terminal")
	}
	if PurchaseOrderStatusReceived.IsTerminal() {
		t.Fatal("received purchase order should remain refundable, not terminal")
	}
}

func TestTerminalStatusesTerminalState(t *testing.T) {
	if !TerminalStatusDeregistered.IsTerminal() || !TerminalStatusExpired.IsTerminal() {
		t.Fatal("deregistered and expired terminals should be terminal")
	}
	if TerminalStatusError.IsTerminal() {
		t.Fatal("error terminal status should be recoverable")
	}

	if !TerminalTxStatusFinalised.IsTerminal() || !TerminalTxStatusOverrideResolved.IsTerminal() {
		t.Fatal("finalised and override_resolved terminal transactions should be terminal")
	}
	if TerminalTxStatusOverridePending.IsTerminal() {
		t.Fatal("override_pending terminal transaction should require resolution")
	}
}
