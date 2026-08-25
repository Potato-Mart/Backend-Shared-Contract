package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
)

// CountryCodeRef is the code-only country relationship persisted by catalog
// records. Display names are resolved from the canonical geography master.
type CountryCodeRef struct {
	Code geography.CountryCode `json:"code"`
}
