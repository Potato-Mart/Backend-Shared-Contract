package payments_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/contracts/payments"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

func TestTerminalTransactionJSONRoundTripWithHistory(t *testing.T) {
	occurredAt := time.Date(2026, 6, 17, 10, 20, 0, 0, time.UTC)
	tx := payments.TerminalTransaction{
		ID:         "ttx_1",
		TerminalID: "term_1",
		OrderID:    "ord_1",
		PaymentID:  "pay_1",
		ProviderReference: &payments.PaymentReference{
			Mx51: &payments.Mx51PaymentReference{
				TransactionID: "mx_tx_1",
			},
		},
		Type:            enums.TerminalTxTypePurchase,
		Requested:       payments.Amounts{Currency: "AUD", PurchaseMinor: 10000},
		Status:          enums.TerminalTxStatusFinalised,
		FinancialStatus: enums.TerminalTxFinancialStatusApproved,
		Result:          payments.Amounts{Currency: "AUD", PurchaseMinor: 10000, AuthorizedMinor: 10000},
		ProviderResult:  "approved",
		ProviderData: common.Metadata{
			"terminal_batch": "batch_1",
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

	if decoded.ProviderReference == nil || decoded.ProviderReference.Mx51 == nil || decoded.ProviderReference.Mx51.TransactionID != "mx_tx_1" {
		t.Fatalf("provider reference did not round-trip: %+v", decoded.ProviderReference)
	}
	if decoded.Result.AuthorizedMinor != 10000 || decoded.FinancialStatus != enums.TerminalTxFinancialStatusApproved {
		t.Fatalf("result/financial status did not round-trip: %+v", decoded)
	}
	if decoded.ProviderData["terminal_batch"] != "batch_1" {
		t.Fatalf("provider data did not round-trip: %+v", decoded.ProviderData)
	}
	if len(decoded.History) != 1 || decoded.History[0].Changes[0].ToValue != "finalised" {
		t.Fatalf("history did not round-trip: %+v", decoded.History)
	}
}

func TestTerminalTransactionJSONLegacyPayloadWithoutHistory(t *testing.T) {
	var tx payments.TerminalTransaction
	if err := json.Unmarshal([]byte(`{"id":"ttx_1","terminal_id":"term_1","type":"purchase","status":"pending"}`), &tx); err != nil {
		t.Fatalf("unmarshal legacy terminal transaction: %v", err)
	}
	if tx.History != nil {
		t.Fatalf("legacy payload history = %+v, want nil", tx.History)
	}
}
