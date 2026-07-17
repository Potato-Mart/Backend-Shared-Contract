package membership_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/contracts/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/contracts/product"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/membership"
)

func TestMembershipAccountAndTierRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 23, 0, 0, 0, 0, time.UTC)
	account := membership.MembershipAccount{
		ID:      "mem_1",
		Owner:   membership.MembershipOwnerRef{OwnerType: membershipenum.MembershipOwnerTypeRetailCustomer, OwnerID: "retail_1"},
		TierKey: "gold",
		Status:  membershipenum.MembershipAccountStatusActive,
		Wallet: membership.MembershipWalletSummary{
			TotalPoints:     1200,
			ReservedPoints:  200,
			AvailablePoints: 1000,
			ExpiringPoints:  300,
			CalculatedAt:    now,
		},
		EnrolledAt: now,
	}
	tier := membership.MembershipTier{
		TierKey:             "gold",
		Label:               "Gold",
		QualificationMetric: membershipenum.MembershipTierMetricAnnualSpend,
		MinQualifyingSpend:  common.Money{AmountMinor: 100000, Currency: "AUD"},
		PointMultiplier:     2,
		IsActive:            true,
		IsSystem:            true,
	}

	payload, err := json.Marshal(struct {
		Account membership.MembershipAccount `json:"account"`
		Tier    membership.MembershipTier    `json:"tier"`
	}{Account: account, Tier: tier})
	if err != nil {
		t.Fatalf("marshal membership account/tier: %v", err)
	}

	var decoded struct {
		Account membership.MembershipAccount `json:"account"`
		Tier    membership.MembershipTier    `json:"tier"`
	}
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal membership account/tier: %v", err)
	}
	if decoded.Account.Wallet.AvailablePoints != 1000 || decoded.Tier.QualificationMetric != membershipenum.MembershipTierMetricAnnualSpend {
		t.Fatalf("membership account/tier did not round-trip: %+v", decoded)
	}
}

func TestPointLedgerBucketsAndReservationRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 23, 0, 0, 0, 0, time.UTC)
	expiresSoon := time.Date(2026, 7, 7, 0, 0, 0, 0, time.UTC)
	expiresLater := time.Date(2026, 10, 15, 0, 0, 0, 0, time.UTC)
	owner := membership.MembershipOwnerRef{OwnerType: membershipenum.MembershipOwnerTypeRetailCustomer, OwnerID: "retail_1"}

	entry := membership.PointLedgerEntry{
		ID:                  "redeem_1",
		MembershipAccountID: "mem_1",
		Owner:               owner,
		Delta:               -40,
		Reason:              membershipenum.MembershipPointReasonRedeem,
		BalanceAfter:        60,
		Allocations: []membership.PointAllocation{
			{LedgerEntryID: "earn_1", Points: 30, ExpiresAt: &expiresSoon},
			{LedgerEntryID: "earn_2", Points: 10, ExpiresAt: &expiresLater},
		},
		CreatedAt: now,
	}
	breakdown := membership.PointBalanceBreakdown{
		MembershipAccountID: "mem_1",
		Owner:               owner,
		TotalPoints:         100,
		ReservedPoints:      40,
		AvailablePoints:     60,
		ExpiringPoints:      100,
		Buckets: []membership.PointBucket{
			{Points: 30, ExpiresAt: &expiresSoon, SourceLedgerEntryID: "earn_1", Reason: membershipenum.MembershipPointReasonOrder},
			{Points: 70, ExpiresAt: &expiresLater, SourceLedgerEntryID: "earn_2", Reason: membershipenum.MembershipPointReasonSignupBonus},
		},
		CalculatedAt: now,
	}
	reservation := membership.PointReservation{
		ID:                  "res_1",
		MembershipAccountID: "mem_1",
		Owner:               owner,
		Points:              40,
		Status:              membershipenum.PointReservationStatusReserved,
		Reason:              membershipenum.MembershipPointReasonRedeem,
		RedemptionType:      membershipenum.MembershipRedemptionTypeCheckoutDiscount,
		Allocations:         entry.Allocations,
		ExpiresAt:           now.Add(10 * time.Minute),
		CreatedAt:           now,
	}

	payload, err := json.Marshal(struct {
		Entry       membership.PointLedgerEntry      `json:"entry"`
		Breakdown   membership.PointBalanceBreakdown `json:"breakdown"`
		Reservation membership.PointReservation      `json:"reservation"`
	}{Entry: entry, Breakdown: breakdown, Reservation: reservation})
	if err != nil {
		t.Fatalf("marshal point contracts: %v", err)
	}

	var decoded struct {
		Entry       membership.PointLedgerEntry      `json:"entry"`
		Breakdown   membership.PointBalanceBreakdown `json:"breakdown"`
		Reservation membership.PointReservation      `json:"reservation"`
	}
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal point contracts: %v", err)
	}
	if len(decoded.Entry.Allocations) != 2 || len(decoded.Breakdown.Buckets) != 2 || decoded.Reservation.Points != 40 {
		t.Fatalf("point contracts did not round-trip: %+v", decoded)
	}
}

