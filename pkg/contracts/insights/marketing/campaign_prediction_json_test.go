package marketing_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/insights/marketing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/insights/marketing/marketing_enums"
)

func TestInsightsMarketingContainsPredictionNotOperationalCampaign(t *testing.T) {
	payload, err := json.Marshal(marketing.CampaignPrediction{PredictionKey: "pred-1", CampaignCode: "spring-2026", Revision: 1, Status: marketing_enums.CampaignPredictionStatusReady, AlgorithmVersion: "v1", PredictedAt: time.Date(2026, 8, 24, 0, 0, 0, 0, time.UTC)})
	if err != nil {
		t.Fatalf("marshal prediction: %v", err)
	}
	if !strings.Contains(string(payload), `"campaign_code":"spring-2026"`) || strings.Contains(string(payload), "recipient") || strings.Contains(string(payload), "channel") {
		t.Fatalf("unexpected insights prediction: %s", payload)
	}
}
