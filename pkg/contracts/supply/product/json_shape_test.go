package product

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/customers/retail/retail_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
)

func TestProductJSONUsesCanonicalComponents(t *testing.T) {
	effectiveFrom := time.Date(2026, 8, 4, 0, 0, 0, 0, time.UTC)
	value := Product{
		SKUCode: "A0001",
		Content: ProductContent{
			Name:         "Taxed product",
			Descriptions: []localization.LocalizedDescription{{Language: "en", Description: "Localized description"}},
		},
		Classification: ProductClassification{
			CategorySKUCode: "A",
			BrandRef:        &classification.BrandRef{ID: "brand_1", Slug: "localized-brand", Name: []localization.LocalizedName{{Language: "en", Name: "Localized brand"}}},
			CollectionRef:   &classification.CollectionRef{ID: "col_frozen", Slug: "frozen", Name: []localization.LocalizedName{{Language: "en", Name: "Frozen"}}},
			CategoryTags:    []classification.CategoryTagRef{{ID: "tag_hotpot", Slug: "hotpot", Name: []localization.LocalizedName{{Language: "en", Name: "Hotpot"}}}},
		},
		Packaging: ProductPackaging{
			Taxed: true,
			PackageOptions: []ProductPackageOption{{
				ID: "pkg_each", Code: "EACH", ProductSKUCode: "A0001", HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1, IsCanonical: true, IsActive: true, EffectiveFrom: effectiveFrom,
			}},
		},
		Supply: &classification.ProductSupply{Supplier: &classification.ProductSupplierRef{Code: "sup_1", Name: "Supplier"}},
	}

	body, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	for _, want := range []string{`"content":{"name":"Taxed product"`, `"classification":{"category_sku_code":"A"`, `"packaging":{"package_options"`, `"taxed":true`, `"metrics":{}`, `"supply":{"supplier":{"code":"sup_1","name":"Supplier"}}`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("Product JSON = %s, want %s", body, want)
		}
	}
	for _, legacy := range []string{`"name":"Taxed product"`, `"category_sku_code":"A"`, `"brand_ref"`, `"collection_ref"`, `"category_tags"`, `"taxed"`} {
		// Canonical component fields are expected only inside their component.
		if strings.Count(string(body), legacy) != 1 {
			t.Fatalf("Product JSON = %s, want one nested %s", body, legacy)
		}
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal product JSON: %v", err)
	}
	for _, component := range []string{"content", "classification", "packaging", "commerce", "metrics"} {
		if _, ok := got[component].(map[string]any); !ok {
			t.Fatalf("Product JSON = %s, missing %s component", body, component)
		}
	}
	for _, legacyTopLevel := range []string{"name", "description", "category_sku_code", "brand_ref", "collection", "category_tags", "taxed", "storage_type"} {
		if _, ok := got[legacyTopLevel]; ok {
			t.Fatalf("Product JSON = %s, retained flat %s", body, legacyTopLevel)
		}
	}
}

func TestProductSellingJSONLivesInCommerce(t *testing.T) {
	body, err := json.Marshal(Product{
		SKUCode: "A0001",
		Commerce: ProductCommerce{Selling: &Selling{
			Channels:   []commerce_enums.OrderType{commerce_enums.OrderTypeB2B, commerce_enums.OrderTypePOS},
			BuyerTypes: []retail_enums.BuyerType{retail_enums.BuyerTypeWholesaleOrganisation},
			Visibility: product_enums.PriceVisibilityWholesaleApprovedOnly,
		}},
	})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}

	var got Product
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal product: %v", err)
	}
	if got.Commerce.Selling == nil || len(got.Commerce.Selling.Channels) != 2 || got.Commerce.Selling.BuyerTypes[0] != retail_enums.BuyerTypeWholesaleOrganisation {
		t.Fatalf("selling rules did not round-trip: %+v", got.Commerce.Selling)
	}
	if strings.Contains(string(body), `"selling":{`) && !strings.Contains(string(body), `"commerce":{"selling":`) {
		t.Fatalf("selling must be nested in commerce: %s", body)
	}
}

func TestProductCommerceOmitsEmptySelling(t *testing.T) {
	body, err := json.Marshal(Product{SKUCode: "A0001"})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	if strings.Contains(string(body), `"selling"`) {
		t.Fatalf("empty selling should be omitted, got %s", body)
	}
}

func TestProductPackageCommerceIsSafeProjection(t *testing.T) {
	asOf := time.Date(2026, 8, 9, 0, 0, 0, 0, time.UTC)
	promotionApplications := []promotion.PromotionApplication{
		{PromotionID: "prm_first", PromotionKind: "discount", PromotionRevision: 1, RelationID: "rel_first", AppliedAt: asOf},
		{PromotionID: "prm_second", PromotionKind: "bundle", PromotionRevision: 2, RelationID: "rel_second", AppliedAt: asOf},
	}
	body, err := json.Marshal(Product{
		SKUCode: "A0001",
		Commerce: ProductCommerce{Packages: []ProductPackageCommerce{{
			PackageOptionID:       "pkg_each",
			PackagePrice:          money.Money{AmountMinor: 1200, Currency: "AUD"},
			TaxAmount:             money.Money{AmountMinor: 109, Currency: "AUD"},
			StockState:            product_enums.StorefrontStockStateInStock,
			PromotionApplications: promotionApplications,
			AsOf:                  asOf,
		}}},
	})
	if err != nil {
		t.Fatalf("marshal product package commerce: %v", err)
	}
	for _, want := range []string{`"packages":[{`, `"package_option_id":"pkg_each"`, `"stock_state":"in_stock"`, `"promotion_applications":[{`, `"promotion_id":"prm_first"`, `"promotion_id":"prm_second"`, `"as_of":"2026-08-09T00:00:00Z"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("Product package commerce JSON = %s, want %s", body, want)
		}
	}
	firstIndex := strings.Index(string(body), `"promotion_id":"prm_first"`)
	secondIndex := strings.Index(string(body), `"promotion_id":"prm_second"`)
	if firstIndex < 0 || secondIndex < firstIndex {
		t.Fatalf("promotion applications must preserve stacking order: %s", body)
	}
	for _, forbidden := range []string{`"depot_code"`, `"location_id"`, `"source_bucket_id"`, `"source_stock_unit_id"`, `"available_base_units"`, `"available_package_count"`, `"audit"`} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("Product package commerce leaked %s: %s", forbidden, body)
		}
	}
}
