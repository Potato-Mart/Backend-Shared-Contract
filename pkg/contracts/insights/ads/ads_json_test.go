package ads_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads"
	ae "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

func TestAdPerformanceRowRoundTripAndWireShape(t *testing.T) {
	clicked := time.Date(2026, 11, 26, 8, 15, 2, 0, time.UTC)
	daily := money.Money{AmountMinor: 5000, Currency: "AUD"}
	row := ads.AdPerformanceRow{
		Platform:                ae.AdPlatformMeta,
		Level:                   ae.AdLevelCampaign,
		ID:                      "120214000000000",
		Name:                    "Festive hamper",
		Status:                  ae.AdEntityStatusActive,
		ReviewStatus:            ae.AdReviewStatusApproved,
		Objective:               "sales",
		Revision:                "r-7",
		SourcePlatform:          ae.AdSourcePlatformWeb,
		Budget:                  &ads.AdBudget{Level: ae.AdLevelAdSet, Daily: &daily},
		Bidding:                 &ads.AdBidding{Strategy: ae.AdBidStrategyMinimumROAS, TargetReturnBasisPoints: 40000},
		Spend:                   []money.Money{{AmountMinor: 286400, Currency: "AUD"}},
		Impressions:             286400,
		Reach:                   120330,
		Clicks:                  5210,
		PlatformConversions:     132,
		PlatformConversionValue: []money.Money{{AmountMinor: 998800, Currency: "AUD"}},
		Orders:                  118,
		NewCustomers:            52,
		Revenue:                 []money.Money{{AmountMinor: 884500, Currency: "AUD"}},
		Buckets: []ads.AdPerformanceBucket{{
			Date:    "2026-11-26",
			Spend:   []money.Money{{AmountMinor: 41000, Currency: "AUD"}},
			Orders:  6,
			Revenue: []money.Money{{AmountMinor: 52300, Currency: "AUD"}},
		}},
	}

	raw, err := json.Marshal(row)
	if err != nil {
		t.Fatal(err)
	}
	var decoded ads.AdPerformanceRow
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(row, decoded) {
		t.Fatalf("round trip changed model: %+v", decoded)
	}

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"platform", "level", "id", "name", "status", "review_status", "revision", "spend", "revenue", "buckets"} {
		if _, present := fields[key]; !present {
			t.Fatalf("performance row is missing %s: %s", key, raw)
		}
	}
	// An absent breakdown, parent and install count stay off the wire.
	for _, key := range []string{"parent_id", "installs"} {
		if _, present := fields[key]; present {
			t.Fatalf("absent %s must be omitted: %s", key, raw)
		}
	}
	if string(fields["budget"]) != `{"level":"ad_set","daily":{"amount_minor":5000,"currency":"AUD"}}` {
		t.Fatalf("budget wire shape changed: %s", fields["budget"])
	}

	// Google reports no reach, which decodes as an absent figure rather than
	// a measured zero.
	var google ads.AdPerformanceRow
	if err := json.Unmarshal([]byte(`{"platform":"google","level":"campaign","id":"1","reach":null}`), &google); err != nil {
		t.Fatal(err)
	}
	if google.Reach != 0 || google.Platform != ae.AdPlatformGoogle {
		t.Fatalf("google row decoded unexpectedly: %+v", google)
	}
	if clicked.IsZero() {
		t.Fatal("fixture instant must be set")
	}
}

