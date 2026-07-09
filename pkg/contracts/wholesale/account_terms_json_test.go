package wholesale_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/wholesale"
)

func TestWholesaleAccountTermsJSONShape(t *testing.T) {
	resp := wholesale.WholesaleAccountTerms{
		OrganisationCode:     "org_123",
		OrganisationAccessID: "access_123",
		UserID:               "user_123",
		BusinessName:         "Potato Mart Cafe",
		TradingName:          "Cafe PM",
		PaymentTermsLabel:    "Net 14",
		Terms: wholesale.WholesaleTerms{
			AccountCheckoutEnabled: true,
			DueDays:                14,
			PurchaseOrderRequired:  true,
			PaymentInstructions:    "Pay by due date.",
			CreditLimit:            &common.Money{AmountMinor: 500000, Currency: "AUD"},
		},
		AllowedCheckoutMethods: []string{"card", "bank_transfer"},
		InvoicePaymentMethods:  []string{"card"},
	}

	raw, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	for _, key := range []string{
		"organisation_code",
		"organisation_access_id",
		"user_id",
		"business_name",
		"trading_name",
		"payment_terms_label",
		"terms",
		"allowed_checkout_methods",
		"invoice_payment_methods",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("missing json key %q in %s", key, raw)
		}
	}
	if _, ok := got["customer_id"]; ok {
		t.Fatalf("WholesaleAccountTerms JSON should not include separate customer_id: %s", raw)
	}

	terms, ok := got["terms"].(map[string]any)
	if !ok {
		t.Fatalf("terms is not an object in %s", raw)
	}
	for _, key := range []string{
		"account_checkout_enabled",
		"due_days",
		"purchase_order_required",
		"payment_instructions",
		"credit_limit",
	} {
		if _, ok := terms[key]; !ok {
			t.Fatalf("terms missing json key %q in %s", key, raw)
		}
	}
}

func TestInvoiceCardPaymentJSONShape(t *testing.T) {
	req := wholesale.InvoiceCardPaymentCommand{IdempotencyKey: "idem_123"}
	reqRaw, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}
	if !json.Valid(reqRaw) {
		t.Fatalf("invalid request JSON: %s", reqRaw)
	}

	resp := wholesale.InvoiceCardPaymentResult{
		InvoiceID:     "inv_123",
		InvoiceNumber: "INV-123",
		OrderNumber:   "MAMA260703ABC123",
		Total:         common.Money{AmountMinor: 12500, Currency: "AUD"},
		Payment: wholesale.InvoiceCardPaymentSession{
			Provider:          "stripe",
			ProviderSessionID: "pi_123",
			ClientKey:         "pi_123_secret_456",
			SessionData:       `{"status":"requires_payment_method"}`,
		},
	}

	raw, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal response: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}

	for _, key := range []string{
		"invoice_id",
		"invoice_number",
		"order_number",
		"total",
		"payment",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("missing json key %q in %s", key, raw)
		}
	}
}
