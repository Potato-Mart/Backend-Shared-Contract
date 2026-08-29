package pkg_test

import (
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestDomainPackageLayout verifies the domain package paths after every
// production model has been split into its own source file.
func TestDomainPackageLayout(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	requiredFiles := map[string]string{
		"contracts/identity/authorisation/role.go":                                "authorisation",
		"contracts/customers/group/group_order_manager_application.go":            "group",
		"contracts/customers/preference/retail_customer_receipt_preferences.go":   "preference",
		"contracts/orders/cart/cart.go":                                           "cart",
		"contracts/orders/buyer/buyer_context.go":                                 "buyer",
		"contracts/orders/group_order/group_order_context.go":                     "group_order",
		"contracts/orders/fulfilment/order_packing_progress.go":                   "fulfilment",
		"contracts/orders/subscription/subscription_plan.go":                      "subscription",
		"contracts/payments/payment/payment.go":                                   "payment",
		"contracts/payments/provider/stripe_payment_reference.go":                 "provider",
		"contracts/payments/merchant/merchant_legal_profile.go":                   "merchant",
		"contracts/payments/receipt/receipt_snapshot.go":                          "receipt",
		"contracts/payments/register/register.go":                                 "register",
		"contracts/pricing/coupon/coupon.go":                                      "coupon",
		"contracts/pricing/special/voucher.go":                                    "special",
		"contracts/pricing/wallet/balance/customer_wallet.go":                     "balance",
		"contracts/pricing/wallet/ledger/point_ledger_entry.go":                   "ledger",
		"contracts/pricing/wallet/points/points_summary.go":                       "points",
		"contracts/pricing/wallet/giftcard/gift_card.go":                          "giftcard",
		"contracts/pricing/wallet/reward/reward_redemption.go":                    "reward",
		"contracts/pricing/wallet/reservation/checkout_benefit_reservation.go":    "reservation",
		"contracts/notification/core/notification.go":                             "core",
		"contracts/notification/delivery/notification_delivery.go":                "delivery",
		"contracts/insights/customer/retail_customer_analytics_profile.go":        "customer",
		"contracts/insights/sales/order_item_fact.go":                             "sales",
		"contracts/marketing/audience/audience.go":                                "audience",
		"contracts/marketing/message/marketing_message.go":                        "message",
		"contracts/supply/catalogue/classification/brand.go":                      "classification",
		"contracts/supply/catalogue/favourite/favourite_list.go":                  "favourite",
		"contracts/supply/catalogue/listing/market_listing.go":                    "listing",
		"contracts/supply/catalogue/product/product.go":                           "product",
		"contracts/supply/catalogue/review/review.go":                             "review",
		"contracts/supply/catalogue/wish/wish_proposal.go":                        "wish",
		"contracts/supply/compliance/manufacturer_declaration.go":                 "compliance",
		"contracts/supply/forecasting/sku_demand_forecast.go":                     "forecasting",
		"contracts/supply/fulfilment/outbound_shipment.go":                        "fulfilment",
		"contracts/supply/inventory/stock_movement.go":                            "inventory",
		"contracts/supply/procurement/supplier_invoice.go":                        "procurement",
		"contracts/supply/warehouse/inbound_receipt.go":                           "warehouse",
		"contracts/pubsub/envelope/event_envelope.go":                             "envelope",
		"contracts/pubsub/routing/event_type.go":                                  "routing",
		"contracts/pubsub/orders/order_created_event.go":                          "orders",
		"contracts/pubsub/payments/payment_captured_event.go":                     "payments",
		"contracts/pubsub/supply/inventory_lot_received_event.go":                 "supply",
		"contracts/pubsub/customers/customer_registered_event.go":                 "customers",
		"contracts/pubsub/pricing/price_changed_event.go":                         "pricing",
		"contracts/pubsub/notification/notification_preferences_changed_event.go": "notification",
	}

	for relativePath, wantPackage := range requiredFiles {
		path := filepath.Join(pkgRoot, filepath.FromSlash(relativePath))
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.PackageClauseOnly)
		if err != nil {
			t.Errorf("required model file %s is missing or invalid: %v", relativePath, err)
			continue
		}
		if got := file.Name.Name; got != wantPackage {
			t.Errorf("%s package = %s, want %s", relativePath, got, wantPackage)
		}
	}

	for _, relativePath := range []string{
		"contracts/supply/category",
		"contracts/supply/favourite",
		"contracts/supply/import_compliance",
		"contracts/supply/product/brand.go",
		"contracts/supply/product/category_tag.go",
		"contracts/supply/product/collection.go",
		"contracts/supply/product/supply.go",
		"contracts/supply/purchase/supplier.go",
		"contracts/supply/warehouse/packing.go",
		"contracts/supply/warehouse/picking.go",
		"contracts/supply/warehouse/reservation.go",
		"contracts/supply/warehouse/inbound.go",
		"contracts/supply/warehouse/stock_movement.go",
		"contracts/supply/warehouse/shipment.go",
		"contracts/pricing/promotion/coupon.go",
		"contracts/pricing/promotion/promotion_enums/coupon_type.go",
		"contracts/pricing/promotion/effective_promotion.go",
		"contracts/pricing/promotion/group_order_discount.go",
		"contracts/pricing/promotion/receipt_offer.go",
		"contracts/pricing/promotion/shared.go",
		"contracts/pricing/promotion/storefront_promotion.go",
		"contracts/supply/classification/sku.go",
		"contracts/pricing/promotion/package_pricing.go",
		"contracts/supply/product/availability.go",
		"contracts/supply/product/detail_image.go",
		"contracts/supply/product/offer.go",
		"contracts/orders/order/volume_discount.go",
	} {
		path := filepath.Join(pkgRoot, filepath.FromSlash(relativePath))
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Errorf("retired path must not exist: %s", relativePath)
		}
	}
}

func TestSourcesRejectRetiredComplianceImportPath(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	retiredIdentifier := "contracts/supply/import" + "_compliance"
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			return nil
		}
		if strings.HasPrefix(relativePkgPath(t, pkgRoot, path), "test/") {
			return nil
		}
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.Contains(string(contents), retiredIdentifier) {
			t.Errorf("%s retains retired %s identifier", relativePkgPath(t, pkgRoot, path), retiredIdentifier)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan source layout: %v", err)
	}
}
