package retail

// RetailCustomerProfileCompletion is a computed read projection describing how
// complete a retail customer's profile is.
type RetailCustomerProfileCompletion struct {
	Percent         int      `json:"percent"`
	CompletedFields []string `json:"completed_fields,omitempty"`
	MissingFields   []string `json:"missing_fields,omitempty"`
}
