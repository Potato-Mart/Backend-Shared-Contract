package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security/security_enums"
)

func TestSecurityEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "security.AlertLevel", valid: []stringEnum{security_enums.AlertLevelOK, security_enums.AlertLevelWarning, security_enums.AlertLevelCritical, security_enums.AlertLevelExpired}, invalid: security_enums.AlertLevel("__invalid__")},
		{name: "security.AuditOutcome", valid: []stringEnum{security_enums.AuditOutcomeSuccess, security_enums.AuditOutcomeFailure, security_enums.AuditOutcomeDenied}, invalid: security_enums.AuditOutcome("__invalid__")},
		{name: "security.AuthAssuranceLevel", valid: []stringEnum{security_enums.AuthAssuranceLevel1, security_enums.AuthAssuranceLevel2, security_enums.AuthAssuranceLevel3}, invalid: security_enums.AuthAssuranceLevel("__invalid__")},
		{name: "security.AuthMethod", valid: []stringEnum{security_enums.AuthMethodPassword, security_enums.AuthMethodMFA, security_enums.AuthMethodPasskey, security_enums.AuthMethodSSO, security_enums.AuthMethodRefreshToken, security_enums.AuthMethodAPIKey}, invalid: security_enums.AuthMethod("__invalid__")},
		{name: "security.DataClassification", valid: []stringEnum{security_enums.DataClassificationPublic, security_enums.DataClassificationInternal, security_enums.DataClassificationConfidential, security_enums.DataClassificationRestricted}, invalid: security_enums.DataClassification("__invalid__")},
		{name: "security.DataProtectionBasis", valid: []stringEnum{security_enums.DataProtectionBasisNotApplicable, security_enums.DataProtectionBasisConsent, security_enums.DataProtectionBasisContract, security_enums.DataProtectionBasisLegalObligation, security_enums.DataProtectionBasisLegitimateInterest}, invalid: security_enums.DataProtectionBasis("__invalid__")},
		{name: "security.SecurityEventSeverity", valid: []stringEnum{security_enums.SecurityEventSeverityInfo, security_enums.SecurityEventSeverityLow, security_enums.SecurityEventSeverityMedium, security_enums.SecurityEventSeverityHigh, security_enums.SecurityEventSeverityCritical}, invalid: security_enums.SecurityEventSeverity("__invalid__")},
		{name: "security.SecurityEventStatus", valid: []stringEnum{security_enums.SecurityEventStatusDetected, security_enums.SecurityEventStatusTriaged, security_enums.SecurityEventStatusInvestigating, security_enums.SecurityEventStatusContained, security_enums.SecurityEventStatusResolved, security_enums.SecurityEventStatusFalsePositive}, invalid: security_enums.SecurityEventStatus("__invalid__")},
		{name: "security.SecurityRiskLevel", valid: []stringEnum{security_enums.SecurityRiskLevelLow, security_enums.SecurityRiskLevelMedium, security_enums.SecurityRiskLevelHigh, security_enums.SecurityRiskLevelCritical}, invalid: security_enums.SecurityRiskLevel("__invalid__")},
	})
}
