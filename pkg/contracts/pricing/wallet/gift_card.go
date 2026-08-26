package wallet

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/benefit"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/wallet/wallet_enums"
)

// GiftCard is a stored-value instrument with a re-spendable balance. The
// GiftCardTransaction ledger is the source of truth for CommittedBalance. Live
// checkout reservations are summarized by ReservedBalance and subtracted to
// produce AvailableBalance. It is referenced everywhere by Code (the business
// key), never by ID.
type GiftCard struct {
	ID               string           `json:"id"`
	Code             string           `json:"code"`
	Owner            benefit.OwnerRef `json:"owner"`
	CommittedBalance money.Money      `json:"committed_balance"`
	ReservedBalance  money.Money      `json:"reserved_balance"`
	AvailableBalance money.Money      `json:"available_balance"`
	InitialValue     money.Money      `json:"initial_value"`
	// ReplacesGiftCardCode links a refund-issued replacement to an original
	// source card that could not be reactivated because it expired or was voided.
	ReplacesGiftCardCode string                      `json:"replaces_gift_card_code,omitempty"`
	Status               wallet_enums.GiftCardStatus `json:"status"`
	IssuedAt             time.Time                   `json:"issued_at"`
	ActivatedAt          *time.Time                  `json:"activated_at,omitempty"`
	ExpiresAt            *time.Time                  `json:"expires_at,omitempty"`
	Note                 string                      `json:"note,omitempty"`
	History              []security.HistoryEntry     `json:"history,omitempty"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	audit.AuditFields
}
