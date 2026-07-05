package marketing_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/marketing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

func TestMarketingCampaignRecipientJSONGroupsContactChannels(t *testing.T) {
	recipient := marketing.MarketingCampaignRecipient{
		ID:             "recipient_1",
		CampaignID:     "campaign_1",
		CustomerNumber: "customer_1",
		Contacts: common.ContactChannels{
			Email:  "buyer@example.com",
			Phone:  "+61000000000",
			LineID: "line_1",
		},
		CustomerName: "Retail Customer",
		Status:       enums.MarketingRecipientStatusPending,
		CreatedAt:    time.Date(2026, 6, 18, 6, 30, 0, 0, time.UTC),
	}

	payload, err := json.Marshal(recipient)
	if err != nil {
		t.Fatalf("marshal marketing recipient: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal marketing recipient JSON: %v", err)
	}

	for _, key := range []string{"id", "campaign_id", "customer_number", "customer_name", "status", "created_at"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("MarketingCampaignRecipient JSON missing top-level %q: %s", key, payload)
		}
	}
	if _, ok := got["contacts"]; !ok {
		t.Fatalf("MarketingCampaignRecipient JSON missing nested contacts: %s", payload)
	}
	for _, key := range []string{"email", "phone", "line_id"} {
		if _, ok := got[key]; ok {
			t.Fatalf("MarketingCampaignRecipient JSON should not include flat contact key %q: %s", key, payload)
		}
	}

	var decoded marketing.MarketingCampaignRecipient
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal marketing recipient: %v", err)
	}
	if decoded.Contacts.Email != "buyer@example.com" || decoded.Contacts.LineID != "line_1" {
		t.Fatalf("contacts did not round-trip: %+v", decoded.Contacts)
	}
}
