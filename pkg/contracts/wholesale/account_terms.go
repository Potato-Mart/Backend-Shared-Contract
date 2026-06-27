package wholesale

import "github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"

const (
	PathWholesaleAccountTermsMe       = "/v1/wholesale-account/me"
	PathInternalWholesaleAccountTerms = "/v1/internal/wholesale/account-terms"
)

// WholesaleAccountTerms is the buyer-facing terms projection used by
// approved wholesale storefront checkout flows.
type WholesaleAccountTerms struct {
	OrganisationID         string                 `json:"organisation_id"`
	OrganisationAccessID   string                 `json:"organisation_access_id"`
	CustomerID             string                 `json:"customer_id"`
	UserID                 string                 `json:"user_id,omitempty"`
	BusinessName           string                 `json:"business_name,omitempty"`
	TradingName            string                 `json:"trading_name,omitempty"`
	PaymentTermsLabel      string                 `json:"payment_terms_label,omitempty"`
	Terms                  WholesaleTerms         `json:"terms"`
	DefaultBillingAddress  *common.ContactAddress `json:"default_billing_address,omitempty"`
	DefaultShippingAddress *common.ContactAddress `json:"default_shipping_address,omitempty"`
	AllowedCheckoutMethods []string               `json:"allowed_checkout_methods"`
	InvoicePaymentMethods  []string               `json:"invoice_payment_methods"`
}

// InvoiceCardPaymentCommand starts/reuses a full-balance card payment for an
// owner-gated invoice. The server calculates the payable amount.
type InvoiceCardPaymentCommand struct {
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

// InvoiceCardPaymentSession mirrors the Stripe-compatible payment session
// shape returned by Commerce without exposing any client-supplied amount.
type InvoiceCardPaymentSession struct {
	Provider          string `json:"provider"`
	ProviderSessionID string `json:"provider_session_id"`
	SessionData       string `json:"session_data,omitempty"`
	ClientKey         string `json:"client_key,omitempty"`
	ExpiresAt         string `json:"expires_at,omitempty"`
	ReturnURL         string `json:"return_url,omitempty"`
	SuccessURL        string `json:"success_url,omitempty"`
	CancelURL         string `json:"cancel_url,omitempty"`
}

// InvoiceCardPaymentResult contains invoice/order references plus the
// payment session needed by a storefront client to confirm card payment.
type InvoiceCardPaymentResult struct {
	InvoiceID     string `json:"invoice_id"`
	InvoiceNumber string `json:"invoice_number"`
	OrderNumber   string `json:"order_number"`
	// Deprecated: use OrderNumber.
	OrderID string                    `json:"order_id,omitempty"`
	Total   common.Money              `json:"total"`
	Payment InvoiceCardPaymentSession `json:"payment"`
}
