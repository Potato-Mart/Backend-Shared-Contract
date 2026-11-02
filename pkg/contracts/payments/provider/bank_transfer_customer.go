package provider

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/party"

// BankTransferCustomer identifies the customer a bank-transfer payment was
// taken from, as staff bound it at the till. PartyRef carries the customer id
// and the name, phone and email shown at confirmation time; CustomerNumber is
// the retail customer's business key when one is known.
type BankTransferCustomer struct {
	party.PartyRef
	CustomerNumber string `json:"customer_number,omitempty"`
}
