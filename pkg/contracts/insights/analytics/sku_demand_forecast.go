package analytics

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security/security_enums"
	"time"
)

// SKUDemandForecast is one depot- and channel-qualified demand prediction.
type SKUDemandForecast struct {
	SKUCode    string `json:"sku_code"`
	MarketCode string `json:"market_code"`
	DepotCode  string `json:"depot_code"`
	// CountryCode is the denormalized country the forecast is attributed to,
	// so a country-scoped principal is filtered by a plain indexed match. An
	// empty value provides no geographic evidence and is handled fail-closed
	// by the consumer.
	CountryCode          geography.CountryCode    `json:"country_code,omitempty"`
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
