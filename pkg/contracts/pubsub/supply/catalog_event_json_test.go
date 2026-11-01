package supply_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/supply"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/product/product_enums"
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

func TestCatalogProductStatusChangedEventJSONShape(t *testing.T) {
	changedAt := time.Date(2026, 11, 1, 7, 8, 9, 0, time.UTC)
	value := event.CatalogProductStatusChangedEvent{
		SKUCode:          "A00001",
		PreviousStatus:   product_enums.ProductStatusActive,
		Status:           product_enums.ProductStatusArchived,
		PreviousRevision: 6,
		Revision:         7,
		ChangedAt:        changedAt,
		ChangedBy:        "admin_1",
	}
	shape := marshalObject(t, value)
	want := map[string]any{
		"sku_code":          "A00001",
		"previous_status":   "active",
		"status":            "archived",
		"previous_revision": float64(6),
		"revision":          float64(7),
		"changed_at":        "2026-11-01T07:08:09Z",
		"changed_by":        "admin_1",
	}
	if len(shape) != len(want) {
		t.Fatalf("product status changed JSON keys = %+v, want %+v", shape, want)
	}
	for key, wantValue := range want {
		if shape[key] != wantValue {
			t.Fatalf("product status changed JSON %q = %#v, want %#v", key, shape[key], wantValue)
		}
	}
	if _, ok := shape["occurred_at"]; ok {
		t.Fatalf("product status payload must rely on envelope occurred_at: %+v", shape)
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal product status changed event: %v", err)
	}
	var decoded event.CatalogProductStatusChangedEvent
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal product status changed event: %v", err)
	}
	if decoded != value {
		t.Fatalf("product status changed event did not round-trip: %+v", decoded)
	}

	first := marshalObject(t, event.CatalogProductStatusChangedEvent{
		SKUCode:   "A00001",
		Status:    product_enums.ProductStatusDraft,
		Revision:  1,
		ChangedAt: changedAt,
	})
	for _, key := range []string{"previous_status", "changed_by"} {
		if _, ok := first[key]; ok {
			t.Fatalf("first status write must omit %q: %+v", key, first)
		}
	}
	for _, key := range []string{"sku_code", "status", "previous_revision", "revision", "changed_at"} {
		if _, ok := first[key]; !ok {
			t.Fatalf("product status changed JSON missing %q: %+v", key, first)
		}
	}
	if first["previous_revision"] != float64(0) {
		t.Fatalf("first status write previous_revision = %#v, want 0", first["previous_revision"])
	}
}
