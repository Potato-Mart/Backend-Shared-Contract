package courier

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/temporal"
)

// DeliveryServiceWindow is the legacy configured recurring local-time window,
// not a provider availability promise. It is retained for compatibility and is
// deprecated for new courier availability configuration. DaysOfWeek uses
// 0=Sunday through 6=Saturday;
// an empty list supplies no windows. Times are HH:MM in the IANA Timezone.
// StartTime must precede EndTime on the same day. Empty ServiceAreaCodes applies
// to all enabled areas in CountryCode. ExcludedDates are local calendar dates.
// Effective dates are inclusive. MinimumLeadTimeMinutes is non-negative elapsed
// time before the window start. Absent Fee is unknown/inherited, never free.
type DeliveryServiceWindow struct {
	Code                   string                `json:"code"`
	CountryCode            geography.CountryCode `json:"country_code"`
	ServiceAreaCodes       []string              `json:"service_area_codes,omitempty"`
	DaysOfWeek             []int                 `json:"days_of_week"`
	StartTime              temporal.TimeOfDay    `json:"start_time"`
	EndTime                temporal.TimeOfDay    `json:"end_time"`
	Timezone               string                `json:"timezone"`
	Label                  string                `json:"label,omitempty"`
	ExcludedDates          []temporal.Date       `json:"excluded_dates,omitempty"`
	EffectiveFrom          *temporal.Date        `json:"effective_from,omitempty"`
	EffectiveUntil         *temporal.Date        `json:"effective_until,omitempty"`
	MinimumLeadTimeMinutes int                   `json:"minimum_lead_time_minutes"`
	Fee                    *money.Money          `json:"fee,omitempty"`
	Enabled                bool                  `json:"enabled"`
}
