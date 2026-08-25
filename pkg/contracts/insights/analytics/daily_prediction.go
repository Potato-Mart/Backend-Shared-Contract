package analytics

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/temporal"

// DailyPrediction is one element in SKUDemandForecast.PredictedDaily.
type DailyPrediction struct {
	Date temporal.Date `json:"date"`
	Qty  float64       `json:"qty"`
}
