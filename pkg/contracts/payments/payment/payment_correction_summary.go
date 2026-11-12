package payment

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

// PaymentCorrectionSummary marks a payment whose amount or method staff
// corrected in place. Count is the number of corrections applied;
// LastCorrectedAt and LastCorrectedBy identify the latest one.
// OriginalAmount and OriginalMethod keep the values recorded before the first
// correction. Warnings lists the side effects the latest correction did not
// perform, such as an unchanged gateway charge. Per-correction detail lives in
// the payment's History.
type PaymentCorrectionSummary struct {
	Count           int                                      `json:"count"`
	LastCorrectedAt time.Time                                `json:"last_corrected_at"`
	LastCorrectedBy security.ActorRef                        `json:"last_corrected_by"`
	OriginalAmount  money.Money                              `json:"original_amount"`
	OriginalMethod  payment_enums.PaymentMethod              `json:"original_method"`
	Warnings        []payment_enums.PaymentCorrectionWarning `json:"warnings,omitempty"`
}
