package product

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/product"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/security"
)

func TestProductStorefrontMerchandisingJSONShape(t *testing.T) {
	start := time.Date(2026, 7, 7, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2026, 7, 21, 0, 0, 0, 0, time.UTC)
	record := Product{
		SKUCode:         "A00084",
		CategorySKUCode: "A00084",
		Name:            "Preorder product",
		StorefrontMerchandising: &StorefrontMerchandising{
			Preorder: &PreorderPolicy{
				Enabled:                true,
				StartsAt:               &start,
				ExpectedAvailableAt:    &expected,
				ScheduleTimezone:       "Australia/Melbourne",
				MaxQuantityPerOrder:    3,
				MaxQuantityPerCustomer: 6,
				Labels:                 []common.LocalizedName{{Language: "en", Name: "Preorder"}},
			},
			SoonExpiry: &SoonExpiryMerchandisingPolicy{
				Enabled:             true,
				WindowDays:          30,
				ShowExactExpiryDate: true,
				Labels:              []common.LocalizedName{{Language: "en", Name: "Short dated"}},
			},
		},
	}

	body, err := json.Marshal(record)
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
		t.Fatalf("expected_available_at = %v, want UTC timestamp (%s)", preorder["expected_available_at"], body)
	}
	if preorder["schedule_timezone"] != "Australia/Melbourne" {
		t.Fatalf("schedule_timezone = %v, want Australia/Melbourne (%s)", preorder["schedule_timezone"], body)
	}
	if !strings.Contains(string(body), `"soon_expiry"`) {
		t.Fatalf("product expiry merchandising policy missing: %s", body)
	}
}

func TestStorefrontDisplayJSONShape(t *testing.T) {
	expectedAt := time.Date(2026, 7, 12, 0, 0, 0, 0, time.UTC)
	display := StorefrontDisplay{
		Preorder: &StorefrontPreorderDisplay{
			Available:           true,
			Status:              productenum.StorefrontPreorderStatusOpen,
			ExpectedAvailableAt: &expectedAt,
			ScheduleTimezone:    "Australia/Melbourne",
			MaxQuantityPerOrder: 2,
			Labels:              []common.LocalizedName{{Language: "en", Name: "Preorder now"}},
		},
		Expiry: &StorefrontExpiryDisplay{
			SoonExpiry:          true,
			Status:              StorefrontExpiryStatusSoonExpiry,
			AlertLevel:          securityenum.AlertLevelCritical,
			DaysToExpiry:        intPtr(5),
			WindowDays:          30,
			ShowExactExpiryDate: true,
			Labels:              []common.LocalizedName{{Language: "en", Name: "Use soon"}},
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
	if preorder["schedule_timezone"] != "Australia/Melbourne" {
		t.Fatalf("preorder display lost schedule timezone: %s", body)
	}
	expiry, ok := got["expiry"].(map[string]any)
	if !ok || expiry["soon_expiry"] != true || expiry["alert_level"] != string(securityenum.AlertLevelCritical) || expiry["days_to_expiry"] != float64(5) {
		t.Fatalf("expiry display JSON mismatch: %s", body)
	}
}

func intPtr(value int) *int { return &value }
