package access

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/identity/access/access_enums"
)

// StaffGeoScope is the persisted geographic grant held by one workforce
// principal. It is never carried by customer principals.
//
// Level decides which of the remaining fields are meaningful: global carries
// none, country carries CountryCode, market carries MarketCodes (and the
// country they belong to), and depot carries DepotCodes. Depots are the only
// site identity in the platform, so there is no store code.
//
// Each backend resolves and enforces the scope independently; this module
// only records it.
type StaffGeoScope struct {
	Level       access_enums.ScopeLevel `json:"level"`
	CountryCode geography.CountryCode   `json:"country_code,omitempty"`
	MarketCodes []string                `json:"market_codes,omitempty"`
	DepotCodes  []string                `json:"depot_codes,omitempty"`
}
