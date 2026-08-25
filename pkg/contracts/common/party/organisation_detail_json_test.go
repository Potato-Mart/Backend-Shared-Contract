package party

import (
	"encoding/json"
	"testing"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security"
)

func TestOrganisationDetailJSONUsesObjectMediaLogo(t *testing.T) {
	payload, err := json.Marshal(OrganisationDetail{
		PartyRef: PartyRef{ID: "org_1", Name: "Potato Mart"},
		Logo: &security.ObjectMedia{
			Code: "media_logo_1",
			URL:  "https://cdn.example.test/organisations/potato-mart.png",
		},
	})
	if err != nil {
		t.Fatalf("marshal organisation detail: %v", err)
	}

	var shape map[string]any
	if err := json.Unmarshal(payload, &shape); err != nil {
		t.Fatalf("unmarshal organisation detail: %v", err)
	}
	logo, ok := shape["logo"].(map[string]any)
	if !ok || logo["code"] != "media_logo_1" || logo["url"] != "https://cdn.example.test/organisations/potato-mart.png" {
		t.Fatalf("OrganisationDetail logo JSON = %s", payload)
	}
	if _, exists := shape["logo_url"]; exists {
		t.Fatalf("OrganisationDetail retained logo_url: %s", payload)
	}
}

func TestOrganisationDetailOmitsEmptyObjectMediaLogo(t *testing.T) {
	payload, err := json.Marshal(OrganisationDetail{PartyRef: PartyRef{ID: "org_1"}})
	if err != nil {
		t.Fatalf("marshal organisation detail: %v", err)
	}
	var shape map[string]any
	if err := json.Unmarshal(payload, &shape); err != nil {
		t.Fatalf("unmarshal organisation detail: %v", err)
	}
	if _, exists := shape["logo"]; exists {
		t.Fatalf("empty logo should be omitted: %s", payload)
	}
}
