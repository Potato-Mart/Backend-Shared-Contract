package order_test

import (
	"encoding/json"
	"strings"
	"testing"

	sales "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/wallet/wallet_enums"
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
	}
	zero := 0
	for _, status := range statuses {
		status := status
		t.Run(status.String(), func(t *testing.T) {
			payload, err := json.Marshal(sales.CustomerPaymentSummary{
				PointsAwardStatus:   status,
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
