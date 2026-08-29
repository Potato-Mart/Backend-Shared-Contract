package retail_test

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/party"
	customers "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/retail"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/retail/retail_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notifications/notification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/shipping"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/event"
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
			Name:     party.PersonName{DisplayName: "Retail Customer"},
			Contacts: party.ContactChannels{Email: "retail@example.com"},
			Gender:   retail_enums.CustomerGenderNonBinary,
		},
		Lifecycle: customers.RetailCustomerLifecycle{
			Status: retail_enums.CustomerStatusActive,
		},
		Commerce: customers.RetailCustomerCommerceProfile{TotalOrders: 2},
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
		DefaultBilling: &party.ContactAddress{
			Address: &geography.Address{Label: "Billing", Line1: "1 Account St", Locality: "Springvale", AdministrativeArea: &geography.AdministrativeAreaRef{Code: "AU-VIC"}, PostalCode: "3171", Country: geography.CountryRef{Code: "AU"}},
		},
		ShippingAddresses: []party.ContactAddress{
			{
				ID:      "addr_123",
				Contact: &party.Recipient{Name: "Retail Customer"},
				Address: &geography.Address{Label: "Home", Line1: "2 Account St", Locality: "Springvale", AdministrativeArea: &geography.AdministrativeAreaRef{Code: "AU-VIC"}, PostalCode: "3171", Country: geography.CountryRef{Code: "AU"}},
			},
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
		"commerce",
		"referral",
		"profile_completion",
		"default_billing",
		"shipping_addresses",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("RetailCustomer JSON missing %q: %s", key, payload)
		}
	}
	for _, removed := range []string{"identity", "recent_orders", "market_code", "country_code", "marketing", "analytics"} {
		if _, ok := got[removed]; ok {
			t.Fatalf("RetailCustomer JSON retained removed %q: %s", removed, payload)
		}
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
	addresses := got["shipping_addresses"].([]any)
	address := addresses[0].(map[string]any)
	if address["id"] != "addr_123" {
		t.Fatalf("RetailCustomer saved address id = %v, want addr_123: %s", address["id"], payload)
	}
}

func TestRetailCustomerDeliveryPreferenceJSONShape(t *testing.T) {
	customer := customers.RetailCustomer{
		ID:                    "customer_1",
		BillingSameAsDelivery: true,
		PreferredDeliverySlot: &shipping.PreferredDeliverySlot{
			Date: "2026-08-01", SlotID: "slot_7", ScheduleRevision: 7,
		},
	}
	payload, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("marshal retail customer delivery preference: %v", err)
	}
	for _, field := range []string{`"billing_same_as_delivery":true`, `"preferred_delivery_slot"`, `"schedule_revision":7`} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("retail customer delivery preference missing %s: %s", field, payload)
		}
	}
}

func TestRetailCustomerReceiptPreferencesJSONShape(t *testing.T) {
	electedAt := time.Date(2026, 8, 13, 1, 2, 3, 0, time.UTC)
	customer := customers.RetailCustomer{
		ID: "retail_123",
		BasicInfo: customers.RetailCustomerBasicInfo{
			Name:                   party.PersonName{DisplayName: "Retail Customer"},
			Contacts:               party.ContactChannels{Email: "retail@example.com"},
			PreferredContactMethod: retail_enums.PreferredContactMethodPhone,
		},
		ReceiptPreferences: &customers.RetailCustomerReceiptPreferences{
			Formats:   []retail_enums.ReceiptFormat{retail_enums.ReceiptFormatElectronic, retail_enums.ReceiptFormatPaper},
			UpdatedAt: &electedAt,
			Source:    "account_preferences",
		},
	}

	payload, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("marshal retail customer receipt preferences: %v", err)
	}
	for _, field := range []string{
		`"preferred_contact_method":"phone"`,
		`"receipt_preferences"`,
		`"formats":["electronic","paper"]`,
		`"updated_at":"2026-08-13T01:02:03Z"`,
		`"source":"account_preferences"`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("retail customer receipt preferences missing %s: %s", field, payload)
		}
	}

	var decoded customers.RetailCustomer
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal retail customer receipt preferences: %v", err)
	}
	if decoded.BasicInfo.PreferredContactMethod != retail_enums.PreferredContactMethodPhone ||
		decoded.ReceiptPreferences == nil || len(decoded.ReceiptPreferences.Formats) != 2 ||
		decoded.ReceiptPreferences.Formats[1] != retail_enums.ReceiptFormatPaper ||
		decoded.ReceiptPreferences.UpdatedAt == nil || !decoded.ReceiptPreferences.UpdatedAt.Equal(electedAt) {
		t.Fatalf("receipt preferences did not round-trip: %+v", decoded.ReceiptPreferences)
	}

	summary := customers.RetailCustomerSummary{
		ID:     "retail_123",
		Status: retail_enums.CustomerStatusActive,
		ReceiptPreferences: &customers.RetailCustomerReceiptPreferences{
			Formats: []retail_enums.ReceiptFormat{retail_enums.ReceiptFormatElectronic},
		},
	}
	summaryPayload, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("marshal retail customer summary: %v", err)
	}
	if !strings.Contains(string(summaryPayload), `"receipt_preferences":{"formats":["electronic"]}`) {
		t.Fatalf("summary receipt preferences missing: %s", summaryPayload)
	}
}

func TestRetailCustomerReceiptPreferencesOmittedWhenAbsent(t *testing.T) {
	customer := customers.RetailCustomer{ID: "retail_123"}

	payload, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("marshal retail customer: %v", err)
	}
	for _, absent := range []string{"receipt_preferences", "preferred_contact_method"} {
		if strings.Contains(string(payload), absent) {
			t.Fatalf("absent %s must be omitted: %s", absent, payload)
		}
	}
}

func TestNotificationPreferencesChangedEventContainsOnlyChangedIdentifiers(t *testing.T) {
	changedAt := time.Date(2026, 7, 30, 2, 3, 4, 0, time.UTC)
	preferencesEvent := event.NotificationPreferencesChangedEvent{
		UserID:              "user_1",
		CustomerNumber:      "RC-1",
		PreferencesRevision: 3,
		ChangedTopicCodes:   []string{"order_status"},
		ChangedChannels:     []notification_enums.NotificationChannel{notification_enums.NotificationChannelPush},
		Source:              "account_preferences",
		ChangedAt:           changedAt,
		RequestID:           "req_1",
	}

	payload, err := json.Marshal(preferencesEvent)
	if err != nil {
		t.Fatalf("marshal notification preferences changed event: %v", err)
	}
	for _, field := range []string{
		`"user_id":"user_1"`,
		`"preferences_revision":3`,
		`"changed_topic_codes":["order_status"]`,
		`"changed_channels":["push"]`,
		`"source":"account_preferences"`,
		`"changed_at":"2026-07-30T02:03:04Z"`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("notification preferences event missing %s: %s", field, payload)
		}
	}

	var decoded event.NotificationPreferencesChangedEvent
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal notification preferences changed event: %v", err)
	}
	if decoded.UserID != "user_1" || len(decoded.ChangedChannels) != 1 || decoded.ChangedChannels[0] != notification_enums.NotificationChannelPush || !decoded.ChangedAt.Equal(changedAt) {
		t.Fatalf("notification preferences event did not round-trip: %+v", decoded)
	}
}
