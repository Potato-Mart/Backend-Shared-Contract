package enums_test

import (
	"reflect"
	"testing"

	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/account"
	campaignenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/campaign"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/customer"
	identityenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/identity"
	marketingenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/marketing"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/membership"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/payment"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/product"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/promotion"
	purchaseenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/purchase"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/sales"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/security"
	shippingenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/shipping"
	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/wallet"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/warehouse"
	wholesaleenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/wholesale"
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
		{name: "accountenum.AccountStatus", valid: []stringEnum{accountenum.AccountStatusPending, accountenum.AccountStatusActive, accountenum.AccountStatusSuspended, accountenum.AccountStatusClosed, accountenum.AccountStatusDeleted}, invalid: accountenum.AccountStatus("__invalid__")},
		{name: "securityenum.AlertLevel", valid: []stringEnum{securityenum.AlertLevelOK, securityenum.AlertLevelWarning, securityenum.AlertLevelCritical, securityenum.AlertLevelExpired}, invalid: securityenum.AlertLevel("__invalid__")},
		{name: "securityenum.AuditOutcome", valid: []stringEnum{securityenum.AuditOutcomeSuccess, securityenum.AuditOutcomeFailure, securityenum.AuditOutcomeDenied}, invalid: securityenum.AuditOutcome("__invalid__")},
		{name: "campaignenum.CampaignCustomerType", valid: []stringEnum{campaignenum.CampaignCustomerTypeGuest, campaignenum.CampaignCustomerTypeRetail, campaignenum.CampaignCustomerTypeWholesale}, invalid: campaignenum.CampaignCustomerType("__invalid__")},
		{name: "campaignenum.CampaignPlacement", valid: []stringEnum{campaignenum.CampaignPlacementTopBanner, campaignenum.CampaignPlacementHomeHero, campaignenum.CampaignPlacementModal, campaignenum.CampaignPlacementCheckoutNotice, campaignenum.CampaignPlacementProductNotice}, invalid: campaignenum.CampaignPlacement("__invalid__")},
		{name: "campaignenum.CampaignPlatform", valid: []stringEnum{campaignenum.CampaignPlatformWeb, campaignenum.CampaignPlatformMobile}, invalid: campaignenum.CampaignPlatform("__invalid__")},
		{name: "campaignenum.CampaignSeverity", valid: []stringEnum{campaignenum.CampaignSeverityInfo, campaignenum.CampaignSeveritySuccess, campaignenum.CampaignSeverityWarning, campaignenum.CampaignSeverityCritical}, invalid: campaignenum.CampaignSeverity("__invalid__")},
		{name: "productenum.MediaStatus", valid: []stringEnum{productenum.MediaStatusPending, productenum.MediaStatusActive, productenum.MediaStatusDeleted}, invalid: productenum.MediaStatus("__invalid__")},
		{name: "warehouseenum.PackingDamageHandling", valid: []stringEnum{warehouseenum.PackingDamageReplaceFromStock, warehouseenum.PackingDamageShortShipRefund}, invalid: warehouseenum.PackingDamageHandling("__invalid__")},
		{name: "warehouseenum.PackingSessionStatus", valid: []stringEnum{warehouseenum.PackingSessionStatusPending, warehouseenum.PackingSessionStatusPacking, warehouseenum.PackingSessionStatusPacked, warehouseenum.PackingSessionStatusSyncPending, warehouseenum.PackingSessionStatusResolved}, invalid: warehouseenum.PackingSessionStatus("__invalid__")},
		{name: "accountenum.AccountType", valid: []stringEnum{accountenum.AccountTypeAdminUser, accountenum.AccountTypeGeneralCustomer, accountenum.AccountTypeWholesaleCustomer}, invalid: accountenum.AccountType("__invalid__")},
		{name: "securityenum.AuthAssuranceLevel", valid: []stringEnum{securityenum.AuthAssuranceLevel1, securityenum.AuthAssuranceLevel2, securityenum.AuthAssuranceLevel3}, invalid: securityenum.AuthAssuranceLevel("__invalid__")},
		{name: "identityenum.AuthIdentityProvider", valid: []stringEnum{identityenum.AuthIdentityProviderPassword, identityenum.AuthIdentityProviderGoogle, identityenum.AuthIdentityProviderApple, identityenum.AuthIdentityProviderAzureAD, identityenum.AuthIdentityProviderOkta, identityenum.AuthIdentityProviderPasskey, identityenum.AuthIdentityProviderServiceToken, identityenum.AuthIdentityProviderLine, identityenum.AuthIdentityProviderDiscord, identityenum.AuthIdentityProviderMicrosoft, identityenum.AuthIdentityProviderOIDC}, invalid: identityenum.AuthIdentityProvider("__invalid__")},
		{name: "identityenum.AuthIdentityStatus", valid: []stringEnum{identityenum.AuthIdentityStatusActive, identityenum.AuthIdentityStatusDisabled, identityenum.AuthIdentityStatusRevoked}, invalid: identityenum.AuthIdentityStatus("__invalid__")},
		{name: "securityenum.AuthMethod", valid: []stringEnum{securityenum.AuthMethodPassword, securityenum.AuthMethodMFA, securityenum.AuthMethodPasskey, securityenum.AuthMethodSSO, securityenum.AuthMethodRefreshToken, securityenum.AuthMethodAPIKey}, invalid: securityenum.AuthMethod("__invalid__")},
		{name: "warehouseenum.CameraProjection", valid: []stringEnum{warehouseenum.CameraPerspective, warehouseenum.CameraOrthographic}, invalid: warehouseenum.CameraProjection("__invalid__")},
		{name: "customerenum.ChurnRisk", valid: []stringEnum{customerenum.ChurnRiskLow, customerenum.ChurnRiskMedium, customerenum.ChurnRiskHigh}, invalid: customerenum.ChurnRisk("__invalid__")},
		{name: "promotionenum.CouponAppliesTo", valid: []stringEnum{promotionenum.CouponAppliesToAll, promotionenum.CouponAppliesToSpecificProducts, promotionenum.CouponAppliesToSpecificCategoryTags}, invalid: promotionenum.CouponAppliesTo("__invalid__")},
		{name: "promotionenum.CouponSource", valid: []stringEnum{promotionenum.CouponSourceManual, promotionenum.CouponSourceRFMComeback, promotionenum.CouponSourceBirthday, promotionenum.CouponSourceReferral, promotionenum.CouponSourceSignupBonus, promotionenum.CouponSourceCampaign}, invalid: promotionenum.CouponSource("__invalid__")},
		{name: "customerenum.CustomerActivityType", valid: []stringEnum{customerenum.CustomerActivityTypeNote, customerenum.CustomerActivityTypeCall, customerenum.CustomerActivityTypeEmail, customerenum.CustomerActivityTypeSMS, customerenum.CustomerActivityTypeLine, customerenum.CustomerActivityTypeOrder, customerenum.CustomerActivityTypeComplaint, customerenum.CustomerActivityTypeReturn, customerenum.CustomerActivityTypeRefund, customerenum.CustomerActivityTypePointsAdjust, customerenum.CustomerActivityTypeTierChange, customerenum.CustomerActivityTypeStatusChange, customerenum.CustomerActivityTypeReferral, customerenum.CustomerActivityTypeCampaign}, invalid: customerenum.CustomerActivityType("__invalid__")},
		{name: "customerenum.CustomerAcquisitionSource", valid: []stringEnum{customerenum.CustomerAcquisitionSourceOnline, customerenum.CustomerAcquisitionSourcePOS, customerenum.CustomerAcquisitionSourceImport, customerenum.CustomerAcquisitionSourceManual, customerenum.CustomerAcquisitionSourcePhone}, invalid: customerenum.CustomerAcquisitionSource("__invalid__")},
		{name: "customerenum.CustomerIdentityKind", valid: []stringEnum{customerenum.CustomerIdentityKindPhone, customerenum.CustomerIdentityKindEmail, customerenum.CustomerIdentityKindLine, customerenum.CustomerIdentityKindMemberCard, customerenum.CustomerIdentityKindPOSID, customerenum.CustomerIdentityKindExternal}, invalid: customerenum.CustomerIdentityKind("__invalid__")},
		{name: "customerenum.CustomerStatus", valid: []stringEnum{customerenum.CustomerStatusActive, customerenum.CustomerStatusInactive, customerenum.CustomerStatusBlocked, customerenum.CustomerStatusClosed}, invalid: customerenum.CustomerStatus("__invalid__")},
		{name: "customerenum.CustomerTier", valid: []stringEnum{customerenum.CustomerTierStandard, customerenum.CustomerTierSilver, customerenum.CustomerTierGold, customerenum.CustomerTierPlatinum}, invalid: customerenum.CustomerTier("__invalid__")},
		{name: "warehouseenum.DamageStage", valid: []stringEnum{warehouseenum.DamageStageInbound, warehouseenum.DamageStagePicking, warehouseenum.DamageStagePacking, warehouseenum.DamageStageStorage}, invalid: warehouseenum.DamageStage("__invalid__")},
		{name: "securityenum.DataClassification", valid: []stringEnum{securityenum.DataClassificationPublic, securityenum.DataClassificationInternal, securityenum.DataClassificationConfidential, securityenum.DataClassificationRestricted}, invalid: securityenum.DataClassification("__invalid__")},
		{name: "securityenum.DataProtectionBasis", valid: []stringEnum{securityenum.DataProtectionBasisNotApplicable, securityenum.DataProtectionBasisConsent, securityenum.DataProtectionBasisContract, securityenum.DataProtectionBasisLegalObligation, securityenum.DataProtectionBasisLegitimateInterest}, invalid: securityenum.DataProtectionBasis("__invalid__")},
		{name: "shippingenum.DeliveryMethod", valid: []stringEnum{shippingenum.DeliveryMethodDelivery, shippingenum.DeliveryMethodPickup, shippingenum.DeliveryMethodOutsourced}, invalid: shippingenum.DeliveryMethod("__invalid__")},
		{name: "shippingenum.DeliveryRegion", valid: []stringEnum{shippingenum.DeliveryRegionLocalMelbourne, shippingenum.DeliveryRegionRegionalVIC, shippingenum.DeliveryRegionInterstate}, invalid: shippingenum.DeliveryRegion("__invalid__")},
		{name: "identityenum.DeviceType", valid: []stringEnum{identityenum.DeviceTypeDesktop, identityenum.DeviceTypeMobile, identityenum.DeviceTypeTablet, identityenum.DeviceTypeAPI}, invalid: identityenum.DeviceType("__invalid__")},
		{name: "promotionenum.DiscountScope", valid: []stringEnum{promotionenum.DiscountScopeAll, promotionenum.DiscountScopeCategoryTag, promotionenum.DiscountScopeProduct}, invalid: promotionenum.DiscountScope("__invalid__")},
		{name: "promotionenum.DiscountType", valid: []stringEnum{promotionenum.DiscountTypePercentage, promotionenum.DiscountTypeFixedAmount, promotionenum.DiscountTypeFreeShipping, promotionenum.DiscountTypeFixedPrice}, invalid: promotionenum.DiscountType("__invalid__")},
		{name: "salesenum.FulfillmentStatus", valid: []stringEnum{salesenum.FulfillmentStatusUnfulfilled, salesenum.FulfillmentStatusPickingPrinted, salesenum.FulfillmentStatusPacking, salesenum.FulfillmentStatusPacked, salesenum.FulfillmentStatusPartial, salesenum.FulfillmentStatusFulfilled}, invalid: salesenum.FulfillmentStatus("__invalid__")},
		{name: "warehouseenum.InboundReceiptStatus", valid: []stringEnum{warehouseenum.InboundReceiptStatusDraft, warehouseenum.InboundReceiptStatusConfirmed}, invalid: warehouseenum.InboundReceiptStatus("__invalid__")},
		{name: "identityenum.IdentityDomain", valid: []stringEnum{identityenum.IdentityDomainCustomer, identityenum.IdentityDomainWorkforce, identityenum.IdentityDomainPartner, identityenum.IdentityDomainService}, invalid: identityenum.IdentityDomain("__invalid__")},
		{name: "warehouseenum.LayoutNodeType", valid: []stringEnum{warehouseenum.LayoutNodeZone, warehouseenum.LayoutNodeAisle, warehouseenum.LayoutNodeRack, warehouseenum.LayoutNodeShelf, warehouseenum.LayoutNodeBin}, invalid: warehouseenum.LayoutNodeType("__invalid__")},
		{name: "marketingenum.MarketingCampaignStatus", valid: []stringEnum{marketingenum.MarketingCampaignStatusDraft, marketingenum.MarketingCampaignStatusSending, marketingenum.MarketingCampaignStatusSent, marketingenum.MarketingCampaignStatusPartial, marketingenum.MarketingCampaignStatusFailed, marketingenum.MarketingCampaignStatusCancelled, marketingenum.MarketingCampaignStatusExported}, invalid: marketingenum.MarketingCampaignStatus("__invalid__")},
		{name: "marketingenum.MarketingChannel", valid: []stringEnum{marketingenum.MarketingChannelEmail, marketingenum.MarketingChannelSMS, marketingenum.MarketingChannelLine, marketingenum.MarketingChannelExport}, invalid: marketingenum.MarketingChannel("__invalid__")},
		{name: "marketingenum.MarketingRecipientStatus", valid: []stringEnum{marketingenum.MarketingRecipientStatusPending, marketingenum.MarketingRecipientStatusSent, marketingenum.MarketingRecipientStatusDelivered, marketingenum.MarketingRecipientStatusBounced, marketingenum.MarketingRecipientStatusOpened, marketingenum.MarketingRecipientStatusClicked, marketingenum.MarketingRecipientStatusFailed, marketingenum.MarketingRecipientStatusUnsubscribed}, invalid: marketingenum.MarketingRecipientStatus("__invalid__")},
		{name: "membershipenum.MembershipAccountStatus", valid: []stringEnum{membershipenum.MembershipAccountStatusActive, membershipenum.MembershipAccountStatusSuspended, membershipenum.MembershipAccountStatusClosed}, invalid: membershipenum.MembershipAccountStatus("__invalid__")},
		{name: "membershipenum.MembershipOwnerType", valid: []stringEnum{membershipenum.MembershipOwnerTypeRetailCustomer, membershipenum.MembershipOwnerTypeWholesaleOrganisation}, invalid: membershipenum.MembershipOwnerType("__invalid__")},
		{name: "membershipenum.MembershipPointReason", valid: []stringEnum{membershipenum.MembershipPointReasonOrder, membershipenum.MembershipPointReasonBirthday, membershipenum.MembershipPointReasonRedeem, membershipenum.MembershipPointReasonRewardRedeem, membershipenum.MembershipPointReasonAdminAdjust, membershipenum.MembershipPointReasonExpired, membershipenum.MembershipPointReasonReferral, membershipenum.MembershipPointReasonSignupBonus, membershipenum.MembershipPointReasonTierUpgrade, membershipenum.MembershipPointReasonManual}, invalid: membershipenum.MembershipPointReason("__invalid__")},
		{name: "membershipenum.MembershipPromotionTarget", valid: []stringEnum{membershipenum.MembershipPromotionTargetAll, membershipenum.MembershipPromotionTargetWholesale, membershipenum.MembershipPromotionTargetRetail, membershipenum.MembershipPromotionTargetTierSpecific}, invalid: membershipenum.MembershipPromotionTarget("__invalid__")},
		{name: "membershipenum.MembershipRedemptionType", valid: []stringEnum{membershipenum.MembershipRedemptionTypeCheckoutDiscount, membershipenum.MembershipRedemptionTypeRewardCatalog}, invalid: membershipenum.MembershipRedemptionType("__invalid__")},
		{name: "membershipenum.MembershipRewardRedemptionStatus", valid: []stringEnum{membershipenum.MembershipRewardRedemptionStatusReserved, membershipenum.MembershipRewardRedemptionStatusRedeemed, membershipenum.MembershipRewardRedemptionStatusCancelled, membershipenum.MembershipRewardRedemptionStatusExpired}, invalid: membershipenum.MembershipRewardRedemptionStatus("__invalid__")},
		{name: "membershipenum.MembershipRewardType", valid: []stringEnum{membershipenum.MembershipRewardTypeOrderDiscount, membershipenum.MembershipRewardTypeProduct, membershipenum.MembershipRewardTypeFreeShipping, membershipenum.MembershipRewardTypeVoucher}, invalid: membershipenum.MembershipRewardType("__invalid__")},
		{name: "membershipenum.MembershipTierMetric", valid: []stringEnum{membershipenum.MembershipTierMetricAnnualSpend, membershipenum.MembershipTierMetricLifetimeSpend, membershipenum.MembershipTierMetricManual}, invalid: membershipenum.MembershipTierMetric("__invalid__")},
		{name: "membershipenum.MemberSubscriptionStatus", valid: []stringEnum{membershipenum.MemberSubscriptionStatusActive, membershipenum.MemberSubscriptionStatusPaused, membershipenum.MemberSubscriptionStatusCancelled}, invalid: membershipenum.MemberSubscriptionStatus("__invalid__")},
		{name: "warehouseenum.ModelFormat", valid: []stringEnum{warehouseenum.ModelFormatGLB, warehouseenum.ModelFormatGLTF, warehouseenum.ModelFormatOBJ, warehouseenum.ModelFormatFBX, warehouseenum.ModelFormatUSDZ}, invalid: warehouseenum.ModelFormat("__invalid__")},
		{name: "salesenum.OrderSourceDeviceType", valid: []stringEnum{salesenum.OrderSourceDeviceTypeIOS, salesenum.OrderSourceDeviceTypeAndroid, salesenum.OrderSourceDeviceTypePC, salesenum.OrderSourceDeviceTypeMobileWeb, salesenum.OrderSourceDeviceTypeTablet, salesenum.OrderSourceDeviceTypePos, salesenum.OrderSourceDeviceTypeManual, salesenum.OrderSourceDeviceTypePhone, salesenum.OrderSourceDeviceTypeVR}, invalid: salesenum.OrderSourceDeviceType("__invalid__")},
		{name: "salesenum.OrderType", valid: []stringEnum{salesenum.OrderTypeOnline, salesenum.OrderTypePOS, salesenum.OrderTypeB2B, salesenum.OrderTypeRelay, salesenum.OrderTypeManual, salesenum.OrderTypeImport}, invalid: salesenum.OrderType("__invalid__")},
		{name: "customerenum.BuyerType", valid: []stringEnum{customerenum.BuyerTypeGuestRetail, customerenum.BuyerTypeRetailCustomer, customerenum.BuyerTypeWholesaleOrganisation}, invalid: customerenum.BuyerType("__invalid__")},
		{name: "productenum.PriceAudience", valid: []stringEnum{productenum.PriceAudienceRetail, productenum.PriceAudienceWholesale}, invalid: productenum.PriceAudience("__invalid__")},
		{name: "productenum.PriceVisibility", valid: []stringEnum{productenum.PriceVisibilityPublic, productenum.PriceVisibilityLoginRequired, productenum.PriceVisibilityWholesaleApprovedOnly, productenum.PriceVisibilityHidden}, invalid: productenum.PriceVisibility("__invalid__")},
		{name: "shippingenum.FulfilmentIntent", valid: []stringEnum{shippingenum.FulfilmentIntentDelivery, shippingenum.FulfilmentIntentPickup, shippingenum.FulfilmentIntentInStoreCarry}, invalid: shippingenum.FulfilmentIntent("__invalid__")},
		{name: "wholesaleenum.OrganisationAccessStatus", valid: []stringEnum{wholesaleenum.OrganisationAccessStatusPending, wholesaleenum.OrganisationAccessStatusActive, wholesaleenum.OrganisationAccessStatusSuspended, wholesaleenum.OrganisationAccessStatusRevoked}, invalid: wholesaleenum.OrganisationAccessStatus("__invalid__")},
		{name: "wholesaleenum.WholesaleBuyerRole", valid: []stringEnum{wholesaleenum.WholesaleBuyerRoleOwner, wholesaleenum.WholesaleBuyerRoleBuyer, wholesaleenum.WholesaleBuyerRoleFinance, wholesaleenum.WholesaleBuyerRoleReadOnly}, invalid: wholesaleenum.WholesaleBuyerRole("__invalid__")},
		{name: "wholesaleenum.WholesalePermission", valid: []stringEnum{wholesaleenum.WholesalePermissionProductsView, wholesaleenum.WholesalePermissionCartWrite, wholesaleenum.WholesalePermissionCheckoutSubmit, wholesaleenum.WholesalePermissionOrdersViewOwn, wholesaleenum.WholesalePermissionOrdersViewOrg, wholesaleenum.WholesalePermissionOrdersReorder, wholesaleenum.WholesalePermissionInvoicesViewOwn, wholesaleenum.WholesalePermissionInvoicesViewOrg, wholesaleenum.WholesalePermissionInvoicesPay, wholesaleenum.WholesalePermissionAccountView, wholesaleenum.WholesalePermissionTeamView, wholesaleenum.WholesalePermissionFavouritesWrite, wholesaleenum.WholesalePermissionOrderListsViewOwn, wholesaleenum.WholesalePermissionOrderListsWriteOwn, wholesaleenum.WholesalePermissionOrderListsViewOrg, wholesaleenum.WholesalePermissionOrderListsWriteOrg}, invalid: wholesaleenum.WholesalePermission("__invalid__")},
		{name: "warehouseenum.OutboundShipmentStatus", valid: []stringEnum{warehouseenum.OutboundShipmentStatusPacked, warehouseenum.OutboundShipmentStatusDispatched}, invalid: warehouseenum.OutboundShipmentStatus("__invalid__")},
		{name: "warehouseenum.PackingDiscrepancyKind", valid: []stringEnum{warehouseenum.PackingDiscrepancyKindShortage, warehouseenum.PackingDiscrepancyKindOverweight, warehouseenum.PackingDiscrepancyKindDamaged, warehouseenum.PackingDiscrepancyKindPending}, invalid: warehouseenum.PackingDiscrepancyKind("__invalid__")},
		{name: "paymentenum.PaymentMethod", valid: []stringEnum{paymentenum.PaymentMethodCard, paymentenum.PaymentMethodCash, paymentenum.PaymentMethodQR, paymentenum.PaymentMethodBankTransfer, paymentenum.PaymentMethodLinePay, paymentenum.PaymentMethodApplePay, paymentenum.PaymentMethodGooglePay, paymentenum.PaymentMethodECPay, paymentenum.PaymentMethodManual, paymentenum.PaymentMethodEFTPOS, paymentenum.PaymentMethodMOTO, paymentenum.PaymentMethodCashout}, invalid: paymentenum.PaymentMethod("__invalid__")},
		{name: "paymentenum.PaymentRecordStatus", valid: []stringEnum{paymentenum.PaymentRecordStatusPending, paymentenum.PaymentRecordStatusProcessing, paymentenum.PaymentRecordStatusCompleted, paymentenum.PaymentRecordStatusFailed, paymentenum.PaymentRecordStatusCancelled, paymentenum.PaymentRecordStatusRefunded, paymentenum.PaymentRecordStatusAwaitingAction, paymentenum.PaymentRecordStatusUnknown}, invalid: paymentenum.PaymentRecordStatus("__invalid__")},
		{name: "paymentenum.PaymentStatus", valid: []stringEnum{paymentenum.PaymentStatusUnknown, paymentenum.PaymentStatusUnpaid, paymentenum.PaymentStatusPending, paymentenum.PaymentStatusPaid, paymentenum.PaymentStatusPartiallyPaid, paymentenum.PaymentStatusRefunded, paymentenum.PaymentStatusPartialRefunded}, invalid: paymentenum.PaymentStatus("__invalid__")},
		{name: "warehouseenum.PickingItemStatus", valid: []stringEnum{warehouseenum.PickingItemStatusPending, warehouseenum.PickingItemStatusPartial, warehouseenum.PickingItemStatusComplete, warehouseenum.PickingItemStatusSkipped}, invalid: warehouseenum.PickingItemStatus("__invalid__")},
		{name: "warehouseenum.PickingListStatus", valid: []stringEnum{warehouseenum.PickingListStatusPending, warehouseenum.PickingListStatusInProgress, warehouseenum.PickingListStatusComplete, warehouseenum.PickingListStatusCancelled}, invalid: warehouseenum.PickingListStatus("__invalid__")},
		{name: "membershipenum.PointReservationStatus", valid: []stringEnum{membershipenum.PointReservationStatusReserved, membershipenum.PointReservationStatusCommitted, membershipenum.PointReservationStatusCancelled, membershipenum.PointReservationStatusExpired}, invalid: membershipenum.PointReservationStatus("__invalid__")},
		{name: "accountenum.Portal", valid: []stringEnum{accountenum.PortalControl, accountenum.PortalStore, accountenum.PortalPartner}, invalid: accountenum.Portal("__invalid__")},
		{name: "accountenum.PortalAccessStatus", valid: []stringEnum{accountenum.PortalAccessStatusPending, accountenum.PortalAccessStatusActive, accountenum.PortalAccessStatusSuspended, accountenum.PortalAccessStatusRevoked}, invalid: accountenum.PortalAccessStatus("__invalid__")},
		{name: "salesenum.PreorderStatus", valid: []stringEnum{salesenum.PreorderStatusRequested, salesenum.PreorderStatusAccepted, salesenum.PreorderStatusRejected, salesenum.PreorderStatusCancelled, salesenum.PreorderStatusConverted, salesenum.PreorderStatusFulfilled, salesenum.PreorderStatusExpired}, invalid: salesenum.PreorderStatus("__invalid__")},
		{name: "productenum.StorefrontPreorderStatus", valid: []stringEnum{productenum.StorefrontPreorderStatusUnavailable, productenum.StorefrontPreorderStatusUpcoming, productenum.StorefrontPreorderStatusOpen, productenum.StorefrontPreorderStatusClosed, productenum.StorefrontPreorderStatusSoldOut}, invalid: productenum.StorefrontPreorderStatus("__invalid__")},
		{name: "productenum.StorefrontExpiryStatus", valid: []stringEnum{productenum.StorefrontExpiryStatusNotApplicable, productenum.StorefrontExpiryStatusSoonExpiry, productenum.StorefrontExpiryStatusExpired}, invalid: productenum.StorefrontExpiryStatus("__invalid__")},
		{name: "promotionenum.PromotionAddonTrigger", valid: []stringEnum{promotionenum.PromotionAddonTriggerAmount, promotionenum.PromotionAddonTriggerRequiredProducts}, invalid: promotionenum.PromotionAddonTrigger("__invalid__")},
		{name: "promotionenum.PromotionClass", valid: []stringEnum{promotionenum.PromotionClassNormal, promotionenum.PromotionClassSpecialCampaign}, invalid: promotionenum.PromotionClass("__invalid__")},
		{name: "promotionenum.PromotionDiscountTarget", valid: []stringEnum{promotionenum.PromotionDiscountTargetCart, promotionenum.PromotionDiscountTargetRequiredItems}, invalid: promotionenum.PromotionDiscountTarget("__invalid__")},
		{name: "promotionenum.PromotionQtyMode", valid: []stringEnum{promotionenum.PromotionQtyModePerProduct, promotionenum.PromotionQtyModeCombined}, invalid: promotionenum.PromotionQtyMode("__invalid__")},
		{name: "promotionenum.PromotionType", valid: []stringEnum{promotionenum.PromotionTypeAutoDiscount, promotionenum.PromotionTypeSpendGift, promotionenum.PromotionTypeAddonPurchase, promotionenum.PromotionTypeBOGO, promotionenum.PromotionTypeBundle, promotionenum.PromotionTypeTieredPricing}, invalid: promotionenum.PromotionType("__invalid__")},
		{name: "productenum.ProductStatus", valid: []stringEnum{productenum.ProductStatusDraft, productenum.ProductStatusActive, productenum.ProductStatusArchived, productenum.ProductStatusDiscontinued}, invalid: productenum.ProductStatus("__invalid__")},
		{name: "purchaseenum.PurchaseOrderStatus", valid: []stringEnum{purchaseenum.PurchaseOrderStatusDraft, purchaseenum.PurchaseOrderStatusSubmitted, purchaseenum.PurchaseOrderStatusConfirmed, purchaseenum.PurchaseOrderStatusPartiallyReceived, purchaseenum.PurchaseOrderStatusReceived, purchaseenum.PurchaseOrderStatusCancelled, purchaseenum.PurchaseOrderStatusRefunded}, invalid: purchaseenum.PurchaseOrderStatus("__invalid__")},
		{name: "paymentenum.RecoveryDecision", valid: []stringEnum{paymentenum.RecoveryDecisionPending, paymentenum.RecoveryDecisionApproved, paymentenum.RecoveryDecisionDeclined}, invalid: paymentenum.RecoveryDecision("__invalid__")},
		{name: "productenum.SalesPerformance", valid: []stringEnum{productenum.SalesPerformanceHot, productenum.SalesPerformanceNormal, productenum.SalesPerformanceSlow}, invalid: productenum.SalesPerformance("__invalid__")},
		{name: "salesenum.SalesOrderStatus", valid: []stringEnum{salesenum.SalesOrderStatusPending, salesenum.SalesOrderStatusConfirmed, salesenum.SalesOrderStatusPaid, salesenum.SalesOrderStatusProcessing, salesenum.SalesOrderStatusPicking, salesenum.SalesOrderStatusPacked, salesenum.SalesOrderStatusShipped, salesenum.SalesOrderStatusDelivered, salesenum.SalesOrderStatusCompleted, salesenum.SalesOrderStatusCancelled, salesenum.SalesOrderStatusRefunded}, invalid: salesenum.SalesOrderStatus("__invalid__")},
		{name: "securityenum.SecurityEventSeverity", valid: []stringEnum{securityenum.SecurityEventSeverityInfo, securityenum.SecurityEventSeverityLow, securityenum.SecurityEventSeverityMedium, securityenum.SecurityEventSeverityHigh, securityenum.SecurityEventSeverityCritical}, invalid: securityenum.SecurityEventSeverity("__invalid__")},
		{name: "securityenum.SecurityEventStatus", valid: []stringEnum{securityenum.SecurityEventStatusDetected, securityenum.SecurityEventStatusTriaged, securityenum.SecurityEventStatusInvestigating, securityenum.SecurityEventStatusContained, securityenum.SecurityEventStatusResolved, securityenum.SecurityEventStatusFalsePositive}, invalid: securityenum.SecurityEventStatus("__invalid__")},
		{name: "securityenum.SecurityRiskLevel", valid: []stringEnum{securityenum.SecurityRiskLevelLow, securityenum.SecurityRiskLevelMedium, securityenum.SecurityRiskLevelHigh, securityenum.SecurityRiskLevelCritical}, invalid: securityenum.SecurityRiskLevel("__invalid__")},
		{name: "paymentenum.SettlementType", valid: []stringEnum{paymentenum.SettlementTypeSettlement, paymentenum.SettlementTypeEnquiry}, invalid: paymentenum.SettlementType("__invalid__")},
		{name: "warehouseenum.ShapeType", valid: []stringEnum{warehouseenum.ShapeBox, warehouseenum.ShapeCylinder, warehouseenum.ShapeSphere, warehouseenum.ShapePlane, warehouseenum.ShapeCustom}, invalid: warehouseenum.ShapeType("__invalid__")},
		{name: "shippingenum.ShippingRateName", valid: []stringEnum{shippingenum.ShippingRateNameStandard, shippingenum.ShippingRateNameExpress, shippingenum.ShippingRateNamePickup}, invalid: shippingenum.ShippingRateName("__invalid__")},
		{name: "warehouseenum.StockMovementType", valid: []stringEnum{warehouseenum.StockMovementTypePurchaseReceipt, warehouseenum.StockMovementTypeSaleReserve, warehouseenum.StockMovementTypeSaleCommit, warehouseenum.StockMovementTypeSaleRelease, warehouseenum.StockMovementTypeAdjustment, warehouseenum.StockMovementTypeDamage, warehouseenum.StockMovementTypeReturn, warehouseenum.StockMovementTypeTransferIn, warehouseenum.StockMovementTypeTransferOut, warehouseenum.StockMovementTypeStocktake}, invalid: warehouseenum.StockMovementType("__invalid__")},
		{name: "warehouseenum.StorageType", valid: []stringEnum{warehouseenum.StorageDry, warehouseenum.StorageChilled, warehouseenum.StorageFrozen}, invalid: warehouseenum.StorageType("__invalid__")},
		{name: "paymentenum.TerminalConnectionMode", valid: []stringEnum{paymentenum.TerminalConnectionModeCloudSync, paymentenum.TerminalConnectionModeCloudAsync, paymentenum.TerminalConnectionModeLocal}, invalid: paymentenum.TerminalConnectionMode("__invalid__")},
		{name: "paymentenum.TerminalProvider", valid: []stringEnum{paymentenum.TerminalProviderMx51}, invalid: paymentenum.TerminalProvider("__invalid__")},
		{name: "paymentenum.TerminalRefundType", valid: []stringEnum{paymentenum.TerminalRefundTypeReferenced, paymentenum.TerminalRefundTypeUnreferenced}, invalid: paymentenum.TerminalRefundType("__invalid__")},
		{name: "paymentenum.TerminalStatus", valid: []stringEnum{paymentenum.TerminalStatusRegistered, paymentenum.TerminalStatusActive, paymentenum.TerminalStatusDeregistered, paymentenum.TerminalStatusExpired, paymentenum.TerminalStatusError}, invalid: paymentenum.TerminalStatus("__invalid__")},
		{name: "paymentenum.TerminalTxFinancialStatus", valid: []stringEnum{paymentenum.TerminalTxFinancialStatusApproved, paymentenum.TerminalTxFinancialStatusDeclined, paymentenum.TerminalTxFinancialStatusCancelled, paymentenum.TerminalTxFinancialStatusUnknown}, invalid: paymentenum.TerminalTxFinancialStatus("__invalid__")},
		{name: "paymentenum.TerminalTxStatus", valid: []stringEnum{paymentenum.TerminalTxStatusUnknown, paymentenum.TerminalTxStatusPending, paymentenum.TerminalTxStatusAwaitingAction, paymentenum.TerminalTxStatusFinalised, paymentenum.TerminalTxStatusOverridePending, paymentenum.TerminalTxStatusOverrideResolved}, invalid: paymentenum.TerminalTxStatus("__invalid__")},
		{name: "paymentenum.TerminalTxType", valid: []stringEnum{paymentenum.TerminalTxTypePurchase, paymentenum.TerminalTxTypeRefund, paymentenum.TerminalTxTypeReversal, paymentenum.TerminalTxTypeCashout, paymentenum.TerminalTxTypePurchaseWithCashout, paymentenum.TerminalTxTypeMOTO, paymentenum.TerminalTxTypeSettlement, paymentenum.TerminalTxTypeSettlementEnquiry}, invalid: paymentenum.TerminalTxType("__invalid__")},
		{name: "identityenum.UserPreferredLanguage", valid: []stringEnum{identityenum.PreferredLanguageEnglish, identityenum.PreferredLanguageTraditionalChinese, identityenum.PreferredLanguageSimplifiedChinese}, invalid: identityenum.UserPreferredLanguage("__invalid__")},
		{name: "identityenum.UserRole", valid: []stringEnum{identityenum.UserRoleSuperAdmin, identityenum.UserRoleAdmin, identityenum.UserRoleSales, identityenum.UserRoleWarehouse, identityenum.UserRoleWarehouseOperator, identityenum.UserRoleMarketing, identityenum.UserRoleCustomer}, invalid: identityenum.UserRole("__invalid__")},
		{name: "promotionenum.VolumeDiscountAppliesTo", valid: []stringEnum{promotionenum.VolumeDiscountAppliesToAll, promotionenum.VolumeDiscountAppliesToWholesale, promotionenum.VolumeDiscountAppliesToRetail}, invalid: promotionenum.VolumeDiscountAppliesTo("__invalid__")},
		{name: "wholesaleenum.WholesaleApplicationState", valid: []stringEnum{wholesaleenum.WholesaleApplicationStateMissing, wholesaleenum.WholesaleApplicationStatePending, wholesaleenum.WholesaleApplicationStateApproved, wholesaleenum.WholesaleApplicationStateRejected, wholesaleenum.WholesaleApplicationStateSuspended}, invalid: wholesaleenum.WholesaleApplicationState("__invalid__")},
		{name: "wholesaleenum.WholesaleOrganisationStatus", valid: []stringEnum{wholesaleenum.WholesaleOrganisationStatusPending, wholesaleenum.WholesaleOrganisationStatusApproved, wholesaleenum.WholesaleOrganisationStatusSuspended, wholesaleenum.WholesaleOrganisationStatusRejected, wholesaleenum.WholesaleOrganisationStatusClosed}, invalid: wholesaleenum.WholesaleOrganisationStatus("__invalid__")},
		{name: "warehouseenum.WMSDraftStatus", valid: []stringEnum{warehouseenum.WMSDraftStatusDraft, warehouseenum.WMSDraftStatusSubmitted, warehouseenum.WMSDraftStatusCancelled}, invalid: warehouseenum.WMSDraftStatus("__invalid__")},
		{name: "warehouseenum.WMSDraftType", valid: []stringEnum{warehouseenum.WMSDraftTypeInbound, warehouseenum.WMSDraftTypeOutbound}, invalid: warehouseenum.WMSDraftType("__invalid__")},
		{name: "walletenum.WalletInstrumentType", valid: []stringEnum{walletenum.WalletInstrumentTypePoints, walletenum.WalletInstrumentTypeGiftCard, walletenum.WalletInstrumentTypeVoucher, walletenum.WalletInstrumentTypeCoupon, walletenum.WalletInstrumentTypeReward}, invalid: walletenum.WalletInstrumentType("__invalid__")},
		{name: "walletenum.WalletExportFormat", valid: []stringEnum{walletenum.WalletExportFormatJSON, walletenum.WalletExportFormatCSVZip}, invalid: walletenum.WalletExportFormat("__invalid__")},
		{name: "walletenum.WalletExportStatus", valid: []stringEnum{walletenum.WalletExportStatusPending, walletenum.WalletExportStatusRunning, walletenum.WalletExportStatusCompleted, walletenum.WalletExportStatusFailed, walletenum.WalletExportStatusExpired}, invalid: walletenum.WalletExportStatus("__invalid__")},
		{name: "walletenum.GiftCardStatus", valid: []stringEnum{walletenum.GiftCardStatusActive, walletenum.GiftCardStatusPartiallyRedeemed, walletenum.GiftCardStatusDepleted, walletenum.GiftCardStatusExpired, walletenum.GiftCardStatusVoid}, invalid: walletenum.GiftCardStatus("__invalid__")},
		{name: "walletenum.GiftCardTransactionReason", valid: []stringEnum{walletenum.GiftCardTransactionReasonIssue, walletenum.GiftCardTransactionReasonRedeem, walletenum.GiftCardTransactionReasonRefund, walletenum.GiftCardTransactionReasonTopUp, walletenum.GiftCardTransactionReasonExpire, walletenum.GiftCardTransactionReasonAdjust}, invalid: walletenum.GiftCardTransactionReason("__invalid__")},
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
	allowed := map[accountenum.AccountType]accountenum.Portal{
		accountenum.AccountTypeAdminUser:         accountenum.PortalControl,
		accountenum.AccountTypeGeneralCustomer:   accountenum.PortalStore,
		accountenum.AccountTypeWholesaleCustomer: accountenum.PortalPartner,
	}

	for accountType, portal := range allowed {
		if !accountType.IsAllowedInPortal(portal) {
			t.Fatalf("%s should be allowed in %s", accountType, portal)
		}
	}

	rejected := []struct {
		accountType accountenum.AccountType
		portal      accountenum.Portal
	}{
		{accountenum.AccountTypeAdminUser, accountenum.PortalStore},
		{accountenum.AccountTypeAdminUser, accountenum.PortalPartner},
		{accountenum.AccountTypeGeneralCustomer, accountenum.PortalControl},
		{accountenum.AccountTypeGeneralCustomer, accountenum.PortalPartner},
		{accountenum.AccountTypeWholesaleCustomer, accountenum.PortalControl},
		{accountenum.AccountTypeWholesaleCustomer, accountenum.PortalStore},
	}

	for _, tt := range rejected {
		if tt.accountType.IsAllowedInPortal(tt.portal) {
			t.Fatalf("%s should not be allowed in %s", tt.accountType, tt.portal)
		}
	}
}

