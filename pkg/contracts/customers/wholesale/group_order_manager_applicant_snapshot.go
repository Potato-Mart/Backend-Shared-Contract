package wholesale

type GroupOrderManagerApplicantSnapshot struct {
	UserID               string `json:"user_id"`
	RetailAccountID      string `json:"retail_account_id"`
	RetailCustomerNumber string `json:"retail_customer_number"`
	Name                 string `json:"name,omitempty"`
	Email                string `json:"email,omitempty"`
}
