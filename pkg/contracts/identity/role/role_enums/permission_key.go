package role_enums

// PermissionKey is a wire-stable workforce RBAC permission identifier.
// Identity owns the catalogue metadata and role-to-permission policy; this
// contract keeps the vocabulary shared without coupling consumers to that
// service's persistence or bootstrap implementation.
type PermissionKey string

const (
	PermissionKeyUserRead   PermissionKey = "user.read"
	PermissionKeyUserWrite  PermissionKey = "user.write"
	PermissionKeyUserDelete PermissionKey = "user.delete"

	PermissionKeyRetailCustomerRead  PermissionKey = "retail_customer.read"
	PermissionKeyRetailCustomerWrite PermissionKey = "retail_customer.write"

	PermissionKeySupplierRead  PermissionKey = "supplier.read"
	PermissionKeySupplierWrite PermissionKey = "supplier.write"

	PermissionKeyRoleRead  PermissionKey = "role.read"
	PermissionKeyRoleWrite PermissionKey = "role.write"

	PermissionKeyMembershipRead  PermissionKey = "membership.read"
	PermissionKeyMembershipWrite PermissionKey = "membership.write"

	PermissionKeyWalletRead  PermissionKey = "wallet.read"
	PermissionKeyWalletWrite PermissionKey = "wallet.write"

	PermissionKeyGiftCardPolicyRead  PermissionKey = "gift_card_policy.read"
	PermissionKeyGiftCardPolicyWrite PermissionKey = "gift_card_policy.write"

	PermissionKeyPromotionRead    PermissionKey = "promotion.read"
	PermissionKeyPromotionWrite   PermissionKey = "promotion.write"
	PermissionKeyPromotionPublish PermissionKey = "promotion.publish"
	PermissionKeyCouponRead       PermissionKey = "coupon.read"
	PermissionKeyCouponWrite      PermissionKey = "coupon.write"

	PermissionKeyPriceRead    PermissionKey = "price.read"
	PermissionKeyPriceApprove PermissionKey = "price.approve"

	PermissionKeyMarketingRead        PermissionKey = "marketing.read"
	PermissionKeyMarketingSend        PermissionKey = "marketing.send"
	PermissionKeyAnalyticsSalesRead   PermissionKey = "analytics.sales.read"
	PermissionKeyAnalyticsProductRead PermissionKey = "analytics.product.read"
	PermissionKeyAnalyticsExport      PermissionKey = "analytics.export"

	PermissionKeyMediaRead   PermissionKey = "media.read"
	PermissionKeyMediaUpload PermissionKey = "media.upload"
	PermissionKeyMediaDelete PermissionKey = "media.delete"

	PermissionKeySettingsRead  PermissionKey = "settings.read"
	PermissionKeySettingsWrite PermissionKey = "settings.write"

	PermissionKeyAuditRead PermissionKey = "audit.read"

	PermissionKeySecurityRead  PermissionKey = "security.read"
	PermissionKeySecurityWrite PermissionKey = "security.write"

	PermissionKeyAccessLogRead PermissionKey = "access_log.read"

	PermissionKeyWholesaleRead  PermissionKey = "wholesale.read"
	PermissionKeyWholesaleWrite PermissionKey = "wholesale.write"

	PermissionKeyProductRead     PermissionKey = "product.read"
	PermissionKeyProductWrite    PermissionKey = "product.write"
	PermissionKeyProductDelete   PermissionKey = "product.delete"
	PermissionKeySKURead         PermissionKey = "sku.read"
	PermissionKeySKUWrite        PermissionKey = "sku.write"
	PermissionKeySKUProductsRead PermissionKey = "sku.products.read"
	PermissionKeyCategoryRead    PermissionKey = "category.read"
	PermissionKeyCategoryWrite   PermissionKey = "category.write"
	PermissionKeyCollectionRead  PermissionKey = "collection.read"
	PermissionKeyCollectionWrite PermissionKey = "collection.write"

	PermissionKeyDepotRead     PermissionKey = "depot.read"
	PermissionKeyDepotWrite    PermissionKey = "depot.write"
	PermissionKeyLocationRead  PermissionKey = "location.read"
	PermissionKeyLocationWrite PermissionKey = "location.write"

	PermissionKeyStockRead     PermissionKey = "stock.read"
	PermissionKeyStockAdjust   PermissionKey = "stock.adjust"
	PermissionKeyStockReserve  PermissionKey = "stock.reserve"
	PermissionKeyStockTransfer PermissionKey = "stock.transfer"

	PermissionKeyPickingRead   PermissionKey = "picking.read"
	PermissionKeyPickingWrite  PermissionKey = "picking.write"
	PermissionKeyPackingRead   PermissionKey = "packing.read"
	PermissionKeyPackingWrite  PermissionKey = "packing.write"
	PermissionKeyShipmentRead  PermissionKey = "shipment.read"
	PermissionKeyShipmentWrite PermissionKey = "shipment.write"

	PermissionKeyPurchaseRead    PermissionKey = "purchase.read"
	PermissionKeyPurchaseWrite   PermissionKey = "purchase.write"
	PermissionKeyPurchasePublish PermissionKey = "purchase.publish"
	PermissionKeyReceiptRead     PermissionKey = "receipt.read"
	PermissionKeyReceiptWrite    PermissionKey = "receipt.write"

	PermissionKeyExpiryRead    PermissionKey = "expiry.read"
	PermissionKeyExpiryRun     PermissionKey = "expiry.run"
	PermissionKeyForecastRead  PermissionKey = "forecast.read"
	PermissionKeyForecastWrite PermissionKey = "forecast.write"
	PermissionKeyInboundRead   PermissionKey = "inbound.read"
	PermissionKeyInboundWrite  PermissionKey = "inbound.write"
	PermissionKeyDamageRead    PermissionKey = "damage.read"
	PermissionKeyDamageWrite   PermissionKey = "damage.write"
	PermissionKeyWMSDraftRead  PermissionKey = "wmsdraft.read"
	PermissionKeyWMSDraftWrite PermissionKey = "wmsdraft.write"
	PermissionKeyLayoutRead    PermissionKey = "layout.read"
	PermissionKeyLayoutWrite   PermissionKey = "layout.write"

	PermissionKeyReviewRead     PermissionKey = "review.read"
	PermissionKeyReviewModerate PermissionKey = "review.moderate"
	PermissionKeyWishRead       PermissionKey = "wish.read"
	PermissionKeyWishManage     PermissionKey = "wish.manage"

	PermissionKeyOrderRead        PermissionKey = "order.read"
	PermissionKeyOrderWrite       PermissionKey = "order.write"
	PermissionKeyOrderCancel      PermissionKey = "order.cancel"
	PermissionKeyPaymentRead      PermissionKey = "payment.read"
	PermissionKeyPaymentCapture   PermissionKey = "payment.capture"
	PermissionKeyRefundRead       PermissionKey = "refund.read"
	PermissionKeyRefundWrite      PermissionKey = "refund.write"
	PermissionKeyRefundRequest    PermissionKey = "refund.request"
	PermissionKeyRefundApprove    PermissionKey = "refund.approve"
	PermissionKeyInvoiceRead      PermissionKey = "invoice.read"
	PermissionKeyInvoiceWrite     PermissionKey = "invoice.write"
	PermissionKeyCartManage       PermissionKey = "cart.manage"
	PermissionKeyTerminalManage   PermissionKey = "terminal.manage"
	PermissionKeyTerminalTransact PermissionKey = "terminal.transact"
	PermissionKeyCommerceConfig   PermissionKey = "commerce.config"
	PermissionKeyPreorderRead     PermissionKey = "preorder.read"
	PermissionKeyPreorderWrite    PermissionKey = "preorder.write"

	PermissionKeyPOSAccess        PermissionKey = "pos.access"
	PermissionKeyPOSSessionManage PermissionKey = "pos.session.manage"
)

