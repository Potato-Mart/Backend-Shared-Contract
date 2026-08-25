package membership_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/wallet"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/wallet/wallet_enums"
)

func TestMembershipAccountComposesWalletOwnedPointsSummary(t *testing.T) {
	now := time.Date(2026, 8, 24, 0, 0, 0, 0, time.UTC)
	account := membership.MembershipAccount{
		ID: "RC-1", Status: membership_enums.MembershipAccountStatusActive,
		Wallet: wallet.PointsSummary{AvailablePoints: 120, CalculatedAt: now}, EnrolledAt: now,
	}
	payload, err := json.Marshal(account)
	if err != nil {
		t.Fatalf("marshal membership account: %v", err)
	}
	if !strings.Contains(string(payload), `"wallet":{"total_points":0,"reserved_points":0,"available_points":120`) {
		t.Fatalf("membership must compose wallet summary: %s", payload)
	}
}

func TestMembershipPolicyAndRewardCatalogStayInMembership(t *testing.T) {
	reward := membership.Reward{Code: "reward-1", Type: wallet_enums.RewardTypeVoucher, PointsCost: 100, IsActive: true}
	payload, err := json.Marshal(struct {
		Policy membership.PointsPolicy `json:"policy"`
		Reward membership.Reward       `json:"reward"`
	}{Policy: membership.PointsPolicy{PointsPerMinorUnit: 1, RedemptionStepPoints: 100}, Reward: reward})
	if err != nil {
		t.Fatalf("marshal membership policy and reward: %v", err)
	}
	if !strings.Contains(string(payload), `"points_per_minor_unit":1`) || !strings.Contains(string(payload), `"type":"VOUCHER"`) {
		t.Fatalf("unexpected membership payload: %s", payload)
	}
}

func TestQualifyingSpendLedgerEntryExcludesPersistenceRetryKey(t *testing.T) {
	now := time.Date(2026, 8, 25, 0, 0, 0, 0, time.UTC)
	entry := membership.QualifyingSpendLedgerEntry{
		ID:                 "spend_1",
		CustomerNumber:     "RC-1",
		Amount:             money.Money{AmountMinor: 1250, Currency: "AUD"},
		Reason:             membership_enums.QualifyingSpendReasonOrderPaid,
		RelatedOrderNumber: "order_1",
		OccurredAt:         now,
		CreatedAt:          now,
	}
	payload, err := json.Marshal(entry)
	if err != nil {
		t.Fatalf("marshal qualifying spend entry: %v", err)
	}
	if strings.Contains(string(payload), `"idempotency_key"`) {
		t.Fatalf("qualifying spend entry must not expose persistence retry state: %s", payload)
	}
	if _, exists := reflect.TypeOf(membership.QualifyingSpendLedgerEntry{}).FieldByName("IdempotencyKey"); exists {
		t.Fatal("qualifying spend entry must not expose IdempotencyKey")
	}
	if !strings.Contains(string(payload), `"related_order_number":"order_1"`) || !strings.Contains(string(payload), `"occurred_at"`) {
		t.Fatalf("qualifying spend entry lost business reference or timestamp: %s", payload)
	}
}
