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

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	order "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/pos"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/promotion/promotion_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/quote"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/quote/quote_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/listing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestV27PromotionScopeGrammarRoundTripsAllSelectorsAndQuantityRanges(t *testing.T) {
	maximum := int64(12)
	scope := promotion.PromotionScope{
		MatchMode: promotion_enums.PromotionMatchModeAll,
		Groups: []promotion.PromotionScopeGroup{
			{
				MatchMode:        promotion_enums.PromotionMatchModeAny,
				SKUCodes:         []string{"POTATO-A", "POTATO-B"},
				MinimumBaseUnits: 3,
				MaximumBaseUnits: &maximum,
			},
			{
				MatchMode:          promotion_enums.PromotionMatchModeAll,
				CollectionCodes:    []string{"collection-weekly", "collection-seasonal"},
				CategoryTagCodes:   []string{"tag-organic", "tag-local"},
				PackageOptionCodes: []string{"package-each", "package-case"},
				MinimumBaseUnits:   1,
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
	if first := got.Groups[0]; first.MatchMode != promotion_enums.PromotionMatchModeAny || len(first.SKUCodes) != 2 || first.MaximumBaseUnits == nil || *first.MaximumBaseUnits != 12 {
		t.Fatalf("product quantity-pool group changed: %+v", first)
	}
	if second := got.Groups[1]; second.MatchMode != promotion_enums.PromotionMatchModeAll || len(second.CollectionCodes) != 2 || len(second.CategoryTagCodes) != 2 || len(second.PackageOptionCodes) != 2 || second.MaximumBaseUnits != nil {
		t.Fatalf("collection/tag/package group or unlimited maximum changed: %+v", second)
	}
}

func TestV27PriceSnapshotFreezesRuleTaxAndEligibilityEvidence(t *testing.T) {
	capturedAt := time.Date(2026, 8, 12, 9, 30, 0, 0, time.UTC)
	value := quote.PriceSnapshot{
		QuoteID: "quote-1", LineID: "line-1", SKUCode: "sku-potato-a", MarketCode: "market-au",
		PriceBookCode: "book-au-online", PriceBookRevision: 5,
		PriceEntryID: "entry-1", PriceEntryRevision: 8,
		Currency: "AUD", CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2},
		ListUnitPrice:  money.Money{AmountMinor: 800, Currency: "AUD"},
		DiscountAmount: money.Money{AmountMinor: 120, Currency: "AUD"},
		TaxableBase:    money.Money{AmountMinor: 680, Currency: "AUD"},
		TaxAmount:      money.Money{AmountMinor: 62, Currency: "AUD"},
		LineTotal:      money.Money{AmountMinor: 680, Currency: "AUD"},
		Tax: quote.TaxSnapshot{
			TaxCategoryCode: "tax-au-gst", TaxRuleID: "rule-au-gst", TaxRuleRevision: 2,
			InclusionBasis:    pricebook_enums.PriceTaxInclusionInclusive,
			RateNumerator:     1,
			RateDenominator:   11,
			TaxableBase:       money.Money{AmountMinor: 680, Currency: "AUD"},
			AllocatedTax:      money.Money{AmountMinor: 62, Currency: "AUD"},
			CalculationSource: quote_enums.TaxCalculationSourceInclusiveExtraction,
			RoundingMethod:    quote_enums.TaxRoundingMethodSumExactThenRound,
		},
		AppliedRules: []quote.AppliedPriceRule{
			{Kind: "group_15", Exclusive: true, FactorNumerator: 85, FactorDenominator: 100, AmountBefore: money.Money{AmountMinor: 800, Currency: "AUD"}, AmountAfter: money.Money{AmountMinor: 680, Currency: "AUD"}, AppliedAt: capturedAt},
		},
		Rounding: quote.RoundingEvidence{
			Mode: quote_enums.RoundingModeHalfUp, PriceEnding: pricebook_enums.PriceEndingPolicyNone,
			Exponent: 2, ExactNumerator: 680, ExactDenominator: 1,
			RoundedAmount: money.Money{AmountMinor: 680, Currency: "AUD"},
		},
		Eligibility: listing.SaleEligibilitySnapshot{
			MarketCode: "market-au", SKUCode: "sku-potato-a", ListingRevision: 3,
			TaxCategoryCode: "tax-au-gst", DepotCode: "SYD-01",
			StockLocation:      warehouse.StockLocationRef{DepotCode: "SYD-01", LocationCode: "A-1"},
			Condition:          warehouse_enums.InventoryConditionGood,
			Disposition:        warehouse_enums.InventoryDispositionStandardSellable,
			AvailableBaseUnits: 40,
			InventoryRevision:  12,
			ValidityToken:      "eligibility-token-1", ValidUntil: capturedAt, CapturedAt: capturedAt,
		},
		Fingerprint: "fingerprint-1", CapturedAt: capturedAt,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal price snapshot: %v", err)
	}
	var got quote.PriceSnapshot
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal price snapshot: %v", err)
	}
	if got.Tax.RateNumerator != 1 || got.Tax.RateDenominator != 11 {
		t.Fatalf("exact tax rate did not round-trip: %+v", got.Tax)
	}
	if len(got.AppliedRules) != 1 || got.AppliedRules[0].Kind != "group_15" || !got.AppliedRules[0].Exclusive {
		t.Fatalf("exclusive rule evidence changed: %+v", got.AppliedRules)
	}
	if got.Eligibility.ValidityToken != "eligibility-token-1" || got.Eligibility.ListingRevision != 3 {
		t.Fatalf("eligibility evidence changed: %+v", got.Eligibility)
	}
	if got.CurrencyExponent.Exponent != 2 || got.Currency != "AUD" {
		t.Fatalf("currency evidence changed: %+v", got.CurrencyExponent)
	}
	for _, retired := range []string{"accepted_package_pricing", "package_pricing_id", "package_pricing_revision"} {
		if strings.Contains(string(payload), retired) {
			t.Fatalf("price snapshot retains retired package-pricing wire fragment %q: %s", retired, payload)
		}
	}
}

