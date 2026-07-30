package sales_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/sales"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/membership"
)

func TestCustomerPaymentSummaryPointAwardAndDebtRoundTrip(t *testing.T) {
	gross := 20
	debtRepaid := 5
	netCredit := 15
	remainingDebt := 3
	summary := sales.CustomerPaymentSummary{
		OrderNumber:         "MAMA260730ABC123",
		PointsAwardStatus:   membershipenum.PointAwardStatusAwarded,
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
	if decoded.PointsAwardStatus != membershipenum.PointAwardStatusAwarded ||
		decoded.PointsEarned == nil || *decoded.PointsEarned != gross ||
		decoded.PointsAppliedToDebt == nil || *decoded.PointsAppliedToDebt != debtRepaid ||
		decoded.PointsNetCredited == nil || *decoded.PointsNetCredited != netCredit ||
		decoded.PointDebtRemaining == nil || *decoded.PointDebtRemaining != remainingDebt {
		t.Fatalf("customer payment point summary did not round-trip: %+v", decoded)
	}
}

func TestCustomerPaymentSummaryPointAwardStatusesAndKnownZero(t *testing.T) {
	statuses := []membershipenum.PointAwardStatus{
		membershipenum.PointAwardStatusIneligible,
		membershipenum.PointAwardStatusDisabled,
		membershipenum.PointAwardStatusPending,
		membershipenum.PointAwardStatusAwarded,
		membershipenum.PointAwardStatusFailed,
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
