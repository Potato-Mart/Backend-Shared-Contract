package payments_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/payments"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/shared"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/payment"
)

func TestTerminalTransactionJSONRoundTripWithHistory(t *testing.T) {
	occurredAt := time.Date(2026, 6, 17, 10, 20, 0, 0, time.UTC)
	tx := payments.TerminalTransaction{
		ID:         "ttx_1",
		TerminalID: "term_1",
		PaymentID:  "pay_1",
		ProviderReference: &payments.PaymentReference{
			Mx51: &payments.Mx51PaymentReference{
				TransactionID: "mx_tx_1",
			},
		},
		ProviderDetails: &payments.TerminalProviderDetails{
			TerminalID: "provider_term_1",
		},
		OperationContext: &payments.ProviderOperationContext{
			RequestID:         "provider_req_1",
			MerchantReference: "merchant_ref_1",
			IdempotencyKey:    "idem_1",
		},
		Type:            paymentenum.TerminalTxTypePurchase,
		Requested:       payments.Amounts{Currency: "AUD", PurchaseMinor: 10000},
		Status:          paymentenum.TerminalTxStatusFinalised,
		FinancialStatus: paymentenum.TerminalTxFinancialStatusApproved,
		Result:          payments.Amounts{Currency: "AUD", PurchaseMinor: 10000, AuthorizedMinor: 10000},
		ProviderResult:  "approved",
		ProviderData: common.Metadata{
			"terminal_batch": "batch_1",
		},
		Payloads: &payments.ProviderPayloads{
			Request:             json.RawMessage(`{"sale":"request"}`),
			Response:            json.RawMessage(`{"sale":"response"}`),
			DisplayNotification: json.RawMessage(`{"display":"approved"}`),
		},
		History: []shared.HistoryEntry{
			{
				OccurredAt: occurredAt,
				Type:       "provider_update",
				Changes: []shared.HistoryChange{
					{Field: "status", FromValue: "pending", ToValue: "finalised"},
				},
			},
		},
	}

	payload, err := json.Marshal(tx)
	if err != nil {
		t.Fatalf("marshal terminal transaction: %v", err)
	}

	var decoded payments.TerminalTransaction
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
	if decoded.Result.AuthorizedMinor != 10000 || decoded.FinancialStatus != paymentenum.TerminalTxFinancialStatusApproved {
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
	terminal := payments.Terminal{
		ID:       "term_1",
		TenantID: "tenant_1",
		Provider: paymentenum.TerminalProviderMx51,
		Status:   paymentenum.TerminalStatusActive,
		ProviderDetails: &payments.TerminalProviderDetails{
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

	info := payments.TerminalConnectionInfo{
		TerminalID:      "term_1",
		Provider:        paymentenum.TerminalProviderMx51,
		ProviderDetails: &payments.TerminalProviderDetails{TerminalID: "provider_term_1"},
		Status:          paymentenum.TerminalStatusActive,
		Connected:       true,
		ProviderStatus:  "online",
		CheckedAt:       time.Date(2026, 6, 18, 6, 0, 0, 0, time.UTC),
	}
	if payload, err = json.Marshal(info); err != nil {
		t.Fatalf("marshal terminal connection info: %v", err)
	}
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal terminal connection info JSON: %v", err)
	}
	for _, key := range []string{"terminal_id", "provider", "provider_details", "status", "connected", "checked_at"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("TerminalConnectionInfo JSON missing %q: %s", key, payload)
		}
	}
}

func TestSettlementJSONGroupsProviderSupportFields(t *testing.T) {
	settlement := payments.Settlement{
		ID:                   "settlement_1",
		TerminalID:           "term_1",
		ProviderSettlementID: "provider_settlement_1",
		ProviderDetails:      &payments.TerminalProviderDetails{MerchantID: "merchant_1", TerminalID: "provider_term_1"},
		OperationContext:     &payments.ProviderOperationContext{RequestID: "provider_req_1", IdempotencyKey: "idem_1"},
		Type:                 paymentenum.SettlementTypeSettlement,
		Status:               paymentenum.TerminalTxStatusFinalised,
		FinancialStatus:      paymentenum.TerminalTxFinancialStatusApproved,
		Totals:               payments.SettlementTotals{Currency: "AUD", TotalMinor: 10000},
		Payloads:             &payments.ProviderPayloads{Request: json.RawMessage(`{"settle":"request"}`)},
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
