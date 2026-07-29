package apiresponseenum

type Code string

const (
	// Generic
	CodeInternal           Code = "INTERNAL_ERROR"
	CodeBadRequest         Code = "BAD_REQUEST"
	CodeValidation         Code = "VALIDATION_ERROR"
	CodeNotFound           Code = "NOT_FOUND"
	CodeConflict           Code = "CONFLICT"
	CodeUnauthorized       Code = "UNAUTHORIZED"
	CodeForbidden          Code = "FORBIDDEN"
	CodeTooManyRequests    Code = "TOO_MANY_REQUESTS"
	CodeRequestTooLarge    Code = "REQUEST_BODY_TOO_LARGE"
	CodeServiceUnavailable Code = "SERVICE_UNAVAILABLE"

	// Auth
	CodeAuthInvalidCredentials    Code = "AUTH_INVALID_CREDENTIALS"
	CodeAuthInvalidToken          Code = "AUTH_INVALID_TOKEN"
	CodeAuthExpiredToken          Code = "AUTH_EXPIRED_TOKEN"
	CodeAuthWrongPortal           Code = "AUTH_WRONG_PORTAL"
	CodeAuthAccountDisabled       Code = "AUTH_ACCOUNT_DISABLED"
	CodeAuthMFARequired           Code = "AUTH_MFA_REQUIRED"
	CodeAuthReauthRequired        Code = "AUTH_REAUTH_REQUIRED"
	CodeEmailVerificationRequired Code = "EMAIL_VERIFICATION_REQUIRED"
	CodeSecurityPolicyViolation   Code = "SECURITY_POLICY_VIOLATION"

	// Identity/access
	CodeIdentityAccountTypeNotAllowed            Code = "IDENTITY_ACCOUNT_TYPE_NOT_ALLOWED"
	CodeIdentityPortalAccessDenied               Code = "IDENTITY_PORTAL_ACCESS_DENIED"
	CodeIdentityPortalAccessRevoked              Code = "IDENTITY_PORTAL_ACCESS_REVOKED"
	CodeIdentityAccountSuspended                 Code = "IDENTITY_ACCOUNT_SUSPENDED"
	CodeIdentityAuthIdentityDisabled             Code = "IDENTITY_AUTH_IDENTITY_DISABLED"
	CodeIdentityMFARequired                      Code = "IDENTITY_MFA_REQUIRED"
	CodeIdentityWholesaleOrganisationNotApproved Code = "IDENTITY_WHOLESALE_ORGANISATION_NOT_APPROVED"
	CodeIdentityOrganisationAccessRequired       Code = "IDENTITY_ORGANISATION_ACCESS_REQUIRED"

	// Membership
	CodeMembershipNotFound                Code = "MEMBERSHIP_NOT_FOUND"
	CodeMembershipInactive                Code = "MEMBERSHIP_INACTIVE"
	CodeMembershipInsufficientPoints      Code = "MEMBERSHIP_INSUFFICIENT_POINTS"
	CodeMembershipRewardUnavailable       Code = "MEMBERSHIP_REWARD_UNAVAILABLE"
	CodeMembershipPointReservationExpired Code = "MEMBERSHIP_POINT_RESERVATION_EXPIRED"

	// User
	CodeUserEmailTaken Code = "USER_EMAIL_TAKEN"

	// Products
	CodeSKUCodeTaken             Code = "SKU_CODE_TAKEN"
	CodeStorageLocationCodeTaken Code = "STORAGE_LOCATION_CODE_TAKEN"
	CodeProductCodeTaken         Code = "PRODUCT_CODE_TAKEN"
	CodeStorageMismatch          Code = "STORAGE_MISMATCH"

	// Inventory
	CodeInsufficientStock Code = "INSUFFICIENT_STOCK"

	// Order
	CodeOrderInvalidTransition Code = "ORDER_INVALID_TRANSITION"
	CodeOrderTerminal          Code = "ORDER_TERMINAL"
	CodeOrderEmpty             Code = "ORDER_EMPTY"
	CodeCartNotActive          Code = "CART_NOT_ACTIVE"

	// Discount
	CodeDiscountNotFound     Code = "DISCOUNT_NOT_FOUND"
	CodeDiscountInactive     Code = "DISCOUNT_INACTIVE"
	CodeDiscountNotStarted   Code = "DISCOUNT_NOT_STARTED"
	CodeDiscountExpired      Code = "DISCOUNT_EXPIRED"
	CodeDiscountExhausted    Code = "DISCOUNT_EXHAUSTED"
	CodeDiscountMinNotMet    Code = "DISCOUNT_MIN_NOT_MET"
	CodeDiscountCodeTaken    Code = "DISCOUNT_CODE_TAKEN"
	CodeDiscountInapplicable Code = "DISCOUNT_INAPPLICABLE"

	// Payment terminal (EFTPOS). Codes are wire-stable so POS clients
	// can switch on them to drive retry, status-check, and recovery UX.
	CodeTerminalNotConnected    Code = "TERMINAL_NOT_CONNECTED"
	CodeTerminalNotRegistered   Code = "TERMINAL_NOT_REGISTERED"
	CodeTerminalOutcomeUnknown  Code = "TERMINAL_OUTCOME_UNKNOWN"
	CodeTerminalSignatureFail   Code = "TERMINAL_SIGNATURE_FAIL"
	CodeTerminalTimeout         Code = "TERMINAL_TIMEOUT"
	CodeTerminalRequestRejected Code = "TERMINAL_REQUEST_REJECTED"
	CodeTerminalBusy            Code = "TERMINAL_BUSY"
)

func (c Code) IsValid() bool {
	switch c {
	case CodeInternal, CodeBadRequest, CodeValidation, CodeNotFound,
		CodeConflict, CodeUnauthorized, CodeForbidden, CodeTooManyRequests,
		CodeRequestTooLarge, CodeServiceUnavailable, CodeAuthInvalidCredentials, CodeAuthInvalidToken,
		CodeAuthExpiredToken, CodeAuthWrongPortal, CodeAuthAccountDisabled,
		CodeAuthMFARequired, CodeAuthReauthRequired, CodeSecurityPolicyViolation,
		CodeEmailVerificationRequired,
		CodeIdentityAccountTypeNotAllowed, CodeIdentityPortalAccessDenied,
		CodeIdentityPortalAccessRevoked, CodeIdentityAccountSuspended,
		CodeIdentityAuthIdentityDisabled, CodeIdentityMFARequired,
		CodeIdentityWholesaleOrganisationNotApproved,
		CodeIdentityOrganisationAccessRequired, CodeMembershipNotFound,
		CodeMembershipInactive, CodeMembershipInsufficientPoints,
		CodeMembershipRewardUnavailable, CodeMembershipPointReservationExpired,
		CodeUserEmailTaken, CodeSKUCodeTaken, CodeStorageLocationCodeTaken,
		CodeProductCodeTaken, CodeStorageMismatch, CodeInsufficientStock,
		CodeOrderInvalidTransition, CodeOrderTerminal, CodeOrderEmpty, CodeCartNotActive,
		CodeDiscountNotFound, CodeDiscountInactive, CodeDiscountNotStarted,
		CodeDiscountExpired, CodeDiscountExhausted, CodeDiscountMinNotMet,
		CodeDiscountCodeTaken, CodeDiscountInapplicable,
		CodeTerminalNotConnected, CodeTerminalNotRegistered,
		CodeTerminalOutcomeUnknown, CodeTerminalSignatureFail,
		CodeTerminalTimeout, CodeTerminalRequestRejected, CodeTerminalBusy:
		return true
	}
	return false
}

func (c Code) String() string { return string(c) }
