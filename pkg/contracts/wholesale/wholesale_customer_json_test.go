package wholesale_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/contracts/wholesale"
	wholesaleenum "github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/enums/wholesale"
)

func TestWholesaleCustomerJSONShape(t *testing.T) {
	customer := wholesale.WholesaleCustomer{
		OrganisationDetail: common.OrganisationDetail{
			PartyRef: common.PartyRef{
				ID:    "org_123",
				Code:  "WO-123",
				Name:  "Potato Buyer Co",
				Email: "accounts@example.com",
			},
			TradingName: "Buyer Co",
			LegalName:   "Potato Buyer Pty Ltd",
			ABN:         "12345678901",
		},
		PrincipalUserID:             "org_user_123",
		PrincipalAccountID:          "org_acct_123",
		PrimaryAuthIdentityID:       "auth_123",
		AuthIdentityIDs:             []string{"auth_123"},
		PrimaryOrganisationAccessID: "access_123",
		Status:                      wholesaleenum.WholesaleOrganisationStatusApproved,
		TierKey:                     "standard",
		PriceTier:                   1,
	}

	payload, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("Marshal WholesaleCustomer: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("Unmarshal WholesaleCustomer JSON: %v", err)
	}

	for _, key := range []string{
		"id",
		"code",
		"name",
		"trading_name",
		"legal_name",
		"abn",
		"principal_user_id",
		"principal_account_id",
		"primary_auth_identity_id",
		"auth_identity_ids",
		"primary_organisation_access_id",
		"status",
		"tier_key",
		"price_tier",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("WholesaleCustomer JSON missing %q: %s", key, payload)
		}
	}
	if _, ok := got["identity"]; ok {
		t.Fatalf("WholesaleCustomer JSON should not include nested identity: %s", payload)
	}
	for _, key := range []string{"customer_number", "basic_info", "commercial", "account_profile", "organisation_access_id", "primary_wholesale_customer_id"} {
		if _, ok := got[key]; ok {
			t.Fatalf("WholesaleCustomer JSON should not include person-profile key %q: %s", key, payload)
		}
	}
	if _, ok := got["membership_id"]; ok {
		t.Fatalf("WholesaleCustomer JSON should not include legacy membership_id: %s", payload)
	}
}
