package enums_test

import (
	security "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/security"
	"testing"
)

func TestSecurityEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "security.AlertLevel", valid: []stringEnum{security.AlertLevelOK, security.AlertLevelWarning, security.AlertLevelCritical, security.AlertLevelExpired}, invalid: security.AlertLevel("__invalid__")},
		{name: "security.AuditOutcome", valid: []stringEnum{security.AuditOutcomeSuccess, security.AuditOutcomeFailure, security.AuditOutcomeDenied}, invalid: security.AuditOutcome("__invalid__")},
		{name: "security.AuthAssuranceLevel", valid: []stringEnum{security.AuthAssuranceLevel1, security.AuthAssuranceLevel2, security.AuthAssuranceLevel3}, invalid: security.AuthAssuranceLevel("__invalid__")},
		{name: "security.AuthMethod", valid: []stringEnum{security.AuthMethodPassword, security.AuthMethodMFA, security.AuthMethodPasskey, security.AuthMethodSSO, security.AuthMethodRefreshToken, security.AuthMethodAPIKey}, invalid: security.AuthMethod("__invalid__")},
		{name: "security.DataClassification", valid: []stringEnum{security.DataClassificationPublic, security.DataClassificationInternal, security.DataClassificationConfidential, security.DataClassificationRestricted}, invalid: security.DataClassification("__invalid__")},
		{name: "security.DataProtectionBasis", valid: []stringEnum{security.DataProtectionBasisNotApplicable, security.DataProtectionBasisConsent, security.DataProtectionBasisContract, security.DataProtectionBasisLegalObligation, security.DataProtectionBasisLegitimateInterest}, invalid: security.DataProtectionBasis("__invalid__")},
		{name: "security.SecurityEventSeverity", valid: []stringEnum{security.SecurityEventSeverityInfo, security.SecurityEventSeverityLow, security.SecurityEventSeverityMedium, security.SecurityEventSeverityHigh, security.SecurityEventSeverityCritical}, invalid: security.SecurityEventSeverity("__invalid__")},
		{name: "security.SecurityEventStatus", valid: []stringEnum{security.SecurityEventStatusDetected, security.SecurityEventStatusTriaged, security.SecurityEventStatusInvestigating, security.SecurityEventStatusContained, security.SecurityEventStatusResolved, security.SecurityEventStatusFalsePositive}, invalid: security.SecurityEventStatus("__invalid__")},
		{name: "security.SecurityRiskLevel", valid: []stringEnum{security.SecurityRiskLevelLow, security.SecurityRiskLevelMedium, security.SecurityRiskLevelHigh, security.SecurityRiskLevelCritical}, invalid: security.SecurityRiskLevel("__invalid__")},
	})
}
