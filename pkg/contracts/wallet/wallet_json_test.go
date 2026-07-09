package wallet_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/wallet"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/membership"
	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/wallet"
)

func TestCustomerWalletRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	owner := membership.MembershipOwnerRef{OwnerType: membershipenum.MembershipOwnerTypeRetailCustomer, OwnerID: "retail_1"}
	bal := common.Money{AmountMinor: 2500, Currency: "AUD"}
	w := wallet.CustomerWallet{
		Owner:               owner,
		MembershipAccountID: "mem_1",
		Instruments: []wallet.WalletInstrument{
			{Type: walletenum.WalletInstrumentTypeGiftCard, Code: "GC-1", Status: "active", Balance: &bal},
			{Type: walletenum.WalletInstrumentTypePoints, Code: "mem_1"},
		},
		Summary:      wallet.CustomerWalletSummary{AvailablePoints: 1000, GiftCardBalanceTotal: bal, ActiveGiftCards: 1},
		CalculatedAt: now,
	}

	payload, err := json.Marshal(w)
	if err != nil {
		t.Fatalf("marshal wallet: %v", err)
	}
	if !strings.Contains(string(payload), `"gift_card_balance_total"`) {
		t.Fatalf("wallet JSON missing gift_card_balance_total: %s", payload)
	}

	var decoded wallet.CustomerWallet
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal wallet: %v", err)
	}
	if decoded.Owner.OwnerID != "retail_1" || len(decoded.Instruments) != 2 ||
		decoded.Summary.GiftCardBalanceTotal.AmountMinor != 2500 {
		t.Fatalf("wallet did not round-trip: %+v", decoded)
	}
	if decoded.Instruments[0].Balance == nil || decoded.Instruments[0].Balance.AmountMinor != 2500 {
		t.Fatalf("gift-card instrument balance did not round-trip: %+v", decoded.Instruments[0])
	}
}

func TestGiftCardRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	owner := membership.MembershipOwnerRef{OwnerType: membershipenum.MembershipOwnerTypeWholesaleOrganisation, OwnerID: "org_1"}
	gc := wallet.GiftCard{
		ID:           "gc_1",
		Code:         "GC-1",
		Owner:        owner,
		Balance:      common.Money{AmountMinor: 5000, Currency: "AUD"},
		InitialValue: common.Money{AmountMinor: 10000, Currency: "AUD"},
		Status:       walletenum.GiftCardStatusPartiallyRedeemed,
		IssuedAt:     now,
	}

	payload, err := json.Marshal(gc)
	if err != nil {
		t.Fatalf("marshal gift card: %v", err)
	}
	var decoded wallet.GiftCard
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal gift card: %v", err)
	}
	if decoded.Code != "GC-1" || decoded.Balance.AmountMinor != 5000 ||
		decoded.Status != walletenum.GiftCardStatusPartiallyRedeemed {
		t.Fatalf("gift card did not round-trip: %+v", decoded)
	}
}

func TestGiftCardTransactionRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	tx := wallet.GiftCardTransaction{
		ID:                 "gctx_1",
		GiftCardCode:       "GC-1",
		Delta:              common.Money{AmountMinor: -5000, Currency: "AUD"},
		BalanceAfter:       common.Money{AmountMinor: 5000, Currency: "AUD"},
		Reason:             walletenum.GiftCardTransactionReasonRedeem,
		RelatedOrderNumber: "MAMA260703ABC123",
		CreatedAt:          now,
	}

	payload, err := json.Marshal(tx)
	if err != nil {
		t.Fatalf("marshal gift card tx: %v", err)
	}
	var decoded wallet.GiftCardTransaction
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal gift card tx: %v", err)
	}
	if decoded.Reason != walletenum.GiftCardTransactionReasonRedeem ||
		decoded.BalanceAfter.AmountMinor != 5000 || decoded.RelatedOrderNumber != "MAMA260703ABC123" {
		t.Fatalf("gift card tx did not round-trip: %+v", decoded)
	}
}

func TestWalletExportRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 5, 0, 0, 0, 0, time.UTC)
	owner := membership.MembershipOwnerRef{OwnerType: membershipenum.MembershipOwnerTypeRetailCustomer, OwnerID: "RC-1"}
	export := wallet.WalletExport{
		SchemaVersion: wallet.WalletExportSchemaVersion,
		Owner:         owner,
		Membership: wallet.WalletExportMembership{
			Account: &membership.MembershipAccount{
				ID:     "mem_1",
				Owner:  owner,
				Status: membershipenum.MembershipAccountStatusActive,
				Wallet: membership.MembershipWalletSummary{
					AvailablePoints: 150,
					CalculatedAt:    now,
				},
				EnrolledAt: now,
			},
			PointLedger: []membership.PointLedgerEntry{{
				ID:                  "pledger_1",
				MembershipAccountID: "mem_1",
				Owner:               owner,
				Delta:               150,
				BalanceAfter:        150,
				Remaining:           150,
				CreatedAt:           now,
			}},
		},
		History: []wallet.WalletExportHistoryEvent{{
			At:          now,
			Type:        walletenum.WalletInstrumentTypePoints,
			Code:        "mem_1",
			DeltaPoints: 150,
			Status:      "earned",
		}},
		Summary:     wallet.WalletExportSummary{AvailablePoints: 150},
		GeneratedAt: now,
		Filters: wallet.WalletExportFilters{
			IncludeHistory:  true,
			IncludeExpired:  true,
			IncludeRedeemed: true,
			IncludeVoided:   true,
		},
		RecordCounts: wallet.WalletExportRecordCounts{
			MembershipPointLedgerEntries: 1,
			HistoryEvents:                1,
		},
	}

	payload, err := json.Marshal(export)
	if err != nil {
		t.Fatalf("marshal wallet export: %v", err)
	}
	if !strings.Contains(string(payload), `"schema_version":"wallet_export_v1"`) {
		t.Fatalf("wallet export payload missing schema version: %s", payload)
	}
	if !strings.Contains(string(payload), `"point_ledger"`) {
		t.Fatalf("wallet export payload missing point ledger: %s", payload)
	}

	var decoded wallet.WalletExport
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal wallet export: %v", err)
	}
	if decoded.SchemaVersion != wallet.WalletExportSchemaVersion ||
		decoded.Owner.OwnerID != "RC-1" ||
		len(decoded.Membership.PointLedger) != 1 ||
		decoded.RecordCounts.HistoryEvents != 1 {
		t.Fatalf("wallet export did not round-trip: %+v", decoded)
	}
}
