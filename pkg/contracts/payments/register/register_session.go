package register

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal"
	pos_enums "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/register/register_enums"
	"time"
)

// RegisterSession is one trading session a register runs, from open to
// close-out.
//
// A session belongs to the register, not to an operator: every operator
// working that register shares the open session, and every workforce rank with
// POS access may open, manage, and close it. A register has at most one OPEN
// session at a time, and may run several sequential sessions per business
// date: a new session can open only after the previous one closed. Sequence
// numbers a register's sessions within one business date from 1, so
// (register_id, business_date, sequence) is the record's natural key. Records
// written before sequential sessions omit sequence and are the first session
// of their business date.
type RegisterSession struct {
	ID           string                  `json:"id"`
	RegisterID   string                  `json:"register_id"`
	TerminalID   string                  `json:"terminal_id,omitempty"`
	DepotCode    string                  `json:"depot_code"`
	MarketCode   string                  `json:"market_code,omitempty"`
	CountryCode  geography.CountryCode   `json:"country_code,omitempty"`
	BusinessDate temporal.Date           `json:"business_date"`
	Sequence     int64                   `json:"sequence,omitempty"`
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
	security.DataProtectionFields
}
