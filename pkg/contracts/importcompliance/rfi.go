package importcompliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
	importcomplianceenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/importcompliance"
)

// RFIRecord is a revisioned Request for Inspection record. External submission
// state changes are append-only evidence supplied by a user; the model does not
// imply that opening a portal or downloading an artifact submitted the request.
type RFIRecord struct {
	ID                     string                                  `json:"id"`
	Revision               RevisionMetadata                        `json:"revision"`
	Channel                importcomplianceenum.RFIChannel         `json:"channel"`
	CurrentSubmissionState importcomplianceenum.RFISubmissionState `json:"current_submission_state"`
	QuarantineNumber       string                                  `json:"quarantine_number,omitempty"`
	AirwayBill             string                                  `json:"airway_bill,omitempty"`
	AvailableFrom          common.Date                             `json:"available_from"`
	RequestedDate          common.Date                             `json:"requested_date"`
	Comments               string                                  `json:"comments,omitempty"`
	BookingAgent           RFIBookingAgent                         `json:"booking_agent"`
	InspectionLocation     RFIInspectionLocation                   `json:"inspection_location"`
	InspectionDirection    string                                  `json:"inspection_direction,omitempty"`
	RequestedTime          importcomplianceenum.RFIRequestedTime   `json:"requested_time"`
	Overtime               bool                                    `json:"overtime"`
	EmailSubjectPrefix     string                                  `json:"email_subject_prefix,omitempty"`
	EmailBody              string                                  `json:"email_body,omitempty"`
	AttachmentMediaIDs     []string                                `json:"attachment_media_ids,omitempty"`
	SubmissionEvents       []RFIExternalEvent                      `json:"submission_events,omitempty"`
	Evidence               []EvidenceReference                     `json:"evidence,omitempty"`
	Artifacts              []ArtifactReference                     `json:"artifacts,omitempty"`

	common.AuditFields
}

type RFIBookingAgent struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

type RFIInspectionLocation struct {
	BusinessNameAndAANumber string `json:"business_name_and_aa_number,omitempty"`
	PremiseAddress          string `json:"premise_address,omitempty"`
	OpeningHours            string `json:"opening_hours,omitempty"`
	ContactName             string `json:"contact_name,omitempty"`
	ContactPhone            string `json:"contact_phone,omitempty"`
	PrivateResidence        bool   `json:"private_residence"`
}

type RFIExternalEvent struct {
	ID                string                                  `json:"id"`
	State             importcomplianceenum.RFISubmissionState `json:"state"`
	ExternalReference string                                  `json:"external_reference,omitempty"`
	OccurredAt        time.Time                               `json:"occurred_at"`
	RecordedBy        string                                  `json:"recorded_by,omitempty"`
	Note              string                                  `json:"note,omitempty"`
	Evidence          []EvidenceReference                     `json:"evidence,omitempty"`
}
