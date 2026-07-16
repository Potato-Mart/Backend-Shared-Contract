package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
)

// ReceiptOffer is the buyer/POS-safe projection of one active promotion. It
// intentionally omits internal rule details, discount configuration, usage
// counters, source metadata, and authoring copy.
type ReceiptOffer struct {
	ID              string                 `json:"id"`
	ReceiptMessages []common.LocalizedName `json:"receipt_messages"`
	StartsAt        *time.Time             `json:"starts_at,omitempty"`
	ExpiresAt       *time.Time             `json:"expires_at,omitempty"`
	Priority        int                    `json:"priority"`
}
