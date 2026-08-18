package classification

import "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"

// CountryCodeRef is the code-only country relationship persisted by catalog
// records. Display names are resolved from the canonical geography master.
type CountryCodeRef struct {
	Code geography.CountryCode `json:"code"`
}

// ObjectMediaRef is the code-only media relationship persisted by catalog
// records. Render URLs are resolved from the managed media asset.
type ObjectMediaRef struct {
	Code string `json:"code"`
}
