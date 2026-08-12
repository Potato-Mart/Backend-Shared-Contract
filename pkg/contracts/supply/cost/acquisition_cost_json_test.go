package cost

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
)

func TestBaseAcquisitionCostIsTaxExclusiveAndRevisioned(t *testing.T) {
	effectiveFrom := time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC)
	payload, err := json.Marshal(BaseAcquisitionCost{
		ID: "cost_1", SKUID: "sku_a00001", Currency: "AUD",
		Amount:   money.Money{AmountMinor: 550, Currency: "AUD"},
		Revision: 5, EffectiveFrom: effectiveFrom,
	})
	if err != nil {
		t.Fatalf("marshal base acquisition cost: %v", err)
	}
	for _, want := range []string{`"sku_id":"sku_a00001"`, `"amount":{"amount_minor":550,"currency":"AUD"}`, `"revision":5`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("BaseAcquisitionCost JSON = %s, want %s", payload, want)
		}
	}
	for _, forbidden := range []string{`"tax_amount"`, `"market_id"`, `"price_book_id"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("acquisition cost leaked %s: %s", forbidden, payload)
		}
	}
}

func TestDepotCarryingCostDerivesTheAverageFromExactIntegers(t *testing.T) {
	asOf := time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC)
	average := money.Money{AmountMinor: 550, Currency: "AUD"}
	stocked, err := json.Marshal(DepotCarryingCost{
		ID: "carry_1", SKUID: "sku_a00001", DepotCode: "AU-VIC-MEL-DC-01",
		Currency: "AUD", CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2},
		TotalCarryingCostMinor: 66000, BaseUnitQuantity: 120,
		CurrentAverageUnitCost: &average, Revision: 3, AsOf: asOf,
	})
	if err != nil {
		t.Fatalf("marshal carrying cost: %v", err)
	}
	for _, want := range []string{
		`"total_carrying_cost_minor":66000`, `"base_unit_quantity":120`,
		`"current_average_unit_cost":{"amount_minor":550,"currency":"AUD"}`,
	} {
		if !strings.Contains(string(stocked), want) {
			t.Fatalf("DepotCarryingCost JSON = %s, want %s", stocked, want)
		}
	}

	// At zero on-hand quantity the carrying cost is zero and the current
	// average is absent rather than zero.
	empty, err := json.Marshal(DepotCarryingCost{
		ID: "carry_1", SKUID: "sku_a00001", DepotCode: "AU-VIC-MEL-DC-01",
		Currency: "AUD", TotalCarryingCostMinor: 0, BaseUnitQuantity: 0, Revision: 4, AsOf: asOf,
	})
	if err != nil {
		t.Fatalf("marshal empty carrying cost: %v", err)
	}
	if strings.Contains(string(empty), `"current_average_unit_cost"`) {
		t.Fatalf("zero-stock carrying cost must omit the average: %s", empty)
	}
}

func TestCarryingCostMovementReversalIsLinkedAndIdempotent(t *testing.T) {
	occurredAt := time.Date(2026, 8, 12, 1, 2, 3, 0, time.UTC)
	payload, err := json.Marshal(CarryingCostMovement{
		ID: "movement_2", SKUID: "sku_a00001", DepotCode: "AU-VIC-MEL-DC-01",
		ReferenceType: "supplier_return", ReferenceID: "return_1",
		IdempotencyKey: "supplier_return:return_1", ReversesMovementID: "movement_1",
		BaseUnitDelta: -12, CarryingCostDelta: -6600,
		BalanceBaseUnitsAfter: 108, BalanceCarryingCostMinorAfter: 59400,
		Currency: "AUD", Revision: 6, OccurredAt: occurredAt,
	})
	if err != nil {
		t.Fatalf("marshal carrying cost movement: %v", err)
	}
	for _, want := range []string{
		`"idempotency_key":"supplier_return:return_1"`,
		`"reverses_movement_id":"movement_1"`,
		`"carrying_cost_minor_delta":-6600`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("CarryingCostMovement JSON = %s, want %s", payload, want)
		}
	}
}
