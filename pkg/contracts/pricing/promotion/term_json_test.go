package promotion

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"
)

func TestPromotionTermsUseExactlyOneTypedValueArm(t *testing.T) {
	stringValue := "gift"
	integerValue := int64(3)
	basisPointsValue := int64(1_000)
	booleanValue := true
	timestampValue := time.Date(2026, 8, 9, 0, 0, 0, 0, time.UTC)
	terms := []PromotionTerm{
		{Key: "gift_kind", StringValue: &stringValue},
		{Key: "minimum_units", IntegerValue: &integerValue},
		{Key: "discount_basis_points", BasisPointsValue: &basisPointsValue},
		{Key: "requires_membership", BooleanValue: &booleanValue},
		{Key: "fixed_discount", MoneyValue: &money.Money{AmountMinor: 500, Currency: "AUD"}},
		{Key: "preorder_release_at", TimestampValue: &timestampValue},
	}
	for _, term := range terms {
		if got := promotionTermValueArmCount(term); got != 1 {
			t.Fatalf("term %q has %d typed value arms, want exactly one", term.Key, got)
		}
	}
	if promotionTermValueArmCount(PromotionTerm{Key: "invalid"}) == 1 || promotionTermValueArmCount(PromotionTerm{Key: "invalid", StringValue: &stringValue, BooleanValue: &booleanValue}) == 1 {
		t.Fatal("test helper must distinguish zero and multiple typed value arms")
	}

	body, err := json.Marshal(terms)
	if err != nil {
		t.Fatalf("marshal terms: %v", err)
	}
	var got []PromotionTerm
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal terms: %v", err)
	}
	for _, term := range got {
		if count := promotionTermValueArmCount(term); count != 1 {
			t.Fatalf("round-tripped term %q has %d typed value arms, want exactly one", term.Key, count)
		}
	}
}

func promotionTermValueArmCount(term PromotionTerm) int {
	count := 0
	for _, present := range []bool{
		term.StringValue != nil,
		term.IntegerValue != nil,
		term.BasisPointsValue != nil,
		term.BooleanValue != nil,
		term.MoneyValue != nil,
		term.TimestampValue != nil,
	} {
		if present {
			count++
		}
	}
	return count
}