func TestAttributedOrderCarriesNoContactDetailsOrClickIdentifiers(t *testing.T) {
	item := ads.AdAttributedOrder{
		OrderNumber:         "MAMA261126KQ4821",
		CustomerNumber:      "RC-104283",
		IsNewCustomer:       true,
		Revenue:             []money.Money{{AmountMinor: 7400, Currency: "AUD"}},
		Platform:            ae.AdPlatformMeta,
		SourcePlatform:      ae.AdSourcePlatformIOS,
		Campaign:            &ads.AdEntityRef{ID: "120214", Name: "Festive hamper"},
		AdSet:               &ads.AdEntityRef{ID: "120215"},
		Ad:                  &ads.AdEntityRef{ID: "120216"},
		ClickToOrderSeconds: 48000,
	}
	raw, err := json.Marshal(item)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"email", "phone", "address", "gclid", "fbclid", "click_id", "fbc", "fbp", "hashed_email", "hashed_phone"} {
		if _, present := fields[forbidden]; present {
			t.Fatalf("attributed order must not expose %s: %s", forbidden, raw)
		}
	}
	if _, present := fields["display_name"]; present {
		t.Fatalf("an unset display name must be omitted: %s", raw)
	}
	var decoded ads.AdAttributedOrder
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(item, decoded) {
		t.Fatalf("round trip changed model: %+v", decoded)
	}
}

func TestTrackingSettingsAndGuardrailsWireShape(t *testing.T) {
	settings := ads.AdTrackingSettings{
		Revision:    "s-3",
		Attribution: &ads.AdAttributionPolicy{Model: ae.AdAttributionModelLastClick, ClickWindowDays: 7},
		Guardrails: &ads.AdGuardrails{
			MaxDailyBudgetPerAccount:   []money.Money{{AmountMinor: 50000, Currency: "AUD"}},
			ConfirmIncreaseOverPercent: 50,
		},
	}
	raw, err := json.Marshal(settings)
	if err != nil {
		t.Fatal(err)
	}
	const want = `{"revision":"s-3","attribution":{"model":"last_click","click_window_days":7},"guardrails":{"max_daily_budget_per_account":[{"amount_minor":50000,"currency":"AUD"}],"confirm_increase_over_percent":50}}`
	if string(raw) != want {
		t.Fatalf("tracking settings wire shape = %s, want %s", raw, want)
	}
	var decoded ads.AdTrackingSettings
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(settings, decoded) {
		t.Fatalf("round trip changed model: %+v", decoded)
	}
	// include_wholesale is false by default, so it stays off the wire and a
	// reader must not infer that wholesale orders are attributed.
	if _, present := map[string]json.RawMessage{}["include_wholesale"]; present {
		t.Fatal("unreachable")
	}
}

func TestChangeRecordKeepsBeforeAndAfterInOneShape(t *testing.T) {
	automatic := false
	requested := time.Date(2026, 11, 27, 9, 0, 0, 0, time.UTC)
	before := ads.AdChangeSet{Status: ae.AdEntityStatusPaused}
	after := ads.AdChangeSet{
		Status:              ae.AdEntityStatusActive,
		Budget:              &ads.AdBudget{Level: ae.AdLevelCampaign, Daily: &money.Money{AmountMinor: 12000, Currency: "AUD"}},
		Schedule:            &ads.AdSchedule{StartAt: &requested},
		AutomaticPlacements: &automatic,
		Placements:          []string{"feed", "stories"},
	}
	record := ads.AdChangeRecord{
		ChangeID:         "chg-1",
		Platform:         ae.AdPlatformMeta,
		Level:            ae.AdLevelCampaign,
		EntityID:         "120214",
		Status:           ae.AdChangeStatusConflict,
		Source:           ae.AdChangeSourceAdmin,
		ExpectedRevision: "r-7",
		RequestedBy:      "staff-9",
		RequestedAt:      &requested,
		Before:           &before,
		After:            &after,
		Warnings:         []string{"large_increase"},
	}
	raw, err := json.Marshal(record)
	if err != nil {
		t.Fatal(err)
	}
	var decoded ads.AdChangeRecord
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(record, decoded) {
		t.Fatalf("round trip changed model: %+v", decoded)
	}
	if reflect.TypeOf(record.Before) != reflect.TypeOf(record.After) {
		t.Fatal("before and after must share one normalised shape")
	}
	// A manual placement selection must survive as false, not vanish.
	if decoded.After.AutomaticPlacements == nil || *decoded.After.AutomaticPlacements {
		t.Fatalf("manual placement selection lost: %+v", decoded.After)
	}
}