// IsValid reports whether key is one of the shared workforce permission keys.
// The switch avoids a runtime allocation and a mutable global catalogue.
func (k PermissionKey) IsValid() bool {
	switch k {
	case PermissionKeyUserRead, PermissionKeyUserWrite, PermissionKeyUserDelete,
		PermissionKeyRetailCustomerRead, PermissionKeyRetailCustomerWrite,
		PermissionKeySupplierRead, PermissionKeySupplierWrite,
		PermissionKeyRoleRead, PermissionKeyRoleWrite,
		PermissionKeyMembershipRead, PermissionKeyMembershipWrite,
		PermissionKeyWalletRead, PermissionKeyWalletWrite,
		PermissionKeyGiftCardPolicyRead, PermissionKeyGiftCardPolicyWrite,
		PermissionKeyPromotionRead, PermissionKeyPromotionWrite, PermissionKeyPromotionPublish,
		PermissionKeyCouponRead, PermissionKeyCouponWrite,
		PermissionKeyPriceRead, PermissionKeyPriceApprove,
		PermissionKeyMarketingRead, PermissionKeyMarketingSend,
		PermissionKeyAnalyticsSalesRead, PermissionKeyAnalyticsProductRead, PermissionKeyAnalyticsExport,
		PermissionKeyMediaRead, PermissionKeyMediaUpload, PermissionKeyMediaDelete,
		PermissionKeySettingsRead, PermissionKeySettingsWrite,
		PermissionKeyAuditRead,
		PermissionKeySecurityRead, PermissionKeySecurityWrite,
		PermissionKeyAccessLogRead,
		PermissionKeyWholesaleRead, PermissionKeyWholesaleWrite,
		PermissionKeyProductRead, PermissionKeyProductWrite, PermissionKeyProductDelete,
		PermissionKeySKURead, PermissionKeySKUWrite, PermissionKeySKUProductsRead,
		PermissionKeyCategoryRead, PermissionKeyCategoryWrite,
		PermissionKeyCollectionRead, PermissionKeyCollectionWrite,
		PermissionKeyDepotRead, PermissionKeyDepotWrite,
		PermissionKeyLocationRead, PermissionKeyLocationWrite,
		PermissionKeyStockRead, PermissionKeyStockAdjust, PermissionKeyStockReserve, PermissionKeyStockTransfer,
		PermissionKeyPickingRead, PermissionKeyPickingWrite,
		PermissionKeyPackingRead, PermissionKeyPackingWrite,
		PermissionKeyShipmentRead, PermissionKeyShipmentWrite,
		PermissionKeyPurchaseRead, PermissionKeyPurchaseWrite, PermissionKeyPurchasePublish,
		PermissionKeyReceiptRead, PermissionKeyReceiptWrite,
		PermissionKeyExpiryRead, PermissionKeyExpiryRun,
		PermissionKeyForecastRead, PermissionKeyForecastWrite,
		PermissionKeyInboundRead, PermissionKeyInboundWrite,
		PermissionKeyDamageRead, PermissionKeyDamageWrite,
		PermissionKeyWMSDraftRead, PermissionKeyWMSDraftWrite,
		PermissionKeyLayoutRead, PermissionKeyLayoutWrite,
		PermissionKeyReviewRead, PermissionKeyReviewModerate,
		PermissionKeyWishRead, PermissionKeyWishManage,
		PermissionKeyOrderRead, PermissionKeyOrderWrite, PermissionKeyOrderCancel,
		PermissionKeyPaymentRead, PermissionKeyPaymentCapture,
		PermissionKeyRefundRead, PermissionKeyRefundWrite, PermissionKeyRefundRequest, PermissionKeyRefundApprove,
		PermissionKeyInvoiceRead, PermissionKeyInvoiceWrite,
		PermissionKeyCartManage,
		PermissionKeyTerminalManage, PermissionKeyTerminalTransact,
		PermissionKeyCommerceConfig,
		PermissionKeyPreorderRead, PermissionKeyPreorderWrite,
		PermissionKeyPOSAccess, PermissionKeyPOSSessionManage:
		return true
	default:
		return false
	}
}

func (k PermissionKey) String() string { return string(k) }
