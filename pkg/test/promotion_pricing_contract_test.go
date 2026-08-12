package pkg_test

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	order "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/pos"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/promotion/promotion_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse"
)

func TestV27PromotionScopeGrammarRoundTripsAllSelectorsAndQuantityRanges(t *testing.T) {
	maximum := int64(12)
	scope := promotion.PromotionScope{
		MatchMode: promotion_enums.PromotionMatchModeAll,
		Groups: []promotion.PromotionScopeGroup{
			{
				MatchMode:        promotion_enums.PromotionMatchModeAny,
				ProductSKUCodes:  []string{"POTATO-A", "POTATO-B"},
				MinimumBaseUnits: 3,
				MaximumBaseUnits: &maximum,
			},
			{
				MatchMode:        promotion_enums.PromotionMatchModeAll,
				CollectionIDs:    []string{"collection-weekly", "collection-seasonal"},
				CategoryTagIDs:   []string{"tag-organic", "tag-local"},
				PackageOptionIDs: []string{"package-each", "package-case"},
				MinimumBaseUnits: 1,
			},
		},
	}

	if !v27PromotionScopeStructurallyUsable(scope) {
		t.Fatal("populated restricted scope must be structurally usable")
	}
	if v27PromotionScopeStructurallyUsable(promotion.PromotionScope{MatchMode: promotion_enums.PromotionMatchModeAll}) {
		t.Fatal("an empty restricted scope must not be structurally usable")
	}
	if !v27PromotionScopeStructurallyUsable(promotion.PromotionScope{Unrestricted: true, MatchMode: promotion_enums.PromotionMatchModeAny}) {
		t.Fatal("an explicitly unrestricted scope must be structurally usable without groups")
	}

	payload, err := json.Marshal(scope)
	if err != nil {
		t.Fatalf("marshal promotion scope: %v", err)
	}
	var got promotion.PromotionScope
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal promotion scope: %v", err)
	}
	if got.MatchMode != promotion_enums.PromotionMatchModeAll || len(got.Groups) != 2 {
		t.Fatalf("outer ALL scope changed: %+v", got)
	}
	if first := got.Groups[0]; first.MatchMode != promotion_enums.PromotionMatchModeAny || len(first.ProductSKUCodes) != 2 || first.MaximumBaseUnits == nil || *first.MaximumBaseUnits != 12 {
		t.Fatalf("product quantity-pool group changed: %+v", first)
	}
	if second := got.Groups[1]; second.MatchMode != promotion_enums.PromotionMatchModeAll || len(second.CollectionIDs) != 2 || len(second.CategoryTagIDs) != 2 || len(second.PackageOptionIDs) != 2 || second.MaximumBaseUnits != nil {
		t.Fatalf("collection/tag/package group or unlimited maximum changed: %+v", second)
	}
}

