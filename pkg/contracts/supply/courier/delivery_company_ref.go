package courier

// DeliveryCompanyRef is a customer-safe identity/display snapshot. Code is the
// stable company-instance key; Integration is the open adapter code. Neither
// is a credential. Name is display text and must never be used as a routing key.
// Revision freezes the company's configuration used for a selection or booking.
type DeliveryCompanyRef struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Integration string `json:"integration"`
	Revision    int64  `json:"revision"`
}
