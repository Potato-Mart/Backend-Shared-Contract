package market

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/market/market_enums"
)

func TestMarketJSONSeparatesCommerceFromGeography(t *testing.T) {
	value := Market{
		ID:               "market_au",
		Code:             "AU",
		Name:             "Australia",
		CountryCode:      "AU",
		DefaultCurrency:  "AUD",
		CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2},
		DefaultLocale:    "en-AU",
		DefaultTimezone:  "Australia/Melbourne",
		Status:           market_enums.MarketStatusActive,
		Revision:         1,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal market: %v", err)
	}
	for _, want := range []string{
		`"id":"market_au"`, `"code":"AU"`, `"country_code":"AU"`,
		`"default_currency":"AUD"`, `"currency_exponent":{"currency":"AUD","exponent":2}`,
		`"status":"active"`, `"revision":1`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("Market JSON = %s, want %s", payload, want)
		}
	}
	for _, forbidden := range []string{`"sku_id"`, `"price"`, `"amount_minor"`, `"depot_code"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("Market JSON leaked %s: %s", forbidden, payload)
		}
	}

	var decoded Market
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal market: %v", err)
	}
	if decoded.ID == string(decoded.CountryCode) {
		t.Fatal("MarketID and CountryCode must remain separate concepts")
	}
	if decoded.CurrencyExponent.Exponent != 2 {
		t.Fatalf("currency exponent did not round-trip: %+v", decoded.CurrencyExponent)
	}
}
