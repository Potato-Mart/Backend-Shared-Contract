package enums_test

import (
	"testing"

	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/security"
)

func TestSecurityEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "securityenum.AlertLevel", valid: []stringEnum{securityenum.AlertLevelOK, securityenum.AlertLevelWarning, securityenum.AlertLevelCritical, securityenum.AlertLevelExpired}, invalid: securityenum.AlertLevel("__invalid__")},
		{name: "securityenum.AuditOutcome", valid: []stringEnum{securityenum.AuditOutcomeSuccess, securityenum.AuditOutcomeFailure, securityenum.AuditOutcomeDenied}, invalid: securityenum.AuditOutcome("__invalid__")},
		{name: "securityenum.AuthAssuranceLevel", valid: []stringEnum{securityenum.AuthAssuranceLevel1, securityenum.AuthAssuranceLevel2, securityenum.AuthAssuranceLevel3}, invalid: securityenum.AuthAssuranceLevel("__invalid__")},
		{name: "securityenum.AuthMethod", valid: []stringEnum{securityenum.AuthMethodPassword, securityenum.AuthMethodMFA, securityenum.AuthMethodPasskey, securityenum.AuthMethodSSO, securityenum.AuthMethodRefreshToken, securityenum.AuthMethodAPIKey}, invalid: securityenum.AuthMethod("__invalid__")},
		{name: "securityenum.DataClassification", valid: []stringEnum{securityenum.DataClassificationPublic, securityenum.DataClassificationInternal, securityenum.DataClassificationConfidential, securityenum.DataClassificationRestricted}, invalid: securityenum.DataClassification("__invalid__")},
		{name: "securityenum.DataProtectionBasis", valid: []stringEnum{securityenum.DataProtectionBasisNotApplicable, securityenum.DataProtectionBasisConsent, securityenum.DataProtectionBasisContract, securityenum.DataProtectionBasisLegalObligation, securityenum.DataProtectionBasisLegitimateInterest}, invalid: securityenum.DataProtectionBasis("__invalid__")},
		{name: "securityenum.SecurityEventSeverity", valid: []stringEnum{securityenum.SecurityEventSeverityInfo, securityenum.SecurityEventSeverityLow, securityenum.SecurityEventSeverityMedium, securityenum.SecurityEventSeverityHigh, securityenum.SecurityEventSeverityCritical}, invalid: securityenum.SecurityEventSeverity("__invalid__")},
		{name: "securityenum.SecurityEventStatus", valid: []stringEnum{securityenum.SecurityEventStatusDetected, securityenum.SecurityEventStatusTriaged, securityenum.SecurityEventStatusInvestigating, securityenum.SecurityEventStatusContained, securityenum.SecurityEventStatusResolved, securityenum.SecurityEventStatusFalsePositive}, invalid: securityenum.SecurityEventStatus("__invalid__")},
		{name: "securityenum.SecurityRiskLevel", valid: []stringEnum{securityenum.SecurityRiskLevelLow, securityenum.SecurityRiskLevelMedium, securityenum.SecurityRiskLevelHigh, securityenum.SecurityRiskLevelCritical}, invalid: securityenum.SecurityRiskLevel("__invalid__")},
	})
}
