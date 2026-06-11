package common

// ContactAddress pairs an optional contact with an optional postal address.
// It is the shared shape for shipping/billing targets and saved address
// books across customer, order, and fulfilment contracts.
type ContactAddress struct {
	Contact *Recipient `json:"contact,omitempty"`
	Address *Address   `json:"address,omitempty"`
}
