package common

type Recipient struct {
	Name  string `json:"name"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}
