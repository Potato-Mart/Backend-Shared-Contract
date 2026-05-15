package payments

import (
	"encoding/json"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// TerminalTransaction is one card-terminal interaction. It is the
// long-lived state behind an Adyen Terminal API request: created when
// the POS submits the request, updated from the synchronous response,
// asynchronous event notification, or transaction status request, and
// finalised when the provider outcome is known.
//
// Relationship to sales.Payment: a TerminalTransaction is the
// instrument-level record (one row per terminal interaction). A
// sales.Payment is the order-level record (one row per "money received
// against the order"). For terminal-backed payments they are usually
// 1:1; for other methods (cash, bank transfer) only Payment exists.
type TerminalTransaction struct {
	ID         string `json:"id"`
	TerminalID string `json:"terminal_id"`
	OrderID    string `json:"order_id,omitempty"`
	PaymentID  string `json:"payment_id,omitempty"`

	// ProviderTransactionID is the canonical transaction identifier
	// returned by the terminal provider. For Adyen Terminal API this is
	// POIData.POITransactionID.TransactionID, usually
	// "<tenderReference>.<pspReference>".
	ProviderTransactionID string `json:"provider_transaction_id,omitempty"`
	PSPReference          string `json:"psp_reference,omitempty"`
	TenderReference       string `json:"tender_reference,omitempty"`

	// Adyen/nexo request identifiers. ServiceID must be unique for the
	// POIID for at least 48 hours and is limited to 1-10 alphanumeric
	// characters by Adyen. SaleTransactionID is the merchant reference
	// shown in Customer Area and Adyen reports.
	ServiceID         string `json:"service_id,omitempty"`
	SaleID            string `json:"sale_id,omitempty"`
	POIID             string `json:"poi_id,omitempty"`
	SaleTransactionID string `json:"sale_transaction_id,omitempty"`

	Type           enums.TerminalTxType `json:"type"`
	Requested      Amounts              `json:"requested"`
	ReceiptOptions ReceiptOptions       `json:"receipt_options,omitempty"`

	Status          enums.TerminalTxStatus          `json:"status"`
	FinancialStatus enums.TerminalTxFinancialStatus `json:"financial_status,omitempty"`
	Result          Amounts                         `json:"result"`

	Message                    string `json:"message,omitempty"`
	ProviderResult             string `json:"provider_result,omitempty"`
	ProviderErrorCondition     string `json:"provider_error_condition,omitempty"`
	ProviderAdditionalResponse string `json:"provider_additional_response,omitempty"`
	MerchantReceipt            string `json:"merchant_receipt,omitempty"`
	CustomerReceipt            string `json:"customer_receipt,omitempty"`

	// Raw provider payloads are stored for diagnostics and future
	// compatibility. They should not be used as the source of truth for
	// business decisions once the normalized fields above are populated.
	ProviderRequest      json.RawMessage `json:"provider_request,omitempty"`
	ProviderResponse     json.RawMessage `json:"provider_response,omitempty"`
	ProviderNotification json.RawMessage `json:"provider_notification,omitempty"`
	DisplayNotification  json.RawMessage `json:"display_notification,omitempty"`

	// RecoveryDecision captures the merchant's manual answer when the
	// provider cannot determine the outcome after status checks.
	RecoveryDecision enums.RecoveryDecision `json:"recovery_decision,omitempty"`

	// IdempotencyKey lets the POS safely retry a Create call after a
	// network blip without double-creating the transaction. It is the
	// caller's responsibility to generate and persist; the contract
	// only guarantees the field exists.
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	LastStatusCheckedAt *time.Time `json:"last_status_checked_at,omitempty"`
	FinalisedAt         *time.Time `json:"finalised_at,omitempty"`

	Metadata common.Metadata `json:"metadata,omitempty"`

	common.AuditFields
}

// CreateTerminalTransactionRequest is the input the POS submits to
// start a new terminal transaction. Exactly one of the *Details pointers
// must be non-nil; the type is implied by which one is set and recorded
// explicitly on the resulting TerminalTransaction.Type.
type CreateTerminalTransactionRequest struct {
	TerminalID     string                       `json:"terminal_id"`
	OrderID        string                       `json:"order_id,omitempty"`
	PaymentID      string                       `json:"payment_id,omitempty"`
	Currency       string                       `json:"currency"`
	ConnectionMode enums.TerminalConnectionMode `json:"connection_mode,omitempty"`
	IdempotencyKey string                       `json:"idempotency_key,omitempty"`

	// Optional Adyen/nexo identifiers. Leave empty to let the service
	// generate them from its configured SaleID and request sequence.
	ServiceID         string `json:"service_id,omitempty"`
	SaleTransactionID string `json:"sale_transaction_id,omitempty"`

	ReceiptOptions     ReceiptOptions  `json:"receipt_options,omitempty"`
	TenderOptions      []string        `json:"tender_options,omitempty"`
	SaleToAcquirerData common.Metadata `json:"sale_to_acquirer_data,omitempty"`

	Purchase            *PurchaseDetails            `json:"purchase_details,omitempty"`
	Refund              *RefundDetails              `json:"refund_details,omitempty"`
	Cashout             *CashoutDetails             `json:"cashout_details,omitempty"`
	PurchaseWithCashout *PurchaseWithCashoutDetails `json:"purchase_with_cashout_details,omitempty"`
	MOTO                *MOTODetails                `json:"moto_details,omitempty"`
}

// PurchaseDetails models a standard purchase. Tip and surcharge are
// optional and MUST be passed as separate fields, never folded into
// PurchaseAmountMinor. For Adyen, the service layer converts these
// minor-unit values to AmountsReq.RequestedAmount, TipAmount, and
// provider-specific surcharge data.
//
// A nil pointer for TipMinor / SurchargeMinor means "let the terminal or
// provider configuration decide"; a *int64 set to 0 means "explicitly
// suppress this component".
type PurchaseDetails struct {
	PurchaseAmountMinor int64  `json:"purchase_amount_minor"`
	TipMinor            *int64 `json:"tip_minor,omitempty"`
	SurchargeMinor      *int64 `json:"surcharge_minor,omitempty"`
}

// RefundDetails models both Adyen referenced and unreferenced refunds.
// Referenced refunds require the original provider transaction id
// ("tenderReference.pspReference") or enough references to reconstruct
// it. Unreferenced refunds only require the amount and a card presented
// on the terminal.
type RefundDetails struct {
	RefundType                    enums.TerminalRefundType `json:"refund_type,omitempty"`
	RefundAmountMinor             int64                    `json:"refund_amount_minor"`
	OriginalProviderTransactionID string                   `json:"original_provider_transaction_id,omitempty"`
	OriginalPSPReference          string                   `json:"original_psp_reference,omitempty"`
	OriginalTenderReference       string                   `json:"original_tender_reference,omitempty"`
}

// CashoutDetails models a cashout-only transaction (customer withdraws
// cash without making a purchase). Adyen calls this cashback without a
// purchase and represents it as AmountsReq.CashBackAmount.
type CashoutDetails struct {
	CashoutAmountMinor int64  `json:"cashout_amount_minor"`
	SurchargeMinor     *int64 `json:"surcharge_minor,omitempty"`
}

// PurchaseWithCashoutDetails combines a purchase and a cashout in one
// transaction. For Adyen, RequestedAmount is the purchase plus cashout
// total and CashBackAmount carries the cashout component.
type PurchaseWithCashoutDetails struct {
	PurchaseAmountMinor int64  `json:"purchase_amount_minor"`
	CashoutAmountMinor  *int64 `json:"cashout_amount_minor,omitempty"`
	SurchargeMinor      *int64 `json:"surcharge_minor,omitempty"`
}

// MOTODetails models a Mail Order / Telephone Order card-not-present
// transaction initiated through the terminal. For Adyen Terminal API the
// service layer sends tenderOption=MOTO in SaleToAcquirerData.
type MOTODetails struct {
	MOTOAmountMinor int64  `json:"moto_amount_minor"`
	SurchargeMinor  *int64 `json:"surcharge_minor,omitempty"`
}

// CheckTerminalTransactionStatusRequest asks the backend to refresh a
// terminal transaction outcome using the provider's status mechanism,
// such as Adyen's TransactionStatusRequest.
type CheckTerminalTransactionStatusRequest struct {
	TerminalTransactionID string `json:"terminal_transaction_id"`
}

// ResolveOverrideRequest closes a transaction stuck in OverridePending
// status with the merchant's manual decision.
type ResolveOverrideRequest struct {
	TerminalTransactionID string                 `json:"terminal_transaction_id"`
	Decision              enums.RecoveryDecision `json:"decision"`
	Note                  string                 `json:"note,omitempty"`
}
