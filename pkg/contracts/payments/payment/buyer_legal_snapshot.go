package payment

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/payments/payment/payment_enums"
	"time"
)

// BuyerLegalSnapshot is the buyer identity frozen onto a document. A market
// may require it above a threshold, for example an Australian tax invoice for
// a total of at least one thousand dollars.
type BuyerLegalSnapshot struct {
	Name                      string                             `json:"name"`
	BusinessNumberScheme      payment_enums.BusinessNumberScheme `json:"business_number_scheme,omitempty"`
	BusinessNumber            string                             `json:"business_number,omitempty"`
	BusinessNumberSchemeLabel string                             `json:"business_number_scheme_label,omitempty"`
	Address                   *geography.Address                 `json:"address,omitempty"`
	FrozenAt                  time.Time                          `json:"frozen_at"`
}
