package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// AdPerformanceTotals is the aggregate of a performance read, and of the
// preceding comparison period when one was asked for.
//
// Money is a list with one entry per currency, because ad accounts in
// different currencies are never added together. PlatformConversions and
// PlatformConversionValue are what the platform itself reports, including
// view-through, cross-device and modelled conversions; Orders, NewCustomers
// and Revenue are our own attributed figures from paid orders. The two never
// have to agree, and a reader must not treat either as a correction of the
// other. Derived rates are left to the reader, which skips them when the
// currencies differ.
type AdPerformanceTotals struct {
	Spend                   []money.Money `json:"spend,omitempty"`
	Impressions             int64         `json:"impressions,omitempty"`
	Clicks                  int64         `json:"clicks,omitempty"`
	Installs                int64         `json:"installs,omitempty"`
	PlatformConversions     int64         `json:"platform_conversions,omitempty"`
	PlatformConversionValue []money.Money `json:"platform_conversion_value,omitempty"`
	Orders                  int64         `json:"orders,omitempty"`
	NewCustomers            int64         `json:"new_customers,omitempty"`
	Revenue                 []money.Money `json:"revenue,omitempty"`
}
