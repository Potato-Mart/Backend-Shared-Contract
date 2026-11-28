package ads

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdSyncState says how current one platform's figures are, and in which time
// zone and currency that platform reports them.
//
// AccountTimeZone is the ad account's own IANA zone, which decides how the
// platform cuts its days; it is not always the market time zone, and the
// mismatch is surfaced as a health check rather than silently corrected.
// AccountCurrency is the currency every spend figure from this platform is
// denominated in. SyncedAt is the last successful sync, so a reader can tell
// stale numbers from missing ones. LastErrorCode is our own code for the last
// failure and never carries platform error text.
type AdSyncState struct {
	Platform        ads_enums.AdPlatform `json:"platform,omitempty"`
	SyncedAt        *time.Time           `json:"synced_at,omitempty"`
	AccountTimeZone string               `json:"account_time_zone,omitempty"`
	AccountCurrency money.CurrencyCode   `json:"account_currency,omitempty"`
	LastErrorCode   string               `json:"last_error_code,omitempty"`
}
