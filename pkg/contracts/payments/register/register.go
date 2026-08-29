package register

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
)

// Register is one physical or virtual point-of-sale register standing in a
// depot. TerminalID links the register to its paired payment terminal when
// one is bound.
type Register struct {
	ID          string                `json:"id"`
	DepotCode   string                `json:"depot_code"`
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	TerminalID  string                `json:"terminal_id,omitempty"`
	Name        string                `json:"name"`
	Status      string                `json:"status,omitempty"`

	audit.AuditFields
}
