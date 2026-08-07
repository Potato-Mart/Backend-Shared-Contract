package security_enums

// DataClassification labels information according to its confidentiality
// and handling requirements.
type DataClassification string

const (
	DataClassificationPublic       DataClassification = "public"
	DataClassificationInternal     DataClassification = "internal"
	DataClassificationConfidential DataClassification = "confidential"
	DataClassificationRestricted   DataClassification = "restricted"
)

func (c DataClassification) IsValid() bool {
	switch c {
	case DataClassificationPublic, DataClassificationInternal,
		DataClassificationConfidential, DataClassificationRestricted:
		return true
	}
	return false
}

func (c DataClassification) String() string { return string(c) }

// DataProtectionBasis records the business/legal basis for processing a
// record that may contain personal information.
type DataProtectionBasis string

const (
	DataProtectionBasisNotApplicable      DataProtectionBasis = "not_applicable"
	DataProtectionBasisConsent            DataProtectionBasis = "consent"
	DataProtectionBasisContract           DataProtectionBasis = "contract"
	DataProtectionBasisLegalObligation    DataProtectionBasis = "legal_obligation"
	DataProtectionBasisLegitimateInterest DataProtectionBasis = "legitimate_interest"
)

func (b DataProtectionBasis) IsValid() bool {
	switch b {
	case DataProtectionBasisNotApplicable, DataProtectionBasisConsent,
		DataProtectionBasisContract, DataProtectionBasisLegalObligation,
		DataProtectionBasisLegitimateInterest:
		return true
	}
	return false
}

func (b DataProtectionBasis) String() string { return string(b) }

// SecurityRiskLevel is a coarse risk label used by sessions, devices,
// requests, and security events.
type SecurityRiskLevel string

const (
	SecurityRiskLevelLow      SecurityRiskLevel = "low"
	SecurityRiskLevelMedium   SecurityRiskLevel = "medium"
	SecurityRiskLevelHigh     SecurityRiskLevel = "high"
	SecurityRiskLevelCritical SecurityRiskLevel = "critical"
)

func (r SecurityRiskLevel) IsValid() bool {
	switch r {
	case SecurityRiskLevelLow, SecurityRiskLevelMedium,
		SecurityRiskLevelHigh, SecurityRiskLevelCritical:
		return true
	}
	return false
}

func (r SecurityRiskLevel) String() string { return string(r) }

// AuthMethod records the primary authentication mechanism used for a login
// or privileged action.
type AuthMethod string

const (
	AuthMethodPassword     AuthMethod = "password"
	AuthMethodMFA          AuthMethod = "mfa"
	AuthMethodPasskey      AuthMethod = "passkey"
	AuthMethodSSO          AuthMethod = "sso"
	AuthMethodRefreshToken AuthMethod = "refresh_token"
	AuthMethodAPIKey       AuthMethod = "api_key"
)

func (m AuthMethod) IsValid() bool {
	switch m {
	case AuthMethodPassword, AuthMethodMFA, AuthMethodPasskey,
		AuthMethodSSO, AuthMethodRefreshToken, AuthMethodAPIKey:
		return true
	}
	return false
}

func (m AuthMethod) String() string { return string(m) }

// AuthAssuranceLevel follows the common AAL1/AAL2/AAL3 vocabulary for
// how strongly a user was authenticated.
type AuthAssuranceLevel string

const (
	AuthAssuranceLevel1 AuthAssuranceLevel = "aal1"
	AuthAssuranceLevel2 AuthAssuranceLevel = "aal2"
	AuthAssuranceLevel3 AuthAssuranceLevel = "aal3"
)

func (a AuthAssuranceLevel) IsValid() bool {
	switch a {
	case AuthAssuranceLevel1, AuthAssuranceLevel2, AuthAssuranceLevel3:
		return true
	}
	return false
}

func (a AuthAssuranceLevel) String() string { return string(a) }

type SecurityEventSeverity string

const (
	SecurityEventSeverityInfo     SecurityEventSeverity = "info"
	SecurityEventSeverityLow      SecurityEventSeverity = "low"
	SecurityEventSeverityMedium   SecurityEventSeverity = "medium"
	SecurityEventSeverityHigh     SecurityEventSeverity = "high"
	SecurityEventSeverityCritical SecurityEventSeverity = "critical"
)

func (s SecurityEventSeverity) IsValid() bool {
	switch s {
	case SecurityEventSeverityInfo, SecurityEventSeverityLow,
		SecurityEventSeverityMedium, SecurityEventSeverityHigh,
		SecurityEventSeverityCritical:
		return true
	}
	return false
}

func (s SecurityEventSeverity) String() string { return string(s) }

type SecurityEventStatus string

const (
	SecurityEventStatusDetected      SecurityEventStatus = "detected"
	SecurityEventStatusTriaged       SecurityEventStatus = "triaged"
	SecurityEventStatusInvestigating SecurityEventStatus = "investigating"
	SecurityEventStatusContained     SecurityEventStatus = "contained"
	SecurityEventStatusResolved      SecurityEventStatus = "resolved"
	SecurityEventStatusFalsePositive SecurityEventStatus = "false_positive"
)

func (s SecurityEventStatus) IsValid() bool {
	switch s {
	case SecurityEventStatusDetected, SecurityEventStatusTriaged,
		SecurityEventStatusInvestigating, SecurityEventStatusContained,
		SecurityEventStatusResolved, SecurityEventStatusFalsePositive:
		return true
	}
	return false
}

func (s SecurityEventStatus) String() string { return string(s) }
