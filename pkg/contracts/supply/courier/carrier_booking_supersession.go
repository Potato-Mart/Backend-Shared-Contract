package courier

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
)

// CarrierBookingSupersession records that staff switched an order to another
// carrier after this booking attempt, so the attempt no longer represents the
// order's carrier job. The attempt keeps its own outcome.
//
// At and By identify the switch, and Reason is the staff reason.
// ReplacementCompany is the company switched to, if one was chosen.
// CancellationOutcome is the result of the automatic cancel request for this
// booking; ManualActionID names the CarrierManualAction opened when that
// request was not configured or did not succeed.
type CarrierBookingSupersession struct {
	At                  time.Time                                `json:"at"`
	By                  security.ActorRef                        `json:"by"`
	Reason              string                                   `json:"reason,omitempty"`
	ReplacementCompany  *DeliveryCompanyRef                      `json:"replacement_company,omitempty"`
	CancellationOutcome courier_enums.CarrierCancellationOutcome `json:"cancellation_outcome"`
	ManualActionID      string                                   `json:"manual_action_id,omitempty"`
}
