package wallet_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/benefit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/wallet"
	benefitenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/benefit"
	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/wallet"
)

func TestCustomerWalletRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	bal := common.Money{AmountMinor: 2500, Currency: "AUD"}
	reserved := common.Money{AmountMinor: 500, Currency: "AUD"}
	available := common.Money{AmountMinor: 2000, Currency: "AUD"}
	w := wallet.CustomerWallet{
		CustomerNumber: "RC-20260727-ABCDEF",
		Instruments: []wallet.WalletInstrument{
			{Type: walletenum.WalletInstrumentTypeGiftCard, Code: "GC-1", Status: "active", CommittedBalance: &bal, ReservedBalance: &reserved, AvailableBalance: &available},
			{Type: walletenum.WalletInstrumentTypePoints, Code: "mem_1"},
		},
		Summary: wallet.CustomerWalletSummary{
			AvailablePoints: 1000,
			PointsPolicy: &membership.PointsPolicy{
				PointsPerMinorUnit: 2, MinimumEligibleBalance: 1000,
				RedemptionStepPoints: 200, MaximumRedemptionPoints: 1000,
			},
			GiftCardCommittedBalanceTotal: bal,
			GiftCardReservedBalanceTotal:  reserved,
			GiftCardAvailableBalanceTotal: available,
			ActiveGiftCards:               1,
		},
		CalculatedAt: now,
	}

	payload, err := json.Marshal(w)
	if err != nil {
		t.Fatalf("marshal wallet: %v", err)
	}
	if strings.Contains(string(payload), `"balance":`) || !strings.Contains(string(payload), `"available_balance"`) {
		t.Fatalf("wallet JSON must expose explicit gift-card balances: %s", payload)
	}
	for _, field := range []string{`"points_per_minor_unit":2`, `"minimum_eligible_balance":1000`, `"redemption_step_points":200`, `"maximum_redemption_points":1000`} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("wallet points policy missing %s: %s", field, payload)
		}
	}

	var decoded wallet.CustomerWallet
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal wallet: %v", err)
	}
	if decoded.CustomerNumber != "RC-20260727-ABCDEF" || len(decoded.Instruments) != 2 ||
		decoded.Summary.GiftCardAvailableBalanceTotal.AmountMinor != 2000 {
		t.Fatalf("wallet did not round-trip: %+v", decoded)
	}
	if decoded.Instruments[0].AvailableBalance == nil || decoded.Instruments[0].AvailableBalance.AmountMinor != 2000 {
		t.Fatalf("gift-card instrument balance did not round-trip: %+v", decoded.Instruments[0])
	}
}

func TestMembershipPassContentRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 17, 1, 2, 3, 0, time.UTC)
	content := wallet.MembershipPassContent{
		CustomerNumber:  "RC-20260717-ABCDEF",
		TierKey:         "standard",
		AvailablePoints: 125,
		Barcode: wallet.MembershipPassBarcode{
			Format:        walletenum.WalletPassBarcodeFormatCode128,
			Value:         "RC-20260717-ABCDEF",
			AlternateText: "RC-20260717-ABCDEF",
		},
		GeneratedAt: now,
	}

	payload, err := json.Marshal(content)
	if err != nil {
		t.Fatalf("marshal membership pass content: %v", err)
	}
	for _, key := range []string{
		`"customer_number":"RC-20260717-ABCDEF"`,
		`"tier_key":"standard"`,
		`"available_points":125`,
		`"format":"code_128"`,
		`"value":"RC-20260717-ABCDEF"`,
		`"alternate_text":"RC-20260717-ABCDEF"`,
		`"generated_at":"2026-07-17T01:02:03Z"`,
	} {
		if !strings.Contains(string(payload), key) {
			t.Fatalf("membership pass JSON missing %s: %s", key, payload)
		}
	}
	for _, forbidden := range []string{`"save_jwt"`, `"pkpass"`, `"class_id"`, `"object_id"`, `"platform"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("provider-neutral membership pass must not contain %s: %s", forbidden, payload)
		}
	}

	var decoded wallet.MembershipPassContent
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal membership pass content: %v", err)
	}
	if decoded.CustomerNumber != content.CustomerNumber || decoded.Barcode.Value != content.CustomerNumber ||
		decoded.Barcode.Format != walletenum.WalletPassBarcodeFormatCode128 || !decoded.GeneratedAt.Equal(now) {
		t.Fatalf("membership pass content did not round-trip: %+v", decoded)
	}

	optionalPayload, err := json.Marshal(wallet.MembershipPassContent{
		CustomerNumber: "RC-1",
		Barcode: wallet.MembershipPassBarcode{
			Format: walletenum.WalletPassBarcodeFormatCode128,
			Value:  "RC-1",
		},
		GeneratedAt: now,
	})
	if err != nil {
		t.Fatalf("marshal membership pass without optional fields: %v", err)
	}
	for _, optional := range []string{`"tier_key"`, `"alternate_text"`} {
		if strings.Contains(string(optionalPayload), optional) {
			t.Fatalf("zero-value optional field %s must be omitted: %s", optional, optionalPayload)
		}
	}
}

func TestGiftCardRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	owner := benefit.OwnerRef{OwnerType: benefitenum.OwnerTypeRetailCustomer, OwnerID: "retail_1"}
	gc := wallet.GiftCard{
		ID:                   "gc_1",
		Code:                 "GC-1",
		Owner:                owner,
		CommittedBalance:     common.Money{AmountMinor: 5000, Currency: "AUD"},
		ReservedBalance:      common.Money{AmountMinor: 1000, Currency: "AUD"},
		AvailableBalance:     common.Money{AmountMinor: 4000, Currency: "AUD"},
		InitialValue:         common.Money{AmountMinor: 10000, Currency: "AUD"},
		ReplacesGiftCardCode: "GC-ORIGINAL",
		Status:               walletenum.GiftCardStatusPartiallyRedeemed,
		IssuedAt:             now,
	}

	payload, err := json.Marshal(gc)
	if err != nil {
		t.Fatalf("marshal gift card: %v", err)
	}
	var decoded wallet.GiftCard
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal gift card: %v", err)
	}
	if decoded.Code != "GC-1" || decoded.CommittedBalance.AmountMinor != 5000 ||
		decoded.ReservedBalance.AmountMinor != 1000 || decoded.AvailableBalance.AmountMinor != 4000 ||
		decoded.ReplacesGiftCardCode != "GC-ORIGINAL" ||
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
		ReservationID:      "benefit_res_1",
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
		decoded.BalanceAfter.AmountMinor != 5000 || decoded.ReservationID != "benefit_res_1" ||
		decoded.RelatedOrderNumber != "MAMA260703ABC123" {
		t.Fatalf("gift card tx did not round-trip: %+v", decoded)
	}
}

func TestVoucherReservationRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 15, 2, 0, 0, 0, time.UTC)
	expires := now.Add(30 * time.Minute)
	voucher := wallet.Voucher{
		ID:                       "voucher_1",
		Code:                     "VOUCHER-1",
		Owner:                    benefit.OwnerRef{OwnerType: benefitenum.OwnerTypeRetailCustomer, OwnerID: "RC-1"},
		Value:                    &common.Money{AmountMinor: 1500, Currency: "AUD"},
		Status:                   walletenum.VoucherStatusReserved,
		SourceRewardCode:         "REWARD-15",
		SourceRewardRedemptionID: "reward_redemption_1",
		IssuedAt:                 now,
		ReservationID:            "benefit_res_1",
		ReservedAt:               &now,
		ReservationExpiresAt:     &expires,
	}

	payload, err := json.Marshal(voucher)
	if err != nil {
		t.Fatalf("marshal reserved voucher: %v", err)
	}
	for _, key := range []string{`"status":"reserved"`, `"reservation_id":"benefit_res_1"`, `"reservation_expires_at"`, `"source_reward_redemption_id":"reward_redemption_1"`, `"value"`} {
		if !strings.Contains(string(payload), key) {
			t.Fatalf("reserved voucher JSON missing %s: %s", key, payload)
		}
	}

	var decoded wallet.Voucher
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal reserved voucher: %v", err)
	}
	if decoded.Status != walletenum.VoucherStatusReserved || decoded.ReservationExpiresAt == nil {
		t.Fatalf("reserved voucher did not round-trip: %+v", decoded)
	}
}

func TestCheckoutBenefitReservationRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 15, 2, 0, 0, 0, time.UTC)
	record := wallet.CheckoutBenefitReservation{
		ID:             "benefit_res_1",
		IdempotencyKey: "checkout-key-1",
		Owner:          benefit.OwnerRef{OwnerType: benefitenum.OwnerTypeRetailCustomer, OwnerID: "RC-1"},
		Voucher: &wallet.VoucherBenefitReservation{
			VoucherCode:   "VOUCHER-1",
			AppliedAmount: common.Money{AmountMinor: 1500, Currency: "AUD"},
			Status:        walletenum.CheckoutBenefitReservationStatusReserved,
		},
		GiftCards: []wallet.GiftCardBenefitReservation{
			{
				GiftCardCode:   "GC-FIRST",
				AppliedAmount:  common.Money{AmountMinor: 2000, Currency: "AUD"},
				RefundedAmount: common.Money{Currency: "AUD"},
				Status:         walletenum.CheckoutBenefitReservationStatusReserved,
			},
			{
				GiftCardCode:   "GC-SECOND",
				AppliedAmount:  common.Money{AmountMinor: 500, Currency: "AUD"},
				RefundedAmount: common.Money{Currency: "AUD"},
				Status:         walletenum.CheckoutBenefitReservationStatusReserved,
			},
		},
		Status:    walletenum.CheckoutBenefitReservationStatusReserved,
		ExpiresAt: now.Add(30 * time.Minute),
		CreatedAt: now,
		UpdatedAt: now,
	}

	payload, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("marshal checkout benefit reservation: %v", err)
	}
	if strings.Index(string(payload), "GC-FIRST") >= strings.Index(string(payload), "GC-SECOND") {
		t.Fatalf("gift-card priority order changed: %s", payload)
	}

	var decoded wallet.CheckoutBenefitReservation
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal checkout benefit reservation: %v", err)
	}
	if decoded.IdempotencyKey != "checkout-key-1" || decoded.Voucher == nil || len(decoded.GiftCards) != 2 {
		t.Fatalf("checkout benefit reservation did not round-trip: %+v", decoded)
	}
}
