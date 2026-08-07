package terminal_test

import (
	"encoding/json"
	"testing"
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/security"
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	payment "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/payments/payment"
	settlement "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/payments/settlement"
	terminal "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/payments/terminal"
)

func TestTerminalTransactionJSONRoundTripWithHistory(t *testing.T) {
	occurredAt := time.Date(2026, 6, 17, 10, 20, 0, 0, time.UTC)
	tx := terminal.TerminalTransaction{
		ID:         "ttx_1",
		TerminalID: "term_1",
		PaymentID:  "pay_1",
		ProviderReference: &payment.PaymentReference{
			Mx51: &payment.Mx51PaymentReference{
				TransactionID: "mx_tx_1",
			},
		},
		ProviderDetails: &terminal.TerminalProviderDetails{
			TerminalID: "provider_term_1",
		},
		OperationContext: &terminal.ProviderOperationContext{
			RequestID:         "provider_req_1",
			MerchantReference: "merchant_ref_1",
			IdempotencyKey:    "idem_1",
		},
		Type:            terminal.TerminalTxTypePurchase,
		Requested:       terminal.Amounts{Currency: "AUD", PurchaseMinor: 10000},
		Status:          terminal.TerminalTxStatusFinalised,
		FinancialStatus: terminal.TerminalTxFinancialStatusApproved,
		Result:          terminal.Amounts{Currency: "AUD", PurchaseMinor: 10000, AuthorizedMinor: 10000},
		ProviderResult:  "approved",
		ProviderData: common.Metadata{
			"terminal_batch": "batch_1",
		},
		Payloads: &terminal.ProviderPayloads{
			Request:             json.RawMessage(`{"sale":"request"}`),
			Response:            json.RawMessage(`{"sale":"response"}`),
			DisplayNotification: json.RawMessage(`{"display":"approved"}`),
		},
		History: []security.HistoryEntry{
			{
				OccurredAt: occurredAt,
				Type:       "provider_update",
				Changes: []security.HistoryChange{
					{Field: "status", FromValue: "pending", ToValue: "finalised"},
				},
			},
		},
	}

	payload, err := json.Marshal(tx)
	if err != nil {
		t.Fatalf("marshal terminal transaction: %v", err)
	}

	var decoded terminal.TerminalTransaction
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal terminal transaction: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal terminal transaction JSON: %v", err)
	}

	for _, key := range []string{"id", "terminal_id", "type", "status", "financial_status", "requested", "result"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("TerminalTransaction JSON missing top-level %q: %s", key, payload)
		}
	}
	for _, key := range []string{"provider_details", "operation_context", "provider_payloads"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("TerminalTransaction JSON missing nested provider support %q: %s", key, payload)
		}
	}
	for _, key := range []string{"provider_request_id", "provider_terminal_id", "merchant_reference", "idempotency_key", "provider_request", "provider_response", "display_notification"} {
		if _, ok := got[key]; ok {
			t.Fatalf("TerminalTransaction JSON should not include flat provider key %q: %s", key, payload)
		}
	}

	if decoded.ProviderReference == nil || decoded.ProviderReference.Mx51 == nil || decoded.ProviderReference.Mx51.TransactionID != "mx_tx_1" {
		t.Fatalf("provider reference did not round-trip: %+v", decoded.ProviderReference)
	}
	if decoded.ProviderDetails == nil || decoded.ProviderDetails.TerminalID != "provider_term_1" {
		t.Fatalf("provider details did not round-trip: %+v", decoded.ProviderDetails)
	}
	if decoded.OperationContext == nil || decoded.OperationContext.IdempotencyKey != "idem_1" {
		t.Fatalf("operation context did not round-trip: %+v", decoded.OperationContext)
	}
	if decoded.Payloads == nil || len(decoded.Payloads.Request) == 0 || len(decoded.Payloads.DisplayNotification) == 0 {
		t.Fatalf("provider payloads did not round-trip: %+v", decoded.Payloads)
	}
	if decoded.Result.AuthorizedMinor != 10000 || decoded.FinancialStatus != terminal.TerminalTxFinancialStatusApproved {
		t.Fatalf("result/financial status did not round-trip: %+v", decoded)
	}
	if decoded.ProviderData["terminal_batch"] != "batch_1" {
		t.Fatalf("provider data did not round-trip: %+v", decoded.ProviderData)
	}
	if len(decoded.History) != 1 || decoded.History[0].Changes[0].ToValue != "finalised" {
		t.Fatalf("history did not round-trip: %+v", decoded.History)
	}
}

func TestTerminalProviderDetailsJSONShapes(t *testing.T) {
	terminal := terminal.Terminal{
		ID:       "term_1",
		TenantID: "tenant_1",
		Provider: terminal.TerminalProviderMx51,
		Status:   terminal.TerminalStatusActive,
		ProviderDetails: &terminal.TerminalProviderDetails{
			MerchantID: "merchant_1",
			StoreID:    "store_1",
			TerminalID: "provider_term_1",
			DeviceID:   "device_1",
			Nickname:   "Front Counter",
			BaseURL:    "https://terminal.example.com",
		},
	}

	payload, err := json.Marshal(terminal)
	if err != nil {
		t.Fatalf("marshal terminal: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal terminal JSON: %v", err)
	}
	for _, key := range []string{"id", "tenant_id", "provider", "status"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("Terminal JSON missing top-level %q: %s", key, payload)
		}
	}
	if _, ok := got["provider_details"]; !ok {
		t.Fatalf("Terminal JSON missing provider_details: %s", payload)
	}
	for _, key := range []string{"provider_merchant_id", "provider_store_id", "provider_terminal_id", "provider_device_id", "terminal_nickname", "provider_base_url"} {
		if _, ok := got[key]; ok {
			t.Fatalf("Terminal JSON should not include flat provider key %q: %s", key, payload)
		}
	}

}

func TestSettlementJSONGroupsProviderSupportFields(t *testing.T) {
	settlement := settlement.Settlement{
		ID:                   "settlement_1",
		TerminalID:           "term_1",
		ProviderSettlementID: "provider_settlement_1",
		ProviderDetails:      &terminal.TerminalProviderDetails{MerchantID: "merchant_1", TerminalID: "provider_term_1"},
		OperationContext:     &terminal.ProviderOperationContext{RequestID: "provider_req_1", IdempotencyKey: "idem_1"},
		Type:                 settlement.SettlementTypeSettlement,
		Status:               terminal.TerminalTxStatusFinalised,
		FinancialStatus:      terminal.TerminalTxFinancialStatusApproved,
		Totals:               settlement.SettlementTotals{Currency: "AUD", TotalMinor: 10000},
		Payloads:             &terminal.ProviderPayloads{Request: json.RawMessage(`{"settle":"request"}`)},
	}

	payload, err := json.Marshal(settlement)
	if err != nil {
		t.Fatalf("marshal settlement: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal settlement JSON: %v", err)
	}
	for _, key := range []string{"id", "terminal_id", "provider_settlement_id", "type", "status", "totals"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("Settlement JSON missing top-level %q: %s", key, payload)
		}
	}
	for _, key := range []string{"provider_details", "operation_context", "provider_payloads"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("Settlement JSON missing nested provider support %q: %s", key, payload)
		}
	}
	for _, key := range []string{"provider_merchant_id", "provider_terminal_id", "provider_request_id", "idempotency_key", "provider_request"} {
		if _, ok := got[key]; ok {
			t.Fatalf("Settlement JSON should not include flat provider key %q: %s", key, payload)
		}
	}

}
