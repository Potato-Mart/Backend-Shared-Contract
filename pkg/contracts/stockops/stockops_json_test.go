package stockops_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/stockops"
)

func TestReservationCommandJSONShape(t *testing.T) {
	command := stockops.ReservationCommand{
		RefType:        "order",
		OrderNumber:    "PO-1",
		Lines:          []stockops.ReservationLine{{ProductSKUCode: "SKU-1", Qty: 2}},
		IdempotencyKey: "reserve:PO-1:SKU-1",
	}
	payload, err := json.Marshal(command)
	if err != nil {
		t.Fatalf("marshal command: %v", err)
	}
	for _, want := range []string{`"order_number":"PO-1"`, `"qty":2`, `"idempotency_key":"reserve:PO-1:SKU-1"`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("reservation JSON missing %s: %s", want, payload)
		}
	}
}

func TestPackingSettlementCommandJSONShape(t *testing.T) {
	command := stockops.PackingSettlementCommand{
		OrderNumber: "PO-1",
		Lines: []stockops.PackingSettlementLine{{
			ProductSKUCode: "SKU-1",
			SaleQty:        1,
			ReleaseQty:     1,
		}},
		IdempotencyKey: "packing:PO-1:v1",
	}
	payload, err := json.Marshal(command)
	if err != nil {
		t.Fatalf("marshal command: %v", err)
	}
	for _, want := range []string{`"sale_qty":1`, `"release_qty":1`, `"idempotency_key":"packing:PO-1:v1"`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("settlement JSON missing %s: %s", want, payload)
		}
	}
	if stockops.PathPackingSettlement != "/v1/internal/stock/packing-settlement" {
		t.Fatalf("packing settlement path = %q", stockops.PathPackingSettlement)
	}
}
