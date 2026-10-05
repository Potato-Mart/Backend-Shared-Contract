package supply_test

import (
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/money"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pubsub/supply"
)

func TestCatalogBaseCostChangedEventJSONShape(t *testing.T) {
	now := time.Date(2026, 8, 12, 7, 8, 9, 0, time.UTC)
	previous := money.Money{AmountMinor: 500, Currency: "AUD"}
	shape := marshalObject(t, event.CatalogBaseCostChangedEvent{
		SKUCode:          "sku_a00001",
		Currency:         "AUD",
		PreviousAmount:   &previous,
		Amount:           money.Money{AmountMinor: 550, Currency: "AUD"},
		PreviousRevision: 4,
		Revision:         5,
		SourceType:       "supplier_invoice",
		SourceID:         "invoice_1",
		EffectiveFrom:    now,
	})
	for _, key := range []string{"sku_code", "currency", "previous_amount", "amount", "previous_revision", "revision", "effective_from"} {
		if _, ok := shape[key]; !ok {
			t.Fatalf("base cost changed JSON missing %q: %+v", key, shape)
		}
	}
	if shape["revision"] != float64(5) {
		t.Fatalf("base cost changed identity did not marshal: %+v", shape)
	}
	if _, ok := shape["occurred_at"]; ok {
		t.Fatalf("base cost payload must rely on envelope occurred_at: %+v", shape)
	}
}
