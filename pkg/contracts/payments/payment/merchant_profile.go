package payment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/payments/payment/payment_enums"
)

// MerchantLegalProfile is the Payments-owned issuer identity for one market,
// effective dated so a historical document always renders the identity that
// was in force when it was issued.
//
// This is a model only. Real legal name, business number, address, and
// registration values are supplied through protected execution input and never
// live in the contract.
type MerchantLegalProfile struct {
	ID         string `json:"id"`
	MarketCode string `json:"market_code"`

	LegalName   string `json:"legal_name"`
	TradingName string `json:"trading_name,omitempty"`

	BusinessNumberScheme payment_enums.BusinessNumberScheme `json:"business_number_scheme"`
	BusinessNumber       string                             `json:"business_number"`
	// BusinessNumberSchemeLabel names the register when the scheme is
	// "other".
	BusinessNumberSchemeLabel string `json:"business_number_scheme_label,omitempty"`

	TaxRegistrationStatus payment_enums.TaxRegistrationStatus `json:"tax_registration_status"`
	TaxRegistrationNumber string                              `json:"tax_registration_number,omitempty"`
	TaxRegisteredFrom     *temporal.Date                      `json:"tax_registered_from,omitempty"`
	TaxRegisteredUntil    *temporal.Date                      `json:"tax_registered_until,omitempty"`

	Address geography.Address `json:"address"`
	Contact party.Recipient   `json:"contact"`

	Status        payment_enums.MerchantProfileStatus `json:"status"`
	EffectiveFrom time.Time                           `json:"effective_from"`
	EffectiveTo   *time.Time                          `json:"effective_to,omitempty"`
	Revision      int64                               `json:"revision"`

	audit.AuditFields
}

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
