package wallet_enums

import "testing"

func TestGiftCardIssuanceSources(t *testing.T) {
	for _, source := range []GiftCardIssuanceSource{GiftCardIssuanceSourcePurchase, GiftCardIssuanceSourceMembershipReward, GiftCardIssuanceSourceRefundReplacement} {
		if !source.IsValid() || source.String() != string(source) {
			t.Fatal(source)
		}
	}
	for _, source := range []GiftCardIssuanceSource{"", "reward", "refund", "unknown"} {
		if source.IsValid() {
			t.Fatalf("unknown issuance source accepted: %s", source)
		}
	}
}
