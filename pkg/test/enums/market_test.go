package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/pricing/market/market_enums"
)

func TestMarketEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "marketenum.MarketStatus", valid: []stringEnum{market_enums.MarketStatusDraft, market_enums.MarketStatusActive, market_enums.MarketStatusSuspended, market_enums.MarketStatusRetired}, invalid: market_enums.MarketStatus("__invalid__")},
	})
}
