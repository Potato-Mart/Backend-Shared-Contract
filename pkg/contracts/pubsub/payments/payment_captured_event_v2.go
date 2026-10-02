package payments

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/payments/payment/payment_enums"
)

// PaymentCapturedEventV2 is payment.captured event_version v2. AggregateID is
// the order number. It carries authenticated immutable timing evidence for the
// captured Amount, preserving unknown exact time across delayed webhook/replay.
// EventEnvelope.OccurredAt and receipt/processing time are not capture proof.
type PaymentCapturedEventV2 struct {
	PaymentID         string                        `json:"payment_id"`
	OrderID           string                        `json:"order_id,omitempty"`
	OrderNumber       string                        `json:"order_number"`
	Method            payment_enums.PaymentMethod   `json:"method,omitempty"`
	Amount            money.Money                   `json:"amount"`
	ProviderSessionID string                        `json:"provider_session_id,omitempty"`
	MarketCode        string                        `json:"market_code,omitempty"`
	CountryCode       geography.CountryCode         `json:"country_code,omitempty"`
	DepotCode         string                        `json:"depot_code,omitempty"`
	CaptureTiming     payment.CaptureTimingEvidence `json:"capture_timing"`
	RequestID         string                        `json:"request_id,omitempty"`
}
