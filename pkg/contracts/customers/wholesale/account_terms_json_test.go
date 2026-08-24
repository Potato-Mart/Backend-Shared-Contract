package wholesale_test

import (
	"encoding/json"

	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/customers/wholesale"
)

func TestWholesaleTermsJSONShape(t *testing.T) {
	terms := wholesale.WholesaleTerms{
		AccountCheckoutEnabled: true,
		DueDays:                14,
		PurchaseOrderRequired:  true,
		PaymentInstructions:    "Pay by due date.",
		CreditLimit:            &money.Money{AmountMinor: 500000, Currency: "AUD"},
	}
	raw, err := json.Marshal(terms)
	if err != nil {
		t.Fatalf("marshal terms: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal terms: %v", err)
	}
	for _, key := range []string{"account_checkout_enabled", "due_days", "purchase_order_required", "payment_instructions", "credit_limit"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("terms missing json key %q in %s", key, raw)
		}
	}
}