func TestV27PromotionApplicationsPreserveOrderedResolvedSKUEvidence(t *testing.T) {
	appliedAt := time.Date(2026, 8, 12, 9, 30, 0, 0, time.UTC)
	applications := []promotion.PromotionApplication{
		{PromotionID: "promotion-membership", PromotionKind: "membership_discount", PromotionRevision: 2, RelationID: "relation-membership", ResolvedTargetSKUCodes: []string{"sku-potato-a"}, AppliedAt: appliedAt},
		{PromotionID: "promotion-bundle", PromotionKind: "future_bundle_kind", PromotionRevision: 7, RelationID: "relation-bundle", ResolvedQualifierSKUCodes: []string{"sku-potato-a", "sku-potato-b"}, ResolvedTargetSKUCodes: []string{"sku-potato-a"}, AppliedAt: appliedAt},
	}

	payload, err := json.Marshal(applications)
	if err != nil {
		t.Fatalf("marshal promotion applications: %v", err)
	}
	var got []promotion.PromotionApplication
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal promotion applications: %v", err)
	}
	if len(got) != 2 || got[0].PromotionID != "promotion-membership" || got[1].PromotionID != "promotion-bundle" {
		t.Fatalf("promotion stacking order changed: %+v", got)
	}
	if got[1].RelationID != "relation-bundle" || len(got[1].ResolvedQualifierSKUCodes) != 2 || len(got[1].ResolvedTargetSKUCodes) != 1 {
		t.Fatalf("resolved qualifier-to-target evidence changed: %+v", got[1])
	}

	applicationType := reflect.TypeOf(promotion.PromotionApplication{})
	for fieldName, wantTag := range map[string]string{
		"RelationID":                "relation_id",
		"ResolvedQualifierSKUCodes": "resolved_qualifier_sku_codes,omitempty",
		"ResolvedTargetSKUCodes":    "resolved_target_sku_codes,omitempty",
	} {
		field, ok := applicationType.FieldByName(fieldName)
		if !ok || field.Tag.Get("json") != wantTag {
			t.Fatalf("PromotionApplication.%s JSON tag = %q, want %q", fieldName, field.Tag.Get("json"), wantTag)
		}
	}
}

func TestV27CustomerSummariesExposeOnlyFrozenSafePricingFacts(t *testing.T) {
	for _, model := range []reflect.Type{reflect.TypeOf(pos.ReceiptLine{}), reflect.TypeOf(order.OrderLineSummary{})} {
		for _, required := range []string{"SKUCode", "ProductName", "ProductImage", "ProductPackageOption", "CapturedAt", "PackagePrice", "TaxAmount", "DiscountAmount", "PromotionApplications", "Total"} {
			if _, ok := model.FieldByName(required); !ok {
				t.Errorf("%s is missing frozen customer-safe field %s", model, required)
			}
		}
		for _, forbidden := range []string{"AcceptedPackagePricing", "PackagePricing", "PriceSnapshot", "Eligibility", "Components", "StockLocation", "SourceBucketID", "SourceStockUnitID", "AvailablePackageCount", "AvailableBaseUnits", "InventoryRevision"} {
			if _, ok := model.FieldByName(forbidden); ok {
				t.Errorf("%s exposes internal pricing/inventory field %s", model, forbidden)
			}
		}
		for i := 0; i < model.NumField(); i++ {
			if model.Field(i).Type == reflect.TypeOf(quote.PriceSnapshot{}) || model.Field(i).Type == reflect.TypeOf(listing.SaleEligibilitySnapshot{}) {
				t.Errorf("%s embeds raw pricing or inventory evidence through field %s", model, model.Field(i).Name)
			}
		}
	}
}

func TestV27OperationalCategoryTagEvidenceIsLocationQualified(t *testing.T) {
	model := reflect.TypeOf(operations.InventoryCategoryTagEvidence{})
	assertJSONFields(t, model, map[string]string{
		"SKUCode":           "sku_code",
		"PackageOptionCode": "package_option_code",
		"CategoryTag":       "category_tag",
		"StockLocation":     "stock_location",
		"Condition":         "condition",
		"Disposition":       "disposition",
		"DateMark":          "date_mark,omitempty",
		"AsOf":              "as_of",
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
		if !strings.HasPrefix(relative, "contracts/pricing/promotion/") {
			return nil
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
		selectorCount := len(group.SKUCodes) + len(group.CollectionCodes) + len(group.CategoryTagCodes) + len(group.PackageOptionCodes)
		if selectorCount == 0 || group.MinimumBaseUnits < 0 {
			return false
		}
		if group.MaximumBaseUnits != nil && *group.MaximumBaseUnits < group.MinimumBaseUnits {
			return false
		}
	}
	return true
}
