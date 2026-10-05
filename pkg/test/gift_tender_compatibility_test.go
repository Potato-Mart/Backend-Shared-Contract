package pkg_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/payments/receipt"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/wallet/giftcard"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/wallet/reservation"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/wallet/wallet_enums"
	events "github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pubsub/payments"
)

// These containers join existing shared records only for fixture acceptance;
// services retain authority over allocation and refund workflows.
type giftTenderFixture struct {
	FundingCases      []giftTenderCase         `json:"funding_cases"`
	RefundOptionality []giftRefundOptionalCase `json:"refund_optionality"`
}

type giftTenderCase struct {
	Name         string                                 `json:"name"`
	Order        order.Order                            `json:"order"`
	Reservation  reservation.CheckoutBenefitReservation `json:"reservation"`
	Payments     []payment.Payment                      `json:"payments"`
	Summary      payment.CustomerPaymentSummary         `json:"summary"`
	Receipt      receipt.ReceiptSnapshot                `json:"receipt"`
	Cards        []giftcard.GiftCard                    `json:"cards"`
	Transactions []giftcard.GiftCardTransaction         `json:"transactions"`
	Refund       *events.RefundCompletedEvent           `json:"refund,omitempty"`
}

type giftRefundOptionalCase struct {
	Name  string                      `json:"name"`
	Event events.RefundCompletedEvent `json:"event"`
}

