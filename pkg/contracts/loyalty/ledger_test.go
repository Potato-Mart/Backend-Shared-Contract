package loyalty_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/contracts/loyalty"
	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/enums"
)

func TestLoyaltyLedgerLegacyPayloadWithoutAllocations(t *testing.T) {
	payload := []byte(`{"id":"led_1","customer_id":"cust_1","delta":30,"reason":"ORDER","balance_after":30,"remaining":30,"created_at":"2026-06-17T00:00:00Z"}`)

	var entry loyalty.LoyaltyLedgerEntry
	if err := json.Unmarshal(payload, &entry); err != nil {
		t.Fatalf("unmarshal legacy loyalty ledger entry: %v", err)
	}
	if entry.Allocations != nil {
		t.Fatalf("allocations = %+v, want nil", entry.Allocations)
	}
	if entry.Remaining != 30 {
		t.Fatalf("remaining = %d, want 30", entry.Remaining)
	}
}

func TestLoyaltyLedgerAllocationRoundTrip(t *testing.T) {
	expiresSoon := time.Date(2026, 7, 7, 0, 0, 0, 0, time.UTC)
	expiresLater := time.Date(2026, 10, 15, 0, 0, 0, 0, time.UTC)

	entry := loyalty.LoyaltyLedgerEntry{
		ID:           "redeem_1",
		CustomerID:   "cust_1",
		Delta:        -40,
		Reason:       enums.LoyaltyLedgerReasonRedeem,
		BalanceAfter: 60,
		Remaining:    0,
		Allocations: []loyalty.LoyaltyPointAllocation{
			{LedgerEntryID: "earn_1", Points: 30, ExpiresAt: &expiresSoon},
			{LedgerEntryID: "earn_2", Points: 10, ExpiresAt: &expiresLater},
		},
		CreatedAt: time.Date(2026, 6, 17, 0, 0, 0, 0, time.UTC),
	}

	payload, err := json.Marshal(entry)
	if err != nil {
		t.Fatalf("marshal loyalty ledger entry: %v", err)
	}

	var decoded loyalty.LoyaltyLedgerEntry
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal loyalty ledger entry: %v", err)
	}

	if len(decoded.Allocations) != 2 {
		t.Fatalf("allocations length = %d, want 2", len(decoded.Allocations))
	}
	if decoded.Allocations[0].LedgerEntryID != "earn_1" || decoded.Allocations[0].Points != 30 {
		t.Fatalf("first allocation did not round-trip: %+v", decoded.Allocations[0])
	}
	if decoded.Allocations[1].LedgerEntryID != "earn_2" || decoded.Allocations[1].Points != 10 {
		t.Fatalf("second allocation did not round-trip: %+v", decoded.Allocations[1])
	}
}

func TestLoyaltyBalanceBreakdownBuckets(t *testing.T) {
	expiresSoon := time.Date(2026, 7, 7, 0, 0, 0, 0, time.UTC)
	expiresLater := time.Date(2026, 10, 15, 0, 0, 0, 0, time.UTC)
	calculatedAt := time.Date(2026, 6, 17, 0, 0, 0, 0, time.UTC)

	breakdown := loyalty.LoyaltyBalanceBreakdown{
		CustomerID:     "cust_1",
		TotalPoints:    100,
		ExpiringPoints: 100,
		Buckets: []loyalty.LoyaltyPointBucket{
			{
				Points:              30,
				ExpiresAt:           &expiresSoon,
				SourceLedgerEntryID: "earn_1",
				Reason:              enums.LoyaltyLedgerReasonOrder,
				RelatedOrderID:      "ord_1",
				RelatedOrderNumber:  "1001",
			},
			{
				Points:              70,
				ExpiresAt:           &expiresLater,
				SourceLedgerEntryID: "earn_2",
				Reason:              enums.LoyaltyLedgerReasonSignupBonus,
			},
		},
		CalculatedAt: calculatedAt,
	}

	payload, err := json.Marshal(breakdown)
	if err != nil {
		t.Fatalf("marshal loyalty balance breakdown: %v", err)
	}

	var decoded loyalty.LoyaltyBalanceBreakdown
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal loyalty balance breakdown: %v", err)
	}

	if decoded.TotalPoints != 100 || decoded.ExpiringPoints != 100 {
		t.Fatalf("points summary did not round-trip: %+v", decoded)
	}
	if len(decoded.Buckets) != 2 {
		t.Fatalf("buckets length = %d, want 2", len(decoded.Buckets))
	}
	if decoded.Buckets[0].Points != 30 || decoded.Buckets[1].Points != 70 {
		t.Fatalf("bucket points did not round-trip: %+v", decoded.Buckets)
	}
	if decoded.Buckets[0].ExpiresAt == nil || !decoded.Buckets[0].ExpiresAt.Equal(expiresSoon) {
		t.Fatalf("first bucket expiry = %v, want %s", decoded.Buckets[0].ExpiresAt, expiresSoon)
	}
}
