package measurement

import (
	"encoding/json"
	"testing"
)

func TestNetContentJSONKeepsExactIntegerMantissas(t *testing.T) {
	value := NetContent{
		NetQuantity:     Measure{Amount: 125, Exponent: -1, Unit: "g"},
		StandardMeasure: Measure{Amount: 100, Exponent: 0, Unit: "g"},
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal net content: %v", err)
	}
	want := `{"net_quantity":{"amount":125,"exponent":-1,"unit":"g"},"standard_measure":{"amount":100,"exponent":0,"unit":"g"}}`
	if string(payload) != want {
		t.Fatalf("NetContent JSON = %s, want %s", payload, want)
	}

	var decoded NetContent
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal net content: %v", err)
	}
	if decoded.NetQuantity.Amount != 125 || decoded.NetQuantity.Exponent != -1 || decoded.StandardMeasure.Unit != "g" {
		t.Fatalf("net content did not round-trip: %+v", decoded)
	}
}
