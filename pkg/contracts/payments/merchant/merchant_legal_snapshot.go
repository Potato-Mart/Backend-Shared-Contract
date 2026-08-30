package merchant

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/payment/payment_enums"
	"time"
)

// MerchantLegalSnapshot is the issuer identity frozen onto one document at
// issuance. A later profile edit never changes an issued document.
type MerchantLegalSnapshot struct {
	ProfileID       string `json:"profile_id"`
	ProfileRevision int64  `json:"profile_revision"`
	MarketCode      string `json:"market_code"`

	LegalName                 string                             `json:"legal_name"`
	TradingName               string                             `json:"trading_name,omitempty"`
	BusinessNumberScheme      payment_enums.BusinessNumberScheme `json:"business_number_scheme"`
	BusinessNumber            string                             `json:"business_number"`
	BusinessNumberSchemeLabel string                             `json:"business_number_scheme_label,omitempty"`

	TaxRegistrationStatus payment_enums.TaxRegistrationStatus `json:"tax_registration_status"`
	TaxRegistrationNumber string                              `json:"tax_registration_number,omitempty"`

	Address geography.Address `json:"address"`
	Contact party.Recipient   `json:"contact"`

	FrozenAt time.Time `json:"frozen_at"`
}
