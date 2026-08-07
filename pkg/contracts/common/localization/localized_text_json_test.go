package localization

import (
	"encoding/json"
	"testing"
)

func TestLocalizedTextJSONRoundTrip(t *testing.T) {
	want := LocalizedText{Language: "zh-TW", Text: "產地直送"}

	body, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal localized text: %v", err)
	}
	if string(body) != `{"language":"zh-TW","text":"產地直送"}` {
		t.Fatalf("LocalizedText JSON = %s, want language and text", body)
	}

	var got LocalizedText
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal localized text: %v", err)
	}
	if got != want {
		t.Fatalf("LocalizedText round-trip = %+v, want %+v", got, want)
	}
}
