package wallet_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/benefit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/benefit/benefit_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/wallet"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/wallet/wallet_enums"
)

func TestCustomerWalletOwnsPointsSummaryWithoutMembershipImport(t *testing.T) {
	now := time.Date(2026, 8, 24, 0, 0, 0, 0, time.UTC)
	w := wallet.CustomerWallet{CustomerNumber: "RC-1", Summary: wallet.CustomerWalletSummary{
		Points:                        wallet.PointsSummary{AvailablePoints: 120, PointDebt: 3, CalculatedAt: now},
		GiftCardAvailableBalanceTotal: money.Money{AmountMinor: 500, Currency: "AUD"},
	}, CalculatedAt: now}
	payload, err := json.Marshal(w)
	if err != nil {
		t.Fatalf("marshal wallet: %v", err)
	}
	if strings.Contains(string(payload), "points_policy") || !strings.Contains(string(payload), `"points":{"total_points":0`) {
		t.Fatalf("wallet leaked membership policy or omitted points: %s", payload)
	}
}

func TestOperationalPointsAndRewardRecordsAreWalletOwned(t *testing.T) {
	now := time.Date(2026, 8, 24, 0, 0, 0, 0, time.UTC)
	entry := wallet.PointLedgerEntry{ID: "pl-1", CustomerNumber: "RC-1", Delta: -100, Reason: wallet_enums.PointLedgerReasonRewardRedeem, BalanceAfter: 50, Remaining: 0, CreatedAt: now}
	redemption := wallet.RewardRedemption{ID: "rr-1", CustomerNumber: "RC-1", RewardCode: "reward-1", PointsSpent: 100, Status: wallet_enums.RewardRedemptionStatusRedeemed, CreatedAt: now}
	payload, err := json.Marshal(struct {
		Entry      wallet.PointLedgerEntry `json:"entry"`
		Redemption wallet.RewardRedemption `json:"redemption"`
	}{Entry: entry, Redemption: redemption})
	if err != nil {
		t.Fatalf("marshal wallet operations: %v", err)
	}
	if !strings.Contains(string(payload), `"reason":"REWARD_REDEEM"`) || !strings.Contains(string(payload), `"status":"REDEEMED"`) {
		t.Fatalf("unexpected wallet operations payload: %s", payload)
	}
}

func TestCheckoutBenefitReservationExcludesPersistenceRetryKey(t *testing.T) {
	now := time.Date(2026, 8, 25, 0, 0, 0, 0, time.UTC)
	reservation := wallet.CheckoutBenefitReservation{
		ID:          "benefit_reservation_1",
		Owner:       benefit.OwnerRef{OwnerType: benefit_enums.OwnerTypeRetailCustomer, OwnerID: "RC-1"},
		OrderNumber: "order_1",
		Status:      wallet_enums.CheckoutBenefitReservationStatusReserved,
		ExpiresAt:   now.Add(time.Hour),
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	payload, err := json.Marshal(reservation)
	if err != nil {
		t.Fatalf("marshal checkout benefit reservation: %v", err)
	}
	if strings.Contains(string(payload), `"idempotency_key"`) {
		t.Fatalf("checkout benefit reservation must not expose persistence retry state: %s", payload)
	}
	if _, exists := reflect.TypeOf(wallet.CheckoutBenefitReservation{}).FieldByName("IdempotencyKey"); exists {
		t.Fatal("checkout benefit reservation must not expose IdempotencyKey")
	}
	if !strings.Contains(string(payload), `"order_number":"order_1"`) || !strings.Contains(string(payload), `"expires_at"`) {
		t.Fatalf("checkout benefit reservation lost business reference or timestamp: %s", payload)
	}
}