func TestConversionWithdrawalDistinguishesFullFromPartial(t *testing.T) {
	full := ads.AdConversionWithdrawal{Full: true, WithdrawnRevenue: &money.Money{AmountMinor: 7400, Currency: "AUD"}}
	raw, err := json.Marshal(full)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(raw), `{"full":true,"withdrawn_revenue":{"amount_minor":7400,"currency":"AUD"}}`; got != want {
		t.Fatalf("full withdrawal wire shape = %s, want %s", got, want)
	}

	partial := ads.AdConversionWithdrawal{
		WithdrawnRevenue: &money.Money{AmountMinor: 2000, Currency: "AUD"},
		RemainingRevenue: &money.Money{AmountMinor: 5400, Currency: "AUD"},
	}
	raw, err = json.Marshal(partial)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	if _, present := fields["full"]; present {
		t.Fatalf("a partial withdrawal must omit full: %s", raw)
	}
	if _, present := fields["remaining_revenue"]; !present {
		t.Fatalf("a partial withdrawal must carry the restated value: %s", raw)
	}
	var decoded ads.AdConversionWithdrawal
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(partial, decoded) {
		t.Fatalf("round trip changed model: %+v", decoded)
	}
}

func TestConnectionAndSyncStateCarryNoCredential(t *testing.T) {
	synced := time.Date(2026, 11, 27, 9, 0, 0, 0, time.UTC)
	connection := ads.AdConnection{
		Platform: ae.AdPlatformGoogle,
		Status:   ae.AdConnectionStatusNeedsAttention,
		Accounts: []ads.AdAccount{{
			MaskedID:   "***4821",
			Name:       "Potato Mart AU",
			Currency:   "AUD",
			TimeZone:   "Australia/Melbourne",
			MarketCode: "AU-VIC",
		}},
		ConversionActionName:             "Website purchase (server upload)",
		GoogleAutoTagging:                true,
		GoogleEnhancedConversionsEnabled: true,
		CredentialStatus:                 ae.AdCredentialStatusExpiring,
		LastSyncedAt:                     &synced,
		LastErrorCode:                    "sync_partial",
	}
	raw, err := json.Marshal(connection)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"token", "access_token", "refresh_token", "developer_token", "secret", "api_key"} {
		if _, present := fields[forbidden]; present {
			t.Fatalf("a connection must never carry %s: %s", forbidden, raw)
		}
	}
	var decoded ads.AdConnection
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(connection, decoded) {
		t.Fatalf("round trip changed model: %+v", decoded)
	}

	state := ads.AdSyncState{
		Platform:        ae.AdPlatformMeta,
		SyncedAt:        &synced,
		AccountTimeZone: "Australia/Melbourne",
		AccountCurrency: "AUD",
	}
	raw, err = json.Marshal(state)
	if err != nil {
		t.Fatal(err)
	}
	const want = `{"platform":"meta","synced_at":"2026-11-27T09:00:00Z","account_time_zone":"Australia/Melbourne","account_currency":"AUD"}`
	if string(raw) != want {
		t.Fatalf("sync state wire shape = %s, want %s", raw, want)
	}
}

