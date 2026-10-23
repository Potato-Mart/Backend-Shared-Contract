package pkg_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v46/pkg/contracts/marketing/campaign"
	"github.com/Potato-Mart/Backend-Shared-Contract/v46/pkg/contracts/marketing/campaign/campaign_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v46/pkg/contracts/pubsub/envelope"
	events "github.com/Potato-Mart/Backend-Shared-Contract/v46/pkg/contracts/pubsub/pricing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v46/pkg/contracts/pubsub/routing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v46/pkg/contracts/supply/forecasting"
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
	if len(fixture.CanonicalCases) != 4 {
		t.Fatal("four placement cases required")
	}
	for i, original := range fixture.CanonicalCases {
		t.Run(original.Name, func(t *testing.T) {
			c, _ := roundTripCampaignPlacement[campaignPlacementFixtureCase](t, raw.CanonicalCases[i])
			if len(c.Campaign.Placements) != 1 || string(c.Campaign.Placements[0].Position) != c.Name || !c.Campaign.Placements[0].Position.IsValid() {
				t.Fatal("placement identity lost")
			}
			if !reflect.DeepEqual(c.Campaign.Placements, c.Comparable.Placements) || !reflect.DeepEqual(c.Campaign.Audience, c.Comparable.Audience) {
				t.Fatal("forecast input drifted")
			}
			if len(c.Campaign.Audience) != 1 || c.Campaign.Audience[0] != campaign_enums.CampaignCustomerTypeRetail {
				t.Fatal("audience drifted")
			}
			if c.Campaign.CampaignCode != c.Comparable.CampaignCode || c.Campaign.Revision != 13 || c.Campaign.ActivationRevision != 5 {
				t.Fatal("identity or revisions drifted")
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
			if string(c.Placements[0].Position) != value || string(comparable.Placements[0].Position) != value || c.Placements[0].Position.String() != value || comparable.Placements[0].Position.String() != value {
				t.Fatal("ordinary JSON decoding rewrote an unrecognized placement")
			}
			if c.Placements[0].Position.IsValid() || comparable.Placements[0].Position.IsValid() {
				t.Fatal("retired, empty or unknown placement became semantically valid")
			}
		})
	}
}

func TestCampaignPlacementEventCompatibility(t *testing.T) {
	fixture, raw := loadCampaignPlacementFixture(t)
	if len(fixture.EventCases) != 2 || fixture.EventCases[0].Name != "published" || fixture.EventCases[1].Name != "archived" {
		t.Fatal("event fixture must cover active and archived payloads")
	}
	for i, original := range fixture.EventCases {
		t.Run(original.Name, func(t *testing.T) {
			c, _ := roundTripCampaignPlacement[campaignPlacementEventFixtureCase](t, raw.EventCases[i])
			payload, wire := roundTripCampaignPlacement[events.CampaignChangedEvent](t, c.Envelope.Payload)
			if c.Topic != routing.EventTopicStorefrontEvents || c.Topic.String() != "storefront-events" || c.Envelope.EventType != routing.EventTypeCampaignChanged || c.Envelope.EventType.String() != "campaign.changed" || c.Envelope.EventVersion != "v3" {
				t.Fatal("campaign event topic, type or existing service-owned version changed")
			}
			if c.Envelope.EventID != "event-campaign-"+c.Name || c.Envelope.AggregateID != payload.CampaignCode || payload.CampaignCode != "campaign-home_banner" || !c.Envelope.OccurredAt.Equal(payload.ChangedAt) || !payload.RefetchRequired || payload.ScopeMode != "TARGETED" || payload.ScopeRevision != 3 || payload.ActivationRevision != 5 || payload.ContentRevision != int64(13+i) || payload.Status.String() != c.Name || payload.Published != (c.Name == "published") {
				t.Fatal("campaign event lost identity, activation, revision or refetch evidence")
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
	object["placements"] = json.RawMessage(`[{"position":` + string(placement) + `,"platforms":["web"],"priority":0}]`)
	data, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}
	return data
}
