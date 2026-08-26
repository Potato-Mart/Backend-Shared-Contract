package wholesale_test

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/customers/wholesale"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/customers/wholesale/wholesale_enums"
)

func TestWholesaleOrganisationJSONShape(t *testing.T) {
	approvedAt := time.Date(2026, 6, 18, 5, 0, 0, 0, time.UTC)
	organisation := wholesale.WholesaleOrganisation{
		OrganisationDetail: party.OrganisationDetail{
			PartyRef: party.PartyRef{
				ID:    "org_123",
				Name:  "Potato Buyer Co",
				Email: "accounts@example.com",
				Phone: "+61000000000",
			},
			LegalName: "Potato Buyer Pty Ltd",
			ABN:       "12345678901",
			RegisteredAddress: &party.ContactAddress{
				Address: &geography.Address{
					Label:              "HQ",
					Line1:              "1 Market Street",
					Locality:           "Sydney",
					AdministrativeArea: &geography.AdministrativeAreaRef{Code: "AU-NSW"},
					PostalCode:         "2000",
					Country:            geography.CountryRef{Code: "AU"},
				},
			},
		},
		PrincipalUserID:             "org_user_123",
		PrincipalAccountID:          "org_acct_123",
		PrimaryAuthIdentityID:       "auth_123",
		AuthIdentityIDs:             []string{"auth_123"},
		PrimaryOrganisationAccessID: "access_123",
		Status:                      wholesale_enums.WholesaleOrganisationStatusApproved,
		Approval:                    &audit.LifecycleAction{By: "admin_1", At: &approvedAt, Reason: "verified"},
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
		"principal_user_id",
		"principal_account_id",
		"primary_auth_identity_id",
		"auth_identity_ids",
		"primary_organisation_access_id",
		"status",
		"approval",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("WholesaleOrganisation JSON missing %q: %s", key, payload)
		}
	}
	for _, removed := range []string{"membership_account_id", "tier_key", "price_tier"} {
		if _, ok := got[removed]; ok {
			t.Fatalf("WholesaleOrganisation JSON contains removed field %q: %s", removed, payload)
		}
	}

	for key := range got {
		if strings.HasPrefix(key, "company_") {
			t.Fatalf("WholesaleOrganisation JSON should not include company-prefixed key %q: %s", key, payload)
		}
		if key == "approved_by" || key == "approved_at" {
			t.Fatalf("WholesaleOrganisation JSON should not include flat approval key %q: %s", key, payload)
		}
		if key == "primary_wholesale_customer_id" {
			t.Fatalf("WholesaleOrganisation JSON should not include retired primary customer key %q: %s", key, payload)
		}
	}
}