func TestGiftTenderCompatibility(t *testing.T) {
	fixture, raw := loadGiftTenderFixture(t)
	wants := map[string]struct {
		total, external, refunded int64
		gift, giftRefund          []int64
	}{
		"mixed":          {10001, 6000, 0, []int64{2501, 1500}, []int64{0, 0}},
		"gift_only":      {4001, 0, 0, []int64{2501, 1500}, []int64{0, 0}},
		"external_only":  {10001, 10001, 0, nil, nil},
		"partial_refund": {10001, 6000, 7501, []int64{2501, 1500}, []int64{1, 1500}},
	}
	if len(fixture.FundingCases) != len(wants) {
		t.Fatalf("funding cases: got %d, want %d", len(fixture.FundingCases), len(wants))
	}
	seen := map[string]bool{}
	for i, original := range fixture.FundingCases {
		want, ok := wants[original.Name]
		if !ok || seen[original.Name] {
			t.Fatalf("unexpected or duplicate funding case %q", original.Name)
		}
		seen[original.Name] = true
		t.Run(original.Name, func(t *testing.T) {
			c, wire := roundTripGiftTender(t, original, raw.FundingCases[i])
			assertGiftTenderMoney(t, "commercial total", c.Order.Total, want.total)
			assertGiftTenderMoney(t, "commercial subtotal", c.Order.Subtotal, want.total+999)
			assertGiftTenderMoney(t, "promotion discount", c.Order.DiscountAmount, 999)
			assertGiftTenderMoney(t, "shipping", c.Order.ShippingAmount, 0)
			assertGiftTenderMoney(t, "tax", c.Order.TaxAmount, 0)
			if c.Order.Subtotal.AmountMinor-c.Order.DiscountAmount.AmountMinor+c.Order.ShippingAmount.AmountMinor != c.Order.Total.AmountMinor {
				t.Fatal("gift funding changed the commercial order calculation")
			}
			assertGiftTenderMoneyPointer(t, "captured total", c.Summary.CapturedTotal, want.total)
			assertGiftTenderMoneyPointer(t, "refunded total", c.Summary.RefundedTotal, want.refunded)
			assertGiftTenderMoney(t, "receipt total", c.Receipt.Total, want.total)
			if c.Reservation.OrderNumber != c.Order.OrderNumber || c.Summary.OrderNumber != c.Order.OrderNumber || c.Receipt.OrderNumber != c.Order.OrderNumber {
				t.Fatal("order identity differs between funding records")
			}

			giftCount := len(want.gift)
			paymentCount := giftCount
			if want.external > 0 {
				paymentCount++
			}
			if len(c.Order.GiftCardRedemptions) != giftCount || len(c.Reservation.GiftCards) != giftCount || len(c.Cards) != giftCount {
				t.Fatal("gift allocations differ between order, reservation and cards")
			}
			if len(c.Payments) != paymentCount || len(c.Summary.Allocations) != paymentCount || len(c.Receipt.PaymentRows) != paymentCount {
				t.Fatal("funding row count differs; a zero remainder must not fabricate an external payment")
			}
			var captured, refunded, giftTotal, giftRefundTotal int64
			for j, p := range c.Payments {
				amount, refundAmount := want.external, want.refunded
				kind := payment_enums.CustomerPaymentAllocationKindExternalPayment
				if j < giftCount {
					amount, refundAmount = want.gift[j], want.giftRefund[j]
					kind = payment_enums.CustomerPaymentAllocationKindGiftCard
					giftTotal += amount
					giftRefundTotal += refundAmount
					code := fmt.Sprintf("GC%012d", j+1)
					snapshot, reserved, card := c.Order.GiftCardRedemptions[j], c.Reservation.GiftCards[j], c.Cards[j]
					if snapshot.GiftCardCode != code || reserved.GiftCardCode != code || card.Code != code {
						t.Fatal("ordered card codes differ between funding records")
					}
					if snapshot.ReservationID != c.Reservation.ID || snapshot.WalletTransactionID != fmt.Sprintf("wallet-redeem-%d", j) || snapshot.OccurredAt == nil {
						t.Fatal("order lost reservation or wallet transaction evidence")
					}
					assertGiftTenderMoney(t, "snapshot applied", snapshot.AppliedAmount, amount)
					assertGiftTenderMoney(t, "reservation applied", reserved.AppliedAmount, amount)
					assertGiftTenderMoney(t, "reservation refunded", reserved.RefundedAmount, refundAmount)
					if p.ID != fmt.Sprintf("pay-gift-%d", j) || p.Method != payment_enums.PaymentMethodGiftCard || p.Provider != "WALLET" || p.ProviderReference == nil || p.ProviderReference.Wallet == nil || p.ProviderReference.Stripe != nil {
						t.Fatal("gift funding lost its payment identity or wallet source")
					}
					if p.ProviderReference.Wallet.GiftCardCode != code || p.ProviderReference.Wallet.WalletTransactionID != snapshot.WalletTransactionID {
						t.Fatal("gift payment references differ from the committed order snapshot")
					}
				} else {
					refundAmount -= giftRefundTotal
					if p.ID != "pay-external" || p.Method != payment_enums.PaymentMethodCard || p.Provider != "STRIPE_ONLINE" || p.ProviderReference == nil || p.ProviderReference.Wallet != nil || p.ProviderReference.Stripe == nil || p.ProviderReference.Stripe.PaymentIntentID != "pi_acceptance_"+c.Name {
						t.Fatal("external payment lost its distinct processor source")
					}
				}
				assertGiftTenderMoney(t, "payment amount", p.Amount, amount)
				if p.OrderNumber != c.Order.OrderNumber || p.PaidAt == nil {
					t.Fatal("payment lost order identity or capture timestamp")
				}
				a, row := c.Summary.Allocations[j], c.Receipt.PaymentRows[j]
				assertGiftTenderMoney(t, "customer allocation", a.Amount, amount)
				assertGiftTenderMoney(t, "original receipt row", row.Amount, amount)
				if a.AllocationID != p.ID || row.AllocationID != p.ID || a.Kind != kind || row.Kind != kind || a.Method != p.Method || row.Method != p.Method || a.Provider != p.Provider || row.Provider != p.Provider || !a.OccurredAt.Equal(*p.PaidAt) || !row.OccurredAt.Equal(*p.PaidAt) {
					t.Fatal("payment identity, source kind or capture evidence differs in customer rows")
				}
				if row.Status != "completed" || row.RefundedAmount != nil {
					t.Fatal("issued receipt must preserve original funding rows after a refund")
				}
				if want.refunded == 0 {
					if p.RefundAmount != nil || a.RefundedAmount != nil {
						t.Fatal("unrefunded payment acquired optional refund evidence")
					}
				} else {
					assertGiftTenderMoneyPointer(t, "payment refund", p.RefundAmount, refundAmount)
					assertGiftTenderMoneyPointer(t, "customer allocation refund", a.RefundedAmount, refundAmount)
					if refundAmount < 0 || refundAmount > p.Amount.AmountMinor {
						t.Fatal("refunded allocation exceeds original funding")
					}
				}
				captured += amount
				refunded += refundAmount
			}
			if captured != want.total || giftTotal+want.external != want.total || refunded != want.refunded {
				t.Fatal("funding or refund sources do not reconcile to aggregate amounts")
			}
			assertGiftTenderLedger(t, c, want.gift, want.giftRefund)
			if want.refunded == 0 {
				if c.Refund != nil {
					t.Fatal("sale fixture unexpectedly contains a refund event")
				}
			} else {
				if c.Refund == nil || c.Refund.PaymentID != "pay-external" || c.Refund.BenefitReservationID != c.Reservation.ID || c.Refund.OrderNumber != c.Order.OrderNumber || c.Refund.OrderID != c.Order.ID {
					t.Fatal("refund lost payment, reservation or order identity")
				}
				assertGiftTenderMoney(t, "event refund total", c.Refund.Amount, 7501)
				assertGiftTenderMoneyPointer(t, "event wallet portion", c.Refund.GiftCardRefundAmount, 1501)
				if c.Refund.Amount.AmountMinor-c.Refund.GiftCardRefundAmount.AmountMinor != 6000 {
					t.Fatal("event total and wallet portion no longer distinguish the external refund")
				}
			}
			if giftCount == 0 {
				object := giftTenderRawObject(t, wire)
				for model, field := range map[string]string{"order": "gift_card_redemptions", "reservation": "gift_cards"} {
					if _, present := giftTenderRawObject(t, object[model])[field]; present {
						t.Fatalf("external-only %s retains optional %s", model, field)
					}
				}
			}
		})
	}
}

