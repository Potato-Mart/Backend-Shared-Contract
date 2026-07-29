package promotion_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/contracts/benefit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/contracts/promotion"
	benefitenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/benefit"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/promotion"
)

func TestCouponAssignmentRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 5, 0, 0, 0, 0, time.UTC)
	rec := promotion.CouponAssignment{
		ID:         "ca_1",
		CouponID:   "coupon_1",
		CouponCode: "SAVE10",
		Owner: benefit.OwnerRef{
			OwnerType: benefitenum.OwnerTypeRetailCustomer,
			OwnerID:   "RC-1",
		},
		Source:              promotionenum.CouponSourceCampaign,
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
	if strings.Contains(string(payload), "customer_number") || !strings.Contains(string(payload), `"owner_type":"retail_customer"`) {
		t.Fatalf("assignment payload must use the generalized owner reference: %s", payload)
	}

	var decoded promotion.CouponAssignment
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal coupon assignment: %v", err)
	}
	if decoded.CouponID != "coupon_1" || decoded.CouponCode != "SAVE10" ||
		decoded.Owner.OwnerID != "RC-1" || decoded.RedeemedOrderNumber != "MAMA260703ABC123" {
		t.Fatalf("coupon assignment did not round-trip: %+v", decoded)
	}
}

func TestCouponUsageRoundTripsWholesaleOwner(t *testing.T) {
	now := time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC)
	owner := benefit.OwnerRef{
		OwnerType: benefitenum.OwnerTypeWholesaleOrganisation,
		OwnerID:   "ORG-1",
	}
	record := promotion.CouponUsageRecord{
		ID:                  "usage_1",
		CouponCode:          "WHOLESALE10",
		Owner:               &owner,
		RedeemedOrderNumber: "MAMA260715ABC123",
		DiscountAmount:      common.Money{AmountMinor: 1000, Currency: "AUD"},
		RedeemedAt:          now,
		CreatedAt:           now,
	}

	payload, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("marshal wholesale coupon usage: %v", err)
	}
	if strings.Contains(string(payload), "customer_number") || !strings.Contains(string(payload), `"owner_type":"wholesale_organisation"`) {
		t.Fatalf("coupon usage must use the wholesale owner reference: %s", payload)
	}

	var decoded promotion.CouponUsageRecord
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal wholesale coupon usage: %v", err)
	}
	if decoded.Owner == nil || decoded.Owner.OwnerID != "ORG-1" {
		t.Fatalf("wholesale coupon owner did not round-trip: %+v", decoded)
	}
}