func TestPortalAccountTypeHelpers(t *testing.T) {
	tests := []struct {
		portal      accountenum.Portal
		accountType accountenum.AccountType
	}{
		{accountenum.PortalControl, accountenum.AccountTypeAdminUser},
		{accountenum.PortalStore, accountenum.AccountTypeGeneralCustomer},
		{accountenum.PortalPartner, accountenum.AccountTypeWholesaleCustomer},
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

		want := []accountenum.AccountType{tt.accountType}
		if got := tt.portal.AllowedAccountTypes(); !reflect.DeepEqual(got, want) {
			t.Fatalf("%s allowed account types = %#v, want %#v", tt.portal, got, want)
		}
		if got := accountenum.AccountTypesForPortal(tt.portal); !reflect.DeepEqual(got, want) {
			t.Fatalf("accountenum.AccountTypesForPortal(%s) = %#v, want %#v", tt.portal, got, want)
		}
	}

	if _, ok := accountenum.Portal("__invalid__").RequiredAccountType(); ok {
		t.Fatal("invalid portal should not have a required account type")
	}
	if got := accountenum.Portal("__invalid__").AllowedAccountTypes(); got != nil {
		t.Fatalf("invalid portal allowed account types = %#v, want nil", got)
	}
}

func TestPortalAccessStatusCanAccess(t *testing.T) {
	if !accountenum.PortalAccessStatusActive.CanAccess() {
		t.Fatal("active portal access should allow access")
	}

	for _, status := range []accountenum.PortalAccessStatus{
		accountenum.PortalAccessStatusPending,
		accountenum.PortalAccessStatusSuspended,
		accountenum.PortalAccessStatusRevoked,
	} {
		if status.CanAccess() {
			t.Fatalf("%s should not allow access", status)
		}
	}
}

