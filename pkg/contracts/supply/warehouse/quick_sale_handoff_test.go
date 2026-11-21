package warehouse

import (
	"encoding/json"
	"testing"
	"time"
)

func TestQuickSaleHandoffJSON(t *testing.T) {
	cutoff := time.Date(2026, 11, 21, 13, 0, 0, 0, time.UTC)
	for _, test := range []struct {
		value QuickSaleHandoff
		json  string
	}{
		{QuickSaleHandoff{}, `{}`},
		{QuickSaleHandoff{HandoffBefore: &cutoff}, `{"handoff_before":"2026-11-21T13:00:00Z"}`},
	} {
		got, err := json.Marshal(test.value)
		if err != nil || string(got) != test.json {
			t.Fatalf("JSON = %s, %v, want %s", got, err, test.json)
		}
	}
}