func TestGiftTenderRefundOptionality(t *testing.T) {
	fixture, raw := loadGiftTenderFixture(t)
	wants := map[string]struct {
		total, gift int64
		present     bool
	}{"absent": {100, 0, false}, "known_zero": {100, 0, true}, "positive": {7501, 1501, true}}
	if len(fixture.RefundOptionality) != len(wants) {
		t.Fatal("refund optionality fixture must include absent, known zero and positive")
	}
	seen := map[string]bool{}
	for i, original := range fixture.RefundOptionality {
		want, ok := wants[original.Name]
		if !ok || seen[original.Name] {
			t.Fatalf("unexpected or duplicate refund optionality case %q", original.Name)
		}
		seen[original.Name] = true
		t.Run(original.Name, func(t *testing.T) {
			c, wire := roundTripGiftTender(t, original, raw.RefundOptionality[i])
			assertGiftTenderMoney(t, "refund event total", c.Event.Amount, want.total)
			event := giftTenderRawObject(t, giftTenderRawObject(t, wire)["event"])
			_, present := event["gift_card_refund_amount"]
			if present != want.present || (c.Event.GiftCardRefundAmount != nil) != want.present {
				t.Fatal("absent restoration and a known zero restoration must remain distinct")
			}
			if want.present {
				assertGiftTenderMoneyPointer(t, "wallet restoration", c.Event.GiftCardRefundAmount, want.gift)
				if want.gift > 0 && c.Event.Amount.AmountMinor-c.Event.GiftCardRefundAmount.AmountMinor != 6000 {
					t.Fatal("wallet restoration was conflated with the event refund total")
				}
			}
		})
	}
}

func assertGiftTenderLedger(t *testing.T, c giftTenderCase, applied, restored []int64) {
	t.Helper()
	expectedTransactions := len(applied)
	if c.Refund != nil {
		expectedTransactions += 2
	}
	if len(c.Transactions) != expectedTransactions {
		t.Fatal("unexpected gift ledger transaction count")
	}
	balances := make(map[string]int64, len(c.Cards))
	for i, card := range c.Cards {
		initial := []int64{5000, 3000}[i]
		assertGiftTenderMoney(t, "initial gift value", card.InitialValue, initial)
		assertGiftTenderMoney(t, "committed gift balance", card.CommittedBalance, initial-applied[i]+restored[i])
		assertGiftTenderMoney(t, "reserved gift balance", card.ReservedBalance, 0)
		assertGiftTenderMoney(t, "available gift balance", card.AvailableBalance, card.CommittedBalance.AmountMinor-card.ReservedBalance.AmountMinor)
		wantStatus := wallet_enums.GiftCardStatusPartiallyRedeemed
		if c.Name == "partial_refund" && i == 1 {
			wantStatus = wallet_enums.GiftCardStatusActive
		}
		if card.Status != wantStatus {
			t.Fatalf("card %s status: got %s, want %s", card.Code, card.Status, wantStatus)
		}
		balances[card.Code] = initial
	}
	for i, entry := range c.Transactions {
		cardIndex, delta, reason := i, int64(0), wallet_enums.GiftCardTransactionReasonRedeem
		if i < len(applied) {
			delta = -applied[i]
			if entry.ID != c.Order.GiftCardRedemptions[i].WalletTransactionID {
				t.Fatal("redemption ledger identity differs from order and payment wallet references")
			}
		} else {
			// Inspected Pricing restoration evidence: second card 1500, then first card 1.
			cardIndex = len(applied) - 1 - (i - len(applied))
			delta, reason = restored[cardIndex], wallet_enums.GiftCardTransactionReasonRefund
			if entry.ID != fmt.Sprintf("wallet-refund-%d", cardIndex) || delta <= 0 {
				t.Fatal("restoration ledger identity or positive refund delta changed")
			}
		}
		code := c.Cards[cardIndex].Code
		if entry.GiftCardCode != code || entry.Reason != reason || entry.ReservationID != c.Reservation.ID || entry.RelatedOrderNumber != c.Order.OrderNumber {
			t.Fatal("ledger lost ordered source card, reason or business references")
		}
		assertGiftTenderMoney(t, "signed ledger delta", entry.Delta, delta)
		balances[code] += delta
		assertGiftTenderMoney(t, "ledger balance after", entry.BalanceAfter, balances[code])
	}
	for _, card := range c.Cards {
		if balances[card.Code] != card.CommittedBalance.AmountMinor {
			t.Fatal("gift-card committed balance does not reconcile with signed ledger deltas")
		}
	}
}

