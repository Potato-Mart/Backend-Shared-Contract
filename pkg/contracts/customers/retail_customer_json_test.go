package customers_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/customers"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/membership"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/customer"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/membership"
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
			Gender:   customerenum.CustomerGenderNonBinary,
		},
		Lifecycle: customers.RetailCustomerLifecycle{
			Status: customerenum.CustomerStatusActive,
		},
		Membership: customers.RetailCustomerMembershipProfile{
			MembershipAccountID: "mem_retail_123",
			Summary: &membership.MembershipAccountSummary{
				ID:              "mem_retail_123",
				Owner:           membership.MembershipOwnerRef{OwnerType: membershipenum.MembershipOwnerTypeRetailCustomer, OwnerID: "retail_123"},
				TierKey:         "standard",
				Status:          membershipenum.MembershipAccountStatusActive,
				AvailablePoints: 120,
				TotalPoints:     120,
			},
		},
		Marketing: customers.RetailCustomerMarketingProfile{EmailOptIn: true},
		Commerce:  customers.RetailCustomerCommerceProfile{TotalOrders: 2},
		Referral: &customers.RetailCustomerReferralProfile{
			Code:                      "REF-ADA",
			ReferrerCustomerNumber:    "RC-REF",
			UsedReferralCodeConfirmed: true,
			UsedByCount:               2,
			RewardVouchersIssued:      1,
			RewardVoucherCodes:        []string{"REF-ABC"},
		},
		ProfileCompletion: &customers.RetailCustomerProfileCompletion{
			Percent:         70,
			CompletedFields: []string{"email", "full_name", "date_of_birth", "phone", "gender"},
			MissingFields:   []string{"default_shipping", "default_billing"},
		},
		DefaultBilling: &common.ContactAddress{
			Address: &common.Address{Label: "Billing", Line1: "1 Account St", City: "Springvale", State: "VIC", Postcode: "3171", Country: "AU"},
		},
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
		"referral",
		"profile_completion",
		"default_billing",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("RetailCustomer JSON missing %q: %s", key, payload)
		}
	}
	if _, ok := got["identity"]; ok {
		t.Fatalf("RetailCustomer JSON should not include nested identity: %s", payload)
	}
	if _, ok := got["recent_orders"]; ok {
		t.Fatalf("RetailCustomer JSON should not include embedded recent orders: %s", payload)
	}
	basicInfo, ok := got["basic_info"].(map[string]any)
	if !ok {
		t.Fatalf("RetailCustomer JSON basic_info should be an object: %s", payload)
	}
	if _, ok := basicInfo["customer_number"]; ok {
		t.Fatalf("RetailCustomer JSON should not include nested customer_number: %s", payload)
	}
	if basicInfo["gender"] != "non_binary" {
		t.Fatalf("RetailCustomer JSON gender = %v, want non_binary: %s", basicInfo["gender"], payload)
	}
	referral := got["referral"].(map[string]any)
	if _, ok := referral["credited"]; ok {
		t.Fatalf("RetailCustomer referral should not include credited: %s", payload)
	}
	if referral["used_referral_code_confirmed"] != true || referral["used_by_count"].(float64) != 2 {
		t.Fatalf("RetailCustomer referral did not expose replacement fields: %+v", referral)
	}
}
