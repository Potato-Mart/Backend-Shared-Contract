package analytics

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/insights/analytics/analytics_enums"
)

func TestRetailCustomerAnalyticsProfileJSONShape(t *testing.T) {
	calculatedAt := time.Date(2026, 8, 24, 1, 2, 3, 0, time.UTC)
	profile := RetailCustomerAnalyticsProfile{
		CustomerNumber: "RC-123",
		Score:          "555",
		Segment:        "champions",
		ChurnRisk:      analytics_enums.ChurnRiskLow,
		CalculatedAt:   calculatedAt,
	}

	payload, err := json.Marshal(profile)
	if err != nil {
		t.Fatalf("marshal retail customer analytics profile: %v", err)
	}
	for _, field := range []string{
		`"customer_number":"RC-123"`,
		`"score":"555"`,
		`"segment":"champions"`,
		`"churn_risk":"LOW"`,
		`"calculated_at":"2026-08-24T01:02:03Z"`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("analytics profile JSON missing %s: %s", field, payload)
		}
	}

	var decoded RetailCustomerAnalyticsProfile
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal retail customer analytics profile: %v", err)
	}
	if decoded.CustomerNumber != "RC-123" || decoded.ChurnRisk != analytics_enums.ChurnRiskLow || !decoded.CalculatedAt.Equal(calculatedAt) {
		t.Fatalf("analytics profile did not round-trip: %+v", decoded)
	}
}
