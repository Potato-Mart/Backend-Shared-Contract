package pkg_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/marketing/campaign"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/marketing/campaign/campaign_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pubsub/envelope"
	events "github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pubsub/pricing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pubsub/routing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/supply/forecasting"
)

// Fixture containers are test-only. Campaign authoring, migration, eligibility
// and rendering remain responsibilities of the owning services and clients.
type campaignPlacementFixture struct {
	CanonicalCases    []campaignPlacementFixtureCase      `json:"canonical_cases"`
	InvalidPlacements []string                            `json:"invalid_placements"`
	EventCases        []campaignPlacementEventFixtureCase `json:"event_cases"`
}

type campaignPlacementFixtureCase struct {
	Name       string                              `json:"name"`
	Campaign   campaign.Campaign                   `json:"campaign"`
	Comparable forecasting.CampaignComparableEvent `json:"comparable"`
}

type campaignPlacementEventFixtureCase struct {
	Name     string                 `json:"name"`
	Topic    routing.EventTopic     `json:"topic"`
	Envelope envelope.EventEnvelope `json:"envelope"`
}

type campaignPlacementRawFixture struct {
	CanonicalCases    []json.RawMessage `json:"canonical_cases"`
	InvalidPlacements []string          `json:"invalid_placements"`
	EventCases        []json.RawMessage `json:"event_cases"`
}

func TestCampaignPlacementCanonicalRecords(t *testing.T) {
	fixture, raw := loadCampaignPlacementFixture(t)
	wants := map[string]campaign_enums.CampaignPlacement{
		"announcement_bar": campaign_enums.CampaignPlacementAnnouncementBar,
		"home_banner":      campaign_enums.CampaignPlacementHomeBanner,
		"account_banner":   campaign_enums.CampaignPlacementAccountBanner,
		"home_modal":       campaign_enums.CampaignPlacementHomeModal,
	}
	if len(fixture.CanonicalCases) != len(wants) {
		t.Fatal("placement fixture must include exactly four canonical surfaces")
	}
	seen := map[string]bool{}
	for i, original := range fixture.CanonicalCases {
		want, ok := wants[original.Name]
		if !ok || seen[original.Name] {
			t.Fatalf("unexpected or duplicate canonical placement %q", original.Name)
		}
		seen[original.Name] = true
		t.Run(original.Name, func(t *testing.T) {
			c, wire := roundTripCampaignPlacement[campaignPlacementFixtureCase](t, raw.CanonicalCases[i])
			if c.Campaign.Placement != want || c.Comparable.Placement != want || !want.IsValid() || want.String() != c.Name {
				t.Fatal("canonical placement lost its exact wire value or enum validity")
			}
			if c.Campaign.CampaignCode != "campaign-"+c.Name || c.Comparable.CampaignCode != c.Campaign.CampaignCode || c.Campaign.ID != "campaign-id-"+c.Name {
				t.Fatal("campaign or forecasting identity changed")
			}
			wantTitle := []localization.LocalizedName{{Language: "en", Name: "Autumn offers"}, {Language: "zh-TW", Name: "秋季優惠"}, {Language: "ja", Name: "秋の特典"}}
			if !reflect.DeepEqual(c.Campaign.Title, wantTitle) {
				t.Fatal("localized title content or ordering changed")
			}
			if c.Campaign.MarketCode != "AU-VIC" || c.Campaign.CountryCode != "AU" || c.Campaign.SeriesCode != "autumn-2026" || c.Campaign.NotificationTopicCode != "seasonal_offer" {
				t.Fatal("campaign geography, series or notification identity changed")
			}
			if c.Campaign.StartsAt == nil || c.Campaign.EndsAt == nil || c.Campaign.ActivatedAt == nil || c.Campaign.Revision != 13 || c.Campaign.ActivationRevision != 5 || c.Campaign.CreatedAt.IsZero() || c.Campaign.UpdatedAt.IsZero() || c.Campaign.CreatedBy != "fixture-editor" || c.Campaign.UpdatedBy != "fixture-editor" {
				t.Fatal("campaign lost scheduling, lifecycle or audit evidence")
			}
			if c.Comparable.SeriesCode != c.Campaign.SeriesCode || c.Comparable.MatchSource != "same_series" || !c.Comparable.StartsAt.Equal(*c.Campaign.StartsAt) || !c.Comparable.EndsAt.Equal(*c.Campaign.EndsAt) || c.Comparable.ScheduleTimezone != c.Campaign.ScheduleTimezone || c.Campaign.ScheduleTimezone != "Australia/Sydney" || !reflect.DeepEqual(c.Comparable.Audience, c.Campaign.Audience) || !reflect.DeepEqual(c.Comparable.GeographicScope, c.Campaign.GeographicScope) || !reflect.DeepEqual(c.Comparable.ResolvedSKUCodes, []string{"SKU-CAMPAIGN", "SKU-SECOND"}) {
				t.Fatal("forecasting placement cutover changed historical evidence")
			}
			if c.Campaign.Audience == nil || c.Campaign.Audience.CustomerType != campaign_enums.CampaignCustomerTypeRetail || c.Campaign.GeographicScope.Mode != "TARGETED" || len(c.Campaign.GeographicScope.Targets) != 1 {
				t.Fatal("campaign lost audience or geographic targeting")
			}
			wireCampaign := campaignPlacementRawObject(t, campaignPlacementRawObject(t, wire)["campaign"])
			if c.Name == "account_banner" {
				if c.Campaign.CTA != nil || c.Campaign.Media != nil || c.Campaign.CTAHref != "" || len(c.Campaign.CTAText) != 0 || len(c.Campaign.BenefitRefs) != 0 {
					t.Fatal("omitted presentation fields acquired values")
				}
				for _, key := range []string{"cta", "cta_href", "cta_text", "media", "benefit_refs"} {
					if _, present := wireCampaign[key]; present {
						t.Fatalf("omitted %s was serialized", key)
					}
				}
			} else {
				if c.Campaign.CTA == nil || c.Campaign.CTA.Type != campaign_enums.CampaignCTADestinationProduct || c.Campaign.CTA.SKUCode != "SKU-CAMPAIGN" || c.Campaign.CTAHref != "/products/SKU-CAMPAIGN" || len(c.Campaign.CTAText) != 2 || c.Campaign.Media == nil || c.Campaign.Media.Code != "media-campaign" || c.Campaign.Media.URL != "/v1/storefront/campaigns/"+c.Campaign.ID+"/media" {
					t.Fatal("populated CTA or media evidence changed")
				}
				if len(c.Campaign.BenefitRefs) != 2 || c.Campaign.BenefitRefs[0].Kind != "promotion" || c.Campaign.BenefitRefs[0].Code != "promotion-autumn" || !reflect.DeepEqual(c.Campaign.BenefitRefs[0].Name, wantTitle) || c.Campaign.BenefitRefs[1].Kind != "coupon" || c.Campaign.BenefitRefs[1].Code != "coupon-autumn" || len(c.Campaign.Message) != 2 || !reflect.DeepEqual(c.Campaign.Targets.SKUCodes, []string{"SKU-CAMPAIGN", "SKU-SECOND"}) || len(c.Campaign.Targets.Categories) != 1 {
					t.Fatal("campaign content, benefit references or product targeting changed")
				}
			}
			if c.Name == "home_modal" {
				if c.Campaign.Status != campaign_enums.CampaignStatusArchived || c.Campaign.IsActive || c.Campaign.DeactivatedAt == nil || c.Campaign.ArchivedAt == nil {
					t.Fatal("archived lifecycle evidence changed")
				}
			} else if c.Campaign.Status != campaign_enums.CampaignStatusActive || !c.Campaign.IsActive || c.Campaign.DeactivatedAt != nil || c.Campaign.ArchivedAt != nil {
				t.Fatal("active lifecycle evidence changed")
			}
		})
	}
}

