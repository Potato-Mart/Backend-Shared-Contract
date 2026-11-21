package warehouse_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal/temporal_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
)

func TestDepotPickupHoursJSON(t *testing.T) {
	for _, test := range []struct {
		value warehouse.DepotPickupHours
		json  string
	}{
		{
			warehouse.DepotPickupHours{DayOfWeek: temporal_enums.DayOfWeekMonday, Opens: "09:00", Closes: "17:30"},
			`{"day_of_week":"MONDAY","opens":"09:00","closes":"17:30"}`,
		},
		{
			warehouse.DepotPickupHours{DayOfWeek: temporal_enums.DayOfWeekSunday, Closed: true},
			`{"day_of_week":"SUNDAY","closed":true}`,
		},
	} {
		got, err := json.Marshal(test.value)
		if err != nil || string(got) != test.json {
			t.Fatalf("JSON = %s, %v, want %s", got, err, test.json)
		}
	}
}

func TestDepotPickupFieldsAreOptional(t *testing.T) {
	legacy := marshalDepotObject(t, warehouse.Depot{Code: "AU-VIC-MEL-DC-01"})
	for _, key := range []string{"pickup_hours", "collection_instructions"} {
		if _, ok := legacy[key]; ok {
			t.Fatalf("depot without pickup details emitted %q: %+v", key, legacy)
		}
	}

	shape := marshalDepotObject(t, warehouse.Depot{
		Code: "AU-VIC-MEL-DC-01",
		PickupHours: []warehouse.DepotPickupHours{
			{DayOfWeek: temporal_enums.DayOfWeekMonday, Opens: "09:00", Closes: "17:00"},
		},
		CollectionInstructions: "Ring the bell at the loading dock.",
	})
	hours, ok := shape["pickup_hours"].([]any)
	if !ok || len(hours) != 1 {
		t.Fatalf("depot pickup_hours = %#v", shape["pickup_hours"])
	}
	if got := shape["collection_instructions"]; got != "Ring the bell at the loading dock." {
		t.Fatalf("depot collection_instructions = %#v", got)
	}

	var decoded warehouse.Depot
	raw := `{"code":"AU-VIC-MEL-DC-01","pickup_hours":[{"day_of_week":"SATURDAY","opens":"10:00","closes":"14:00"},{"day_of_week":"SUNDAY","closed":true}],"collection_instructions":"Use gate 2."}`
	if err := json.Unmarshal([]byte(raw), &decoded); err != nil {
		t.Fatalf("decode depot: %v", err)
	}
	if len(decoded.PickupHours) != 2 || decoded.PickupHours[1].DayOfWeek != temporal_enums.DayOfWeekSunday || !decoded.PickupHours[1].Closed || decoded.CollectionInstructions != "Use gate 2." {
		t.Fatalf("decoded depot pickup details = %+v", decoded)
	}
}
