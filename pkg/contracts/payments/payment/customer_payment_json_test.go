package payment_test

import (
	"encoding/json"
	"strings"
	"testing"

	sales "github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/pricing/wallet/wallet_enums"
)

func TestCustomerPaymentSummaryPointAwardAndDebtRoundTrip(t *testing.T) {
	gross := 20
	debtRepaid := 5
	netCredit := 15
	remainingDebt := 3
	summary := sales.CustomerPaymentSummary{
		OrderNumber:         "MAMA260730ABC123",
		PointsAwardStatus:   wallet_enums.PointAwardStatusAwarded,
		PointsEarned:        &gross,
		PointsAppliedToDebt: &debtRepaid,
		PointsNetCredited:   &netCredit,
		PointDebtRemaining:  &remainingDebt,
	}

	payload, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("marshal customer payment summary: %v", err)
	}
	for _, field := range []string{
		`"points_award_status":"awarded"`,
		`"points_earned":20`,
		`"points_applied_to_debt":5`,
		`"points_net_credited":15`,
		`"point_debt_remaining":3`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("customer payment summary missing %s: %s", field, payload)
		}
	}

	var decoded sales.CustomerPaymentSummary
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal customer payment summary: %v", err)
	}
	if decoded.PointsAwardStatus != wallet_enums.PointAwardStatusAwarded ||
		decoded.PointsEarned == nil || *decoded.PointsEarned != gross ||
		decoded.PointsAppliedToDebt == nil || *decoded.PointsAppliedToDebt != debtRepaid ||
		decoded.PointsNetCredited == nil || *decoded.PointsNetCredited != netCredit ||
		decoded.PointDebtRemaining == nil || *decoded.PointDebtRemaining != remainingDebt {
		t.Fatalf("customer payment point summary did not round-trip: %+v", decoded)
	}
}

func TestCustomerPaymentSummaryPointAwardStatusesAndKnownZero(t *testing.T) {
	statuses := []wallet_enums.PointAwardStatus{
		wallet_enums.PointAwardStatusIneligible,
		wallet_enums.PointAwardStatusDisabled,
		wallet_enums.PointAwardStatusPending,
		wallet_enums.PointAwardStatusAwarded,
		wallet_enums.PointAwardStatusFailed,
		wallet_enums.PointAwardStatusHeld,
		wallet_enums.PointAwardStatusReleased,
		wallet_enums.PointAwardStatusDropped,
	}
	zero := 0
	for _, status := range statuses {
		status := status
		if !status.IsValid() {
			t.Fatalf("invalid supported status: %s", status)
		}
		t.Run(status.String(), func(t *testing.T) {
			payload, err := json.Marshal(sales.CustomerPaymentSummary{
				PointsAwardStatus:   status,
				PointsHeld:          &zero,
				PointsEarned:        &zero,
				PointsAppliedToDebt: &zero,
				PointsNetCredited:   &zero,
				PointDebtRemaining:  &zero,
			})
			if err != nil {
				t.Fatalf("marshal %s point award status: %v", status, err)
			}
			for _, field := range []string{
				`"points_award_status":"` + status.String() + `"`,
				`"points_held":0`,
				`"points_earned":0`,
				`"points_applied_to_debt":0`,
				`"points_net_credited":0`,
				`"point_debt_remaining":0`,
			} {
				if !strings.Contains(string(payload), field) {
					t.Fatalf("known-zero point summary missing %s: %s", field, payload)
				}
			}
		})
	}
}

func TestHeldSummaryPreservesUnknownStatusAndOptionalBalance(t *testing.T) {
	for _, raw := range []string{`{"points_award_status":"held","points_held":17}`, `{"points_award_status":"future-status","points_held":0}`} {
		var got sales.CustomerPaymentSummary
		if err := json.Unmarshal([]byte(raw), &got); err != nil {
			t.Fatal(err)
		}
		if got.PointsHeld == nil || got.PointsEarned != nil {
			t.Fatalf("held balance lost or inferred earned: %+v", got)
		}
		out, err := json.Marshal(got)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(out), `"points_award_status":"`+got.PointsAwardStatus.String()+`"`) {
			t.Fatal("status normalized")
		}
	}
	if wallet_enums.PointAwardStatus("future-status").IsValid() {
		t.Fatal("unknown status validated")
	}
}
