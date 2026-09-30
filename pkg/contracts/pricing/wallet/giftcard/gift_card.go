package giftcard

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/pricing/benefit"

	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/pricing/wallet/wallet_enums"
)

// GiftCard is a stored-value instrument with a re-spendable balance. The
// GiftCardTransaction ledger is the source of truth for CommittedBalance. Live
// checkout reservations are summarized by ReservedBalance and subtracted to
// produce AvailableBalance. It is referenced everywhere by Code (the business
// key), never by ID. ID remains the stable internal identity. Code is the single
// customer-facing claim and POS code: GC followed by exactly 12 decimal digits.
// PIN verification belongs to the service: exactly four decimal digits, including
// leading zeros. Neither plaintext PINs nor verifier material belong in this model.
type GiftCard struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	// CustomerNumber records customer binding after claim or POS redemption.
	// The service binds the order customer; owned online use needs no PIN.
	CustomerNumber   string           `json:"customer_number,omitempty"`
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
