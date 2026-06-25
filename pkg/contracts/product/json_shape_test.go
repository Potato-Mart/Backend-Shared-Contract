package product

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
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

func TestProductSellingJSONGroup(t *testing.T) {
	body, err := json.Marshal(Product{
		ID:      "prd_1",
		SKUCode: "A0001",
		SKU:     "A",
		Name:    "Wholesale-only product",
		Selling: &Selling{
			Channels:   []enums.OrderType{enums.OrderTypeB2B, enums.OrderTypePOS},
			BuyerTypes: []enums.BuyerType{enums.BuyerTypeWholesaleOrganisation},
			Visibility: enums.PriceVisibilityWholesaleApprovedOnly,
		},
	})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal product JSON: %v", err)
	}
	selling, ok := got["selling"].(map[string]any)
	if !ok {
		t.Fatalf("Product JSON missing selling group: %s", body)
	}
	if selling["visibility"] != "wholesale_approved_only" {
		t.Fatalf("selling.visibility = %v, want wholesale_approved_only (%s)", selling["visibility"], body)
	}

	var decoded Product
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal product: %v", err)
	}
	if decoded.Selling == nil || len(decoded.Selling.Channels) != 2 || decoded.Selling.BuyerTypes[0] != enums.BuyerTypeWholesaleOrganisation {
		t.Fatalf("selling rules did not round-trip: %+v", decoded.Selling)
	}
}

func TestProductOmitsEmptySelling(t *testing.T) {
	body, err := json.Marshal(Product{ID: "prd_1", SKUCode: "A0001", SKU: "A", Name: "No selling rules"})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	if strings.Contains(string(body), `"selling"`) {
		t.Fatalf("empty selling should be omitted, got %s", body)
	}
}
