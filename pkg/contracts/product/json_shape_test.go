package product

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestProductJSONIncludesTaxed(t *testing.T) {
	body, err := json.Marshal(Product{
		ID:      "prd_1",
		SKUCode: "A0001",
		SKU:     "A",
		Name:    "Taxed product",
		Taxed:   true,
	})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	if !strings.Contains(string(body), `"taxed":true`) {
		t.Fatalf("Product JSON = %s, want taxed true field", body)
	}
}

func TestSnapshotJSONIncludesTaxed(t *testing.T) {
	body, err := json.Marshal(Snapshot{
		ID:      "prd_1",
		SKUCode: "A0001",
		Name:    "Taxed snapshot",
		Taxed:   true,
	})
	if err != nil {
		t.Fatalf("marshal snapshot: %v", err)
	}
	if !strings.Contains(string(body), `"taxed":true`) {
		t.Fatalf("Snapshot JSON = %s, want taxed true field", body)
	}
}
