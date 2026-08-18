package product

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/security"
)

func TestV29ProductStoryUsesLocalizedNameAndMediaCodes(t *testing.T) {
	body, err := json.Marshal(Product{Content: ProductContent{
		Name:   localization.LocalizedName{Language: "en", Name: "Golden Potato"},
		Images: &Images{Cover: &security.ObjectMedia{Code: "MED-COVER"}, Gallery: []security.ObjectMedia{{Code: "MED-GALLERY"}}},
	}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"name":{"language":"en","name":"Golden Potato"}`) || !strings.Contains(string(body), `"code":"MED-COVER"`) {
		t.Fatalf("Product story JSON = %s", body)
	}
}
