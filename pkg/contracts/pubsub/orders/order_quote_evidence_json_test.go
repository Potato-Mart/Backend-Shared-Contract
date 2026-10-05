package orders_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pubsub/orders"
	"testing"
)

func TestPaidQuoteEvidenceCompatibility(t *testing.T) {
	for _, tc := range []struct {
		name, wire  string
		revision    int64
		fingerprint string
	}{
		{name: "legacy", wire: `{}`},
		{name: "numeric", wire: `{"quote_key":"quote-1","quote_revision":4}`, revision: 4},
		{name: "opaque", wire: `{"quote_key":"quote-1","quote_fingerprint":"00ff:opaque/revision-4"}`, fingerprint: "00ff:opaque/revision-4"},
		{name: "both", wire: `{"quote_key":"quote-1","quote_revision":4,"quote_fingerprint":"00ff:opaque/revision-4"}`, revision: 4, fingerprint: "00ff:opaque/revision-4"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var event orders.OrderPaidEvent
			if err := json.Unmarshal([]byte(tc.wire), &event); err != nil {
				t.Fatal(err)
			}
			if event.QuoteRevision != tc.revision || event.QuoteFingerprint != tc.fingerprint {
				t.Fatal("quote evidence changed")
			}
			encoded, err := json.Marshal(event)
			if err != nil {
				t.Fatal(err)
			}
			var got, want map[string]json.RawMessage
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatal(err)
			}
			for _, key := range []string{"quote_key", "quote_revision", "quote_fingerprint"} {
				if string(got[key]) != string(want[key]) {
					t.Fatalf("%s changed: got %s want %s", key, got[key], want[key])
				}
			}
		})
	}
}