func TestV27PackagePricingPreservesOrderedPromotionApplications(t *testing.T) {
	appliedAt := time.Date(2026, 8, 9, 9, 30, 0, 0, time.UTC)
	value := promotion.PackagePricing{
		ID: "pricing-1", Revision: 5, InventoryRevision: 8,
		ProductSKUCode: "POTATO-A", PackageOptionID: "package-each",
		PackagePrice: money.Money{AmountMinor: 800, Currency: "AUD"},
		TaxAmount:    money.Money{AmountMinor: 73, Currency: "AUD"},
		ValidFrom:    appliedAt, Timezone: "Australia/Sydney",
		StockLocation: warehouse.StockLocationRef{DepotCode: "SYD-01", LocationCode: "A-1"},
		PromotionApplications: []promotion.PromotionApplication{
			{PromotionID: "promotion-membership", PromotionKind: "membership_discount", PromotionRevision: 2, RelationID: "relation-membership", ResolvedTargetProductSKUCodes: []string{"POTATO-A"}, AppliedAt: appliedAt},
			{PromotionID: "promotion-bundle", PromotionKind: "future_bundle_kind", PromotionRevision: 7, RelationID: "relation-bundle", ResolvedQualifierProductSKUCodes: []string{"POTATO-A", "POTATO-B"}, ResolvedTargetProductSKUCodes: []string{"POTATO-A"}, AppliedAt: appliedAt},
		},
		CapturedAt: appliedAt,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal package pricing: %v", err)
	}
	var got promotion.PackagePricing
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal package pricing: %v", err)
	}
	if len(got.PromotionApplications) != 2 || got.PromotionApplications[0].PromotionID != "promotion-membership" || got.PromotionApplications[1].PromotionID != "promotion-bundle" {
		t.Fatalf("promotion stacking order changed: %+v", got.PromotionApplications)
	}
	if got.PromotionApplications[1].RelationID != "relation-bundle" || len(got.PromotionApplications[1].ResolvedQualifierProductSKUCodes) != 2 || len(got.PromotionApplications[1].ResolvedTargetProductSKUCodes) != 1 {
		t.Fatalf("resolved qualifier-to-target evidence changed: %+v", got.PromotionApplications[1])
	}

	applicationType := reflect.TypeOf(promotion.PromotionApplication{})
	for fieldName, wantTag := range map[string]string{
		"RelationID":                       "relation_id",
		"ResolvedQualifierProductSKUCodes": "resolved_qualifier_product_sku_codes,omitempty",
		"ResolvedTargetProductSKUCodes":    "resolved_target_product_sku_codes,omitempty",
	} {
		field, ok := applicationType.FieldByName(fieldName)
		if !ok || field.Tag.Get("json") != wantTag {
			t.Fatalf("PromotionApplication.%s JSON tag = %q, want %q", fieldName, field.Tag.Get("json"), wantTag)
		}
	}
}

func TestV27CustomerSummariesExposeOnlyFrozenSafePricingFacts(t *testing.T) {
	for _, model := range []reflect.Type{reflect.TypeOf(pos.ReceiptLine{}), reflect.TypeOf(order.OrderLineSummary{})} {
		for _, required := range []string{"ProductSKUCode", "ProductName", "ProductImage", "ProductPackageOption", "CapturedAt", "PackagePrice", "TaxAmount", "DiscountAmount", "PromotionApplications", "Total"} {
			if _, ok := model.FieldByName(required); !ok {
				t.Errorf("%s is missing frozen customer-safe field %s", model, required)
			}
		}
		for _, forbidden := range []string{"AcceptedPackagePricing", "PackagePricing", "Components", "StockLocation", "SourceBucketID", "SourceStockUnitID", "AvailablePackageCount", "AvailableBaseUnits", "InventoryRevision"} {
			if _, ok := model.FieldByName(forbidden); ok {
				t.Errorf("%s exposes internal pricing/inventory field %s", model, forbidden)
			}
		}
		for i := 0; i < model.NumField(); i++ {
			if model.Field(i).Type == reflect.TypeOf(promotion.PackagePricing{}) {
				t.Errorf("%s embeds raw PackagePricing through field %s", model, model.Field(i).Name)
			}
		}
	}
}

func TestV27OperationalCategoryTagEvidenceIsLocationQualified(t *testing.T) {
	model := reflect.TypeOf(operations.InventoryCategoryTagEvidence{})
	assertJSONFields(t, model, map[string]string{
		"ProductSKUCode":  "product_sku_code",
		"PackageOptionID": "package_option_id",
		"CategoryTag":     "category_tag",
		"StockLocation":   "stock_location",
		"Condition":       "condition",
		"Disposition":     "disposition",
		"DateMark":        "date_mark,omitempty",
		"AsOf":            "as_of",
	})
	field, _ := model.FieldByName("StockLocation")
	if field.Type != reflect.TypeOf(warehouse.StockLocationRef{}) {
		t.Fatalf("InventoryCategoryTagEvidence.StockLocation type = %s, want warehouse.StockLocationRef", field.Type)
	}
}

