package common

// PartyRef is a lightweight reference to a person or organisation involved
// in a transaction (order customer, supplier, etc.). It carries just enough
// to display and contact the party without joining the full record.
type PartyRef struct {
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}
