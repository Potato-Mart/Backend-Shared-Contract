package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
)

// GiftCardIssuedEvent requests delivery of the email for a captured gift-card
// purchase after the stored-value instrument has been committed. IssuanceID is
// the idempotency key. ClaimCode is present only when the recipient did not yet
// have a verified retail customer account and must never be persisted outside
// the protected notification delivery record.
//
// Amount stays the face and purchase evidence: it is what the buyer was
// charged. BonusAmountMinor reports any promotional bonus granted on top of it
// in the same currency, so the charged amount and the issued balance remain
// independently auditable. It is absent when the denomination carries no bonus,
// and absent on every event published before v27.3.0.
type GiftCardIssuedEvent struct {
	IssuanceID                string      `json:"issuance_id"`
	DenominationPolicyVersion int         `json:"denomination_policy_version,omitempty"`
	RecipientEmail            string      `json:"recipient_email"`
	RecipientName             string      `json:"recipient_name"`
	SenderName                string      `json:"sender_name"`
	Amount                    money.Money `json:"amount"`
	BonusAmountMinor          int64       `json:"bonus_amount_minor,omitempty"`
	Message                   string      `json:"message,omitempty"`
	ClaimCode                 string      `json:"claim_code,omitempty"`
	Locale                    string      `json:"locale,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	IssuedAt    time.Time             `json:"issued_at"`
}

// VoucherClaimIssuedEvent requests delivery of a claim invitation without
// placing claim material on Pub/Sub. DeliveryHandle is audience-bound to the
// Customers service, which exchanges it with Pricing immediately before send.
type VoucherClaimIssuedEvent struct {
	IssuanceID     string `json:"issuance_id"`
	DeliveryHandle string `json:"delivery_handle"`
}
