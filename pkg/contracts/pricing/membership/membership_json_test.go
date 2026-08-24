package membership_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/wallet"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/wallet/wallet_enums"
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
