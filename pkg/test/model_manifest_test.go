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

// v28ModelPackageManifest classifies every production package. Adding an
// exported type changes the digest below and requires an explicit manifest
// review instead of silently expanding the contract surface.
var v28ModelPackageManifest = map[string]string{
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
	"contracts/identity/deletion":                                "record,snapshot",
	"contracts/identity/deletion/deletion_enums":                 "enum",
	"contracts/identity/role":                                    "entity,event,record",
	"contracts/identity/role/role_enums":                         "enum",
	"contracts/insights/analytics":                               "record",
	"contracts/insights/marketing":                               "entity,record",
	"contracts/insights/marketing/marketing_enums":               "enum",
	"contracts/marketing":                                        "entity,event,record,value",
	"contracts/marketing/marketing_enums":                        "enum",
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
	"contracts/supply/review":                                    "entity,record,value",
	"contracts/supply/review/review_enums":                       "enum",
	"contracts/supply/warehouse":                                 "entity,event,record,snapshot,value",
	"contracts/supply/warehouse/warehouse_enums":                 "enum",
	"contracts/supply/wish":                                      "entity,record,value",
	"contracts/supply/wish/wish_enums":                           "enum",
	"versioning":                                                 "module-metadata",
}

// Reviewed for the v28.0.0 RBAC and geographic-scoping cut-over. The digest
// captures the complete exported model manifest, so the change below was
// classified explicitly before the digest was moved.
//
// v28.0.0 adds three exported types and removes two, with no manifest
// classification change:
//
//   - access.StaffGeoScope — the persisted workforce geographic grant, a
//     record under contracts/identity/access's existing "record" class.
//   - access_enums.ScopeLevel — global/country/market/depot, an enum under
//     contracts/identity/access/access_enums' existing "enum" class.
//   - pos.RegisterSession and pos.SessionTotalsSnapshot replace
//     pos.RegisterShift and pos.ShiftTotalsSnapshot, keeping the
//     contracts/orders/pos "entity" and "snapshot" classes. pos_enums
//     .SessionStatus likewise replaces ShiftStatus inside the existing "enum"
//     class.
//
// Every other v28 change is a field on an existing type — the flat scope
// claims, Role.rank, the gift-card policy's market/country/audit fields, and
// the denormalized market_code/country_code/depot_code geography — and fields do
// not enter this digest.
//
// Re-reviewed for the pre-tag v28.0.0 correction pass, which deliberately
// leaves this digest UNCHANGED. The digest hashes package|class|TypeName
// triples only, so none of that pass's changes can reach it:
//
//   - Dropping the `sales` workforce rank removes a constant, not a type;
//     role_enums.UserRole itself is untouched.
//   - membership.MemberSubscription.market_code/country_code,
//     analytics.OrderItemFact/RefundItemFact/MetricRollup/SKUDemandForecast
//     .country_code, and MetricRollup.market_code are all fields on types that
//     already exist and already carry their manifest class.
//
// An unmoved digest is therefore the correct, reviewed outcome here rather
// than a skipped gate: this comment is the review record, and the guards that
// actually police the change are the workforce wire lock in pkg/test/enums
// and the removed-identifier set in hard_cutover_test.go.
const v28ExportedTypeManifestDigest = "bea3ad8d5f4da90a6125c0f767f275354a399c16ccd5bec50a552826dea2b817"

func TestV28ExportedTypesMatchModelManifest(t *testing.T) {
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
		class, classified := v28ModelPackageManifest[packagePath]
		if !classified {
			t.Errorf("%s is not classified in the v28 model manifest", packagePath)
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
		t.Fatalf("read v28 model manifest: %v", err)
	}
	for packagePath := range v28ModelPackageManifest {
		if !seenPackages[packagePath] {
			t.Errorf("manifest package %s has no production source", packagePath)
		}
	}
	sort.Strings(entries)
	sum := sha256.Sum256([]byte(strings.Join(entries, "\n")))
	got := hex.EncodeToString(sum[:])
	if got != v28ExportedTypeManifestDigest {
		t.Fatalf("exported model manifest changed: got %s; classify the change and update the reviewed digest", got)
	}
}
