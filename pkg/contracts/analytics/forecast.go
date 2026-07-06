package analytics

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/common"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/enums/security"
)

// SKUDemandForecast holds the latest AI-generated stock-out prediction for
// a single SKU. Only one row per SKU is kept (upsert on sku); no history.
type SKUDemandForecast struct {
	SKUCode              string    `json:"sku_code"`
	ProductName          string    `json:"product_name,omitempty"`
	ComputedAt           time.Time `json:"computed_at"`
	HistoryDays          int       `json:"history_days"`
	ForecastHorizonDays  int       `json:"forecast_horizon_days"`
	PredictedDemandTotal float64   `json:"predicted_demand_total"`
	// PredictedDaily is an ordered list of per-day demand predictions
	// covering the forecast horizon.
	PredictedDaily    []DailyPrediction `json:"predicted_daily"`
	CurrentStockAtRun float64           `json:"current_stock_at_run"`
	// DaysUntilStockout is nil when no stock-out is predicted within the horizon.
	DaysUntilStockout *float64                `json:"days_until_stockout,omitempty"`
	AlertLevel        securityenum.AlertLevel `json:"alert_level"`
	Algorithm         string                  `json:"algorithm"`
	AlgorithmParams   common.Metadata         `json:"algorithm_params,omitempty"`
}

// DailyPrediction is one element in SKUDemandForecast.PredictedDaily.
type DailyPrediction struct {
	Date common.Date `json:"date"`
	Qty  float64     `json:"qty"`
}
