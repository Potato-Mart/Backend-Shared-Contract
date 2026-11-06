package supply_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	event "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/supply"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/supply/supply_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification/classification_enums"
)

func TestPackingFactJSONRoundTripsStageLinesAndStorageTotals(t *testing.T) {
	started := time.Date(2026, 11, 6, 1, 0, 0, 0, time.UTC)
	completed := started.Add(4 * time.Minute)
	fact := event.PackingFact{
		FactID: "packing:SO-1:2:0:AMBIENT_CHILLED", OrderNumber: "SO-1", FulfillmentGeneration: 2, ReopenCycle: 0,
		MarketCode: "MEL", DepotCode: "MAMA", Stage: supply_enums.PackingStageGroupAmbientChilled,
		PackedBy: "staff_1", StartedAt: &started, CompletedAt: &completed, PackedAt: completed.Add(time.Minute),
		PackageCount: 2,
		Lines: []event.PackingFactLine{
			{OrderItemID: "item_1", SKUCode: "SKU-1", StorageType: classification_enums.StorageAmbient, TotalBaseUnits: 24},
			{OrderItemID: "item_2", SKUCode: "SKU-2", StorageType: classification_enums.StorageChilled, TotalBaseUnits: 3},
		},
		StorageTotals: []event.PackingFactStorageTotal{
			{StorageType: classification_enums.StorageAmbient, TotalBaseUnits: 24, LineCount: 1},
			{StorageType: classification_enums.StorageChilled, TotalBaseUnits: 3, LineCount: 1},
		},
	}
	payload, err := json.Marshal(fact)
	if err != nil {
		t.Fatalf("marshal packing fact: %v", err)
	}
	for _, want := range []string{
		`"fact_id":"packing:SO-1:2:0:AMBIENT_CHILLED"`, `"reopen_cycle":0`, `"stage":"AMBIENT_CHILLED"`,
		`"started_at":"2026-11-06T01:00:00Z"`, `"storage_type":"CHILLED"`, `"total_base_units":24`, `"package_count":2`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("packing fact JSON missing %s: %s", want, payload)
		}
	}
	if strings.Contains(string(payload), `"occurred_at"`) {
		t.Fatalf("packing fact must rely on the envelope occurred_at: %s", payload)
	}
	var decoded event.PackingFact
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal packing fact: %v", err)
	}
	if !reflect.DeepEqual(decoded, fact) {
		t.Fatalf("packing fact round trip changed: %+v", decoded)
	}
}

func TestPackingFactOmitsUnrecordedStageTimes(t *testing.T) {
	payload, err := json.Marshal(event.PackingFact{FactID: "f", OrderNumber: "SO-1", Stage: supply_enums.PackingStageGroupFrozen})
	if err != nil {
		t.Fatalf("marshal packing fact: %v", err)
	}
	for _, forbidden := range []string{`"started_at"`, `"completed_at"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("unrecorded %s must be omitted: %s", forbidden, payload)
		}
	}
}
