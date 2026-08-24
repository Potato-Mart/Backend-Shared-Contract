package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/wallet/wallet_enums"
)

// PointBucket is an expiry-aware available points batch.
type PointBucket struct {
	Points              int                            `json:"points"`
	ExpiresAt           *time.Time                     `json:"expires_at,omitempty"`
	SourceLedgerEntryID string                         `json:"source_ledger_entry_id"`
	Reason              wallet_enums.PointLedgerReason `json:"reason"`
	RelatedOrderNumber  string                         `json:"related_order_number,omitempty"`
}
