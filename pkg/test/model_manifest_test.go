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

// modelPackageManifest classifies every production package. Adding an
// exported type changes the digest below and requires an explicit manifest
// review instead of silently expanding the contract surface.
var modelPackageManifest = map[string]string{
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
	"contracts/customers/wholesale":                              "entity,record,snapshot",
	"contracts/customers/wholesale/wholesale_enums":              "enum",
	"contracts/identity/access":                                  "event,record,session",
	"contracts/identity/access/access_enums":                     "enum",
	"contracts/identity/account":                                 "entity,event,record",
	"contracts/identity/account/account_enums":                   "enum",
	"contracts/identity/role":                                    "entity,event,record",
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

// The digest captures package|class|TypeName triples after excluding
// service-local DTO, workflow, persistence, provider-diagnostic, migration,
// and build-metadata surfaces. Field-only changes are locked by JSON-shape and
// retired-symbol tests instead.
const exportedTypeManifestDigest = "225e8b2370d58d9e1d073bd2fd663a95929a6a0c71aa4d24fb0c8e8754a2feff"

func TestExportedTypesMatchModelManifest(t *testing.T) {
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
