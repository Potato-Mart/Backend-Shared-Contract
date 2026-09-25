package order_enums

import (
	"encoding/json"
	"testing"
)

func TestFulfillmentStatusWireValuesRoundTrip(t *testing.T) {
	cases := []struct {
		status   FulfillmentStatus
		wireJSON string
	}{
		{status: FulfillmentStatusCancelled, wireJSON: `"cancelled"`},
		{status: FulfillmentStatusFulfilled, wireJSON: `"fulfilled"`},
	}
	for _, tc := range cases {
		t.Run(string(tc.status), func(t *testing.T) {
			if !tc.status.IsValid() {
				t.Fatalf("%q should be valid", tc.status)
			}
			raw, err := json.Marshal(tc.status)
			if err != nil {
				t.Fatalf("marshal status: %v", err)
			}
			if string(raw) != tc.wireJSON {
				t.Fatalf("status JSON = %s, want %s", raw, tc.wireJSON)
			}
			var decoded FulfillmentStatus
			if err := json.Unmarshal(raw, &decoded); err != nil {
				t.Fatalf("unmarshal status: %v", err)
			}
			if decoded != tc.status || !decoded.IsValid() {
				t.Fatalf("decoded status = %q, want %q", decoded, tc.status)
			}
		})
	}
}
