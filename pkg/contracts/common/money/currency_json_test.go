package money

import (
	"encoding/json"
	"testing"
)

func TestMoneyJSONKeepsMinorUnitsAndTypedCurrency(t *testing.T) {
	payload, err := json.Marshal(Money{AmountMinor: 1234, Currency: "AUD"})
	if err != nil {
		t.Fatalf("marshal money: %v", err)
	}
	if got, want := string(payload), `{"amount_minor":1234,"currency":"AUD"}`; got != want {
		t.Fatalf("Money JSON = %s, want %s", got, want)
	}

	var decoded Money
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal money: %v", err)
	}
	if decoded.AmountMinor != 1234 || decoded.Currency != "AUD" {
		t.Fatalf("money did not round-trip: %+v", decoded)
	}
}

func TestCurrencyExponentCarriesNonTwoDecimalCurrencies(t *testing.T) {
	for _, tc := range []struct {
		exponent CurrencyExponent
		want     string
	}{
		{CurrencyExponent{Currency: "AUD", Exponent: 2}, `{"currency":"AUD","exponent":2}`},
		{CurrencyExponent{Currency: "JPY", Exponent: 0}, `{"currency":"JPY","exponent":0}`},
		{CurrencyExponent{Currency: "BHD", Exponent: 3}, `{"currency":"BHD","exponent":3}`},
	} {
		payload, err := json.Marshal(tc.exponent)
		if err != nil {
			t.Fatalf("marshal currency exponent: %v", err)
		}
		if string(payload) != tc.want {
			t.Fatalf("CurrencyExponent JSON = %s, want %s", payload, tc.want)
		}
	}
}