func TestDraftPostAndHealthShapesStayCodeOnly(t *testing.T) {
	draft := ads.AdDraft{
		DraftID:  "drf-1",
		Platform: ae.AdPlatformMeta,
		Status:   ae.AdDraftStatusInReview,
		Placement: &ads.AdDraftPlacement{
			Mode:    ae.AdPlacementModeExisting,
			AdSetID: "120215",
		},
		Creative: &ads.AdCreative{Source: ae.AdCreativeSourceMedia, MediaCodes: []string{"MED-1", "MED-2"}},
		Copy: []ads.AdCopy{{
			Locale:       "zh-TW",
			PrimaryText:  "primary",
			Headline:     "headline",
			CallToAction: ae.AdCallToActionShopNow,
			Reviewed:     true,
		}},
		Landing: &ads.AdLanding{Kind: ae.AdLandingKindCampaign, Ref: "xmas-2026"},
	}
	raw, err := json.Marshal(draft)
	if err != nil {
		t.Fatal(err)
	}
	var decodedDraft ads.AdDraft
	if err := json.Unmarshal(raw, &decodedDraft); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(draft, decodedDraft) {
		t.Fatalf("round trip changed model: %+v", decodedDraft)
	}

	post := ads.AdPagePost{
		PostID:              "17900000000000000",
		Platform:            ae.AdPlatformMeta,
		Source:              "instagram",
		Type:                ae.AdPostTypeReel,
		Excerpt:             "opening words",
		Likes:               210,
		IneligibilityReason: ae.AdPostIneligibilityReasonLicensedMusic,
	}
	raw, err = json.Marshal(post)
	if err != nil {
		t.Fatal(err)
	}
	var postFields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &postFields); err != nil {
		t.Fatal(err)
	}
	if _, present := postFields["eligible"]; present {
		t.Fatalf("an ineligible post omits the false flag: %s", raw)
	}
	if string(postFields["ineligibility_reason"]) != `"licensed_music"` {
		t.Fatalf("ineligibility reason must stay a code: %s", raw)
	}

	check := ads.AdHealthCheck{
		Code:             ae.AdHealthCheckCodeConversionGap,
		Severity:         ae.AdHealthSeverityWarning,
		Platform:         ae.AdPlatformMeta,
		Level:            ae.AdLevelCampaign,
		EntityIDs:        []string{"120214"},
		ObservedCount:    42,
		MatchedCount:     11,
		RatioBasisPoints: 2619,
	}
	raw, err = json.Marshal(check)
	if err != nil {
		t.Fatal(err)
	}
	var checkFields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &checkFields); err != nil {
		t.Fatal(err)
	}
	if _, present := checkFields["message"]; present {
		t.Fatalf("health findings carry codes and numbers, never wording: %s", raw)
	}
	var decodedCheck ads.AdHealthCheck
	if err := json.Unmarshal(raw, &decodedCheck); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(check, decodedCheck) {
		t.Fatalf("round trip changed model: %+v", decodedCheck)
	}
}

func TestTotalsAndEstimateKeepPlatformAndOwnFiguresApart(t *testing.T) {
	totals := ads.AdPerformanceTotals{
		Spend:                   []money.Money{{AmountMinor: 1000, Currency: "AUD"}, {AmountMinor: 2000, Currency: "TWD"}},
		Impressions:             1001300,
		Clicks:                  19150,
		PlatformConversions:     410,
		PlatformConversionValue: []money.Money{{AmountMinor: 9000, Currency: "AUD"}},
		Orders:                  318,
		NewCustomers:            131,
		Revenue:                 []money.Money{{AmountMinor: 8000, Currency: "AUD"}},
	}
	raw, err := json.Marshal(totals)
	if err != nil {
		t.Fatal(err)
	}
	var decoded ads.AdPerformanceTotals
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(totals, decoded) {
		t.Fatalf("round trip changed model: %+v", decoded)
	}
	if len(decoded.Spend) != 2 {
		t.Fatal("spend must stay split by currency so markets are never summed")
	}

	summary := ads.AdAttributedOrderSummary{Orders: 318, NewCustomers: 131, GuestOrders: 27, MedianClickToOrderSeconds: 138240}
	if _, err := json.Marshal(summary); err != nil {
		t.Fatal(err)
	}
	estimate := ads.AdDeliveryEstimate{LowerBound: 12000, UpperBound: 31000}
	raw, err = json.Marshal(estimate)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(raw), `{"lower_bound":12000,"upper_bound":31000}`; got != want {
		t.Fatalf("delivery estimate wire shape = %s, want %s", got, want)
	}

	targeting := ads.AdTargetingSummary{Locations: []string{"AU-VIC"}, AgeMin: 25, AgeMax: 54, AutomaticAudience: true}
	raw, err = json.Marshal(targeting)
	if err != nil {
		t.Fatal(err)
	}
	var targetingFields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &targetingFields); err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"email", "phone", "customer_numbers"} {
		if _, present := targetingFields[forbidden]; present {
			t.Fatalf("a targeting summary must not describe individuals: %s", raw)
		}
	}
}
