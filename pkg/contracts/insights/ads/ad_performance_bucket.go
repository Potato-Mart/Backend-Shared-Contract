package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// AdPerformanceBucket is one day of a performance row, used when a read asks
// for daily granularity.
//
// Date is a plain calendar date, "YYYY-MM-DD", in the ad account time zone
// that the surrounding sync state names, which is not always the market time
// zone. Spend comes from the platform; Orders and Revenue are our own
// attributed figures, so a bucket can carry platform spend with no orders, or
// orders against a day whose spend the platform has since restated.
type AdPerformanceBucket struct {
	Date    string        `json:"date,omitempty"`
	Spend   []money.Money `json:"spend,omitempty"`
	Orders  int64         `json:"orders,omitempty"`
	Revenue []money.Money `json:"revenue,omitempty"`
}
