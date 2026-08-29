package product

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/classification"
)

func TestProductStoryUsesLocalizedNameAndMediaCodes(t *testing.T) {
	body, err := json.Marshal(Product{Content: ProductContent{
		Name:   localization.LocalizedName{Language: "en", Name: "Golden Potato"},
		Images: &Images{Cover: &classification.ObjectMediaRef{Code: "MED-COVER"}, Gallery: []classification.ObjectMediaRef{{Code: "MED-GALLERY"}}},
	}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"name":{"language":"en","name":"Golden Potato"}`) || !strings.Contains(string(body), `"code":"MED-COVER"`) {
		t.Fatalf("Product story JSON = %s", body)
	}
}
