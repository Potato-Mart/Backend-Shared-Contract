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
		{name: "AccountStatus", valid: []stringEnum{AccountStatusPending, AccountStatusActive, AccountStatusSuspended, AccountStatusClosed, AccountStatusDeleted}, invalid: AccountStatus("__invalid__")},
		{name: "AlertLevel", valid: []stringEnum{AlertLevelOK, AlertLevelWarning, AlertLevelCritical, AlertLevelExpired}, invalid: AlertLevel("__invalid__")},
		{name: "AuditOutcome", valid: []stringEnum{AuditOutcomeSuccess, AuditOutcomeFailure, AuditOutcomeDenied}, invalid: AuditOutcome("__invalid__")},
		{name: "CampaignCustomerType", valid: []stringEnum{CampaignCustomerTypeGuest, CampaignCustomerTypeRetail, CampaignCustomerTypeWholesale}, invalid: CampaignCustomerType("__invalid__")},
		{name: "CampaignPlacement", valid: []stringEnum{CampaignPlacementTopBanner, CampaignPlacementHomeHero, CampaignPlacementModal, CampaignPlacementCheckoutNotice, CampaignPlacementProductNotice}, invalid: CampaignPlacement("__invalid__")},
		{name: "CampaignPlatform", valid: []stringEnum{CampaignPlatformWeb, CampaignPlatformMobile}, invalid: CampaignPlatform("__invalid__")},
		{name: "CampaignSeverity", valid: []stringEnum{CampaignSeverityInfo, CampaignSeveritySuccess, CampaignSeverityWarning, CampaignSeverityCritical}, invalid: CampaignSeverity("__invalid__")},
		{name: "MediaStatus", valid: []stringEnum{MediaStatusPending, MediaStatusActive, MediaStatusDeleted}, invalid: MediaStatus("__invalid__")},
		{name: "PackingDamageHandling", valid: []stringEnum{PackingDamageReplaceFromStock, PackingDamageShortShipRefund}, invalid: PackingDamageHandling("__invalid__")},
		{name: "PackingSessionStatus", valid: []stringEnum{PackingSessionStatusPending, PackingSessionStatusPacking, PackingSessionStatusPacked, PackingSessionStatusSyncPending, PackingSessionStatusResolved}, invalid: PackingSessionStatus("__invalid__")},
		{name: "AccountType", valid: []stringEnum{AccountTypeAdminUser, AccountTypeGeneralCustomer, AccountTypeWholesaleCustomer}, invalid: AccountType("__invalid__")},
		{name: "AuthAssuranceLevel", valid: []stringEnum{AuthAssuranceLevel1, AuthAssuranceLevel2, AuthAssuranceLevel3}, invalid: AuthAssuranceLevel("__invalid__")},
		{name: "AuthIdentityProvider", valid: []stringEnum{AuthIdentityProviderPassword, AuthIdentityProviderGoogle, AuthIdentityProviderApple, AuthIdentityProviderAzureAD, AuthIdentityProviderOkta, AuthIdentityProviderPasskey, AuthIdentityProviderServiceToken, AuthIdentityProviderLine, AuthIdentityProviderDiscord, AuthIdentityProviderMicrosoft, AuthIdentityProviderOIDC}, invalid: AuthIdentityProvider("__invalid__")},
		{name: "AuthIdentityStatus", valid: []stringEnum{AuthIdentityStatusActive, AuthIdentityStatusDisabled, AuthIdentityStatusRevoked}, invalid: AuthIdentityStatus("__invalid__")},
		{name: "AuthMethod", valid: []stringEnum{AuthMethodPassword, AuthMethodMFA, AuthMethodPasskey, AuthMethodSSO, AuthMethodRefreshToken, AuthMethodAPIKey}, invalid: AuthMethod("__invalid__")},
		{name: "CameraProjection", valid: []stringEnum{CameraPerspective, CameraOrthographic}, invalid: CameraProjection("__invalid__")},
		{name: "ChurnRisk", valid: []stringEnum{ChurnRiskLow, ChurnRiskMedium, ChurnRiskHigh}, invalid: ChurnRisk("__invalid__")},
		{name: "CouponAppliesTo", valid: []stringEnum{CouponAppliesToAll, CouponAppliesToSpecificProducts, CouponAppliesToSpecificCategories}, invalid: CouponAppliesTo("__invalid__")},
		{name: "CouponSource", valid: []stringEnum{CouponSourceManual, CouponSourceRFMComeback, CouponSourceBirthday, CouponSourceReferral, CouponSourceSignupBonus, CouponSourceCampaign}, invalid: CouponSource("__invalid__")},
		{name: "CustomerActivityType", valid: []stringEnum{CustomerActivityTypeNote, CustomerActivityTypeCall, CustomerActivityTypeEmail, CustomerActivityTypeSMS, CustomerActivityTypeLine, CustomerActivityTypeOrder, CustomerActivityTypeComplaint, CustomerActivityTypeReturn, CustomerActivityTypeRefund, CustomerActivityTypePointsAdjust, CustomerActivityTypeTierChange, CustomerActivityTypeStatusChange, CustomerActivityTypeReferral, CustomerActivityTypeCampaign}, invalid: CustomerActivityType("__invalid__")},
		{name: "CustomerAcquisitionSource", valid: []stringEnum{CustomerAcquisitionSourceOnline, CustomerAcquisitionSourcePOS, CustomerAcquisitionSourceImport, CustomerAcquisitionSourceManual, CustomerAcquisitionSourcePhone}, invalid: CustomerAcquisitionSource("__invalid__")},
		{name: "CustomerIdentityKind", valid: []stringEnum{CustomerIdentityKindPhone, CustomerIdentityKindEmail, CustomerIdentityKindLine, CustomerIdentityKindMemberCard, CustomerIdentityKindPOSID, CustomerIdentityKindExternal}, invalid: CustomerIdentityKind("__invalid__")},
		{name: "CustomerStatus", valid: []stringEnum{CustomerStatusActive, CustomerStatusInactive, CustomerStatusBlocked, CustomerStatusClosed}, invalid: CustomerStatus("__invalid__")},
		{name: "CustomerTier", valid: []stringEnum{CustomerTierStandard, CustomerTierSilver, CustomerTierGold, CustomerTierPlatinum}, invalid: CustomerTier("__invalid__")},
		{name: "DamageStage", valid: []stringEnum{DamageStageInbound, DamageStagePicking, DamageStagePacking, DamageStageStorage}, invalid: DamageStage("__invalid__")},
		{name: "DataClassification", valid: []stringEnum{DataClassificationPublic, DataClassificationInternal, DataClassificationConfidential, DataClassificationRestricted}, invalid: DataClassification("__invalid__")},
		{name: "DataProtectionBasis", valid: []stringEnum{DataProtectionBasisNotApplicable, DataProtectionBasisConsent, DataProtectionBasisContract, DataProtectionBasisLegalObligation, DataProtectionBasisLegitimateInterest}, invalid: DataProtectionBasis("__invalid__")},
		{name: "DeviceType", valid: []stringEnum{DeviceTypeDesktop, DeviceTypeMobile, DeviceTypeTablet, DeviceTypeAPI}, invalid: DeviceType("__invalid__")},
		{name: "DiscountScope", valid: []stringEnum{DiscountScopeAll, DiscountScopeCategory, DiscountScopeProduct}, invalid: DiscountScope("__invalid__")},
		{name: "DiscountType", valid: []stringEnum{DiscountTypePercentage, DiscountTypeFixedAmount, DiscountTypeFreeShipping, DiscountTypeFixedPrice}, invalid: DiscountType("__invalid__")},
		{name: "FulfillmentStatus", valid: []stringEnum{FulfillmentStatusUnfulfilled, FulfillmentStatusPickingPrinted, FulfillmentStatusPacking, FulfillmentStatusPacked, FulfillmentStatusPartial, FulfillmentStatusFulfilled}, invalid: FulfillmentStatus("__invalid__")},
		{name: "InboundReceiptStatus", valid: []stringEnum{InboundReceiptStatusDraft, InboundReceiptStatusConfirmed}, invalid: InboundReceiptStatus("__invalid__")},
		{name: "IdentityDomain", valid: []stringEnum{IdentityDomainCustomer, IdentityDomainWorkforce, IdentityDomainPartner, IdentityDomainService}, invalid: IdentityDomain("__invalid__")},
		{name: "LayoutNodeType", valid: []stringEnum{LayoutNodeZone, LayoutNodeAisle, LayoutNodeRack, LayoutNodeShelf, LayoutNodeBin}, invalid: LayoutNodeType("__invalid__")},
		{name: "MarketingCampaignStatus", valid: []stringEnum{MarketingCampaignStatusDraft, MarketingCampaignStatusSending, MarketingCampaignStatusSent, MarketingCampaignStatusPartial, MarketingCampaignStatusFailed, MarketingCampaignStatusCancelled, MarketingCampaignStatusExported}, invalid: MarketingCampaignStatus("__invalid__")},
		{name: "MarketingChannel", valid: []stringEnum{MarketingChannelEmail, MarketingChannelSMS, MarketingChannelLine, MarketingChannelExport}, invalid: MarketingChannel("__invalid__")},
		{name: "MarketingRecipientStatus", valid: []stringEnum{MarketingRecipientStatusPending, MarketingRecipientStatusSent, MarketingRecipientStatusDelivered, MarketingRecipientStatusBounced, MarketingRecipientStatusOpened, MarketingRecipientStatusClicked, MarketingRecipientStatusFailed, MarketingRecipientStatusUnsubscribed}, invalid: MarketingRecipientStatus("__invalid__")},
		{name: "MembershipAccountStatus", valid: []stringEnum{MembershipAccountStatusActive, MembershipAccountStatusSuspended, MembershipAccountStatusClosed}, invalid: MembershipAccountStatus("__invalid__")},
		{name: "MembershipOwnerType", valid: []stringEnum{MembershipOwnerTypeRetailCustomer, MembershipOwnerTypeWholesaleOrganisation}, invalid: MembershipOwnerType("__invalid__")},
		{name: "MembershipPointReason", valid: []stringEnum{MembershipPointReasonOrder, MembershipPointReasonBirthday, MembershipPointReasonRedeem, MembershipPointReasonRewardRedeem, MembershipPointReasonAdminAdjust, MembershipPointReasonExpired, MembershipPointReasonReferral, MembershipPointReasonSignupBonus, MembershipPointReasonTierUpgrade, MembershipPointReasonManual}, invalid: MembershipPointReason("__invalid__")},
		{name: "MembershipPromotionTarget", valid: []stringEnum{MembershipPromotionTargetAll, MembershipPromotionTargetWholesale, MembershipPromotionTargetRetail, MembershipPromotionTargetTierSpecific}, invalid: MembershipPromotionTarget("__invalid__")},
		{name: "MembershipRedemptionType", valid: []stringEnum{MembershipRedemptionTypeCheckoutDiscount, MembershipRedemptionTypeRewardCatalog}, invalid: MembershipRedemptionType("__invalid__")},
		{name: "MembershipRewardRedemptionStatus", valid: []stringEnum{MembershipRewardRedemptionStatusReserved, MembershipRewardRedemptionStatusRedeemed, MembershipRewardRedemptionStatusCancelled, MembershipRewardRedemptionStatusExpired}, invalid: MembershipRewardRedemptionStatus("__invalid__")},
		{name: "MembershipRewardType", valid: []stringEnum{MembershipRewardTypeOrderDiscount, MembershipRewardTypeProduct, MembershipRewardTypeFreeShipping, MembershipRewardTypeVoucher}, invalid: MembershipRewardType("__invalid__")},
		{name: "MembershipTierMetric", valid: []stringEnum{MembershipTierMetricAnnualSpend, MembershipTierMetricLifetimeSpend, MembershipTierMetricManual}, invalid: MembershipTierMetric("__invalid__")},
		{name: "MemberSubscriptionStatus", valid: []stringEnum{MemberSubscriptionStatusActive, MemberSubscriptionStatusPaused, MemberSubscriptionStatusCancelled}, invalid: MemberSubscriptionStatus("__invalid__")},
		{name: "ModelFormat", valid: []stringEnum{ModelFormatGLB, ModelFormatGLTF, ModelFormatOBJ, ModelFormatFBX, ModelFormatUSDZ}, invalid: ModelFormat("__invalid__")},
		{name: "OrderSourceDeviceType", valid: []stringEnum{OrderSourceDeviceTypeIOS, OrderSourceDeviceTypeAndroid, OrderSourceDeviceTypePC, OrderSourceDeviceTypeMobileWeb, OrderSourceDeviceTypeTablet, OrderSourceDeviceTypePos, OrderSourceDeviceTypeManual, OrderSourceDeviceTypePhone, OrderSourceDeviceTypeVR}, invalid: OrderSourceDeviceType("__invalid__")},
		{name: "OrderType", valid: []stringEnum{OrderTypeOnline, OrderTypePOS, OrderTypeB2B, OrderTypeRelay, OrderTypeManual, OrderTypeImport}, invalid: OrderType("__invalid__")},
		{name: "BuyerType", valid: []stringEnum{BuyerTypeGuestRetail, BuyerTypeRetailCustomer, BuyerTypeWholesaleOrganisation}, invalid: BuyerType("__invalid__")},
		{name: "PriceAudience", valid: []stringEnum{PriceAudienceRetail, PriceAudienceWholesale}, invalid: PriceAudience("__invalid__")},
		{name: "PriceVisibility", valid: []stringEnum{PriceVisibilityPublic, PriceVisibilityLoginRequired, PriceVisibilityWholesaleApprovedOnly, PriceVisibilityHidden}, invalid: PriceVisibility("__invalid__")},
		{name: "FulfilmentIntent", valid: []stringEnum{FulfilmentIntentDelivery, FulfilmentIntentPickup, FulfilmentIntentInStoreCarry}, invalid: FulfilmentIntent("__invalid__")},
		{name: "OrganisationAccessStatus", valid: []stringEnum{OrganisationAccessStatusPending, OrganisationAccessStatusActive, OrganisationAccessStatusSuspended, OrganisationAccessStatusRevoked}, invalid: OrganisationAccessStatus("__invalid__")},
		{name: "WholesaleBuyerRole", valid: []stringEnum{WholesaleBuyerRoleOwner, WholesaleBuyerRoleBuyer, WholesaleBuyerRoleFinance, WholesaleBuyerRoleReadOnly}, invalid: WholesaleBuyerRole("__invalid__")},
		{name: "WholesalePermission", valid: []stringEnum{WholesalePermissionCatalogueView, WholesalePermissionCartWrite, WholesalePermissionCheckoutSubmit, WholesalePermissionOrdersViewOwn, WholesalePermissionOrdersViewOrg, WholesalePermissionOrdersReorder, WholesalePermissionInvoicesViewOwn, WholesalePermissionInvoicesViewOrg, WholesalePermissionInvoicesPay, WholesalePermissionAccountView, WholesalePermissionTeamView, WholesalePermissionFavouritesWrite, WholesalePermissionOrderListsViewOwn, WholesalePermissionOrderListsWriteOwn, WholesalePermissionOrderListsViewOrg, WholesalePermissionOrderListsWriteOrg}, invalid: WholesalePermission("__invalid__")},
		{name: "OutboundShipmentStatus", valid: []stringEnum{OutboundShipmentStatusPacked, OutboundShipmentStatusDispatched}, invalid: OutboundShipmentStatus("__invalid__")},
		{name: "PackingDiscrepancyKind", valid: []stringEnum{PackingDiscrepancyKindShortage, PackingDiscrepancyKindOverweight, PackingDiscrepancyKindDamaged, PackingDiscrepancyKindPending}, invalid: PackingDiscrepancyKind("__invalid__")},
		{name: "PaymentMethod", valid: []stringEnum{PaymentMethodCard, PaymentMethodCash, PaymentMethodQR, PaymentMethodBankTransfer, PaymentMethodLinePay, PaymentMethodApplePay, PaymentMethodGooglePay, PaymentMethodECPay, PaymentMethodManual, PaymentMethodEFTPOS, PaymentMethodMOTO, PaymentMethodCashout}, invalid: PaymentMethod("__invalid__")},
		{name: "PaymentRecordStatus", valid: []stringEnum{PaymentRecordStatusPending, PaymentRecordStatusProcessing, PaymentRecordStatusCompleted, PaymentRecordStatusFailed, PaymentRecordStatusCancelled, PaymentRecordStatusRefunded, PaymentRecordStatusAwaitingAction, PaymentRecordStatusUnknown}, invalid: PaymentRecordStatus("__invalid__")},
		{name: "PaymentStatus", valid: []stringEnum{PaymentStatusUnknown, PaymentStatusUnpaid, PaymentStatusPending, PaymentStatusPaid, PaymentStatusPartiallyPaid, PaymentStatusRefunded, PaymentStatusPartialRefunded}, invalid: PaymentStatus("__invalid__")},
		{name: "PickingItemStatus", valid: []stringEnum{PickingItemStatusPending, PickingItemStatusPartial, PickingItemStatusComplete, PickingItemStatusSkipped}, invalid: PickingItemStatus("__invalid__")},
		{name: "PickingListStatus", valid: []stringEnum{PickingListStatusPending, PickingListStatusInProgress, PickingListStatusComplete, PickingListStatusCancelled}, invalid: PickingListStatus("__invalid__")},
		{name: "PointReservationStatus", valid: []stringEnum{PointReservationStatusReserved, PointReservationStatusCommitted, PointReservationStatusCancelled, PointReservationStatusExpired}, invalid: PointReservationStatus("__invalid__")},
		{name: "Portal", valid: []stringEnum{PortalControl, PortalStore, PortalPartner}, invalid: Portal("__invalid__")},
		{name: "PortalAccessStatus", valid: []stringEnum{PortalAccessStatusPending, PortalAccessStatusActive, PortalAccessStatusSuspended, PortalAccessStatusRevoked}, invalid: PortalAccessStatus("__invalid__")},
		{name: "PromotionAddonTrigger", valid: []stringEnum{PromotionAddonTriggerAmount, PromotionAddonTriggerRequiredProducts}, invalid: PromotionAddonTrigger("__invalid__")},
		{name: "PromotionClass", valid: []stringEnum{PromotionClassNormal, PromotionClassSpecialCampaign}, invalid: PromotionClass("__invalid__")},
		{name: "PromotionDiscountTarget", valid: []stringEnum{PromotionDiscountTargetCart, PromotionDiscountTargetRequiredItems}, invalid: PromotionDiscountTarget("__invalid__")},
		{name: "PromotionQtyMode", valid: []stringEnum{PromotionQtyModePerProduct, PromotionQtyModeCombined}, invalid: PromotionQtyMode("__invalid__")},
		{name: "PromotionType", valid: []stringEnum{PromotionTypeAutoDiscount, PromotionTypeSpendGift, PromotionTypeAddonPurchase, PromotionTypeBOGO, PromotionTypeBundle, PromotionTypeTieredPricing}, invalid: PromotionType("__invalid__")},
		{name: "ProductStatus", valid: []stringEnum{ProductStatusDraft, ProductStatusActive, ProductStatusArchived, ProductStatusDiscontinued}, invalid: ProductStatus("__invalid__")},
		{name: "PurchaseOrderStatus", valid: []stringEnum{PurchaseOrderStatusDraft, PurchaseOrderStatusSubmitted, PurchaseOrderStatusConfirmed, PurchaseOrderStatusPartiallyReceived, PurchaseOrderStatusReceived, PurchaseOrderStatusCancelled, PurchaseOrderStatusRefunded}, invalid: PurchaseOrderStatus("__invalid__")},
		{name: "RecoveryDecision", valid: []stringEnum{RecoveryDecisionPending, RecoveryDecisionApproved, RecoveryDecisionDeclined}, invalid: RecoveryDecision("__invalid__")},
		{name: "SalesPerformance", valid: []stringEnum{SalesPerformanceHot, SalesPerformanceNormal, SalesPerformanceSlow}, invalid: SalesPerformance("__invalid__")},
		{name: "SalesOrderStatus", valid: []stringEnum{SalesOrderStatusPending, SalesOrderStatusConfirmed, SalesOrderStatusPaid, SalesOrderStatusProcessing, SalesOrderStatusPicking, SalesOrderStatusPacked, SalesOrderStatusShipped, SalesOrderStatusDelivered, SalesOrderStatusCompleted, SalesOrderStatusCancelled, SalesOrderStatusRefunded}, invalid: SalesOrderStatus("__invalid__")},
		{name: "SecurityEventSeverity", valid: []stringEnum{SecurityEventSeverityInfo, SecurityEventSeverityLow, SecurityEventSeverityMedium, SecurityEventSeverityHigh, SecurityEventSeverityCritical}, invalid: SecurityEventSeverity("__invalid__")},
		{name: "SecurityEventStatus", valid: []stringEnum{SecurityEventStatusDetected, SecurityEventStatusTriaged, SecurityEventStatusInvestigating, SecurityEventStatusContained, SecurityEventStatusResolved, SecurityEventStatusFalsePositive}, invalid: SecurityEventStatus("__invalid__")},
		{name: "SecurityRiskLevel", valid: []stringEnum{SecurityRiskLevelLow, SecurityRiskLevelMedium, SecurityRiskLevelHigh, SecurityRiskLevelCritical}, invalid: SecurityRiskLevel("__invalid__")},
		{name: "SettlementType", valid: []stringEnum{SettlementTypeSettlement, SettlementTypeEnquiry}, invalid: SettlementType("__invalid__")},
		{name: "ShapeType", valid: []stringEnum{ShapeBox, ShapeCylinder, ShapeSphere, ShapePlane, ShapeCustom}, invalid: ShapeType("__invalid__")},
		{name: "ShippingRateName", valid: []stringEnum{ShippingRateNameStandard, ShippingRateNameExpress, ShippingRateNamePickup}, invalid: ShippingRateName("__invalid__")},
		{name: "StockMovementType", valid: []stringEnum{StockMovementTypePurchaseReceipt, StockMovementTypeSaleReserve, StockMovementTypeSaleCommit, StockMovementTypeSaleRelease, StockMovementTypeAdjustment, StockMovementTypeDamage, StockMovementTypeReturn, StockMovementTypeTransferIn, StockMovementTypeTransferOut, StockMovementTypeStocktake}, invalid: StockMovementType("__invalid__")},
		{name: "StorageType", valid: []stringEnum{StorageDry, StorageChilled, StorageFrozen}, invalid: StorageType("__invalid__")},
		{name: "TerminalConnectionMode", valid: []stringEnum{TerminalConnectionModeCloudSync, TerminalConnectionModeCloudAsync, TerminalConnectionModeLocal}, invalid: TerminalConnectionMode("__invalid__")},
		{name: "TerminalProvider", valid: []stringEnum{TerminalProviderMx51}, invalid: TerminalProvider("__invalid__")},
		{name: "TerminalRefundType", valid: []stringEnum{TerminalRefundTypeReferenced, TerminalRefundTypeUnreferenced}, invalid: TerminalRefundType("__invalid__")},
		{name: "TerminalStatus", valid: []stringEnum{TerminalStatusRegistered, TerminalStatusActive, TerminalStatusDeregistered, TerminalStatusExpired, TerminalStatusError}, invalid: TerminalStatus("__invalid__")},
		{name: "TerminalTxFinancialStatus", valid: []stringEnum{TerminalTxFinancialStatusApproved, TerminalTxFinancialStatusDeclined, TerminalTxFinancialStatusCancelled, TerminalTxFinancialStatusUnknown}, invalid: TerminalTxFinancialStatus("__invalid__")},
		{name: "TerminalTxStatus", valid: []stringEnum{TerminalTxStatusUnknown, TerminalTxStatusPending, TerminalTxStatusAwaitingAction, TerminalTxStatusFinalised, TerminalTxStatusOverridePending, TerminalTxStatusOverrideResolved}, invalid: TerminalTxStatus("__invalid__")},
		{name: "TerminalTxType", valid: []stringEnum{TerminalTxTypePurchase, TerminalTxTypeRefund, TerminalTxTypeReversal, TerminalTxTypeCashout, TerminalTxTypePurchaseWithCashout, TerminalTxTypeMOTO, TerminalTxTypeSettlement, TerminalTxTypeSettlementEnquiry}, invalid: TerminalTxType("__invalid__")},
		{name: "UserPreferredLanguage", valid: []stringEnum{PreferredLanguageEnglish, PreferredLanguageTraditionalChinese, PreferredLanguageSimplifiedChinese}, invalid: UserPreferredLanguage("__invalid__")},
		{name: "UserRole", valid: []stringEnum{UserRoleSuperAdmin, UserRoleAdmin, UserRoleSales, UserRoleWarehouse, UserRoleWarehouseOperator, UserRoleMarketing, UserRoleCustomer}, invalid: UserRole("__invalid__")},
		{name: "VolumeDiscountAppliesTo", valid: []stringEnum{VolumeDiscountAppliesToAll, VolumeDiscountAppliesToWholesale, VolumeDiscountAppliesToRetail}, invalid: VolumeDiscountAppliesTo("__invalid__")},
		{name: "WholesaleOrganisationStatus", valid: []stringEnum{WholesaleOrganisationStatusPending, WholesaleOrganisationStatusApproved, WholesaleOrganisationStatusSuspended, WholesaleOrganisationStatusRejected, WholesaleOrganisationStatusClosed}, invalid: WholesaleOrganisationStatus("__invalid__")},
		{name: "WMSDraftStatus", valid: []stringEnum{WMSDraftStatusDraft, WMSDraftStatusSubmitted, WMSDraftStatusCancelled}, invalid: WMSDraftStatus("__invalid__")},
		{name: "WMSDraftType", valid: []stringEnum{WMSDraftTypeInbound, WMSDraftTypeOutbound}, invalid: WMSDraftType("__invalid__")},
		{name: "WalletInstrumentType", valid: []stringEnum{WalletInstrumentTypePoints, WalletInstrumentTypeGiftCard, WalletInstrumentTypeVoucher, WalletInstrumentTypeCoupon, WalletInstrumentTypeReward}, invalid: WalletInstrumentType("__invalid__")},
		{name: "GiftCardStatus", valid: []stringEnum{GiftCardStatusActive, GiftCardStatusPartiallyRedeemed, GiftCardStatusDepleted, GiftCardStatusExpired, GiftCardStatusVoid}, invalid: GiftCardStatus("__invalid__")},
		{name: "GiftCardTransactionReason", valid: []stringEnum{GiftCardTransactionReasonIssue, GiftCardTransactionReasonRedeem, GiftCardTransactionReasonRefund, GiftCardTransactionReasonTopUp, GiftCardTransactionReasonExpire, GiftCardTransactionReasonAdjust}, invalid: GiftCardTransactionReason("__invalid__")},
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

