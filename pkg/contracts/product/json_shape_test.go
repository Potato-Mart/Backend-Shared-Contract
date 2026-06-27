package product

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
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
		Brand: []common.LocalizedName{{Language: "en", Name: "Localized brand"}},
		Taxed: true,
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
}

func TestSnapshotJSONIncludesTaxed(t *testing.T) {
	body, err := json.Marshal(Snapshot{
		ID:      "prd_1",
		SKUCode: "A0001",
		Name:    "Taxed snapshot",
		Description: []common.LocalizedDescription{
			{Language: "en", Description: "Snapshot description"},
		},
		Brand:    []common.LocalizedName{{Language: "en", Name: "Snapshot brand"}},
		Supplier: "sup_1",
		Taxed:    true,
	})
	if err != nil {
		t.Fatalf("marshal snapshot: %v", err)
	}
	if !strings.Contains(string(body), `"taxed":true`) {
		t.Fatalf("Snapshot JSON = %s, want taxed true field", body)
	}
	if strings.Contains(string(body), `"vendor"`) || !strings.Contains(string(body), `"supplier":"sup_1"`) {
		t.Fatalf("Snapshot JSON = %s, want supplier and no legacy vendor", body)
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

func TestProductDecodesLegacyDisplayFields(t *testing.T) {
	var got Product
	if err := json.Unmarshal([]byte(`{
		"id":"prd_1",
		"sku_code":"A0001",
		"sku":"A",
		"name":"Legacy",
		"description":"Flat legacy description",
		"brand":"Flat legacy brand",
		"identifiers":{"vendor":"SUP-1"}
	}`), &got); err != nil {
		t.Fatalf("unmarshal legacy product: %v", err)
	}
	if len(got.Description) != 1 || got.Description[0].Description != "Flat legacy description" {
		t.Fatalf("description = %+v, want localized legacy description", got.Description)
	}
	if len(got.Brand) != 1 || got.Brand[0].Name != "Flat legacy brand" {
		t.Fatalf("brand = %+v, want localized legacy brand", got.Brand)
	}
	if got.Identifiers.Supplier != "SUP-1" {
		t.Fatalf("supplier = %q, want vendor backfill", got.Identifiers.Supplier)
	}
}

func TestSnapshotDecodesLegacyBrandAndVendor(t *testing.T) {
	var got Snapshot
	if err := json.Unmarshal([]byte(`{"id":"prd_1","sku_code":"A0001","brand":"Legacy brand","vendor":"SUP-1"}`), &got); err != nil {
		t.Fatalf("unmarshal legacy snapshot: %v", err)
	}
	if len(got.Brand) != 1 || got.Brand[0].Name != "Legacy brand" {
		t.Fatalf("brand = %+v, want localized legacy brand", got.Brand)
	}
	if got.Supplier != "SUP-1" {
		t.Fatalf("supplier = %q, want vendor backfill", got.Supplier)
	}
}
