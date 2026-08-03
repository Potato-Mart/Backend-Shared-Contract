package promotion

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	geographyenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/geography"
)

func TestReceiptOfferOmitsPromotionRuleInternals(t *testing.T) {
	now := time.Date(2026, 7, 15, 1, 2, 3, 0, time.UTC)
	body, err := json.Marshal(ReceiptOffer{
		ID:               "prm_1",
		SeriesKey:        "series_weekend",
		ReceiptMessages:  []common.LocalizedName{{Language: "en", Name: "Weekend special"}},
		StartsAt:         &now,
		ScheduleTimezone: "Australia/Sydney",
		GeographicContext: common.GeographicContext{
			Source:      geographyenum.GeographicContextSourceRetailCustomerProfile,
			CountryCode: "AU", ScopeRevision: 4, RuleRevision: 9,
			EvaluationTimezone: "Australia/Sydney",
		},
		Priority: 50,
	})
	if err != nil {
		t.Fatalf("marshal receipt offers: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal receipt offers: %v", err)
	}
	for _, forbidden := range []string{"name", "description", "discount_type", "discount_value", "usage_limit", "used_count", "source"} {
		if _, exists := got[forbidden]; exists {
			t.Fatalf("buyer-safe receipt offer leaked %q (%s)", forbidden, body)
		}
	}
	if got["starts_at"] != now.Format(time.RFC3339) {
		t.Fatalf("starts_at = %#v, want %q", got["starts_at"], now.Format(time.RFC3339))
	}
	context, ok := got["geographic_context"].(map[string]any)
	if !ok || context["source"] != "RETAIL_CUSTOMER_PROFILE" || got["schedule_timezone"] != "Australia/Sydney" {
		t.Fatalf("receipt offer geographic context mismatch: %s", body)
	}
}
