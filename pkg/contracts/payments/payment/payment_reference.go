package payment

type PaymentReference struct {
	Stripe *StripePaymentReference `json:"stripe,omitempty"`
	Mx51   *Mx51PaymentReference   `json:"mx51,omitempty"`
	Wallet *WalletPaymentReference `json:"wallet,omitempty"`
}
