package wallet_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/wallet"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

func TestCustomerWalletRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	owner := membership.MembershipOwnerRef{OwnerType: enums.MembershipOwnerTypeRetailCustomer, OwnerID: "retail_1"}
	bal := common.Money{AmountMinor: 2500, Currency: "AUD"}
	w := wallet.CustomerWallet{
		Owner:               owner,
		MembershipAccountID: "mem_1",
		Instruments: []wallet.WalletInstrument{
			{Type: enums.WalletInstrumentTypeGiftCard, Code: "GC-1", Status: "active", Balance: &bal},
			{Type: enums.WalletInstrumentTypePoints, Code: "mem_1"},
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
	owner := membership.MembershipOwnerRef{OwnerType: enums.MembershipOwnerTypeWholesaleOrganisation, OwnerID: "org_1"}
	gc := wallet.GiftCard{
		ID:           "gc_1",
		Code:         "GC-1",
		Owner:        owner,
		Balance:      common.Money{AmountMinor: 5000, Currency: "AUD"},
		InitialValue: common.Money{AmountMinor: 10000, Currency: "AUD"},
		Status:       enums.GiftCardStatusPartiallyRedeemed,
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
		decoded.Status != enums.GiftCardStatusPartiallyRedeemed {
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
		Reason:             enums.GiftCardTransactionReasonRedeem,
		RelatedOrderNumber: "ORD-1",
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
	if decoded.Reason != enums.GiftCardTransactionReasonRedeem ||
		decoded.BalanceAfter.AmountMinor != 5000 || decoded.RelatedOrderNumber != "ORD-1" {
		t.Fatalf("gift card tx did not round-trip: %+v", decoded)
	}
}
