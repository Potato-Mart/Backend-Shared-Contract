package product_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/product"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/product"
)

func TestStorefrontProductJSONIsCustomerSafe(t *testing.T) {
	expiry := time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)
	discount := 20
	projection := product.StorefrontProduct{
		SKUCode: "SKU-CODE-1", SKU: "SKU-1", Name: "Product",
		BrandRef: &product.BrandRef{
			ID: "64c13ab08edf48a008793ca1", Slug: "happy-potato",
			Name: []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
		},
		CurrentStock: 12,
		Pricing: product.StorefrontPricing{
			Audience: productenum.PriceAudienceRetail,
			Current:  &common.Money{AmountMinor: 800, Currency: "AUD"},
			Original: &common.Money{AmountMinor: 1000, Currency: "AUD"},
		},
		StorefrontDisplay: product.StorefrontDisplay{},
		ExpiryDate:        &expiry,
		PromotionBadge: &product.StorefrontPromotionBadge{
			PromotionID: "promo_1", DiscountPercent: &discount,
		},
	}
	payload, err := json.Marshal(projection)
	if err != nil {
		t.Fatalf("marshal storefront product: %v", err)
	}
	for _, want := range []string{`"sku_code":"SKU-CODE-1"`, `"sku":"SKU-1"`, `"current_stock":12`, `"audience":"retail"`, `"promotion_badge"`, `"discount_percent":20`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("storefront product JSON = %s, want %s", payload, want)
		}
	}
	for _, forbidden := range []string{`"cost"`, `"offline"`, `"storefront_merchandising"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("storefront product JSON = %s, must omit %s", payload, forbidden)
		}
	}
	if strings.Contains(string(payload), `"brand_key"`) || !strings.Contains(string(payload), `"brand_ref":{"id":"64c13ab08edf48a008793ca1"`) {
		t.Fatalf("storefront product brand JSON = %s", payload)
	}
}
