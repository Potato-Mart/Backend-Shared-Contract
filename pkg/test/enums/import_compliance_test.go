package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

func TestImportComplianceEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "import_compliance_enums.ReviewState", valid: []stringEnum{import_compliance_enums.ReviewStateDraft, import_compliance_enums.ReviewStateInReview, import_compliance_enums.ReviewStateApproved, import_compliance_enums.ReviewStateRejected, import_compliance_enums.ReviewStateArchived}, invalid: import_compliance_enums.ReviewState("__invalid__")},
		{name: "import_compliance_enums.Jurisdiction", valid: []stringEnum{import_compliance_enums.JurisdictionAustralia, import_compliance_enums.JurisdictionTaiwan}, invalid: import_compliance_enums.Jurisdiction("__invalid__")},
		{name: "import_compliance_enums.ImportMode", valid: []stringEnum{import_compliance_enums.ImportModeAirCargo, import_compliance_enums.ImportModeSeaAmbient, import_compliance_enums.ImportModeSeaFrozen}, invalid: import_compliance_enums.ImportMode("__invalid__")},
		{name: "import_compliance_enums.EvidenceKind", valid: []stringEnum{import_compliance_enums.EvidenceKindMedia, import_compliance_enums.EvidenceKindOfficialSource, import_compliance_enums.EvidenceKindInternalRecord, import_compliance_enums.EvidenceKindManualCitation, import_compliance_enums.EvidenceKindAutomationProposal}, invalid: import_compliance_enums.EvidenceKind("__invalid__")},
		{name: "import_compliance_enums.RFIChannel", valid: []stringEnum{import_compliance_enums.RFIChannelBiosecurityPortal, import_compliance_enums.RFIChannelEmailException}, invalid: import_compliance_enums.RFIChannel("__invalid__")},
		{name: "import_compliance_enums.RFISubmissionState", valid: []stringEnum{import_compliance_enums.RFISubmissionStateNotSubmitted, import_compliance_enums.RFISubmissionStateSubmitted, import_compliance_enums.RFISubmissionStateConfirmed, import_compliance_enums.RFISubmissionStateRescheduled, import_compliance_enums.RFISubmissionStateCancelled, import_compliance_enums.RFISubmissionStateClosed}, invalid: import_compliance_enums.RFISubmissionState("__invalid__")},
		{name: "import_compliance_enums.RFIRequestedTime", valid: []stringEnum{import_compliance_enums.RFIRequestedTimeAnytime, import_compliance_enums.RFIRequestedTimeAM, import_compliance_enums.RFIRequestedTimePM}, invalid: import_compliance_enums.RFIRequestedTime("__invalid__")},
		{name: "import_compliance_enums.LabelSize", valid: []stringEnum{import_compliance_enums.LabelSize100x60, import_compliance_enums.LabelSize100x80, import_compliance_enums.LabelSize80x50, import_compliance_enums.LabelSize65x45, import_compliance_enums.LabelSize60x40, import_compliance_enums.LabelSize50x30}, invalid: import_compliance_enums.LabelSize("__invalid__")},
		{name: "import_compliance_enums.LabelOrientation", valid: []stringEnum{import_compliance_enums.LabelOrientationPortrait, import_compliance_enums.LabelOrientationLandscape}, invalid: import_compliance_enums.LabelOrientation("__invalid__")},
		{name: "import_compliance_enums.ArtifactKind", valid: []stringEnum{import_compliance_enums.ArtifactKindPDF, import_compliance_enums.ArtifactKindDOCX, import_compliance_enums.ArtifactKindXLSX, import_compliance_enums.ArtifactKindZIP, import_compliance_enums.ArtifactKindEML}, invalid: import_compliance_enums.ArtifactKind("__invalid__")},
	})
}
