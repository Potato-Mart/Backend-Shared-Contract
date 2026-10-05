package payments_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pubsub/payments"
	"os"
	"reflect"
	"testing"
)

func TestTierRefundProofWireFixtures(t *testing.T) {
	body, err := os.ReadFile("testdata/tier_refund_proof.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixtures map[string]json.RawMessage
	if err := json.Unmarshal(body, &fixtures); err != nil {
		t.Fatal(err)
	}
	for name, wire := range fixtures {
		t.Run(name, func(t *testing.T) {
			var event payments.RefundCompletedEvent
			if err := json.Unmarshal(wire, &event); err != nil {
				t.Fatal(err)
			}
			encoded, err := json.Marshal(event)
			if err != nil {
				t.Fatal(err)
			}
			var want, got map[string]json.RawMessage
			if err := json.Unmarshal(wire, &want); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatal(err)
			}
			for key, expected := range want {
				var a, b any
				if err := json.Unmarshal(expected, &a); err != nil {
					t.Fatal(err)
				}
				if err := json.Unmarshal(got[key], &b); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(a, b) {
					t.Fatalf("%s changed: got %s want %s", key, got[key], expected)
				}
			}
			if name == "legacy" {
				if event.TierRefundProof != nil {
					t.Fatal("missing legacy proof became zero")
				}
				if _, ok := got["tier_refund_proof"]; ok {
					t.Fatal("legacy unexpectedly emits proof")
				}
				return
			}
			proof := event.TierRefundProof
			if proof == nil || !proof.Basis.IsValid() || proof.AllocationVersion != 1 {
				t.Fatal("proof identity lost")
			}
			if name == "quoted_positive" {
				if event.QualifyingSpendReversal.AmountMinor != 1100 || proof.NetQualifyingSpendReversal.AmountMinor != 900 {
					t.Fatal("gross/net money conflated")
				}
			}
			var nested map[string]json.RawMessage
			if err := json.Unmarshal(got["tier_refund_proof"], &nested); err != nil {
				t.Fatal(err)
			}
			if _, ok := nested["net_qualifying_spend_reversal"]; !ok {
				t.Fatal("required money omitted")
			}
			if name == "gift_zero" {
				if proof.Basis != membership_enums.TierRefundProofBasisExcludedGiftPurchase || proof.NetQualifyingSpendReversal.AmountMinor != 0 || proof.NetQualifyingSpendReversal.Currency != "AUD" {
					t.Fatal("gift zero lost")
				}
				for _, key := range []string{"quote_key", "quote_fingerprint"} {
					if _, ok := nested[key]; ok {
						t.Fatal("gift witness fabricated quote ref")
					}
				}
			}
			if name == "quoted_zero" {
				if proof.NetQualifyingSpendReversal.AmountMinor != 0 || proof.QuoteKey == "" || proof.QuoteFingerprint == "" {
					t.Fatal("quoted zero lost original quote")
				}
			}
		})
	}
}

func TestTierRefundProofPreservesOpaqueValuesAndDoesNotInventAuthority(t *testing.T) {
	// Pure serialization preserves even opaque text; services validate native
	// authority. Standard JSON encoding must not trim or silently create a proof.
	wire := `{"basis":"quoted_credit","proof_id":" proof/exact ","proof_fingerprint":" fp:00A ","allocation_id":" alloc ","allocation_version":1,"allocation_fingerprint":" alloc-fp ","quote_key":" q ","quote_fingerprint":" q-fp ","net_qualifying_spend_reversal":{"amount_minor":0,"currency":"AUD"}}`
	var proof membership.TierRefundProof
	if err := json.Unmarshal([]byte(wire), &proof); err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(proof)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != wire {
		t.Fatalf("opaque values changed: %s", encoded)
	}
	var missing payments.RefundCompletedEvent
	if err := json.Unmarshal([]byte(`{"order_number":"ORDER-1","qualifying_spend_reversal":{"amount_minor":0,"currency":"AUD"}}`), &missing); err != nil {
		t.Fatal(err)
	}
	if missing.TierRefundProof != nil {
		t.Fatal("gross zero inferred proof")
	}
}
