package analytics

import "time"

// ForecastAvailabilitySnapshot freezes the quantities used by a forecast.
type ForecastAvailabilitySnapshot struct {
	SellableAvailableBaseUnits int64     `json:"sellable_available_base_units"`
	ReservedBaseUnits          int64     `json:"reserved_base_units"`
	StagedBaseUnits            int64     `json:"staged_base_units"`
	QualityHoldBaseUnits       int64     `json:"quality_hold_base_units"`
	Revision                   int64     `json:"revision"`
	AsOf                       time.Time `json:"as_of"`
}