func TestCampaignPlacementInvalidValuesPreserveRawStrings(t *testing.T) {
	fixture, raw := loadCampaignPlacementFixture(t)
	wants := []string{"top_banner", "home_hero", "modal", "checkout_notice", "product_notice", "", "future_sidebar"}
	if !reflect.DeepEqual(fixture.InvalidPlacements, wants) {
		t.Fatal("invalid fixture must cover every retired placement, empty and future values")
	}
	base := campaignPlacementRawObject(t, raw.CanonicalCases[0])
	for _, value := range fixture.InvalidPlacements {
		t.Run(fmt.Sprintf("%q", value), func(t *testing.T) {
			// Substitute fixture input only; there is no production normalization or alias map.
			campaignJSON := campaignPlacementWithRawValue(t, base["campaign"], value)
			comparableJSON := campaignPlacementWithRawValue(t, base["comparable"], value)
			c, _ := roundTripCampaignPlacement[campaign.Campaign](t, campaignJSON)
			comparable, _ := roundTripCampaignPlacement[forecasting.CampaignComparableEvent](t, comparableJSON)
			if string(c.Placement) != value || string(comparable.Placement) != value || c.Placement.String() != value || comparable.Placement.String() != value {
				t.Fatal("ordinary JSON decoding rewrote an unrecognized placement")
			}
			if c.Placement.IsValid() || comparable.Placement.IsValid() {
				t.Fatal("retired, empty or unknown placement became semantically valid")
			}
		})
	}
}