func TestAccountTypePortalAdmission(t *testing.T) {
	allowed := map[AccountType]Portal{
		AccountTypeAdminUser:         PortalControl,
		AccountTypeGeneralCustomer:   PortalStore,
		AccountTypeWholesaleCustomer: PortalPartner,
	}

	for accountType, portal := range allowed {
		if !accountType.IsAllowedInPortal(portal) {
			t.Fatalf("%s should be allowed in %s", accountType, portal)
		}
	}

	rejected := []struct {
		accountType AccountType
		portal      Portal
	}{
		{AccountTypeAdminUser, PortalStore},
		{AccountTypeAdminUser, PortalPartner},
		{AccountTypeGeneralCustomer, PortalControl},
		{AccountTypeGeneralCustomer, PortalPartner},
		{AccountTypeWholesaleCustomer, PortalControl},
		{AccountTypeWholesaleCustomer, PortalStore},
	}

	for _, tt := range rejected {
		if tt.accountType.IsAllowedInPortal(tt.portal) {
			t.Fatalf("%s should not be allowed in %s", tt.accountType, tt.portal)
		}
	}
}

func TestPortalAccountTypeHelpers(t *testing.T) {
	tests := []struct {
		portal      Portal
		accountType AccountType
	}{
		{PortalControl, AccountTypeAdminUser},
		{PortalStore, AccountTypeGeneralCustomer},
		{PortalPartner, AccountTypeWholesaleCustomer},
	}

	for _, tt := range tests {
		if !tt.portal.RequiresAccountType(tt.accountType) {
			t.Fatalf("%s should require %s", tt.portal, tt.accountType)
		}

		required, ok := tt.portal.RequiredAccountType()
		if !ok {
			t.Fatalf("%s should have a required account type", tt.portal)
		}
		if required != tt.accountType {
			t.Fatalf("%s required account type = %s, want %s", tt.portal, required, tt.accountType)
		}

		want := []AccountType{tt.accountType}
		if got := tt.portal.AllowedAccountTypes(); !reflect.DeepEqual(got, want) {
			t.Fatalf("%s allowed account types = %#v, want %#v", tt.portal, got, want)
		}
		if got := AccountTypesForPortal(tt.portal); !reflect.DeepEqual(got, want) {
			t.Fatalf("AccountTypesForPortal(%s) = %#v, want %#v", tt.portal, got, want)
		}
	}

	if _, ok := Portal("__invalid__").RequiredAccountType(); ok {
		t.Fatal("invalid portal should not have a required account type")
	}
	if got := Portal("__invalid__").AllowedAccountTypes(); got != nil {
		t.Fatalf("invalid portal allowed account types = %#v, want nil", got)
	}
}

