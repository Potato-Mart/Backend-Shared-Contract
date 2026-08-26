package terminal_test

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"

	payment "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/payments/payment"
	settlement "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/payments/settlement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/payments/settlement/settlement_enums"
	terminal "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/payments/terminal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/payments/terminal/terminal_enums"
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
		Type:            terminal_enums.TerminalTxTypePurchase,
		Requested:       terminal.Amounts{Currency: "AUD", PurchaseMinor: 10000},
		Status:          terminal_enums.TerminalTxStatusFinalised,
		FinancialStatus: terminal_enums.TerminalTxFinancialStatusApproved,
		Result:          terminal.Amounts{Currency: "AUD", PurchaseMinor: 10000, AuthorizedMinor: 10000},
		ProviderResult:  "approved",
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
	for _, key := range []string{"provider_details"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("TerminalTransaction JSON missing nested provider support %q: %s", key, payload)
		}
	}
	for _, key := range []string{
		"operation_context", "provider_payloads", "provider_data",
		"receipt_options",
		"provider_request_id", "provider_terminal_id", "merchant_reference", "idempotency_key",
		"provider_request", "provider_response", "display_notification",
	} {
		if _, ok := got[key]; ok {
			t.Fatalf("TerminalTransaction JSON should not include provider transport or opaque data key %q: %s", key, payload)
		}
	}

	if decoded.ProviderReference == nil || decoded.ProviderReference.Mx51 == nil || decoded.ProviderReference.Mx51.TransactionID != "mx_tx_1" {
		t.Fatalf("provider reference did not round-trip: %+v", decoded.ProviderReference)
	}
	if decoded.ProviderDetails == nil || decoded.ProviderDetails.TerminalID != "provider_term_1" {
		t.Fatalf("provider details did not round-trip: %+v", decoded.ProviderDetails)
	}
	if decoded.Result.AuthorizedMinor != 10000 || decoded.FinancialStatus != terminal_enums.TerminalTxFinancialStatusApproved {
		t.Fatalf("result/financial status did not round-trip: %+v", decoded)
	}
	if len(decoded.History) != 1 || decoded.History[0].Changes[0].ToValue != "finalised" {
		t.Fatalf("history did not round-trip: %+v", decoded.History)
	}
}

func TestTerminalProviderDetailsJSONShapes(t *testing.T) {
	terminal := terminal.Terminal{
		ID:       "term_1",
		TenantID: "tenant_1",
		Provider: terminal_enums.TerminalProviderMx51,
		Status:   terminal_enums.TerminalStatusActive,
		ProviderDetails: &terminal.TerminalProviderDetails{
			MerchantID: "merchant_1",
			StoreID:    "store_1",
			TerminalID: "provider_term_1",
			DeviceID:   "device_1",
			Nickname:   "Front Counter",
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
	providerDetails, ok := got["provider_details"].(map[string]any)
	if !ok {
		t.Fatalf("Terminal provider_details has unexpected shape: %#v", got["provider_details"])
	}
	if _, ok := providerDetails["base_url"]; ok {
		t.Fatalf("Terminal provider_details must not expose a provider endpoint: %s", payload)
	}
	for _, key := range []string{
		"provider_merchant_id", "provider_store_id", "provider_terminal_id", "provider_device_id", "terminal_nickname",
		"provider_base_url", "base_url",
	} {
		if _, ok := got[key]; ok {
			t.Fatalf("Terminal JSON should not include flat provider key %q: %s", key, payload)
		}
	}

}

func TestTerminalContractsExcludeProviderTransportAndOpaqueFields(t *testing.T) {
	assertNoFields(t, reflect.TypeOf(terminal.TerminalProviderDetails{}), "BaseURL")
	assertNoFields(t, reflect.TypeOf(terminal.Terminal{}), "Metadata")
	assertNoFields(t, reflect.TypeOf(terminal.TerminalTransaction{}), "OperationContext", "Payloads", "ProviderData", "ReceiptOptions", "Metadata")
	assertNoFields(t, reflect.TypeOf(settlement.Settlement{}), "OperationContext", "Payloads", "Metadata")
	assertPackageDoesNotDeclareType(t, "ReceiptOptions")
}

func assertNoFields(t *testing.T, recordType reflect.Type, names ...string) {
	t.Helper()
	for _, name := range names {
		if _, exists := recordType.FieldByName(name); exists {
			t.Errorf("%s must not expose provider transport or opaque field %s", recordType.Name(), name)
		}
	}
}

func assertPackageDoesNotDeclareType(t *testing.T, typeName string) {
	t.Helper()
	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate terminal contract test source")
	}

	packages, err := parser.ParseDir(token.NewFileSet(), filepath.Dir(testFile), func(info fs.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, 0)
	if err != nil {
		t.Fatalf("parse terminal contract package: %v", err)
	}
	contractPackage, ok := packages["terminal"]
	if !ok {
		t.Fatalf("terminal package not found while checking retired type %s", typeName)
	}
	for _, sourceFile := range contractPackage.Files {
		for _, declaration := range sourceFile.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpecification, ok := specification.(*ast.TypeSpec)
				if ok && typeSpecification.Name.Name == typeName {
					t.Errorf("terminal contract must not declare retired request-control type %s", typeName)
				}
			}
		}
	}
}

func TestSettlementJSONGroupsProviderSupportFields(t *testing.T) {
	settlement := settlement.Settlement{
		ID:                   "settlement_1",
		TerminalID:           "term_1",
		ProviderSettlementID: "provider_settlement_1",
		Type:                 settlement_enums.SettlementTypeSettlement,
		Status:               terminal_enums.TerminalTxStatusFinalised,
		FinancialStatus:      terminal_enums.TerminalTxFinancialStatusApproved,
		Totals:               settlement.SettlementTotals{Currency: "AUD", TotalMinor: 10000},
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
	for _, key := range []string{
		"operation_context", "provider_payloads", "provider_details", "metadata",
		"provider_merchant_id", "provider_terminal_id", "provider_request_id", "idempotency_key", "provider_request",
	} {
		if _, ok := got[key]; ok {
			t.Fatalf("Settlement JSON should not include provider transport key %q: %s", key, payload)
		}
	}

}
