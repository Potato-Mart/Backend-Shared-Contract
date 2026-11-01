package courier

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
)

// CarrierBookingAttempt records one Supply attempt to book a carrier job for
// an order. One booking covers every shipment of the order, listed in
// ShipmentIDs; a retry is a new attempt. DeliveryCompany freezes the company
// and, through its revision, the configuration the request was built from.
//
// CarrierReference is the carrier's job reference, set when Outcome is booked
// or already_booked. ErrorCode, ErrorFields and ErrorMessage describe a
// rejected attempt: ErrorFields names the request fields the carrier
// rejected, and ErrorMessage is redacted, staff-facing text. RequestSummary
// holds redacted, non-secret facts about the request, never credentials or
// the raw provider payload. OverrideReason records why staff booked a carrier
// other than the one on the accepted delivery selection. Edits holds the
// values staff changed before this attempt. Sandbox marks an attempt sent to
// the carrier's test environment; it never books a real job.
//
// ResolvedAt is set when the attempt leaves pending. AuditFields track the
// record's own creation and updates. The record may contain personal data in
// edits, error text and the request summary; Supply labels it through
// DataProtectionFields and owns retention.
type CarrierBookingAttempt struct {
	ID               string                                `json:"id"`
	OrderNumber      string                                `json:"order_number"`
	ShipmentIDs      []string                              `json:"shipment_ids"`
	DeliveryCompany  DeliveryCompanyRef                    `json:"delivery_company"`
	Outcome          courier_enums.CarrierBookingOutcome   `json:"outcome"`
	CarrierReference string                                `json:"carrier_reference,omitempty"`
	RequestSummary   []CarrierBookingRequestSummaryEntry   `json:"request_summary,omitempty"`
	ErrorCode        courier_enums.CarrierBookingErrorCode `json:"error_code,omitempty"`
	ErrorFields      []string                              `json:"error_fields,omitempty"`
	ErrorMessage     string                                `json:"error_message,omitempty"`
	OverrideReason   string                                `json:"override_reason,omitempty"`
	Edits            *CarrierBookingEdits                  `json:"edits,omitempty"`
	Sandbox          bool                                  `json:"sandbox"`
	AttemptedBy      string                                `json:"attempted_by"`
	AttemptedAt      time.Time                             `json:"attempted_at"`
	ResolvedAt       *time.Time                            `json:"resolved_at,omitempty"`
	audit.AuditFields
	security.DataProtectionFields
}
