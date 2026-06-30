package wholesale_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/wholesale"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

func TestWholesaleOrganisationJSONShape(t *testing.T) {
	approvedAt := time.Date(2026, 6, 18, 5, 0, 0, 0, time.UTC)
	organisation := wholesale.WholesaleOrganisation{
		OrganisationDetail: common.OrganisationDetail{
			PartyRef: common.PartyRef{
				ID:    "org_123",
				Name:  "Potato Buyer Co",
				Email: "accounts@example.com",
				Phone: "+61000000000",
			},
			LegalName: "Potato Buyer Pty Ltd",
			ABN:       "12345678901",
			RegisteredAddress: &common.ContactAddress{
				Address: &common.Address{
					Label:    "HQ",
					Line1:    "1 Market Street",
					City:     "Sydney",
					State:    "NSW",
					Postcode: "2000",
					Country:  "AU",
				},
			},
		},
		MembershipAccountID: "mem_org_123",
		Status:              enums.WholesaleOrganisationStatusApproved,
		TierKey:             "standard",
		Approval:            &common.LifecycleAction{By: "admin_1", At: &approvedAt, Reason: "verified"},
	}

	payload, err := json.Marshal(organisation)
	if err != nil {
		t.Fatalf("Marshal WholesaleOrganisation: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("Unmarshal WholesaleOrganisation JSON: %v", err)
	}

	for _, key := range []string{
		"id",
		"name",
		"legal_name",
		"abn",
		"registered_address",
		"membership_account_id",
		"status",
		"tier_key",
		"approval",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("WholesaleOrganisation JSON missing %q: %s", key, payload)
		}
	}

	for key := range got {
		if strings.HasPrefix(key, "company_") {
			t.Fatalf("WholesaleOrganisation JSON should not include company-prefixed key %q: %s", key, payload)
		}
		if key == "approved_by" || key == "approved_at" {
			t.Fatalf("WholesaleOrganisation JSON should not include flat approval key %q: %s", key, payload)
		}
	}
}
