package analytics

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"time"
)

// MetricRollup is one hourly or daily aggregate consumed by Admin reports.
//
// MarketCode and CountryCode are the denormalized geography the aggregate is
// attributed to. A rollup carrying neither is platform-wide and is readable
// only at global scope; a geographically scoped principal filters on these
// two fields and a rollup that does not match is not returned. Both are
// absent on rollups computed before v28.0.0, and an absent value means
// "unattributed" rather than "everywhere".
type MetricRollup struct {
	Metric       string                `json:"metric"`
	Granularity  string                `json:"granularity"`
	WindowStart  time.Time             `json:"window_start"`
	WindowEnd    time.Time             `json:"window_end"`
	Dimension    string                `json:"dimension,omitempty"`
	MarketCode   string                `json:"market_code,omitempty"`
	CountryCode  geography.CountryCode `json:"country_code,omitempty"`
	Count        int64                 `json:"count"`
	Amount       money.Money           `json:"amount"`
	CalculatedAt time.Time             `json:"calculated_at"`
}
