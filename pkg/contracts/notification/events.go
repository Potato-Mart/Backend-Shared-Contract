package notification

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
)

// GiftCardIssuedEvent requests delivery of the email for a captured gift-card
// purchase after the stored-value instrument has been committed. IssuanceID is
// the idempotency key. ClaimCode is present only when the recipient did not yet
// have a verified retail customer account and must never be persisted outside
// the protected notification delivery record.
type GiftCardIssuedEvent struct {
	IssuanceID     string       `json:"issuance_id"`
	RecipientEmail string       `json:"recipient_email"`
	RecipientName  string       `json:"recipient_name"`
	SenderName     string       `json:"sender_name"`
	Amount         common.Money `json:"amount"`
	Message        string       `json:"message,omitempty"`
	ClaimCode      string       `json:"claim_code,omitempty"`
	Locale         string       `json:"locale,omitempty"`
	IssuedAt       time.Time    `json:"issued_at"`
}