func loadGiftTenderFixture(t *testing.T) (giftTenderFixture, struct {
	FundingCases      []json.RawMessage `json:"funding_cases"`
	RefundOptionality []json.RawMessage `json:"refund_optionality"`
}) {
	t.Helper()
	data, err := os.ReadFile("testdata/gift_tender_compatibility.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture giftTenderFixture
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&fixture); err != nil {
		t.Fatalf("decode shared gift tender records: %v", err)
	}
	var raw struct {
		FundingCases      []json.RawMessage `json:"funding_cases"`
		RefundOptionality []json.RawMessage `json:"refund_optionality"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatal(err)
	}
	return fixture, raw
}

func roundTripGiftTender[T any](t *testing.T, original T, raw json.RawMessage) (T, []byte) {
	t.Helper()
	wire, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	var decoded T
	if err := json.Unmarshal(wire, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(original, decoded) {
		t.Fatal("shared records lost identity, ordering, currency or optional evidence on JSON round trip")
	}
	assertGiftTenderRawMoney(t, raw, wire, "fixture")
	return decoded, wire
}

// RawMessage traversal compares integer JSON tokens without ever decoding money
// through float64. Existing money/exponent and secret-exclusion tests remain authoritative.
func assertGiftTenderRawMoney(t *testing.T, before, after json.RawMessage, path string) {
	t.Helper()
	before = bytes.TrimSpace(before)
	if len(before) == 0 {
		t.Fatalf("missing JSON at %s", path)
	}
	switch before[0] {
	case '{':
		original, encoded := giftTenderRawObject(t, before), giftTenderRawObject(t, after)
		if amount, ok := original["amount_minor"]; ok {
			var exact int64
			if err := json.Unmarshal(amount, &exact); err != nil {
				t.Fatalf("noninteger money at %s: %v", path, err)
			}
			if string(bytes.TrimSpace(encoded["amount_minor"])) != strconv.FormatInt(exact, 10) || !bytes.Equal(original["currency"], encoded["currency"]) {
				t.Fatalf("integer amount or currency changed at %s", path)
			}
		}
		for key, value := range original {
			assertGiftTenderRawMoney(t, value, encoded[key], path+"."+key)
		}
	case '[':
		var original, encoded []json.RawMessage
		if err := json.Unmarshal(before, &original); err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(after, &encoded); err != nil {
			t.Fatal(err)
		}
		if len(original) != len(encoded) {
			t.Fatalf("array length changed at %s", path)
		}
		for i := range original {
			assertGiftTenderRawMoney(t, original[i], encoded[i], fmt.Sprintf("%s[%d]", path, i))
		}
	}
}

func giftTenderRawObject(t *testing.T, raw json.RawMessage) map[string]json.RawMessage {
	t.Helper()
	var object map[string]json.RawMessage
	if err := json.Unmarshal(raw, &object); err != nil {
		t.Fatalf("decode raw JSON object: %v", err)
	}
	return object
}

func assertGiftTenderMoney(t *testing.T, label string, got money.Money, want int64) {
	t.Helper()
	if got.AmountMinor != want || got.Currency != "AUD" {
		t.Fatalf("%s: got %d %s, want %d AUD", label, got.AmountMinor, got.Currency, want)
	}
}

func assertGiftTenderMoneyPointer(t *testing.T, label string, got *money.Money, want int64) {
	t.Helper()
	if got == nil {
		t.Fatalf("%s: missing money pointer", label)
	}
	assertGiftTenderMoney(t, label, *got, want)
}
