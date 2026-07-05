package customers_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/customers"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

func TestRetailCustomerJSONShape(t *testing.T) {
	customer := customers.RetailCustomer{
		ID:                    "retail_123",
		CustomerNumber:        "RC-123",
		UserID:                "user_123",
		AccountID:             "acct_123",
		PrimaryAuthIdentityID: "auth_123",
		AuthIdentityIDs:       []string{"auth_123"},
		BasicInfo: customers.RetailCustomerBasicInfo{
			Name:     common.PersonName{DisplayName: "Retail Customer"},
			Contacts: common.ContactChannels{Email: "retail@example.com"},
		},
		Lifecycle: customers.RetailCustomerLifecycle{
			Status: enums.CustomerStatusActive,
		},
		Membership: customers.RetailCustomerMembershipProfile{
			MembershipAccountID: "mem_retail_123",
			Summary: &membership.MembershipAccountSummary{
				ID:              "mem_retail_123",
				Owner:           membership.MembershipOwnerRef{OwnerType: enums.MembershipOwnerTypeRetailCustomer, OwnerID: "retail_123"},
				TierKey:         "standard",
				Status:          enums.MembershipAccountStatusActive,
				AvailablePoints: 120,
				TotalPoints:     120,
			},
		},
		Marketing: customers.RetailCustomerMarketingProfile{EmailOptIn: true},
		Commerce:  customers.RetailCustomerCommerceProfile{TotalOrders: 2},
	}

	payload, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("Marshal RetailCustomer: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("Unmarshal RetailCustomer JSON: %v", err)
	}

	for _, key := range []string{
		"id",
		"customer_number",
		"user_id",
		"account_id",
		"primary_auth_identity_id",
		"auth_identity_ids",
		"basic_info",
		"lifecycle",
		"membership",
		"marketing",
		"commerce",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("RetailCustomer JSON missing %q: %s", key, payload)
		}
	}
	if _, ok := got["identity"]; ok {
		t.Fatalf("RetailCustomer JSON should not include nested identity: %s", payload)
	}
	basicInfo, ok := got["basic_info"].(map[string]any)
	if !ok {
		t.Fatalf("RetailCustomer JSON basic_info should be an object: %s", payload)
	}
	if _, ok := basicInfo["customer_number"]; ok {
		t.Fatalf("RetailCustomer JSON should not include nested customer_number: %s", payload)
	}
}
