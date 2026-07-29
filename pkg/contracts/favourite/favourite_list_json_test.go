package favourite_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/contracts/favourite"
	favouriteenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/favourite"
)

func TestFavouriteListJSONShape(t *testing.T) {
	now := time.Date(2026, 7, 18, 1, 2, 3, 0, time.UTC)
	list := favourite.FavouriteList{
		ID: "list_1",
		Owner: favourite.FavouriteListOwner{
			Type:             favouriteenum.FavouriteListOwnerTypeWholesaleOrganisation,
			OrganisationCode: "ORG-1",
		},
		Name:            "List-1",
		DefaultNameSlot: 1,
		Products: []favourite.FavouriteListProduct{{
			ProductSKUCode: "SKU-1",
			AddedAt:        now,
		}},
	}
	payload, err := json.Marshal(list)
	if err != nil {
		t.Fatalf("marshal favourite list: %v", err)
	}
	for _, want := range []string{`"type":"wholesale_organisation"`, `"organisation_code":"ORG-1"`, `"default_name_slot":1`, `"product_sku_code":"SKU-1"`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("favourite list JSON = %s, want %s", payload, want)
		}
	}
}
