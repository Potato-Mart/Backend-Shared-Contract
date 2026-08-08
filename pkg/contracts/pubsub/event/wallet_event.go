package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/money"
)

// GiftCardIssuedEvent requests delivery of the email for a captured gift-card
// purchase after the stored-value instrument has been committed. IssuanceID is
// the idempotency key. ClaimCode is present only when the recipient did not yet
// have a verified retail customer account and must never be persisted outside
// the protected notification delivery record.
type GiftCardIssuedEvent struct {
	IssuanceID                string      `json:"issuance_id"`
	DenominationPolicyVersion int         `json:"denomination_policy_version,omitempty"`
	RecipientEmail            string      `json:"recipient_email"`
	RecipientName             string      `json:"recipient_name"`
	SenderName                string      `json:"sender_name"`
	Amount                    money.Money `json:"amount"`
	Message                   string      `json:"message,omitempty"`
	ClaimCode                 string      `json:"claim_code,omitempty"`
	Locale                    string      `json:"locale,omitempty"`
	IssuedAt                  time.Time   `json:"issued_at"`
}

// VoucherClaimIssuedEvent requests delivery of a claim invitation without
// placing claim material on Pub/Sub. DeliveryHandle is audience-bound to the
// Customers service, which exchanges it with Pricing immediately before send.
type VoucherClaimIssuedEvent struct {
	IssuanceID     string `json:"issuance_id"`
	DeliveryHandle string `json:"delivery_handle"`
}
