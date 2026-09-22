package courier

// DeliveryCompanyRef is a customer-safe identity/display snapshot. Code is the
// stable company-instance key; Integration retains the api/manual mode. Adapter
// identifies the separate server-side adapter (for example detrack or bcrc).
// Neither is a credential. Name must never be used as a routing key.
// Revision freezes the company's configuration used for a selection or booking.
type DeliveryCompanyRef struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Integration string `json:"integration"`
	Adapter     string `json:"adapter,omitempty"`
	Revision    int64  `json:"revision"`
}
