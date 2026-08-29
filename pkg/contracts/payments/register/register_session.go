package register

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/temporal"
	pos_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/register/register_enums"
	"time"
)

// RegisterSession is the single trading session one register runs for one
// business day, from open to close-out.
//
// A session belongs to the register, not to an operator: every operator
// working that register on that business date shares the same session, and
// every workforce rank with POS access may open, manage, and close it. One
// register therefore has at most one session per business date, and
// (register_id, business_date) is the record's natural key.
type RegisterSession struct {
	ID           string                  `json:"id"`
	RegisterID   string                  `json:"register_id"`
	TerminalID   string                  `json:"terminal_id,omitempty"`
	DepotCode    string                  `json:"depot_code"`
	MarketCode   string                  `json:"market_code,omitempty"`
	CountryCode  geography.CountryCode   `json:"country_code,omitempty"`
	BusinessDate temporal.Date           `json:"business_date"`
	Status       pos_enums.SessionStatus `json:"status"`

	OpenedAt       time.Time  `json:"opened_at"`
	OpenedByUserID string     `json:"opened_by_user_id"`
	ClosedAt       *time.Time `json:"closed_at,omitempty"`
	ClosedByUserID string     `json:"closed_by_user_id,omitempty"`

	OpeningFloat money.Money  `json:"opening_float"`
	ClosingCount *money.Money `json:"closing_count,omitempty"`
	ExpectedCash *money.Money `json:"expected_cash,omitempty"`
	CashVariance *money.Money `json:"cash_variance,omitempty"`

	audit.AuditFields
}
