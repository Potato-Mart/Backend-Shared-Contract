package retail

// RetailCustomerManagementProfile groups CRM fields that are manually edited
// by staff and are never overwritten by sync jobs.
type RetailCustomerManagementProfile struct {
	Notes    string   `json:"notes,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	SalesRep string   `json:"sales_rep,omitempty"`
	CRMTier  string   `json:"crm_tier,omitempty"`
}
