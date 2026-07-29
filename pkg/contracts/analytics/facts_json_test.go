package analytics

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestItemFactsUseBrandID(t *testing.T) {
	for name, value := range map[string]any{
		"order":  OrderItemFact{ProductSKUCode: "A0001", BrandID: "64c13ab08edf48a008793ca1"},
		"refund": RefundItemFact{ProductSKUCode: "A0001", BrandID: "64c13ab08edf48a008793ca1"},
	} {
		t.Run(name, func(t *testing.T) {
			body, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(body), `"brand_id":"64c13ab08edf48a008793ca1"`) || strings.Contains(string(body), `"brand_key"`) {
				t.Fatalf("item fact JSON = %s", body)
			}
		})
	}
}
