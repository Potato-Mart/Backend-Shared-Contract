package pricing

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/wallet/wallet_enums"
	"time"
)

// GiftCardIssuedEventV2 is wallet.gift_card_issued event_version v2, published
// by Pricing on customer-events after a committed purchase, membership reward
// or refund-replacement issuance. AggregateID is the immutable IssuanceID.
// Notification resolves protected delivery material through an authorized
// issuance lookup, validating source, original value and owning geography.
// Card/claim codes, PINs, customer identity and contact data never enter this fact.
type GiftCardIssuedEventV2 struct {
	IssuanceID string                              `json:"issuance_id"`
	Source     wallet_enums.GiftCardIssuanceSource `json:"source"`
	// IssuedValue is the original total credited initial value in its currency,
	// including a purchase bonus when applicable. It is never the buyer charge,
	// current spendable balance, refund payment amount or denomination evidence.
	IssuedValue money.Money `json:"issued_value"`
	// An absent owning market/country means no evidence; scoped consumers fail
	// closed rather than inventing geography from the recipient or source.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	IssuedAt    time.Time             `json:"issued_at"`
}
