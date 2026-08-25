package role_enums

import "testing"

func TestPermissionKeyWireValuesAreLocked(t *testing.T) {
	keys := []struct {
		key  PermissionKey
		wire string
	}{
		{PermissionKeyUserRead, "user.read"},
		{PermissionKeyUserWrite, "user.write"},
		{PermissionKeyUserDelete, "user.delete"},
		{PermissionKeyRetailCustomerRead, "retail_customer.read"},
		{PermissionKeyRetailCustomerWrite, "retail_customer.write"},
		{PermissionKeySupplierRead, "supplier.read"},
		{PermissionKeySupplierWrite, "supplier.write"},
		{PermissionKeyRoleRead, "role.read"},
		{PermissionKeyRoleWrite, "role.write"},
		{PermissionKeyMembershipRead, "membership.read"},
		{PermissionKeyMembershipWrite, "membership.write"},
		{PermissionKeyWalletRead, "wallet.read"},
		{PermissionKeyWalletWrite, "wallet.write"},
		{PermissionKeyGiftCardPolicyRead, "gift_card_policy.read"},
		{PermissionKeyGiftCardPolicyWrite, "gift_card_policy.write"},
		{PermissionKeyPromotionRead, "promotion.read"},
		{PermissionKeyPromotionWrite, "promotion.write"},
		{PermissionKeyPromotionPublish, "promotion.publish"},
		{PermissionKeyCouponRead, "coupon.read"},
		{PermissionKeyCouponWrite, "coupon.write"},
		{PermissionKeyPriceRead, "price.read"},
		{PermissionKeyPriceApprove, "price.approve"},
		{PermissionKeyMarketingRead, "marketing.read"},
		{PermissionKeyMarketingSend, "marketing.send"},
		{PermissionKeyAnalyticsSalesRead, "analytics.sales.read"},
		{PermissionKeyAnalyticsProductRead, "analytics.product.read"},
		{PermissionKeyAnalyticsExport, "analytics.export"},
		{PermissionKeyMediaRead, "media.read"},
		{PermissionKeyMediaUpload, "media.upload"},
		{PermissionKeyMediaDelete, "media.delete"},
		{PermissionKeySettingsRead, "settings.read"},
		{PermissionKeySettingsWrite, "settings.write"},
		{PermissionKeyAuditRead, "audit.read"},
		{PermissionKeySecurityRead, "security.read"},
		{PermissionKeySecurityWrite, "security.write"},
		{PermissionKeyAccessLogRead, "access_log.read"},
		{PermissionKeyWholesaleRead, "wholesale.read"},
		{PermissionKeyWholesaleWrite, "wholesale.write"},
		{PermissionKeyProductRead, "product.read"},
		{PermissionKeyProductWrite, "product.write"},
		{PermissionKeyProductDelete, "product.delete"},
		{PermissionKeySKURead, "sku.read"},
		{PermissionKeySKUWrite, "sku.write"},
		{PermissionKeySKUProductsRead, "sku.products.read"},
		{PermissionKeyCategoryRead, "category.read"},
		{PermissionKeyCategoryWrite, "category.write"},
		{PermissionKeyCollectionRead, "collection.read"},
		{PermissionKeyCollectionWrite, "collection.write"},
		{PermissionKeyDepotRead, "depot.read"},
		{PermissionKeyDepotWrite, "depot.write"},
		{PermissionKeyLocationRead, "location.read"},
		{PermissionKeyLocationWrite, "location.write"},
		{PermissionKeyStockRead, "stock.read"},
		{PermissionKeyStockAdjust, "stock.adjust"},
		{PermissionKeyStockReserve, "stock.reserve"},
		{PermissionKeyStockTransfer, "stock.transfer"},
		{PermissionKeyPickingRead, "picking.read"},
		{PermissionKeyPickingWrite, "picking.write"},
		{PermissionKeyPackingRead, "packing.read"},
		{PermissionKeyPackingWrite, "packing.write"},
		{PermissionKeyShipmentRead, "shipment.read"},
		{PermissionKeyShipmentWrite, "shipment.write"},
		{PermissionKeyPurchaseRead, "purchase.read"},
		{PermissionKeyPurchaseWrite, "purchase.write"},
		{PermissionKeyPurchasePublish, "purchase.publish"},
		{PermissionKeyReceiptRead, "receipt.read"},
		{PermissionKeyReceiptWrite, "receipt.write"},
		{PermissionKeyExpiryRead, "expiry.read"},
		{PermissionKeyExpiryRun, "expiry.run"},
		{PermissionKeyForecastRead, "forecast.read"},
		{PermissionKeyForecastWrite, "forecast.write"},
		{PermissionKeyInboundRead, "inbound.read"},
		{PermissionKeyInboundWrite, "inbound.write"},
		{PermissionKeyDamageRead, "damage.read"},
		{PermissionKeyDamageWrite, "damage.write"},
		{PermissionKeyWMSDraftRead, "wmsdraft.read"},
		{PermissionKeyWMSDraftWrite, "wmsdraft.write"},
		{PermissionKeyLayoutRead, "layout.read"},
		{PermissionKeyLayoutWrite, "layout.write"},
		{PermissionKeyReviewRead, "review.read"},
		{PermissionKeyReviewModerate, "review.moderate"},
		{PermissionKeyWishRead, "wish.read"},
		{PermissionKeyWishManage, "wish.manage"},
		{PermissionKeyOrderRead, "order.read"},
		{PermissionKeyOrderWrite, "order.write"},
		{PermissionKeyOrderCancel, "order.cancel"},
		{PermissionKeyPaymentRead, "payment.read"},
		{PermissionKeyPaymentCapture, "payment.capture"},
		{PermissionKeyRefundRead, "refund.read"},
		{PermissionKeyRefundWrite, "refund.write"},
		{PermissionKeyRefundRequest, "refund.request"},
		{PermissionKeyRefundApprove, "refund.approve"},
		{PermissionKeyInvoiceRead, "invoice.read"},
		{PermissionKeyInvoiceWrite, "invoice.write"},
		{PermissionKeyCartManage, "cart.manage"},
		{PermissionKeyTerminalManage, "terminal.manage"},
		{PermissionKeyTerminalTransact, "terminal.transact"},
		{PermissionKeyCommerceConfig, "commerce.config"},
		{PermissionKeyPreorderRead, "preorder.read"},
		{PermissionKeyPreorderWrite, "preorder.write"},
		{PermissionKeyPOSAccess, "pos.access"},
		{PermissionKeyPOSSessionManage, "pos.session.manage"},
	}

	if len(keys) != 102 {
		t.Fatalf("shared workforce permission set has %d keys; want 102", len(keys))
	}

	seen := make(map[string]struct{}, len(keys))
	for _, entry := range keys {
		if got := entry.key.String(); got != entry.wire {
			t.Errorf("PermissionKey.String() = %q, want %q", got, entry.wire)
		}
		if !entry.key.IsValid() {
			t.Errorf("PermissionKey(%q) must validate", entry.wire)
		}
		if _, duplicate := seen[entry.wire]; duplicate {
			t.Errorf("duplicate workforce permission wire value %q", entry.wire)
		}
		seen[entry.wire] = struct{}{}
	}

	for _, retired := range []string{
		"customer.read",
		"customer.write",
		"wholesale_customer.read",
		"wholesale_customer.write",
		"regular_customer.read",
		"regular_customer.write",
		"pos.shift.manage",
		"__unknown_permission__",
	} {
		key := PermissionKey(retired)
		if key.IsValid() {
			t.Errorf("retired or unknown permission %q still validates", retired)
		}
		if got := key.String(); got != retired {
			t.Errorf("invalid permission String() = %q, want %q", got, retired)
		}
	}
}

func TestPermissionClassificationWireValuesAreLocked(t *testing.T) {
	classifications := map[PermissionClassification]string{
		PermissionClassificationUI:       "ui",
		PermissionClassificationField:    "field-level",
		PermissionClassificationService:  "service-only",
		PermissionClassificationReserved: "intentionally-reserved",
	}
	if len(classifications) != 4 {
		t.Fatalf("permission classification set has %d values; want 4", len(classifications))
	}
	for classification, wire := range classifications {
		if got := classification.String(); got != wire {
			t.Errorf("PermissionClassification.String() = %q, want %q", got, wire)
		}
		if !classification.IsValid() {
			t.Errorf("PermissionClassification(%q) must validate", wire)
		}
	}
	if PermissionClassification("__invalid__").IsValid() {
		t.Error("unknown permission classification validates")
	}
}
