package payment

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/provider"

type PaymentReference struct {
	Stripe *provider.StripePaymentReference `json:"stripe,omitempty"`
	Mx51   *provider.Mx51PaymentReference   `json:"mx51,omitempty"`
	Wallet *provider.WalletPaymentReference `json:"wallet,omitempty"`
}
