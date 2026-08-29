package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

// RFIRecord is a revisioned Request for Inspection record. External submission
// state changes are append-only evidence supplied by a user; the model does not
// imply that opening a portal or downloading an artifact submitted the request.
type RFIRecord struct {
	ID       string           `json:"id"`
	Revision RevisionMetadata `json:"revision"`
	// MarketCode and CountryCode are the denormalized market and country the
	// record belongs to, carried so a geographically scoped staff query is
	// a plain indexed match.
	MarketCode             string                                     `json:"market_code,omitempty"`
	CountryCode            geography.CountryCode                      `json:"country_code,omitempty"`
	Channel                import_compliance_enums.RFIChannel         `json:"channel"`
	CurrentSubmissionState import_compliance_enums.RFISubmissionState `json:"current_submission_state"`
	QuarantineNumber       string                                     `json:"quarantine_number,omitempty"`
	AirwayBill             string                                     `json:"airway_bill,omitempty"`
	AvailableFrom          temporal.Date                              `json:"available_from"`
	RequestedDate          temporal.Date                              `json:"requested_date"`
	Comments               string                                     `json:"comments,omitempty"`
	BookingAgent           RFIBookingAgent                            `json:"booking_agent"`
	InspectionLocation     RFIInspectionLocation                      `json:"inspection_location"`
	InspectionDirection    string                                     `json:"inspection_direction,omitempty"`
	RequestedTime          import_compliance_enums.RFIRequestedTime   `json:"requested_time"`
	Overtime               bool                                       `json:"overtime"`
	EmailSubjectPrefix     string                                     `json:"email_subject_prefix,omitempty"`
	EmailBody              string                                     `json:"email_body,omitempty"`
	AttachmentMediaCodes   []string                                   `json:"attachment_media_codes,omitempty"`
	SubmissionEvents       []RFIExternalEvent                         `json:"submission_events,omitempty"`
	Evidence               []EvidenceReference                        `json:"evidence,omitempty"`
	Artifacts              []ArtifactReference                        `json:"artifacts,omitempty"`

	audit.AuditFields
}
