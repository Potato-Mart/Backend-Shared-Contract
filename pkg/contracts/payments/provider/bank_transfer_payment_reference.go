package provider

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

// BankTransferPaymentReference records the staff confirmation of a payment
// the customer made by bank transfer at the till.
//
// TransferAmount is the amount the payment settles, which is the full amount
// due when it was recorded. ConfirmationStatus records whether staff saw the
// full amount on the evidence (confirmed_full) or only part of it
// (confirmed_partial); a confirmed_partial payment is still completed for the
// full amount and is flagged for finance follow-up. Evidence lists the
// finalized private evidence objects attached in the same write as the
// payment. ConfirmedBy is the user id of the staff member who confirmed it.
type BankTransferPaymentReference struct {
	PayerReference     string                                       `json:"payer_reference,omitempty"`
	Customer           *BankTransferCustomer                        `json:"customer,omitempty"`
	TransferAmount     money.Money                                  `json:"transfer_amount"`
	ConfirmationStatus payment_enums.BankTransferConfirmationStatus `json:"confirmation_status"`
	Evidence           []BankTransferEvidence                       `json:"evidence,omitempty"`
	ConfirmedAt        time.Time                                    `json:"confirmed_at"`
	ConfirmedBy        string                                       `json:"confirmed_by"`
}
