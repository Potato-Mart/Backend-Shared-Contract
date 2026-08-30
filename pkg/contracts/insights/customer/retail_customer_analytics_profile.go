package customer

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/insights/analytics/analytics_enums"
)

// RetailCustomerAnalyticsProfile is the standalone RFM analysis calculated
// for a retail customer. Customer profiles do not embed this derived record.
type RetailCustomerAnalyticsProfile struct {
	CustomerNumber string                    `json:"customer_number"`
	RecencyDays    *int                      `json:"recency_days,omitempty"`
	R              *int                      `json:"r,omitempty"`
	F              *int                      `json:"f,omitempty"`
	M              *int                      `json:"m,omitempty"`
	Score          string                    `json:"score,omitempty"`
	Segment        string                    `json:"segment,omitempty"`
	ChurnRisk      analytics_enums.ChurnRisk `json:"churn_risk,omitempty"`
	AvgRepeatDays  *float64                  `json:"avg_repeat_days,omitempty"`
	CalculatedAt   time.Time                 `json:"calculated_at"`
}
