package courier

// DeliveryCompanyRef is a customer-safe identity/display snapshot. Code is the
// stable company key; Integration retains the api/manual mode. Code is the sole
// public company identity. Name is display-only and must never be used as a key.
// Revision freezes the company's configuration used for a selection or booking.
type DeliveryCompanyRef struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Integration string `json:"integration"`
	Revision    int64  `json:"revision"`
}
