package wallet_test

import (
	"encoding/json"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/benefit"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/wallet"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/benefit/benefit_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/wallet/wallet_enums"
)

func TestCouponAssignmentRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 5, 0, 0, 0, 0, time.UTC)
	rec := wallet.CouponAssignment{
		ID:         "ca_1",
		CouponID:   "coupon_1",
		CouponCode: "SAVE10",
		Owner: benefit.OwnerRef{
			OwnerType: benefit_enums.OwnerTypeRetailCustomer,
			OwnerID:   "RC-1",
		},
		Source:              wallet_enums.CouponSourceCampaign,
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

	var decoded wallet.CouponAssignment
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
		OwnerType: benefit_enums.OwnerTypeWholesaleOrganisation,
		OwnerID:   "ORG-1",
	}
	record := wallet.CouponUsageRecord{
		ID:                  "usage_1",
		CouponCode:          "WHOLESALE10",
		Owner:               &owner,
		RedeemedOrderNumber: "MAMA260715ABC123",
		DiscountAmount:      money.Money{AmountMinor: 1000, Currency: "AUD"},
		RedeemedAt:          now,
		RefundID:            "refund_1",
		RefundedAt:          &now,
		CreatedAt:           now,
	}

	payload, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("marshal wholesale coupon usage: %v", err)
	}
	if strings.Contains(string(payload), "customer_number") || !strings.Contains(string(payload), `"owner_type":"wholesale_organisation"`) {
		t.Fatalf("coupon usage must use the wholesale owner reference: %s", payload)
	}
	unrefundedRecord := record
	unrefundedRecord.RefundID = ""
	unrefundedRecord.RefundedAt = nil
	unrefundedPayload, err := json.Marshal(unrefundedRecord)
	if err != nil {
		t.Fatalf("marshal unrefunded coupon usage: %v", err)
	}
	if strings.Contains(string(unrefundedPayload), "refund_id") || strings.Contains(string(unrefundedPayload), "refunded_at") {
		t.Fatalf("unrefunded coupon usage should omit refund audit fields: %s", unrefundedPayload)
	}

	var decoded wallet.CouponUsageRecord
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal wholesale coupon usage: %v", err)
	}
	if decoded.Owner == nil || decoded.Owner.OwnerID != "ORG-1" ||
		decoded.RefundID != "refund_1" || decoded.RefundedAt == nil || !decoded.RefundedAt.Equal(now) {
		t.Fatalf("wholesale coupon owner did not round-trip: %+v", decoded)
	}
}
