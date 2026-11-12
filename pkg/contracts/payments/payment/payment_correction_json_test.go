package payment_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	sales "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

func TestPaymentCorrectionSummaryRoundTrips(t *testing.T) {
	const fixture = `{"id":"pay_1","order_number":"MAMA261112ABC123","amount":{"amount_minor":4500,"currency":"AUD"},"method":"cash","status":"completed","correction":{"count":2,"last_corrected_at":"2026-11-12T03:04:05Z","last_corrected_by":{"actor_id":"staff_1","actor_email":"staff@example.test","actor_domain":"workforce","actor_role":"countryAdmin"},"original_amount":{"amount_minor":5000,"currency":"AUD"},"original_method":"card","warnings":["gateway_amount_unchanged"]},"created_at":"2026-11-10T00:00:00Z","updated_at":"2026-11-12T03:04:05Z"}`
	var payment sales.Payment
	if err := json.Unmarshal([]byte(fixture), &payment); err != nil {
		t.Fatal(err)
	}
	correction := payment.Correction
	if correction == nil || correction.Count != 2 || correction.LastCorrectedAt.IsZero() || correction.LastCorrectedBy.ActorID != "staff_1" {
		t.Fatalf("correction identity lost: %+v", correction)
	}
	if correction.OriginalAmount.AmountMinor != 5000 || correction.OriginalMethod != payment_enums.PaymentMethodCard {
		t.Fatalf("original values lost: %+v", correction)
	}
	if len(correction.Warnings) != 1 || correction.Warnings[0] != payment_enums.PaymentCorrectionWarningGatewayAmountUnchanged || !correction.Warnings[0].IsValid() {
		t.Fatalf("warnings lost: %+v", correction.Warnings)
	}
	encoded, err := json.Marshal(payment)
	if err != nil {
		t.Fatal(err)
	}
	var got, want map[string]any
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(fixture), &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("corrected payment JSON changed:\n got: %s\nwant: %s", encoded, fixture)
	}
}

func TestUncorrectedPaymentOmitsCorrection(t *testing.T) {
	encoded, err := json.Marshal(sales.Payment{ID: "pay_2", OrderNumber: "MAMA261112ABC123", Method: payment_enums.PaymentMethodCash})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), `"correction"`) {
		t.Fatalf("uncorrected payment JSON = %s, want correction omitted", encoded)
	}
}
