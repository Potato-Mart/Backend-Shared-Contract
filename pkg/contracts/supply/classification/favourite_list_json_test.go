package classification_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/classification/classification_enums"
)

func TestFavouriteListJSONShape(t *testing.T) {
	now := time.Date(2026, 7, 18, 1, 2, 3, 0, time.UTC)
	list := classification.FavouriteList{
		ID: "list_1",
		Owner: classification.FavouriteListOwner{
			Type:             classification_enums.FavouriteListOwnerTypeWholesaleOrganisation,
			OrganisationCode: "ORG-1",
		},
		Name:            "List-1",
		DefaultNameSlot: 1,
		Products: []classification.FavouriteListProduct{{
			SKUCode: "SKU-1",
			AddedAt: now,
		}},
	}
	payload, err := json.Marshal(list)
	if err != nil {
		t.Fatalf("marshal favourite list: %v", err)
	}
	for _, want := range []string{`"type":"wholesale_organisation"`, `"organisation_code":"ORG-1"`, `"default_name_slot":1`, `"sku_code":"SKU-1"`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("favourite list JSON = %s, want %s", payload, want)
		}
	}
}