func TestWholesaleBuyerRolePermissions(t *testing.T) {
	ownerPerms := wholesaleenum.PermissionsForWholesaleBuyerRole(wholesaleenum.WholesaleBuyerRoleOwner)
	if len(ownerPerms) == 0 {
		t.Fatal("owner should receive permissions")
	}
	if !wholesaleenum.HasWholesalePermission(wholesaleenum.WholesalePermissionStrings(ownerPerms), wholesaleenum.WholesalePermissionTeamView) {
		t.Fatal("owner should receive team.view")
	}
	if !wholesaleenum.HasWholesalePermission(wholesaleenum.WholesalePermissionStrings(ownerPerms), wholesaleenum.WholesalePermissionOrderListsWriteOrg) {
		t.Fatal("owner should receive order_lists.write_org")
	}

	buyerPerms := wholesaleenum.WholesalePermissionStrings(wholesaleenum.PermissionsForWholesaleBuyerRole(wholesaleenum.WholesaleBuyerRoleBuyer))
	if wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionInvoicesViewOrg) {
		t.Fatal("buyer should not receive organisation invoice visibility")
	}
	if !wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionCheckoutSubmit) {
		t.Fatal("buyer should receive checkout.submit")
	}
	if !wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionFavouritesWrite) ||
		!wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionOrderListsWriteOwn) ||
		!wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionOrderListsViewOrg) {
		t.Fatal("buyer should receive own list mutation, org list view, and favourites.write")
	}
	if wholesaleenum.HasWholesalePermission(buyerPerms, wholesaleenum.WholesalePermissionOrderListsWriteOrg) {
		t.Fatal("buyer should not receive organisation list mutation")
	}

	financePerms := wholesaleenum.WholesalePermissionStrings(wholesaleenum.PermissionsForWholesaleBuyerRole(wholesaleenum.WholesaleBuyerRoleFinance))
	if !wholesaleenum.HasWholesalePermission(financePerms, wholesaleenum.WholesalePermissionInvoicesPay) {
		t.Fatal("finance should receive invoices.pay")
	}
	if wholesaleenum.HasWholesalePermission(financePerms, wholesaleenum.WholesalePermissionCheckoutSubmit) {
		t.Fatal("finance should not receive checkout.submit")
	}
	if wholesaleenum.HasWholesalePermission(financePerms, wholesaleenum.WholesalePermissionFavouritesWrite) ||
		wholesaleenum.HasWholesalePermission(financePerms, wholesaleenum.WholesalePermissionOrderListsViewOwn) ||
		wholesaleenum.HasWholesalePermission(financePerms, wholesaleenum.WholesalePermissionOrderListsViewOrg) {
		t.Fatal("finance should not receive procurement list permissions")
	}

	readOnlyPerms := wholesaleenum.WholesalePermissionStrings(wholesaleenum.PermissionsForWholesaleBuyerRole(wholesaleenum.WholesaleBuyerRoleReadOnly))
	if !wholesaleenum.HasWholesalePermission(readOnlyPerms, wholesaleenum.WholesalePermissionOrderListsViewOrg) {
		t.Fatal("read-only should receive organisation list view")
	}
	if wholesaleenum.HasWholesalePermission(readOnlyPerms, wholesaleenum.WholesalePermissionFavouritesWrite) ||
		wholesaleenum.HasWholesalePermission(readOnlyPerms, wholesaleenum.WholesalePermissionOrderListsWriteOwn) ||
		wholesaleenum.HasWholesalePermission(readOnlyPerms, wholesaleenum.WholesalePermissionCartWrite) {
		t.Fatal("read-only should not receive procurement mutation or cart permissions")
	}

	if got := wholesaleenum.PermissionsForWholesaleBuyerRole(wholesaleenum.WholesaleBuyerRole("__invalid__")); got != nil {
		t.Fatalf("invalid role permissions = %#v, want nil", got)
	}
	if wholesaleenum.HasWholesalePermission([]string{wholesaleenum.WholesalePermissionTeamView.String()}, wholesaleenum.WholesalePermission("__invalid__")) {
		t.Fatal("invalid required permission should never match")
	}
}

