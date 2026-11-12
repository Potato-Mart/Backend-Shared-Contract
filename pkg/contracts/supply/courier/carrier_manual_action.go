package courier

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
)

// CarrierManualAction records carrier-side work staff must do by hand. Supply
// writes one when a carrier switch supersedes a booking it could not cancel
// automatically; while it is open, the order's shipments show a
// carrier_cancellation_required warning.
//
// DeliveryCompany and CarrierReference identify the booking to cancel, and
// SourceAttemptID names the superseded CarrierBookingAttempt. ShipmentIDs
// lists the order's shipments the booking covered. DepotCode, MarketCode and
// CountryCode are the dispatching depot's scope, denormalized so a
// geographically scoped staff query is a plain indexed match.
//
// Cause is the automatic cancellation outcome that required the action.
// ErrorCode and ErrorMessage describe a failed cancel request; ErrorMessage is
// redacted, staff-facing text. SwitchReason is the staff reason for the
// carrier switch, and ReplacementCompany the company switched to, if any.
//
// ResolvedAt, ResolvedBy and ResolutionNote are set when staff resolve the
// action. AuditFields track the record's own creation and updates.
type CarrierManualAction struct {
	ID                 string                                   `json:"id"`
	OrderNumber        string                                   `json:"order_number"`
	ShipmentIDs        []string                                 `json:"shipment_ids,omitempty"`
	DepotCode          string                                   `json:"depot_code"`
	MarketCode         string                                   `json:"market_code,omitempty"`
	CountryCode        geography.CountryCode                    `json:"country_code,omitempty"`
	Kind               courier_enums.CarrierManualActionKind    `json:"kind"`
	Status             courier_enums.CarrierManualActionStatus  `json:"status"`
	DeliveryCompany    DeliveryCompanyRef                       `json:"delivery_company"`
	CarrierReference   string                                   `json:"carrier_reference,omitempty"`
	SourceAttemptID    string                                   `json:"source_attempt_id,omitempty"`
	Cause              courier_enums.CarrierCancellationOutcome `json:"cause"`
	ErrorCode          courier_enums.CarrierBookingErrorCode    `json:"error_code,omitempty"`
	ErrorMessage       string                                   `json:"error_message,omitempty"`
	SwitchReason       string                                   `json:"switch_reason,omitempty"`
	ReplacementCompany *DeliveryCompanyRef                      `json:"replacement_company,omitempty"`
	ResolvedAt         *time.Time                               `json:"resolved_at,omitempty"`
	ResolvedBy         *security.ActorRef                       `json:"resolved_by,omitempty"`
	ResolutionNote     string                                   `json:"resolution_note,omitempty"`
	audit.AuditFields
}
