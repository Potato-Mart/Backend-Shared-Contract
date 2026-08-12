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

// v27ModelPackageManifest classifies every production package. Adding an
// exported type changes the digest below and requires an explicit manifest
// review instead of silently expanding the contract surface.
var v27ModelPackageManifest = map[string]string{
	"contracts/common/apiresponse/apiresponse_enums":             "enum",
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
	"contracts/customers/campaign":                               "entity,event,record,snapshot",
	"contracts/customers/campaign/campaign_enums":                "enum",
	"contracts/customers/retail":                                 "entity,event,record,snapshot",
	"contracts/customers/retail/retail_enums":                    "enum",
	"contracts/customers/wholesale":                              "entity,record,snapshot",
	"contracts/customers/wholesale/wholesale_enums":              "enum",
	"contracts/identity/access":                                  "claims,event,record,session",
	"contracts/identity/access/access_enums":                     "enum",
	"contracts/identity/account":                                 "entity,event,record",
	"contracts/identity/account/account_enums":                   "enum",
	"contracts/identity/role":                                    "entity,event,record",
	"contracts/identity/role/role_enums":                         "enum",
	"contracts/insights/analytics":                               "record",
	"contracts/insights/marketing":                               "entity,record",
	"contracts/insights/marketing/marketing_enums":               "enum",
	"contracts/notifications/backinstock":                        "entity,event,record",
	"contracts/notifications/backinstock/backinstock_enums":      "enum",
	"contracts/notifications/customer":                           "entity,event,record",
	"contracts/notifications/customer/customer_enums":            "enum",
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
	"contracts/pricing/membership":                               "entity,record,value",
	"contracts/pricing/membership/membership_enums":              "enum",
	"contracts/pricing/promotion":                                "entity,event,record,value",
	"contracts/pricing/promotion/promotion_enums":                "enum",
	"contracts/pricing/wallet":                                   "entity,record,snapshot,value",
	"contracts/pricing/wallet/wallet_enums":                      "enum",
	"contracts/pubsub/envelop":                                   "record",
	"contracts/pubsub/event":                                     "event,record",
	"contracts/pubsub/event/event_enums":                         "enum",
	"contracts/supply/classification":                            "entity,record,value",
	"contracts/supply/classification/classification_enums":       "enum",
	"contracts/supply/import_compliance":                         "entity,record,snapshot,value",
	"contracts/supply/import_compliance/import_compliance_enums": "enum",
	"contracts/supply/operations":                                "entity,event,record,snapshot,value",
	"contracts/supply/product":                                   "entity,event,snapshot,value",
	"contracts/supply/product/product_enums":                     "enum",
	"contracts/supply/purchase":                                  "entity,record",
	"contracts/supply/purchase/purchase_enums":                   "enum",
	"contracts/supply/review":                                    "entity,record,value",
	"contracts/supply/review/review_enums":                       "enum",
	"contracts/supply/warehouse":                                 "entity,event,record,snapshot,value",
	"contracts/supply/warehouse/warehouse_enums":                 "enum",
	"contracts/supply/wish":                                      "entity,record,value",
	"contracts/supply/wish/wish_enums":                           "enum",
	"versioning":                                                 "module-metadata",
}

// Reviewed for the final v27.0.0 object-media, canonical-product, package-layout,
// unified-promotion, and package-pricing release surface. The digest captures
// the complete exported model manifest after the hard cutover.
const v27ExportedTypeManifestDigest = "cd8db873980ca014009903fb145e7f0dc57779551d9375a524cbcf7b4b5de020"

func TestV27ExportedTypesMatchModelManifest(t *testing.T) {
	seenPackages := make(map[string]bool)
	var entries []string
	pkgRoot := sharedContractPkgRoot(t)
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
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
		class, classified := v27ModelPackageManifest[packagePath]
		if !classified {
			t.Errorf("%s is not classified in the v27 model manifest", packagePath)
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
		t.Fatalf("read v27 model manifest: %v", err)
	}
	for packagePath := range v27ModelPackageManifest {
		if !seenPackages[packagePath] {
			t.Errorf("manifest package %s has no production source", packagePath)
		}
	}
	sort.Strings(entries)
	sum := sha256.Sum256([]byte(strings.Join(entries, "\n")))
	got := hex.EncodeToString(sum[:])
	if got != v27ExportedTypeManifestDigest {
		t.Fatalf("exported model manifest changed: got %s; classify the change and update the reviewed digest", got)
	}
}