func TestAccountStatusTerminalState(t *testing.T) {
	for _, status := range []accountenum.AccountStatus{accountenum.AccountStatusClosed, accountenum.AccountStatusDeleted} {
		if !status.IsTerminal() {
			t.Fatalf("%s should be terminal", status)
		}
	}

	for _, status := range []accountenum.AccountStatus{
		accountenum.AccountStatusPending,
		accountenum.AccountStatusActive,
		accountenum.AccountStatusSuspended,
	} {
		if status.IsTerminal() {
			t.Fatalf("%s should not be terminal", status)
		}
	}
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

func TestTerminalStatusesTerminalState(t *testing.T) {
	if !paymentenum.TerminalStatusDeregistered.IsTerminal() || !paymentenum.TerminalStatusExpired.IsTerminal() {
		t.Fatal("deregistered and expired terminals should be terminal")
	}
	if paymentenum.TerminalStatusError.IsTerminal() {
		t.Fatal("error terminal status should be recoverable")
	}

	if !paymentenum.TerminalTxStatusFinalised.IsTerminal() || !paymentenum.TerminalTxStatusOverrideResolved.IsTerminal() {
		t.Fatal("finalised and override_resolved terminal transactions should be terminal")
	}
	if paymentenum.TerminalTxStatusOverridePending.IsTerminal() {
		t.Fatal("override_pending terminal transaction should require resolution")
	}
}
