package wallet_test

import (
	"encoding/json"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/benefit"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/wallet"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/benefit/benefit_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/wallet/wallet_enums"
)

func TestCustomerWalletRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	bal := money.Money{AmountMinor: 2500, Currency: "AUD"}
	reserved := money.Money{AmountMinor: 500, Currency: "AUD"}
	available := money.Money{AmountMinor: 2000, Currency: "AUD"}
	activatedAt := now.Add(time.Hour)
	redeemedAt := now.Add(2 * time.Hour)
	w := wallet.CustomerWallet{
		CustomerNumber: "RC-20260727-ABCDEF",
		Instruments: []wallet.WalletInstrument{
			{
				Type: wallet_enums.WalletInstrumentTypeGiftCard, Code: "GC-1", Status: "active",
				CommittedBalance: &bal, ReservedBalance: &reserved, AvailableBalance: &available,
				IssuedAt: &now, ActivatedAt: &activatedAt,
			},
			{Type: wallet_enums.WalletInstrumentTypePoints, Code: "mem_1"},
			{Type: wallet_enums.WalletInstrumentTypeVoucher, Code: "VOUCHER-1", IssuedAt: &now, RedeemedAt: &redeemedAt},
		},
		Summary: wallet.CustomerWalletSummary{
			AvailablePoints: 1000,
			PointDebt:       25,
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
	for _, field := range []string{
		`"point_debt":25`,
		`"issued_at":"2026-06-30T00:00:00Z"`,
		`"activated_at":"2026-06-30T01:00:00Z"`,
		`"redeemed_at":"2026-06-30T02:00:00Z"`,
		`"points_per_minor_unit":2`,
		`"minimum_eligible_balance":1000`,
		`"redemption_step_points":200`,
		`"maximum_redemption_points":1000`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("wallet points policy missing %s: %s", field, payload)
		}
	}

	var decoded wallet.CustomerWallet
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal wallet: %v", err)
	}
	if decoded.CustomerNumber != "RC-20260727-ABCDEF" || len(decoded.Instruments) != 3 ||
		decoded.Summary.GiftCardAvailableBalanceTotal.AmountMinor != 2000 {
		t.Fatalf("wallet did not round-trip: %+v", decoded)
	}
	if decoded.Instruments[0].AvailableBalance == nil || decoded.Instruments[0].AvailableBalance.AmountMinor != 2000 ||
		decoded.Instruments[0].IssuedAt == nil || decoded.Instruments[0].ActivatedAt == nil ||
		decoded.Instruments[0].RedeemedAt != nil || decoded.Summary.PointDebt != 25 {
		t.Fatalf("gift-card instrument balance did not round-trip: %+v", decoded.Instruments[0])
	}
	if decoded.Instruments[2].RedeemedAt == nil || !decoded.Instruments[2].RedeemedAt.Equal(redeemedAt) {
		t.Fatalf("single-use instrument lifecycle did not round-trip: %+v", decoded.Instruments[2])
	}
}

func TestGiftCardDenominationPolicyRoundTrip(t *testing.T) {
	policy := wallet.GiftCardDenominationPolicy{
		Version:             2,
		Currency:            "AUD",
		AllowedAmountsMinor: []int64{50_000, 80_000, 100_000, 150_000, 200_000},
	}

	payload, err := json.Marshal(policy)
	if err != nil {
		t.Fatalf("marshal gift-card denomination policy: %v", err)
	}
	if string(payload) != `{"version":2,"currency":"AUD","allowed_amounts_minor":[50000,80000,100000,150000,200000]}` {
		t.Fatalf("gift-card denomination policy JSON = %s", payload)
	}

	var decoded wallet.GiftCardDenominationPolicy
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal gift-card denomination policy: %v", err)
	}
	if decoded.Version != 2 || decoded.Currency != "AUD" || len(decoded.AllowedAmountsMinor) != 5 || decoded.AllowedAmountsMinor[1] != 80_000 {
		t.Fatalf("gift-card denomination policy did not round-trip: %+v", decoded)
	}
}

func TestGiftCardDenominationBonusRoundTrip(t *testing.T) {
	// Byte-exact wire form. Bonuses are an ordered slice of pairs, never a map,
	// so consumers may compare two policy revisions byte for byte.
	const wire = `{"version":3,"currency":"AUD","allowed_amounts_minor":[50000,100000,200000],` +
		`"bonus_amounts_minor":[{"amount_minor":50000,"bonus_minor":0},` +
		`{"amount_minor":100000,"bonus_minor":5000},` +
		`{"amount_minor":200000,"bonus_minor":15000}]}`

	var decoded wallet.GiftCardDenominationPolicy
	if err := json.Unmarshal([]byte(wire), &decoded); err != nil {
		t.Fatalf("unmarshal gift-card denomination policy: %v", err)
	}

	payload, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("marshal gift-card denomination policy: %v", err)
	}
	if string(payload) != wire {
		t.Fatalf("gift-card denomination bonus JSON = %s, want %s", payload, wire)
	}
	if len(decoded.BonusAmountsMinor) != 3 ||
		decoded.BonusAmountsMinor[1] != (wallet.GiftCardDenominationBonus{AmountMinor: 100_000, BonusMinor: 5_000}) ||
		decoded.BonusAmountsMinor[2].BonusMinor != 15_000 {
		t.Fatalf("gift-card denomination bonuses did not round-trip: %+v", decoded.BonusAmountsMinor)
	}

	// The buyer is charged AmountMinor; the issued card carries the sum.
	bonus := decoded.BonusAmountsMinor[2]
	if bonus.AmountMinor+bonus.BonusMinor != 215_000 {
		t.Fatalf("issued balance = %d, want 215000", bonus.AmountMinor+bonus.BonusMinor)
	}

	// A policy with no bonuses keeps the pre-v27.3.0 wire form byte for byte.
	legacy, err := json.Marshal(wallet.GiftCardDenominationPolicy{
		Version: 3, Currency: "AUD", AllowedAmountsMinor: []int64{50_000},
	})
	if err != nil {
		t.Fatalf("marshal bonus-free gift-card denomination policy: %v", err)
	}
	if string(legacy) != `{"version":3,"currency":"AUD","allowed_amounts_minor":[50000]}` {
		t.Fatalf("bonus-free denomination policy JSON = %s", legacy)
	}
}

func TestMembershipPassContentRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 17, 1, 2, 3, 0, time.UTC)
	content := wallet.MembershipPassContent{
		CustomerNumber:  "RC-20260717-ABCDEF",
		TierKey:         "standard",
		AvailablePoints: 125,
		Barcode: wallet.MembershipPassBarcode{
			Format:        wallet_enums.WalletPassBarcodeFormatCode128,
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
		decoded.Barcode.Format != wallet_enums.WalletPassBarcodeFormatCode128 || !decoded.GeneratedAt.Equal(now) {
		t.Fatalf("membership pass content did not round-trip: %+v", decoded)
	}

	optionalPayload, err := json.Marshal(wallet.MembershipPassContent{
		CustomerNumber: "RC-1",
		Barcode: wallet.MembershipPassBarcode{
			Format: wallet_enums.WalletPassBarcodeFormatCode128,
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
	owner := benefit.OwnerRef{OwnerType: benefit_enums.OwnerTypeRetailCustomer, OwnerID: "retail_1"}
	gc := wallet.GiftCard{
		ID:                   "gc_1",
		Code:                 "GC-1",
		Owner:                owner,
		CommittedBalance:     money.Money{AmountMinor: 5000, Currency: "AUD"},
		ReservedBalance:      money.Money{AmountMinor: 1000, Currency: "AUD"},
		AvailableBalance:     money.Money{AmountMinor: 4000, Currency: "AUD"},
		InitialValue:         money.Money{AmountMinor: 10000, Currency: "AUD"},
		ReplacesGiftCardCode: "GC-ORIGINAL",
		Status:               wallet_enums.GiftCardStatusPartiallyRedeemed,
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
		decoded.Status != wallet_enums.GiftCardStatusPartiallyRedeemed {
		t.Fatalf("gift card did not round-trip: %+v", decoded)
	}
}

func TestGiftCardTransactionRoundTrip(t *testing.T) {
	now := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	tx := wallet.GiftCardTransaction{
		ID:                 "gctx_1",
		GiftCardCode:       "GC-1",
		Delta:              money.Money{AmountMinor: -5000, Currency: "AUD"},
		BalanceAfter:       money.Money{AmountMinor: 5000, Currency: "AUD"},
		Reason:             wallet_enums.GiftCardTransactionReasonRedeem,
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
	if decoded.Reason != wallet_enums.GiftCardTransactionReasonRedeem ||
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
		Owner:                    benefit.OwnerRef{OwnerType: benefit_enums.OwnerTypeRetailCustomer, OwnerID: "RC-1"},
		Value:                    &money.Money{AmountMinor: 1500, Currency: "AUD"},
		Status:                   wallet_enums.VoucherStatusReserved,
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
	if decoded.Status != wallet_enums.VoucherStatusReserved || decoded.ReservationExpiresAt == nil {
		t.Fatalf("reserved voucher did not round-trip: %+v", decoded)
	}
}

func TestCheckoutBenefitReservationRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 15, 2, 0, 0, 0, time.UTC)
	record := wallet.CheckoutBenefitReservation{
		ID:             "benefit_res_1",
		IdempotencyKey: "checkout-key-1",
		Owner:          benefit.OwnerRef{OwnerType: benefit_enums.OwnerTypeRetailCustomer, OwnerID: "RC-1"},
		Voucher: &wallet.VoucherBenefitReservation{
			VoucherCode:   "VOUCHER-1",
			AppliedAmount: money.Money{AmountMinor: 1500, Currency: "AUD"},
			Status:        wallet_enums.CheckoutBenefitReservationStatusReserved,
		},
		GiftCards: []wallet.GiftCardBenefitReservation{
			{
				GiftCardCode:   "GC-FIRST",
				AppliedAmount:  money.Money{AmountMinor: 2000, Currency: "AUD"},
				RefundedAmount: money.Money{Currency: "AUD"},
				Status:         wallet_enums.CheckoutBenefitReservationStatusReserved,
			},
			{
				GiftCardCode:   "GC-SECOND",
				AppliedAmount:  money.Money{AmountMinor: 500, Currency: "AUD"},
				RefundedAmount: money.Money{Currency: "AUD"},
				Status:         wallet_enums.CheckoutBenefitReservationStatusReserved,
			},
		},
		Status:    wallet_enums.CheckoutBenefitReservationStatusReserved,
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
