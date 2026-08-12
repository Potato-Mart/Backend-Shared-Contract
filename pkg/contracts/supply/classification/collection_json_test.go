package classification_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/classification"
)

func TestCollectionAndCategorySlugsAreOptional(t *testing.T) {
	body, err := json.Marshal(classification.Collection{
		ID:   "col_frozen",
		Slug: "frozen",
		Name: []localization.LocalizedName{{Language: "en", Name: "Frozen"}},
	})
	if err != nil {
		t.Fatalf("marshal collection with slug: %v", err)
	}
	if !strings.Contains(string(body), `"slug":"frozen"`) {
		t.Fatalf("Collection JSON = %s, want slug when present", body)
	}

	body, err = json.Marshal(classification.Collection{
		ID:   "col_frozen",
		Name: []localization.LocalizedName{{Language: "en", Name: "Frozen"}},
		CategoryTags: []classification.CategoryTag{
			{ID: "tag_hotpot", Name: []localization.LocalizedName{{Language: "en", Name: "Hotpot"}}, CollectionID: "col_frozen"},
		},
	})
	if err != nil {
		t.Fatalf("marshal collection: %v", err)
	}
	if strings.Contains(string(body), `"slug"`) {
		t.Fatalf("empty collection/category slugs should be omitted, got %s", body)
	}

	body, err = json.Marshal(classification.CollectionRef{
		ID:   "col_frozen",
		Name: []localization.LocalizedName{{Language: "en", Name: "Frozen"}},
	})
	if err != nil {
		t.Fatalf("marshal collection ref: %v", err)
	}
	if strings.Contains(string(body), `"slug"`) {
		t.Fatalf("empty collection ref slug should be omitted, got %s", body)
	}
}
