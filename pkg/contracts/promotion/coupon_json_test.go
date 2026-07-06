package promotion_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

func TestCouponAssignmentRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 5, 0, 0, 0, 0, time.UTC)
	rec := promotion.CouponAssignment{
		ID:                  "ca_1",
		CouponID:            "coupon_1",
		CouponCode:          "SAVE10",
		CustomerNumber:      "RC-1",
		Source:              enums.CouponSourceCampaign,
		Status:              "redeemed",
		ExpiresAt:           &now,
		RedeemedAt:          &now,
		RedeemedOrderNumber: "MAMA260703ABC123",
		Note:                "campaign issue",
		CreatedAt:           now,
	}

	payload, err := json.Marshal(rec)
	if err != nil {
		t.Fatalf("marshal coupon assignment: %v", err)
	}
	if strings.Contains(string(payload), "customer_coupon") {
		t.Fatalf("assignment payload should not expose customer_coupon naming: %s", payload)
	}

	var decoded promotion.CouponAssignment
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal coupon assignment: %v", err)
	}
	if decoded.CouponID != "coupon_1" || decoded.CouponCode != "SAVE10" ||
		decoded.RedeemedOrderNumber != "MAMA260703ABC123" {
		t.Fatalf("coupon assignment did not round-trip: %+v", decoded)
	}
}
