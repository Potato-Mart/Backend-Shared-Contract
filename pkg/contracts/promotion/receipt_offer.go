package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
)

// ReceiptOffer is the buyer/POS-safe projection of one active promotion. It
// intentionally omits internal rule details, discount configuration, usage
// counters, source metadata, and authoring copy.
type ReceiptOffer struct {
	ID                string                   `json:"id"`
	SeriesKey         string                   `json:"series_key"`
	ReceiptMessages   []common.LocalizedName   `json:"receipt_messages"`
	StartsAt          *time.Time               `json:"starts_at,omitempty"`
	ExpiresAt         *time.Time               `json:"expires_at,omitempty"`
	ScheduleTimezone  string                   `json:"schedule_timezone"`
	GeographicContext common.GeographicContext `json:"geographic_context"`
	Priority          int                      `json:"priority"`
}
