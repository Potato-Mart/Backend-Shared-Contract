package marketing_test

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/insights/marketing"

	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/insights/marketing/marketing_enums"
)

func TestMarketingCampaignIncludesGeographicScopeAndScheduleTimezone(t *testing.T) {
	value := marketing.MarketingCampaign{
		ID: "campaign_1", Name: "National launch",
		Channel:          marketing_enums.MarketingChannelEmail,
		Status:           marketing_enums.MarketingCampaignStatusDraft,
		GeographicScope:  geography.GeographicScope{Mode: geography_enums.GeographicScopeModeGlobal},
		ScheduleTimezone: "Etc/UTC",
	}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal marketing campaign: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal marketing campaign JSON: %v", err)
	}
	scope, ok := got["geographic_scope"].(map[string]any)
	if !ok || scope["mode"] != "GLOBAL" || got["schedule_timezone"] != "Etc/UTC" {
		t.Fatalf("marketing campaign geographic schedule mismatch: %s", payload)
	}
}

func TestMarketingCampaignRecipientJSONGroupsContactChannels(t *testing.T) {
	recipient := marketing.MarketingCampaignRecipient{
		ID:             "recipient_1",
		CampaignID:     "campaign_1",
		CustomerNumber: "customer_1",
		Contacts: party.ContactChannels{
			Email:           "buyer@example.com",
			Phone:           "+61000000000",
			ExternalHandles: map[string]string{"line_id": "line_1"},
		},
		CustomerName: "Retail Customer",
		Status:       marketing_enums.MarketingRecipientStatusPending,
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
	if decoded.Contacts.Email != "buyer@example.com" || decoded.Contacts.ExternalHandles["line_id"] != "line_1" {
		t.Fatalf("contacts did not round-trip: %+v", decoded.Contacts)
	}
}
