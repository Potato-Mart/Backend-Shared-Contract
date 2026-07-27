package product

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/product"
)

func TestProductStorefrontMerchandisingJSONShape(t *testing.T) {
	start := time.Date(2026, 7, 7, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2026, 7, 21, 0, 0, 0, 0, time.UTC)
	product := Product{
		ID:      "prd_preorder",
		SKUCode: "A00084",
		SKU:     "A00084",
		Name:    "Preorder product",
		StorefrontMerchandising: &StorefrontMerchandising{
			Preorder: &PreorderPolicy{
				Enabled:                true,
				StartsAt:               &start,
				ExpectedAvailableAt:    &expected,
				MaxQuantityPerOrder:    3,
				MaxQuantityPerCustomer: 6,
				Labels:                 []common.LocalizedName{{Language: "en", Name: "Preorder"}},
			},
			SoonExpiry: &SoonExpiryMerchandisingPolicy{
				Enabled:             true,
				WindowDays:          14,
				ShowExactExpiryDate: true,
				Labels:              []common.LocalizedName{{Language: "en", Name: "Limited shelf life"}},
			},
		},
	}

	body, err := json.Marshal(product)
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal product JSON: %v", err)
	}
	merchandising, ok := got["storefront_merchandising"].(map[string]any)
	if !ok {
		t.Fatalf("Product JSON missing storefront_merchandising: %s", body)
	}
	preorder, ok := merchandising["preorder"].(map[string]any)
	if !ok || preorder["enabled"] != true || preorder["max_quantity_per_order"] != float64(3) {
		t.Fatalf("Product preorder policy JSON mismatch: %s", body)
	}
	if preorder["expected_available_at"] != "2026-07-21T00:00:00Z" {
		t.Fatalf("expected_available_at = %v, want 2026-07-21T00:00:00Z (%s)", preorder["expected_available_at"], body)
	}
	soonExpiry, ok := merchandising["soon_expiry"].(map[string]any)
	if !ok || soonExpiry["window_days"] != float64(14) || soonExpiry["show_exact_expiry_date"] != true {
		t.Fatalf("Product soon-expiry policy JSON mismatch: %s", body)
	}
}

func TestStorefrontDisplayJSONShape(t *testing.T) {
	expiresAt := time.Date(2026, 7, 12, 0, 0, 0, 0, time.UTC)
	days := 6
	display := StorefrontDisplay{
		Preorder: &StorefrontPreorderDisplay{
			Available:           true,
			Status:              productenum.StorefrontPreorderStatusOpen,
			ExpectedAvailableAt: &expiresAt,
			MaxQuantityPerOrder: 2,
			Labels:              []common.LocalizedName{{Language: "en", Name: "Preorder now"}},
		},
		Expiry: &StorefrontExpiryDisplay{
			SoonExpiry:          true,
			Status:              productenum.StorefrontExpiryStatusSoonExpiry,
			ExpiresAt:           &expiresAt,
			DaysToExpiry:        &days,
			WindowDays:          14,
			ShowExactExpiryDate: true,
			Labels:              []common.LocalizedName{{Language: "en", Name: "Soon expiry"}},
		},
	}

	body, err := json.Marshal(display)
	if err != nil {
		t.Fatalf("marshal storefront display: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal storefront display JSON: %v", err)
	}
	preorder, ok := got["preorder"].(map[string]any)
	if !ok || preorder["available"] != true || preorder["status"] != "open" {
		t.Fatalf("preorder display JSON mismatch: %s", body)
	}
	expiry, ok := got["expiry"].(map[string]any)
	if !ok || expiry["soon_expiry"] != true || expiry["status"] != "soon_expiry" || expiry["days_to_expiry"] != float64(6) {
		t.Fatalf("expiry display JSON mismatch: %s", body)
	}
}
