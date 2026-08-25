package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
)

// GiftCardIssuedEvent records the stored-value issuance fact after a captured
// gift-card purchase. Notification services resolve recipient, claim, locale,
// and rendered content from protected service-owned data; this event never
// carries delivery material or customer contact data.
//
// Amount stays the face and purchase evidence: it is what the buyer was
// charged. BonusAmountMinor reports any promotional bonus granted on top of it
// in the same currency, so the charged amount and the issued balance remain
// independently auditable. It is absent when the denomination carries no bonus.
type GiftCardIssuedEvent struct {
	IssuanceID                string      `json:"issuance_id"`
	DenominationPolicyVersion int         `json:"denomination_policy_version,omitempty"`
	Amount                    money.Money `json:"amount"`
	BonusAmountMinor          int64       `json:"bonus_amount_minor,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. A consumer that persists a geographically scoped record
	// treats an absent value as "no evidence" and fails closed rather than
	// defaulting.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	IssuedAt    time.Time             `json:"issued_at"`
}