func TestRewardRedemptionAndMemberSubscriptionRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 23, 0, 0, 0, 0, time.UTC)
	discount := common.Money{AmountMinor: 500, Currency: "AUD"}
	owner := membership.MembershipOwnerRef{OwnerType: membershipenum.MembershipOwnerTypeRetailCustomer, OwnerID: "retail_1"}
	reward := membership.Reward{
		ID:             "reward_1",
		Code:           "reward_1",
		Name:           "Five dollar discount",
		Type:           membershipenum.MembershipRewardTypeOrderDiscount,
		PointsCost:     500,
		DiscountAmount: &discount,
		IsActive:       true,
	}
	redemption := membership.RewardRedemption{
		ID:                  "redemption_1",
		MembershipAccountID: "mem_1",
		Owner:               owner,
		RewardCode:          "reward_1",
		ReservationID:       "res_1",
		PointsSpent:         500,
		Status:              membershipenum.MembershipRewardRedemptionStatusRedeemed,
		DiscountAmount:      &discount,
		CreatedAt:           now,
	}
	plan := membership.SubscriptionPlan{
		ID:            "plan_1",
		Product:       product.Snapshot{ID: "prod_1", Name: "Weekly Potatoes"},
		UnitPrice:     common.Money{AmountMinor: 1200, Currency: "AUD"},
		FrequencyDays: 7,
		IsActive:      true,
	}
	subscription := membership.MemberSubscription{
		ID:                  "sub_1",
		MembershipAccountID: "mem_1",
		Owner:               owner,
		PlanID:              "plan_1",
		Qty:                 2,
		Status:              membershipenum.MemberSubscriptionStatusActive,
		StartedAt:           now,
		NextOrderAt:         now.AddDate(0, 0, 7),
	}

	payload, err := json.Marshal(struct {
		Reward       membership.Reward             `json:"reward"`
		Redemption   membership.RewardRedemption   `json:"redemption"`
		Plan         membership.SubscriptionPlan   `json:"plan"`
		Subscription membership.MemberSubscription `json:"subscription"`
	}{Reward: reward, Redemption: redemption, Plan: plan, Subscription: subscription})
	if err != nil {
		t.Fatalf("marshal reward/subscription contracts: %v", err)
	}

	var decoded struct {
		Reward       membership.Reward             `json:"reward"`
		Redemption   membership.RewardRedemption   `json:"redemption"`
		Plan         membership.SubscriptionPlan   `json:"plan"`
		Subscription membership.MemberSubscription `json:"subscription"`
	}
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal reward/subscription contracts: %v", err)
	}
	if decoded.Reward.PointsCost != 500 || decoded.Redemption.PointsSpent != 500 ||
		decoded.Plan.FrequencyDays != 7 || decoded.Subscription.Status != membershipenum.MemberSubscriptionStatusActive {
		t.Fatalf("reward/subscription contracts did not round-trip: %+v", decoded)
	}
}
