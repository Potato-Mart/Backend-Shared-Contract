package pkg_test

import (
	"crypto/sha256"
	"encoding/hex"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// v32ModelPackageManifest retains the completed pre-v33 package inventory
// for the migration review. The active v33 inventory follows it.
var v32ModelPackageManifest = map[string]string{
	"contracts/common/audit":                                     "record,value",
	"contracts/common/commerce/commerce_enums":                   "enum",
	"contracts/common/device":                                    "record",
	"contracts/common/geography":                                 "record,value",
	"contracts/common/geography/geography_enums":                 "enum",
	"contracts/common/geometry":                                  "value",
	"contracts/common/identity":                                  "record",
	"contracts/common/identity/identity_enums":                   "enum",
	"contracts/common/localization":                              "value",
	"contracts/common/measurement":                               "value",
	"contracts/common/metadata":                                  "value",
	"contracts/common/money":                                     "value",
	"contracts/common/packaging":                                 "record,value",
	"contracts/common/packaging/packaging_enums":                 "enum",
	"contracts/common/party":                                     "record,value",
	"contracts/common/security":                                  "event,record,value",
	"contracts/common/security/security_enums":                   "enum",
	"contracts/common/temporal":                                  "value",
	"contracts/customers/retail":                                 "entity,event,record,snapshot",
	"contracts/customers/retail/retail_enums":                    "enum",
	"contracts/customers/wholesale":                              "entity,record,snapshot,value",
	"contracts/customers/wholesale/wholesale_enums":              "enum",
	"contracts/identity/access":                                  "event,record,session",
	"contracts/identity/access/access_enums":                     "enum",
	"contracts/identity/account":                                 "entity,event,record",
	"contracts/identity/account/account_enums":                   "enum",
	"contracts/identity/role":                                    "entity,event,record,value",
	"contracts/identity/role/role_enums":                         "enum",
	"contracts/insights/analytics":                               "record",
	"contracts/insights/analytics/analytics_enums":               "enum",
	"contracts/insights/marketing":                               "record",
	"contracts/insights/marketing/marketing_enums":               "enum",
	"contracts/marketing/campaign":                               "entity,record",
	"contracts/marketing/campaign/campaign_enums":                "enum",
	"contracts/marketing/message":                                "entity,record",
	"contracts/marketing/message/message_enums":                  "enum",
	"contracts/notifications":                                    "entity,record,value",
	"contracts/notifications/notification_enums":                 "enum",
	"contracts/orders/order":                                     "entity,event,snapshot,record",
	"contracts/orders/order/order_enums":                         "enum",
	"contracts/orders/pos":                                       "entity,record,snapshot",
	"contracts/orders/pos/pos_enums":                             "enum",
	"contracts/orders/shipping":                                  "record,value",
	"contracts/orders/shipping/shipping_enums":                   "enum",
	"contracts/payments/payment":                                 "entity,event,snapshot,record,value",
	"contracts/payments/payment/payment_enums":                   "enum",
	"contracts/payments/settlement":                              "entity,event,snapshot,record,value",
	"contracts/payments/settlement/settlement_enums":             "enum",
	"contracts/payments/terminal":                                "entity,event,snapshot,record,value",
	"contracts/payments/terminal/terminal_enums":                 "enum",
	"contracts/pricing/benefit":                                  "value",
	"contracts/pricing/benefit/benefit_enums":                    "enum",
	"contracts/pricing/market":                                   "entity",
	"contracts/pricing/market/market_enums":                      "enum",
	"contracts/pricing/membership":                               "entity,record,value",
	"contracts/pricing/membership/membership_enums":              "enum",
	"contracts/pricing/pricebook":                                "entity,record",
	"contracts/pricing/pricebook/pricebook_enums":                "enum",
	"contracts/pricing/promotion":                                "entity,event,record,value",
	"contracts/pricing/promotion/promotion_enums":                "enum",
	"contracts/pricing/quote":                                    "snapshot,value",
	"contracts/pricing/quote/quote_enums":                        "enum",
	"contracts/pricing/wallet":                                   "entity,record,snapshot,value",
	"contracts/pricing/wallet/wallet_enums":                      "enum",
	"contracts/pubsub/envelop":                                   "record",
	"contracts/pubsub/event":                                     "event,record",
	"contracts/pubsub/event/event_enums":                         "enum",
	"contracts/customers/review":                                 "entity,record,value",
	"contracts/customers/review/review_enums":                    "enum",
	"contracts/supply/classification":                            "entity,record,value",
	"contracts/supply/cost":                                      "entity,record",
	"contracts/supply/classification/classification_enums":       "enum",
	"contracts/supply/import_compliance":                         "entity,record,snapshot,value",
	"contracts/supply/import_compliance/import_compliance_enums": "enum",
	"contracts/supply/listing":                                   "entity,record,snapshot,value",
	"contracts/supply/listing/listing_enums":                     "enum",
	"contracts/supply/operations":                                "entity,event,record,snapshot,value",
	"contracts/supply/product":                                   "entity,event,snapshot,value",
	"contracts/supply/product/product_enums":                     "enum",
	"contracts/supply/purchase":                                  "entity,record",
	"contracts/supply/purchase/purchase_enums":                   "enum",
	"contracts/supply/warehouse":                                 "entity,event,record,snapshot,value",
	"contracts/supply/warehouse/warehouse_enums":                 "enum",
	"contracts/supply/wish":                                      "entity,record,value",
	"contracts/supply/wish/wish_enums":                           "enum",
}

// v33ModelPackageManifest classifies every populated v33 production package.
// Adding an exported type changes the digest below and requires an explicit
// manifest review instead of silently expanding the contract surface.
var v33ModelPackageManifest = map[string]string{
	"contracts/common/audit":                                         "record",
	"contracts/common/commerce/commerce_enums":                       "enum",
	"contracts/common/device":                                        "record",
	"contracts/common/geography":                                     "record",
	"contracts/common/geography/geography_enums":                     "enum",
	"contracts/common/geometry":                                      "record",
	"contracts/common/identity":                                      "record",
	"contracts/common/identity/identity_enums":                       "enum",
	"contracts/common/localization":                                  "record",
	"contracts/common/measurement":                                   "record",
	"contracts/common/metadata":                                      "record",
	"contracts/common/money":                                         "record",
	"contracts/common/packaging":                                     "record",
	"contracts/common/packaging/packaging_enums":                     "enum",
	"contracts/common/party":                                         "record",
	"contracts/common/security":                                      "record",
	"contracts/common/security/security_enums":                       "enum",
	"contracts/common/temporal":                                      "record",
	"contracts/customers/group":                                      "record",
	"contracts/customers/group/group_enums":                          "enum",
	"contracts/customers/preference":                                 "record",
	"contracts/customers/preference/preference_enums":                "enum",
	"contracts/customers/retail":                                     "record",
	"contracts/customers/retail/retail_enums":                        "enum",
	"contracts/customers/wholesale":                                  "record",
	"contracts/customers/wholesale/wholesale_enums":                  "enum",
	"contracts/identity/access":                                      "record",
	"contracts/identity/access/access_enums":                         "enum",
	"contracts/identity/account":                                     "record",
	"contracts/identity/account/account_enums":                       "enum",
	"contracts/identity/authorisation":                               "record",
	"contracts/identity/authorisation/role_enums":                    "enum",
	"contracts/insights/analytics":                                   "record",
	"contracts/insights/analytics/analytics_enums":                   "enum",
	"contracts/insights/customer":                                    "record",
	"contracts/insights/sales":                                       "record",
	"contracts/marketing/audience":                                   "record",
	"contracts/marketing/campaign":                                   "record",
	"contracts/marketing/campaign/campaign_enums":                    "enum",
	"contracts/marketing/message":                                    "record",
	"contracts/marketing/message/message_enums":                      "enum",
	"contracts/notification/core":                                    "record",
	"contracts/notification/core/notification_enums":                 "enum",
	"contracts/notification/delivery":                                "record",
	"contracts/notification/email":                                   "record",
	"contracts/notification/preference":                              "record",
	"contracts/notification/push":                                    "record",
	"contracts/notification/sms":                                     "record",
	"contracts/orders/buyer":                                         "record",
	"contracts/orders/cart":                                          "record",
	"contracts/orders/fulfilment":                                    "record",
	"contracts/orders/group_order":                                   "record",
	"contracts/orders/group_order/group_order_enums":                 "enum",
	"contracts/orders/order":                                         "record",
	"contracts/orders/order/order_enums":                             "enum",
	"contracts/orders/shipping":                                      "record",
	"contracts/orders/shipping/shipping_enums":                       "enum",
	"contracts/orders/subscription":                                  "record",
	"contracts/orders/subscription/subscription_enums":               "enum",
	"contracts/payments/merchant":                                    "record",
	"contracts/payments/payment":                                     "record",
	"contracts/payments/payment/payment_enums":                       "enum",
	"contracts/payments/provider":                                    "record",
	"contracts/payments/receipt":                                     "record",
	"contracts/payments/register":                                    "record",
	"contracts/payments/register/register_enums":                     "enum",
	"contracts/payments/settlement":                                  "record",
	"contracts/payments/settlement/settlement_enums":                 "enum",
	"contracts/payments/terminal":                                    "record",
	"contracts/payments/terminal/terminal_enums":                     "enum",
	"contracts/pricing/benefit":                                      "record",
	"contracts/pricing/benefit/benefit_enums":                        "enum",
	"contracts/pricing/coupon":                                       "record",
	"contracts/pricing/market":                                       "record",
	"contracts/pricing/market/market_enums":                          "enum",
	"contracts/pricing/membership":                                   "record",
	"contracts/pricing/membership/membership_enums":                  "enum",
	"contracts/pricing/pricebook":                                    "record",
	"contracts/pricing/pricebook/pricebook_enums":                    "enum",
	"contracts/pricing/promotion":                                    "record",
	"contracts/pricing/promotion/promotion_enums":                    "enum",
	"contracts/pricing/quote":                                        "record",
	"contracts/pricing/quote/quote_enums":                            "enum",
	"contracts/pricing/special":                                      "record",
	"contracts/pricing/wallet/balance":                               "record",
	"contracts/pricing/wallet/giftcard":                              "record",
	"contracts/pricing/wallet/ledger":                                "record",
	"contracts/pricing/wallet/points":                                "record",
	"contracts/pricing/wallet/reservation":                           "record",
	"contracts/pricing/wallet/reward":                                "record",
	"contracts/pricing/wallet/wallet_enums":                          "enum",
	"contracts/pubsub/customers":                                     "record",
	"contracts/pubsub/envelope":                                      "record",
	"contracts/pubsub/notification":                                  "record",
	"contracts/pubsub/orders":                                        "record",
	"contracts/pubsub/payments":                                      "record",
	"contracts/pubsub/pricing":                                       "record",
	"contracts/pubsub/routing":                                       "enum",
	"contracts/pubsub/supply":                                        "record",
	"contracts/supply/catalogue/classification":                      "record",
	"contracts/supply/catalogue/classification/classification_enums": "enum",
	"contracts/supply/catalogue/favourite":                           "record",
	"contracts/supply/catalogue/favourite/favourite_enums":           "enum",
	"contracts/supply/catalogue/listing":                             "record",
	"contracts/supply/catalogue/listing/listing_enums":               "enum",
	"contracts/supply/catalogue/product":                             "record",
	"contracts/supply/catalogue/product/product_enums":               "enum",
	"contracts/supply/catalogue/review":                              "record",
	"contracts/supply/catalogue/review/review_enums":                 "enum",
	"contracts/supply/catalogue/wish":                                "record",
	"contracts/supply/catalogue/wish/wish_enums":                     "enum",
	"contracts/supply/compliance":                                    "record",
	"contracts/supply/compliance/compliance_enums":                   "enum",
	"contracts/supply/forecasting":                                   "record",
	"contracts/supply/forecasting/marketing_enums":                   "enum",
	"contracts/supply/fulfilment":                                    "record",
	"contracts/supply/inventory":                                     "record",
	"contracts/supply/procurement":                                   "record",
	"contracts/supply/procurement/purchase_enums":                    "enum",
	"contracts/supply/warehouse":                                     "record",
	"contracts/supply/warehouse/warehouse_enums":                     "enum",
}

var modelPackageManifest map[string]string

// The digest captures package|class|TypeName triples after excluding
// service-local DTO, workflow, persistence, provider-diagnostic, migration,
// and build-metadata surfaces. Field-only changes are locked by JSON-shape and
// retired-symbol tests instead.
const exportedTypeManifestDigest = "bc50b13d8ef6147e2385bb4697359250d1cbc2da0a02ee452546d8b9de9d464b"

func TestExportedTypesMatchModelManifest(t *testing.T) {
	modelPackageManifest = make(map[string]string, len(v33ModelPackageManifest))
	for packagePath, class := range v33ModelPackageManifest {
		modelPackageManifest[packagePath] = class
	}
	seenPackages := make(map[string]bool)
	var entries []string
	pkgRoot := sharedContractPkgRoot(t)
	contractsRoot := filepath.Join(pkgRoot, "contracts")
	err := filepath.WalkDir(contractsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		packagePath := filepath.ToSlash(filepath.Dir(filepath.FromSlash(relativePkgPath(t, pkgRoot, path))))
		if packagePath == "." {
			return nil
		}
		class, classified := modelPackageManifest[packagePath]
		if !classified {
			t.Errorf("%s is not classified in the model manifest", packagePath)
			return nil
		}
		seenPackages[packagePath] = true

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		for _, decl := range file.Decls {
			general, ok := decl.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, spec := range general.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if ok && typeSpec.Name.IsExported() {
					entries = append(entries, packagePath+"|"+class+"|"+typeSpec.Name.Name)
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("read model manifest: %v", err)
	}
	for packagePath := range modelPackageManifest {
		if !seenPackages[packagePath] {
			t.Errorf("manifest package %s has no production source", packagePath)
		}
	}
	sort.Strings(entries)
	sum := sha256.Sum256([]byte(strings.Join(entries, "\n")))
	got := hex.EncodeToString(sum[:])
	if got != exportedTypeManifestDigest {
		t.Fatalf("exported model manifest changed: got %s; classify the change and update the reviewed digest", got)
	}
}
