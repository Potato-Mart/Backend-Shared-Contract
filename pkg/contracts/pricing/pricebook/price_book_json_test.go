package pricebook

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/wholesale/wholesale_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/product/product_enums"
)

func TestPriceBookOwnsCurrencyChannelAudienceAndPolicies(t *testing.T) {
	validFrom := time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC)
	payload, err := json.Marshal(PriceBook{
		ID: "book_au_online", Code: "AU_ONLINE", Name: "AU online", MarketCode: "market_au",
		Currency: "AUD", CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2},
		Channel: commerce_enums.OrderTypeOnline, Audience: product_enums.PriceAudienceRetail,
		TaxInclusion: pricebook_enums.PriceTaxInclusionInclusive,
		PriceEnding:  pricebook_enums.PriceEndingPolicyCharmNine,
		Status:       pricebook_enums.PriceBookStatusActive,
		ValidFrom:    validFrom, Revision: 2,
	})
	if err != nil {
		t.Fatalf("marshal price book: %v", err)
	}
	for _, want := range []string{
		`"market_code":"market_au"`, `"currency":"AUD"`, `"channel":"online"`,
		`"audience":"retail"`, `"tax_inclusion":"tax_inclusive"`,
		`"price_ending":"charm_nine"`, `"status":"active"`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("PriceBook JSON = %s, want %s", payload, want)
		}
	}
}

func TestPriceEntryIsApprovalGatedAndRevisioned(t *testing.T) {
	validFrom := time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC)
	sourceRevision := int64(7)
	value := PriceEntry{
		ID: "entry_1", PriceBookCode: "book_au_online", SKUCode: "sku_a00001",
		Amount:     money.Money{AmountMinor: 319, Currency: "AUD"},
		Status:     pricebook_enums.PriceEntryStatusDraft,
		Derivation: pricebook_enums.PriceDerivationSuggestedFromBaseCost,
		ValidFrom:  validFrom, SourceBaseCostRevision: &sourceRevision, Revision: 1,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal price entry: %v", err)
	}
	for _, want := range []string{
		`"price_book_code":"book_au_online"`, `"sku_code":"sku_a00001"`,
		`"amount":{"amount_minor":319,"currency":"AUD"}`,
		`"status":"draft"`, `"derivation":"suggested_from_base_cost"`,
		`"source_base_cost_revision":7`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("PriceEntry JSON = %s, want %s", payload, want)
		}
	}
	if strings.Contains(string(payload), `"approval"`) {
		t.Fatalf("an unapproved entry must omit approval evidence: %s", payload)
	}

	var decoded PriceEntry
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal price entry: %v", err)
	}
	if decoded.Status == pricebook_enums.PriceEntryStatusApproved {
		t.Fatal("a draft must never decode as approved")
	}
}

func TestPriceBookAssignmentPopulatesOnlyItsResolutionLevel(t *testing.T) {
	organisation, err := json.Marshal(PriceBookAssignment{
		ID: "assignment_1", MarketCode: "market_au", PriceBookCode: "book_au_supermarket",
		Kind:                 pricebook_enums.PriceBookAssignmentKindOrganisationCategory,
		OrganisationCategory: wholesale_enums.WholesaleOrganisationCategorySupermarket,
		Status:               pricebook_enums.PriceBookStatusActive,
	})
	if err != nil {
		t.Fatalf("marshal assignment: %v", err)
	}
	if !strings.Contains(string(organisation), `"organisation_category":"supermarket"`) {
		t.Fatalf("category assignment JSON = %s", organisation)
	}
	for _, omitted := range []string{`"channel"`, `"organisation_code"`} {
		if strings.Contains(string(organisation), omitted) {
			t.Fatalf("category assignment must omit %s: %s", omitted, organisation)
		}
	}

	channel, err := json.Marshal(PriceBookAssignment{
		ID: "assignment_2", MarketCode: "market_au", PriceBookCode: "book_au_online",
		Kind:    pricebook_enums.PriceBookAssignmentKindChannelDefault,
		Channel: commerce_enums.OrderTypeOnline,
		Status:  pricebook_enums.PriceBookStatusActive,
	})
	if err != nil {
		t.Fatalf("marshal assignment: %v", err)
	}
	if !strings.Contains(string(channel), `"channel":"online"`) || strings.Contains(string(channel), `"organisation_category"`) {
		t.Fatalf("channel assignment JSON = %s", channel)
	}
}
