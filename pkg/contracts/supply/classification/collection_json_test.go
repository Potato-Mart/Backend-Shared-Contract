package classification_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/classification"
)

func TestV30CollectionAndTagsHaveRootSlugsAndCodeOnlyReferences(t *testing.T) {
	collection := classification.Collection{
		ID: "64c13ab08edf48a008793ca2", Code: "COL0001", Slug: "frozen-food",
		Name: []localization.LocalizedName{{Language: "en", Name: "Frozen Food"}},
		CategoryTags: []classification.CategoryTag{{
			ID: "64c13ab08edf48a008793ca3", Code: "TAG0001", Slug: "hotpot",
			Name: []localization.LocalizedName{{Language: "en", Name: "Hotpot"}}, CollectionCode: "COL0001",
		}},
	}
	body, err := json.Marshal(collection)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"code":"COL0001"`, `"slug":"frozen-food"`, `"collection_code":"COL0001"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("Collection JSON = %s, want %s", body, want)
		}
	}

	refBody, err := json.Marshal(classification.CollectionRef{Code: "COL0001"})
	if err != nil {
		t.Fatal(err)
	}
	if string(refBody) != `{"code":"COL0001"}` {
		t.Fatalf("CollectionRef leaked non-business identity: %s", refBody)
	}
}

func TestV3001CollectionIconUsesCodeOnlyCatalogMediaRef(t *testing.T) {
	body, err := json.Marshal(classification.Collection{Icon: &classification.ObjectMediaRef{Code: "MED-COLLECTION"}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"icon":{"code":"MED-COLLECTION"}`) || strings.Contains(string(body), `"url"`) {
		t.Fatalf("Collection icon must persist only the media code: %s", body)
	}
}
