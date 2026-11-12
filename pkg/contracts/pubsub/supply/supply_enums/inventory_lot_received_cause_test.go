package supply_enums

import "testing"

func TestInventoryLotReceivedCause(t *testing.T) {
	for value, want := range map[InventoryLotReceivedCause]string{
		InventoryLotReceivedCauseReceipt:         "receipt",
		InventoryLotReceivedCauseLotSplit:        "lot_split",
		InventoryLotReceivedCauseLotReassignment: "lot_reassignment",
		InventoryLotReceivedCauseCountGain:       "count_gain",
		InventoryLotReceivedCausePOSAddBack:      "pos_add_back",
	} {
		if !value.IsValid() || value.String() != want {
			t.Fatalf("lot received cause %q, want valid %q", value, want)
		}
	}
	for _, value := range []InventoryLotReceivedCause{"", "RECEIPT", "split", "transfer"} {
		if value.IsValid() {
			t.Fatalf("unexpected valid lot received cause %q", value)
		}
	}
}