func TestCampaignPlacementEventCompatibility(t *testing.T) {
	fixture, raw := loadCampaignPlacementFixture(t)
	if len(fixture.EventCases) != 2 || fixture.EventCases[0].Name != "active" || fixture.EventCases[1].Name != "archived" {
		t.Fatal("event fixture must cover active and archived payloads")
	}
	for i, original := range fixture.EventCases {
		t.Run(original.Name, func(t *testing.T) {
			c, _ := roundTripCampaignPlacement[campaignPlacementEventFixtureCase](t, raw.EventCases[i])
			payload, wire := roundTripCampaignPlacement[events.CampaignChangedEvent](t, c.Envelope.Payload)
			if c.Topic != routing.EventTopicStorefrontEvents || c.Topic.String() != "storefront-events" || c.Envelope.EventType != routing.EventTypeCampaignChanged || c.Envelope.EventType.String() != "campaign.changed" || c.Envelope.EventVersion != "v2" {
				t.Fatal("campaign event topic, type or existing service-owned version changed")
			}
			if c.Envelope.EventID != "event-campaign-"+c.Name || c.Envelope.AggregateID != payload.CampaignCode || payload.CampaignCode != "campaign-home_banner" || !c.Envelope.OccurredAt.Equal(payload.ChangedAt) || !payload.RefetchRequired || payload.ScopeMode != "TARGETED" || payload.ScopeRevision != 3 || payload.ActivationRevision != 5 || payload.ContentRevision != int64(13+i) || payload.Status.String() != c.Name || payload.IsActive != (c.Name == "active") {
				t.Fatal("campaign event lost identity, activation, revision or refetch evidence")
			}
			if c.Name == "active" && (payload.SeriesCode != "autumn-2026" || payload.PromotionID != "promotion-autumn") {
				t.Fatal("optional campaign event links changed")
			}
			object := campaignPlacementRawObject(t, wire)
			for _, forbidden := range []string{"placement", "title", "message", "cta", "cta_href", "cta_text", "media", "targets", "audience", "benefit_refs", "provider", "recipient", "customer_number"} {
				if _, present := object[forbidden]; present {
					t.Fatalf("customer-safe campaign event exposed %s", forbidden)
				}
			}
		})
	}
}

func loadCampaignPlacementFixture(t *testing.T) (campaignPlacementFixture, campaignPlacementRawFixture) {
	t.Helper()
	data, err := os.ReadFile("testdata/campaign_placement_contract.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture campaignPlacementFixture
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&fixture); err != nil {
		t.Fatalf("decode campaign placement fixture: %v", err)
	}
	var raw campaignPlacementRawFixture
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatal(err)
	}
	return fixture, raw
}

func roundTripCampaignPlacement[T any](t *testing.T, raw json.RawMessage) (T, []byte) {
	t.Helper()
	// encoding/json compacts RawMessage payload whitespace when marshaling an
	// envelope. Compare typed evidence after the same whitespace-only compaction.
	var compact bytes.Buffer
	if err := json.Compact(&compact, raw); err != nil {
		t.Fatal(err)
	}
	var original T
	decoder := json.NewDecoder(bytes.NewReader(compact.Bytes()))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&original); err != nil {
		t.Fatalf("decode shared campaign record: %v", err)
	}
	wire, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	var decoded T
	if err := json.Unmarshal(wire, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(original, decoded) {
		t.Fatal("campaign record lost typed evidence on JSON round trip")
	}
	assertCampaignPlacementJSONShape(t, raw, wire, "record")
	return decoded, wire
}

// RawMessage preserves exact scalar tokens and array order without floating-point decoding.
func assertCampaignPlacementJSONShape(t *testing.T, before, after json.RawMessage, path string) {
	t.Helper()
	before, after = bytes.TrimSpace(before), bytes.TrimSpace(after)
	if len(before) == 0 || len(after) == 0 {
		t.Fatalf("missing JSON at %s", path)
	}
	if (before[0] == '{' || before[0] == '[') && before[0] != after[0] {
		t.Fatalf("JSON value type changed at %s", path)
	}
	switch before[0] {
	case '{':
		left, right := campaignPlacementRawObject(t, before), campaignPlacementRawObject(t, after)
		if len(left) != len(right) {
			t.Fatalf("JSON field set changed at %s", path)
		}
		for key, value := range left {
			assertCampaignPlacementJSONShape(t, value, right[key], path+"."+key)
		}
	case '[':
		var left, right []json.RawMessage
		if err := json.Unmarshal(before, &left); err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(after, &right); err != nil {
			t.Fatal(err)
		}
		if len(left) != len(right) {
			t.Fatalf("JSON array length changed at %s", path)
		}
		for i := range left {
			assertCampaignPlacementJSONShape(t, left[i], right[i], fmt.Sprintf("%s[%d]", path, i))
		}
	default:
		if !bytes.Equal(before, after) {
			t.Fatalf("JSON scalar changed at %s: got %s, want %s", path, after, before)
		}
	}
}

func campaignPlacementRawObject(t *testing.T, raw json.RawMessage) map[string]json.RawMessage {
	t.Helper()
	var object map[string]json.RawMessage
	if err := json.Unmarshal(raw, &object); err != nil {
		t.Fatalf("decode campaign JSON object: %v", err)
	}
	return object
}

func campaignPlacementWithRawValue(t *testing.T, raw json.RawMessage, value string) []byte {
	t.Helper()
	object := campaignPlacementRawObject(t, raw)
	placement, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	object["placement"] = placement
	data, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}
	return data
}
