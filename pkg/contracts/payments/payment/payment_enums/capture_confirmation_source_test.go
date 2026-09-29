package payment_enums

import "testing"

func TestCaptureConfirmationSource(t *testing.T) {
	for _, source := range []CaptureConfirmationSource{CaptureConfirmationSourceProviderEvent, CaptureConfirmationSourceProviderStatusObservation} {
		if !source.IsValid() || source.String() != string(source) {
			t.Fatal(source)
		}
	}
	for _, source := range []CaptureConfirmationSource{"", "charge_created", "intent_created", "processing_time"} {
		if source.IsValid() {
			t.Fatalf("unauthenticated or unrelated timing source accepted: %s", source)
		}
	}
}
