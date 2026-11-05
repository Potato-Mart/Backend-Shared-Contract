package promotion

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/marketing/audience"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/marketing/audience/audience_enums"
)

func TestPromotionPublicationFixtureRoundTrip(t *testing.T) {
	raw, err := os.ReadFile("testdata/promotion_publication.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var decoded PromotionPublication
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	startsAt := time.Date(2026, 11, 6, 0, 0, 0, 0, time.UTC)
	endsAt := time.Date(2026, 11, 9, 0, 0, 0, 0, time.UTC)
	want := PromotionPublication{
		PromotionID:  "prm_weekend",
		Published:    true,
		Revision:     7,
		MarketCode:   "market_au",
		CampaignCode: "CAM-SPRING",
		Audience: &audience.Audience{
			CustomerType: audience_enums.CustomerTypeRetail,
			Platform:     audience_enums.PlatformMobile,
		},
		StartsAt:   &startsAt,
		EndsAt:     &endsAt,
		Timezone:   "Australia/Sydney",
		ObservedAt: time.Date(2026, 11, 5, 1, 2, 3, 0, time.UTC),
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Fatalf("fixture decoded to %+v, want %+v", decoded, want)
	}
	encoded, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var again PromotionPublication
	if err := json.Unmarshal(encoded, &again); err != nil || !reflect.DeepEqual(again, want) {
		t.Fatalf("round trip changed the publication: %v %+v", err, again)
	}
}

func TestPromotionPublicationKeepsRequiredFieldsAndOmitsAbsentOptionalFields(t *testing.T) {
	encoded, err := json.Marshal(PromotionPublication{
		PromotionID: "prm_unpublished",
		ObservedAt:  time.Date(2026, 11, 5, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var fields map[string]any
	if err := json.Unmarshal(encoded, &fields); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	for _, required := range []string{"promotion_id", "published", "revision", "observed_at"} {
		if _, ok := fields[required]; !ok {
			t.Fatalf("required %s must stay on the wire: %s", required, encoded)
		}
	}
	if fields["published"] != false || fields["revision"] != float64(0) {
		t.Fatalf("unpublished zero revision must encode explicitly: %s", encoded)
	}
	for _, absent := range []string{"market_code", "campaign_code", "audience", "starts_at", "ends_at", "timezone"} {
		if _, ok := fields[absent]; ok {
			t.Fatalf("absent %s must be omitted: %s", absent, encoded)
		}
	}
	if len(fields) != 4 {
		t.Fatalf("unexpected fields on the wire: %s", encoded)
	}
}

func TestPromotionPublicationCarriesNoRulesTermsUsageOrCustomerData(t *testing.T) {
	encoded, err := json.Marshal(PromotionPublication{})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	for _, forbidden := range []string{"relations", "terms", "controls", "scope", "usage", "customer_id", "applications", "content"} {
		if strings.Contains(string(encoded), `"`+forbidden) {
			t.Fatalf("publication exposed %s: %s", forbidden, encoded)
		}
	}
	allowed := map[string]bool{
		"promotion_id": true, "published": true, "revision": true, "market_code": true,
		"campaign_code": true, "audience": true, "starts_at": true, "ends_at": true,
		"timezone": true, "observed_at": true,
	}
	publicationType := reflect.TypeOf(PromotionPublication{})
	for i := 0; i < publicationType.NumField(); i++ {
		name, _, _ := strings.Cut(publicationType.Field(i).Tag.Get("json"), ",")
		if !allowed[name] {
			t.Fatalf("PromotionPublication gained unreviewed field %q", name)
		}
	}
	if publicationType.NumField() != len(allowed) {
		t.Fatalf("PromotionPublication has %d fields, want %d", publicationType.NumField(), len(allowed))
	}
}
