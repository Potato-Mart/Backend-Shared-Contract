package wallet_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
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

func TestRewardRedemptionFreezesOptionalMarketAndCountry(t *testing.T) {
	now := time.Date(2026, 8, 24, 0, 0, 0, 0, time.UTC)
	redemption := wallet.RewardRedemption{
		ID:             "rr-1",
		CustomerNumber: "RC-1",
		MarketCode:     "mkt_au_vic",
		CountryCode:    geography.CountryCode("AU"),
		RewardCode:     "reward-1",
		PointsSpent:    100,
		Status:         wallet_enums.RewardRedemptionStatusRedeemed,
		CreatedAt:      now,
	}
	payload, err := json.Marshal(redemption)
	if err != nil {
		t.Fatalf("marshal reward redemption: %v", err)
	}
	if !strings.Contains(string(payload), `"market_code":"mkt_au_vic"`) || !strings.Contains(string(payload), `"country_code":"AU"`) {
		t.Fatalf("reward redemption omitted geographic evidence: %s", payload)
	}
	emptyPayload, err := json.Marshal(wallet.RewardRedemption{})
	if err != nil {
		t.Fatalf("marshal empty reward redemption: %v", err)
	}
	if strings.Contains(string(emptyPayload), `"market_code"`) || strings.Contains(string(emptyPayload), `"country_code"`) {
		t.Fatalf("empty reward redemption retained optional geography: %s", emptyPayload)
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

func TestRewardRedemptionOutcomeCarriesTypedIssueEvidence(t *testing.T) {
	now := time.Date(2026, 8, 26, 0, 0, 0, 0, time.UTC)
	redemption := wallet.RewardRedemption{
		ID:             "redemption_1",
		CustomerNumber: "RC-1",
		RewardCode:     "reward-partner",
		PointsSpent:    5000,
		Status:         wallet_enums.RewardRedemptionStatusRedeemed,
		Outcome: &wallet.RewardRedemptionOutcome{
			GiftCardCode: "gc-001",
			External: &wallet.ExternalRewardFulfilment{
				ProviderCode:      "partner_stream_plus",
				ExternalReference: "sub_889",
				Status:            wallet_enums.ExternalRewardFulfilmentStatusProvisioned,
				ProvisionedAt:     &now,
			},
		},
		CreatedAt: now,
	}
	payload, err := json.Marshal(redemption)
	if err != nil {
		t.Fatalf("marshal reward redemption: %v", err)
	}
	if !strings.Contains(string(payload), `"customer_number":"RC-1"`) || !strings.Contains(string(payload), `"points_spent":5000`) {
		t.Fatalf("reward redemption lost the redeeming member or the points spent: %s", payload)
	}
	if !strings.Contains(string(payload), `"outcome":{"gift_card_code":"gc-001"`) {
		t.Fatalf("reward redemption lost its issued gift-card evidence: %s", payload)
	}
	if !strings.Contains(string(payload), `"provider_code":"partner_stream_plus"`) || !strings.Contains(string(payload), `"status":"PROVISIONED"`) {
		t.Fatalf("reward redemption lost its external partner fulfilment: %s", payload)
	}

	var decoded wallet.RewardRedemption
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal reward redemption: %v", err)
	}
	if decoded.Outcome == nil || decoded.Outcome.External == nil || decoded.Outcome.External.ExternalReference != "sub_889" {
		t.Fatalf("decoded reward redemption outcome = %#v", decoded.Outcome)
	}

	bare, err := json.Marshal(wallet.RewardRedemption{})
	if err != nil {
		t.Fatalf("marshal empty reward redemption: %v", err)
	}
	if strings.Contains(string(bare), `"outcome"`) {
		t.Fatalf("an unfulfilled reward redemption must omit its outcome: %s", bare)
	}
}
