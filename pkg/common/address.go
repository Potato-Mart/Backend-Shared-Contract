package common

type Address struct {
	Label    string `json:"label"`
	Line1    string `json:"line1"`
	Line2    string `json:"line2,omitempty"`
	City     string `json:"city"`
	State    string `json:"state"`
	Postcode string `json:"postcode"`
	Country  string `json:"country"`
}
