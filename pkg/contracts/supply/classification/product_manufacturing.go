package classification

// ProductManufacturing contains customer-safe product manufacturing details.
// Its fields are optional so partially known declarations remain representable.
type ProductManufacturing struct {
	CompanyName string          `json:"company_name,omitempty"`
	CountryRef  *CountryCodeRef `json:"country_ref,omitempty"`
}