func TestV27PromotionProductionSurfaceRejectsRetiredMechanicsAndDependencies(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	retiredIdentifiers := map[string]struct{}{
		"SellableOffer": {}, "SellableOfferSnapshot": {}, "AcceptedOffer": {},
		"ReceiptOffer": {}, "AppliedPromotion": {}, "EffectivePromotion": {},
		"StorefrontPromotion": {}, "PromotionProduct": {}, "VolumeDiscountTier": {},
		"PointPromotion": {}, "MembershipPromotionTarget": {},
	}
	retiredWireFragments := []string{"same_sale_order", "accepted_offer", "offer_id", "offer_revision", "sellable_offer"}
	offerWord := regexp.MustCompile(`(?i)\boffers?\b`)
	closedPromotionEnums := make(map[string]struct{})

	err := filepath.WalkDir(filepath.Join(pkgRoot, "contracts"), func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		source, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		text := string(source)
		lower := strings.ToLower(text)
		if offerWord.MatchString(text) {
			t.Errorf("%s retains retired offer terminology", path)
		}
		for _, fragment := range retiredWireFragments {
			if strings.Contains(lower, fragment) {
				t.Errorf("%s retains retired wire fragment %q", path, fragment)
			}
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, source, 0)
		if parseErr != nil {
			return parseErr
		}
		ast.Inspect(file, func(node ast.Node) bool {
			identifier, ok := node.(*ast.Ident)
			if ok {
				if _, retired := retiredIdentifiers[identifier.Name]; retired {
					t.Errorf("%s retains retired identifier %s", path, identifier.Name)
				}
			}
			return true
		})

		relative := filepath.ToSlash(relativePkgPath(t, pkgRoot, path))
		if strings.HasPrefix(relative, "contracts/pricing/promotion/") {
			for _, spec := range file.Imports {
				if strings.Contains(spec.Path.Value, "/supply/product") {
					t.Errorf("%s imports product and would reverse the canonical product dependency", path)
				}
			}
			if strings.Contains(lower, "soon_expiry") || strings.Contains(lower, "soon-expiry") || strings.Contains(lower, `"damaged"`) {
				t.Errorf("%s defines operational inventory evidence as a promotion mechanic", path)
			}
		}
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpecification, ok := specification.(*ast.TypeSpec)
				if !ok || !typeSpecification.Name.IsExported() || !strings.Contains(typeSpecification.Name.Name, "Promotion") {
					continue
				}
				underlying, stringBacked := typeSpecification.Type.(*ast.Ident)
				if stringBacked && underlying.Name == "string" {
					closedPromotionEnums[typeSpecification.Name.Name] = struct{}{}
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	var enumNames []string
	for name := range closedPromotionEnums {
		enumNames = append(enumNames, name)
	}
	sort.Strings(enumNames)
	wantEnums := []string{"PromotionMatchMode", "PromotionStatus"}
	if !reflect.DeepEqual(enumNames, wantEnums) {
		t.Fatalf("closed promotion enums = %v, want only %v", enumNames, wantEnums)
	}
}

func v27PromotionScopeStructurallyUsable(scope promotion.PromotionScope) bool {
	if scope.Unrestricted {
		return scope.MatchMode.IsValid()
	}
	if !scope.MatchMode.IsValid() || len(scope.Groups) == 0 {
		return false
	}
	for _, group := range scope.Groups {
		if !group.MatchMode.IsValid() {
			return false
		}
		selectorCount := len(group.ProductSKUCodes) + len(group.CollectionIDs) + len(group.CategoryTagIDs) + len(group.PackageOptionIDs)
		if selectorCount == 0 || group.MinimumBaseUnits < 0 {
			return false
		}
		if group.MaximumBaseUnits != nil && *group.MaximumBaseUnits < group.MinimumBaseUnits {
			return false
		}
	}
	return true
}
