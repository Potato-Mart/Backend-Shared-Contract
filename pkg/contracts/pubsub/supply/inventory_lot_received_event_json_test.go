package supply_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging/packaging_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/supply"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/supply/supply_enums"
)

func TestInventoryLotReceivedEventCarriesCauseAndOriginLot(t *testing.T) {
	receivedAt := time.Date(2026, 11, 12, 1, 2, 3, 0, time.UTC)
	split := event.InventoryLotReceivedEvent{
		LotID: "lot_2", SKUCode: "A00001", DepotCode: "AU-VIC-MEL-DC-01", DestinationBucketID: "bucket_each",
		ReceivedComposition: composition(packaging_enums.PackageHandlingUnitEach, "pkg_each", 6, 1),
		MovementID:          "movement_split", LotRevision: 1, ReceivedAt: receivedAt,
		Cause: supply_enums.InventoryLotReceivedCauseLotSplit, OriginLotID: "lot_1",
	}
	shape := marshalObject(t, split)
	if shape["cause"] != "lot_split" || shape["origin_lot_id"] != "lot_1" {
		t.Fatalf("lot split event JSON = %+v", shape)
	}
	raw, err := json.Marshal(split)
	if err != nil {
		t.Fatal(err)
	}
	var got event.InventoryLotReceivedEvent
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got.Cause != supply_enums.InventoryLotReceivedCauseLotSplit || !got.Cause.IsValid() || got.OriginLotID != "lot_1" {
		t.Fatalf("lot split event round trip = %+v", got)
	}

	// A payload from a publisher that predates Cause decodes with no cause and
	// re-encodes without the new keys.
	legacy := event.InventoryLotReceivedEvent{LotID: "lot_3", SKUCode: "A00001", DepotCode: "AU-VIC-MEL-DC-01", MovementID: "movement_receipt", ReceivedAt: receivedAt}
	shape = marshalObject(t, legacy)
	for _, key := range []string{"cause", "origin_lot_id"} {
		if _, ok := shape[key]; ok {
			t.Fatalf("legacy lot received JSON = %+v, want %s omitted", shape, key)
		}
	}
}
