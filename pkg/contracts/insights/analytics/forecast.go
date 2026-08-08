package analytics

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/commerce/commerce_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/temporal"
)

// SKUDemandForecast is one depot- and channel-qualified demand prediction.
type SKUDemandForecast struct {
	ProductSKUCode       string                   `json:"product_sku_code"`
	DepotCode            string                   `json:"depot_code"`
	Channel              commerce_enums.OrderType `json:"channel"`
	Timezone             string                   `json:"timezone"`
	ProductName          string                   `json:"product_name,omitempty"`
	ComputedAt           time.Time                `json:"computed_at"`
	HistoryDays          int                      `json:"history_days"`
	ForecastHorizonDays  int                      `json:"forecast_horizon_days"`
	PredictedDemandTotal float64                  `json:"predicted_demand_total"`
	// PredictedDaily is an ordered list of per-day demand predictions
	// covering the forecast horizon.
	PredictedDaily    []DailyPrediction            `json:"predicted_daily"`
	AvailabilityAtRun ForecastAvailabilitySnapshot `json:"availability_at_run"`
	// DaysUntilStockout is nil when no stock-out is predicted within the horizon.
	DaysUntilStockout *float64                  `json:"days_until_stockout,omitempty"`
	AlertLevel        security_enums.AlertLevel `json:"alert_level"`
	Algorithm         string                    `json:"algorithm"`
	AlgorithmParams   metadata.Metadata         `json:"algorithm_params,omitempty"`
}

// ForecastAvailabilitySnapshot freezes the quantities used by a forecast.
type ForecastAvailabilitySnapshot struct {
	SellableAvailableBaseUnits int64     `json:"sellable_available_base_units"`
	ReservedBaseUnits          int64     `json:"reserved_base_units"`
	StagedBaseUnits            int64     `json:"staged_base_units"`
	QualityHoldBaseUnits       int64     `json:"quality_hold_base_units"`
	Revision                   int64     `json:"revision"`
	AsOf                       time.Time `json:"as_of"`
}

// DailyPrediction is one element in SKUDemandForecast.PredictedDaily.
type DailyPrediction struct {
	Date temporal.Date `json:"date"`
	Qty  float64       `json:"qty"`
}
