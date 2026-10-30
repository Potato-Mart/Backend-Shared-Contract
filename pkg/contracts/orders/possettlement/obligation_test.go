package possettlement

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

// This fixed vector is also checked independently by Orders and Supply. No
// service-specific canonicalization implementation is imported into the model.
func TestObligationPreservesCrossOwnerCanonicalVector(t *testing.T) {
	raw, err := os.ReadFile("testdata/obligation_v2.json")
	if err != nil {
		t.Fatal(err)
	}
	var value Obligation
	if err := json.Unmarshal(raw, &value); err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	decoder := json.NewDecoder(bytes.NewReader(encoded))
	decoder.UseNumber()
	var fields any
	if err := decoder.Decode(&fields); err != nil {
		t.Fatal(err)
	}
	canonical, err := json.Marshal(fields)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(append([]byte("orders-pos-obligation-v2\x00"), canonical...))
	if got := hex.EncodeToString(sum[:]); got != "c5b9a907c42f3bde231d2d0d9477e67ea044248897c610382551ae67a8cfd54f" {
		t.Fatalf("owner vector changed: %s", got)
	}
	if len(value.Lines) != 4 || value.Lines[0].Reservations[0].Allocation.KnownLotPhysical == nil || value.Lines[1].Reservations[0].Allocation.UnknownIdentityPhysical == nil || value.Lines[2].Reservations[0].Allocation.UnallocatedNoLotPhysical == nil || len(value.Lines[3].Reservations) != 0 {
		t.Fatal("native source partition was lost")
	}
}

func TestOriginalTenderIdentitiesAndBarrierSurviveJSON(t *testing.T) {
	want := Instrument{Kind: "gift_card", Authority: "payments", ID: "payment", BenefitReservationID: "reservation", WalletTransactionID: "debit", GiftCardCode: "fixture-card"}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var got Instrument
	if err := json.Unmarshal(raw, &got); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("original instrument identity lost: %+v %v", got, err)
	}
	barrier := RefundBarrier{OrderNumber: "order", Sequence: 9007199254740993, PrefixSHA256: "prefix", CoverageRevision: "writers-v2"}
	raw, err = json.Marshal(barrier)
	if err != nil {
		t.Fatal(err)
	}
	var decoded RefundBarrier
	if err := json.Unmarshal(raw, &decoded); err != nil || decoded != barrier {
		t.Fatalf("exact refund sequence lost: %+v %v", decoded, err)
	}
}

func TestObligationKeepsZeroTenderArrayIdentity(t *testing.T) {
	nilTender, err := json.Marshal(Obligation{Version: 2, ZeroDue: true})
	if err != nil {
		t.Fatal(err)
	}
	emptyTender, err := json.Marshal(Obligation{Version: 2, ZeroDue: true, Instruments: []Instrument{}})
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(nilTender, emptyTender) {
		t.Fatal("nil and empty tender arrays were conflated")
	}
}
