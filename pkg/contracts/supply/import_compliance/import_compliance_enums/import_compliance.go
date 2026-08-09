// Package import_compliance_enums contains stable wire values used by reusable
// import-compliance records. Workflow policy remains owned by the consuming
// backend.
package import_compliance_enums

type ReviewState string

const (
	ReviewStateDraft    ReviewState = "draft"
	ReviewStateInReview ReviewState = "in_review"
	ReviewStateApproved ReviewState = "approved"
	ReviewStateRejected ReviewState = "rejected"
	ReviewStateArchived ReviewState = "archived"
)

func (s ReviewState) IsValid() bool {
	switch s {
	case ReviewStateDraft, ReviewStateInReview, ReviewStateApproved, ReviewStateRejected, ReviewStateArchived:
		return true
	}
	return false
}

func (s ReviewState) String() string { return string(s) }

type Jurisdiction string

const (
	JurisdictionAustralia Jurisdiction = "AU"
	JurisdictionTaiwan    Jurisdiction = "TW"
)

func (j Jurisdiction) IsValid() bool {
	switch j {
	case JurisdictionAustralia, JurisdictionTaiwan:
		return true
	}
	return false
}

func (j Jurisdiction) String() string { return string(j) }

type ImportMode string

const (
	ImportModeAirCargo   ImportMode = "air_cargo"
	ImportModeSeaAmbient ImportMode = "sea_ambient"
	ImportModeSeaFrozen  ImportMode = "sea_frozen"
)

func (m ImportMode) IsValid() bool {
	switch m {
	case ImportModeAirCargo, ImportModeSeaAmbient, ImportModeSeaFrozen:
		return true
	}
	return false
}

func (m ImportMode) String() string { return string(m) }

type EvidenceKind string

const (
	EvidenceKindMedia              EvidenceKind = "media"
	EvidenceKindOfficialSource     EvidenceKind = "official_source"
	EvidenceKindInternalRecord     EvidenceKind = "internal_record"
	EvidenceKindManualCitation     EvidenceKind = "manual_citation"
	EvidenceKindAutomationProposal EvidenceKind = "automation_proposal"
)

func (k EvidenceKind) IsValid() bool {
	switch k {
	case EvidenceKindMedia, EvidenceKindOfficialSource, EvidenceKindInternalRecord,
		EvidenceKindManualCitation, EvidenceKindAutomationProposal:
		return true
	}
	return false
}

func (k EvidenceKind) String() string { return string(k) }

type RFIChannel string

const (
	RFIChannelBiosecurityPortal RFIChannel = "biosecurity_portal"
	RFIChannelEmailException    RFIChannel = "email_exception"
)

func (c RFIChannel) IsValid() bool {
	switch c {
	case RFIChannelBiosecurityPortal, RFIChannelEmailException:
		return true
	}
	return false
}

func (c RFIChannel) String() string { return string(c) }

type RFISubmissionState string

const (
	RFISubmissionStateNotSubmitted RFISubmissionState = "not_submitted"
	RFISubmissionStateSubmitted    RFISubmissionState = "submitted"
	RFISubmissionStateConfirmed    RFISubmissionState = "confirmed"
	RFISubmissionStateRescheduled  RFISubmissionState = "rescheduled"
	RFISubmissionStateCancelled    RFISubmissionState = "cancelled"
	RFISubmissionStateClosed       RFISubmissionState = "closed"
)

func (s RFISubmissionState) IsValid() bool {
	switch s {
	case RFISubmissionStateNotSubmitted, RFISubmissionStateSubmitted,
		RFISubmissionStateConfirmed, RFISubmissionStateRescheduled,
		RFISubmissionStateCancelled, RFISubmissionStateClosed:
		return true
	}
	return false
}

func (s RFISubmissionState) String() string { return string(s) }

type RFIRequestedTime string

const (
	RFIRequestedTimeAnytime RFIRequestedTime = "anytime"
	RFIRequestedTimeAM      RFIRequestedTime = "am"
	RFIRequestedTimePM      RFIRequestedTime = "pm"
)

func (t RFIRequestedTime) IsValid() bool {
	switch t {
	case RFIRequestedTimeAnytime, RFIRequestedTimeAM, RFIRequestedTimePM:
		return true
	}
	return false
}

func (t RFIRequestedTime) String() string { return string(t) }

type LabelSize string

const (
	LabelSize100x60 LabelSize = "100x60"
	LabelSize100x80 LabelSize = "100x80"
	LabelSize80x50  LabelSize = "80x50"
	LabelSize65x45  LabelSize = "65x45"
	LabelSize60x40  LabelSize = "60x40"
	LabelSize50x30  LabelSize = "50x30"
)

func (s LabelSize) IsValid() bool {
	switch s {
	case LabelSize100x60, LabelSize100x80, LabelSize80x50,
		LabelSize65x45, LabelSize60x40, LabelSize50x30:
		return true
	}
	return false
}

func (s LabelSize) String() string { return string(s) }

type LabelOrientation string

const (
	LabelOrientationPortrait  LabelOrientation = "portrait"
	LabelOrientationLandscape LabelOrientation = "landscape"
)

func (o LabelOrientation) IsValid() bool {
	switch o {
	case LabelOrientationPortrait, LabelOrientationLandscape:
		return true
	}
	return false
}

func (o LabelOrientation) String() string { return string(o) }

type ArtifactKind string

const (
	ArtifactKindPDF  ArtifactKind = "pdf"
	ArtifactKindDOCX ArtifactKind = "docx"
	ArtifactKindXLSX ArtifactKind = "xlsx"
	ArtifactKindZIP  ArtifactKind = "zip"
	ArtifactKindEML  ArtifactKind = "eml"
)

func (k ArtifactKind) IsValid() bool {
	switch k {
	case ArtifactKindPDF, ArtifactKindDOCX, ArtifactKindXLSX, ArtifactKindZIP, ArtifactKindEML:
		return true
	}
	return false
}

func (k ArtifactKind) String() string { return string(k) }
