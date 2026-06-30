package product

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

func TestProductJSONIncludesTaxed(t *testing.T) {
	body, err := json.Marshal(Product{
		ID:      "prd_1",
		SKUCode: "A0001",
		SKU:     "A",
		Name:    "Taxed product",
		Description: []common.LocalizedDescription{
			{Language: "en", Description: "Localized description"},
		},
		Brand:           []common.LocalizedName{{Language: "en", Name: "Localized brand"}},
		Taxed:           true,
		Catalogue:       "winter",
		SupplierCode:    "sup_1",
		PlacingAreaCode: "A1",
	})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	if !strings.Contains(string(body), `"taxed":true`) {
		t.Fatalf("Product JSON = %s, want taxed true field", body)
	}
	if !strings.Contains(string(body), `"description":[`) || !strings.Contains(string(body), `"brand":[`) {
		t.Fatalf("Product JSON = %s, want localized description and brand arrays", body)
	}
	for _, keyValue := range []string{
		`"catalogue":"winter"`,
		`"supplier_code":"sup_1"`,
		`"placing_area_code":"A1"`,
	} {
		if !strings.Contains(string(body), keyValue) {
			t.Fatalf("Product JSON = %s, want flat %s", body, keyValue)
		}
	}
	if strings.Contains(string(body), `"identifiers"`) {
		t.Fatalf("Product JSON = %s, should not include nested identifiers", body)
	}
}

func TestSnapshotJSONIncludesTaxed(t *testing.T) {
	body, err := json.Marshal(Snapshot{
		ID:      "prd_1",
		SKUCode: "A0001",
		Name:    "Taxed snapshot",
		Description: []common.LocalizedDescription{
			{Language: "en", Description: "Snapshot description"},
		},
		Brand:        []common.LocalizedName{{Language: "en", Name: "Snapshot brand"}},
		SupplierCode: "sup_1",
		Taxed:        true,
	})
	if err != nil {
		t.Fatalf("marshal snapshot: %v", err)
	}
	if !strings.Contains(string(body), `"taxed":true`) {
		t.Fatalf("Snapshot JSON = %s, want taxed true field", body)
	}
	if strings.Contains(string(body), `"vendor"`) || !strings.Contains(string(body), `"supplier_code":"sup_1"`) {
		t.Fatalf("Snapshot JSON = %s, want supplier_code and no legacy vendor", body)
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
