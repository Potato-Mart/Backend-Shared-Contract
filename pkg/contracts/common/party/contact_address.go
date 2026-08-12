package party

import "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography"

// ContactAddress pairs an optional contact with an optional postal address.
// It is the shared shape for shipping/billing targets and saved address
// books across customer, order, and fulfilment contracts.
type ContactAddress struct {
	// ID is populated for persisted address-book entries. Inline checkout,
	// order, billing, and fulfilment snapshots may leave it empty.
	ID      string             `json:"id,omitempty"`
	Contact *Recipient         `json:"contact,omitempty"`
	Address *geography.Address `json:"address,omitempty"`
}
