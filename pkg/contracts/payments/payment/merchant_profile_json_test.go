package payment

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/payments/payment/payment_enums"
)

// The values below are structural placeholders only. Real merchant legal
// identity is supplied through protected execution input and never lives in
// the contract or its tests.
func TestMerchantLegalProfileIsEffectiveDatedPerMarket(t *testing.T) {
	effectiveFrom := time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC)
	registeredFrom := temporal.Date("2020-07-20")
	payload, err := json.Marshal(MerchantLegalProfile{
		ID: "profile_1", MarketCode: "market_au",
		LegalName:             "Example Legal Name",
		TradingName:           "Example Trading Name",
		BusinessNumberScheme:  payment_enums.BusinessNumberSchemeABN,
		BusinessNumber:        "00000000000",
		TaxRegistrationStatus: payment_enums.TaxRegistrationStatusRegistered,
		TaxRegisteredFrom:     &registeredFrom,
		Address: geography.Address{
			Line1: "1 Example Street", Locality: "Example Locality", PostalCode: "0000",
			Country: geography.CountryRef{Code: "AU", Name: "Australia"},
		},
		Contact:       party.Recipient{Name: "Example Contact", Email: "contact@example.test"},
		Status:        payment_enums.MerchantProfileStatusActive,
		EffectiveFrom: effectiveFrom, Revision: 1,
	})
	if err != nil {
		t.Fatalf("marshal merchant profile: %v", err)
	}
	for _, want := range []string{
		`"market_code":"market_au"`, `"legal_name":"Example Legal Name"`,
		`"business_number_scheme":"abn"`, `"tax_registration_status":"registered"`,
		`"tax_registered_from":"2020-07-20"`, `"address"`, `"contact"`,
		`"status":"active"`, `"effective_from":"2026-08-12T00:00:00Z"`, `"revision":1`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("MerchantLegalProfile JSON = %s, want %s", payload, want)
		}
	}
}

func TestMerchantAndBuyerSnapshotsFreezeIssuanceIdentity(t *testing.T) {
	frozenAt := time.Date(2026, 8, 12, 4, 5, 6, 0, time.UTC)
	issuer, err := json.Marshal(MerchantLegalSnapshot{
		ProfileID: "profile_1", ProfileRevision: 3, MarketCode: "market_au",
		LegalName:             "Example Legal Name",
		BusinessNumberScheme:  payment_enums.BusinessNumberSchemeABN,
		BusinessNumber:        "00000000000",
		TaxRegistrationStatus: payment_enums.TaxRegistrationStatusRegistered,
		Address: geography.Address{
			Line1: "1 Example Street", Locality: "Example Locality", PostalCode: "0000",
			Country: geography.CountryRef{Code: "AU", Name: "Australia"},
		},
		Contact:  party.Recipient{Name: "Example Contact"},
		FrozenAt: frozenAt,
	})
	if err != nil {
		t.Fatalf("marshal merchant snapshot: %v", err)
	}
	for _, want := range []string{`"profile_revision":3`, `"frozen_at":"2026-08-12T04:05:06Z"`} {
		if !strings.Contains(string(issuer), want) {
			t.Fatalf("MerchantLegalSnapshot JSON = %s, want %s", issuer, want)
		}
	}

	buyer, err := json.Marshal(BuyerLegalSnapshot{
		Name:                 "Example Buyer",
		BusinessNumberScheme: payment_enums.BusinessNumberSchemeABN,
		BusinessNumber:       "11111111111",
		FrozenAt:             frozenAt,
	})
	if err != nil {
		t.Fatalf("marshal buyer snapshot: %v", err)
	}
	if !strings.Contains(string(buyer), `"business_number":"11111111111"`) {
		t.Fatalf("BuyerLegalSnapshot JSON = %s", buyer)
	}

	minimal, err := json.Marshal(BuyerLegalSnapshot{Name: "Example Buyer", FrozenAt: frozenAt})
	if err != nil {
		t.Fatalf("marshal minimal buyer snapshot: %v", err)
	}
	for _, omitted := range []string{`"business_number_scheme"`, `"business_number"`, `"address"`} {
		if strings.Contains(string(minimal), omitted) {
			t.Fatalf("a buyer below the tax-invoice threshold must omit %s: %s", omitted, minimal)
		}
	}
}