func TestPortalAccessStatusCanAccess(t *testing.T) {
	if !PortalAccessStatusActive.CanAccess() {
		t.Fatal("active portal access should allow access")
	}

	for _, status := range []PortalAccessStatus{
		PortalAccessStatusPending,
		PortalAccessStatusSuspended,
		PortalAccessStatusRevoked,
	} {
		if status.CanAccess() {
			t.Fatalf("%s should not allow access", status)
		}
	}
}

func TestWholesaleBuyerRolePermissions(t *testing.T) {
	ownerPerms := PermissionsForWholesaleBuyerRole(WholesaleBuyerRoleOwner)
	if len(ownerPerms) == 0 {
		t.Fatal("owner should receive permissions")
	}
	if !HasWholesalePermission(WholesalePermissionStrings(ownerPerms), WholesalePermissionTeamView) {
		t.Fatal("owner should receive team.view")
	}
	if !HasWholesalePermission(WholesalePermissionStrings(ownerPerms), WholesalePermissionOrderListsWriteOrg) {
		t.Fatal("owner should receive order_lists.write_org")
	}

	buyerPerms := WholesalePermissionStrings(PermissionsForWholesaleBuyerRole(WholesaleBuyerRoleBuyer))
	if HasWholesalePermission(buyerPerms, WholesalePermissionInvoicesViewOrg) {
		t.Fatal("buyer should not receive organisation invoice visibility")
	}
	if !HasWholesalePermission(buyerPerms, WholesalePermissionCheckoutSubmit) {
		t.Fatal("buyer should receive checkout.submit")
	}
	if !HasWholesalePermission(buyerPerms, WholesalePermissionFavouritesWrite) ||
		!HasWholesalePermission(buyerPerms, WholesalePermissionOrderListsWriteOwn) ||
		!HasWholesalePermission(buyerPerms, WholesalePermissionOrderListsViewOrg) {
		t.Fatal("buyer should receive own list mutation, org list view, and favourites.write")
	}
	if HasWholesalePermission(buyerPerms, WholesalePermissionOrderListsWriteOrg) {
		t.Fatal("buyer should not receive organisation list mutation")
	}

	financePerms := WholesalePermissionStrings(PermissionsForWholesaleBuyerRole(WholesaleBuyerRoleFinance))
	if !HasWholesalePermission(financePerms, WholesalePermissionInvoicesPay) {
		t.Fatal("finance should receive invoices.pay")
	}
	if HasWholesalePermission(financePerms, WholesalePermissionCheckoutSubmit) {
		t.Fatal("finance should not receive checkout.submit")
	}
	if HasWholesalePermission(financePerms, WholesalePermissionFavouritesWrite) ||
		HasWholesalePermission(financePerms, WholesalePermissionOrderListsViewOwn) ||
		HasWholesalePermission(financePerms, WholesalePermissionOrderListsViewOrg) {
		t.Fatal("finance should not receive procurement list permissions")
	}

	readOnlyPerms := WholesalePermissionStrings(PermissionsForWholesaleBuyerRole(WholesaleBuyerRoleReadOnly))
	if !HasWholesalePermission(readOnlyPerms, WholesalePermissionOrderListsViewOrg) {
		t.Fatal("read-only should receive organisation list view")
	}
	if HasWholesalePermission(readOnlyPerms, WholesalePermissionFavouritesWrite) ||
		HasWholesalePermission(readOnlyPerms, WholesalePermissionOrderListsWriteOwn) ||
		HasWholesalePermission(readOnlyPerms, WholesalePermissionCartWrite) {
		t.Fatal("read-only should not receive procurement mutation or cart permissions")
	}

	if got := PermissionsForWholesaleBuyerRole(WholesaleBuyerRole("__invalid__")); got != nil {
		t.Fatalf("invalid role permissions = %#v, want nil", got)
	}
	if HasWholesalePermission([]string{WholesalePermissionTeamView.String()}, WholesalePermission("__invalid__")) {
		t.Fatal("invalid required permission should never match")
	}
}

func TestAccountStatusTerminalState(t *testing.T) {
	for _, status := range []AccountStatus{AccountStatusClosed, AccountStatusDeleted} {
		if !status.IsTerminal() {
			t.Fatalf("%s should be terminal", status)
		}
	}

	for _, status := range []AccountStatus{
		AccountStatusPending,
		AccountStatusActive,
		AccountStatusSuspended,
	} {
		if status.IsTerminal() {
			t.Fatalf("%s should not be terminal", status)
		}
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
