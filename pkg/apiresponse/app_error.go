// Package apiresponse defines the canonical error codes that every
// service in the MeatAndMe backend returns. Holding these codes in a
// shared module means API consumers (web, mobile, future
// microservices) can switch on stable string codes instead of matching
// message strings.
package apiresponse

type Code string

const (
	// Generic
	CodeInternal        Code = "INTERNAL_ERROR"
	CodeBadRequest      Code = "BAD_REQUEST"
	CodeValidation      Code = "VALIDATION_ERROR"
	CodeNotFound        Code = "NOT_FOUND"
	CodeConflict        Code = "CONFLICT"
	CodeUnauthorized    Code = "UNAUTHORIZED"
	CodeForbidden       Code = "FORBIDDEN"
	CodeTooManyRequests Code = "TOO_MANY_REQUESTS"
	CodeRequestTooLarge Code = "REQUEST_BODY_TOO_LARGE"

	// Auth
	CodeAuthInvalidCredentials  Code = "AUTH_INVALID_CREDENTIALS"
	CodeAuthInvalidToken        Code = "AUTH_INVALID_TOKEN"
	CodeAuthExpiredToken        Code = "AUTH_EXPIRED_TOKEN"
	CodeAuthWrongPortal         Code = "AUTH_WRONG_PORTAL"
	CodeAuthAccountDisabled     Code = "AUTH_ACCOUNT_DISABLED"
	CodeAuthMFARequired         Code = "AUTH_MFA_REQUIRED"
	CodeAuthReauthRequired      Code = "AUTH_REAUTH_REQUIRED"
	CodeSecurityPolicyViolation Code = "SECURITY_POLICY_VIOLATION"

	// Identity/access
	// CodeIdentityAccountTypeNotAllowed means the account type cannot enter the requested portal.
	CodeIdentityAccountTypeNotAllowed Code = "IDENTITY_ACCOUNT_TYPE_NOT_ALLOWED"
	// CodeIdentityPortalAccessDenied means portal admission failed.
	CodeIdentityPortalAccessDenied Code = "IDENTITY_PORTAL_ACCESS_DENIED"
	// CodeIdentityPortalAccessRevoked means portal access has been revoked.
	CodeIdentityPortalAccessRevoked Code = "IDENTITY_PORTAL_ACCESS_REVOKED"
	// CodeIdentityAccountSuspended means the account/persona is suspended.
	CodeIdentityAccountSuspended Code = "IDENTITY_ACCOUNT_SUSPENDED"
	// CodeIdentityAuthIdentityDisabled means the selected auth identity is disabled.
	CodeIdentityAuthIdentityDisabled Code = "IDENTITY_AUTH_IDENTITY_DISABLED"
	// CodeIdentityMFARequired means stronger authentication is required.
	CodeIdentityMFARequired Code = "IDENTITY_MFA_REQUIRED"
	// CodeIdentityWholesaleOrganisationNotApproved means the wholesale organisation is not approved.
	CodeIdentityWholesaleOrganisationNotApproved Code = "IDENTITY_WHOLESALE_ORGANISATION_NOT_APPROVED"
	// CodeIdentityOrganisationAccessRequired means valid wholesale organisation access is required.
	CodeIdentityOrganisationAccessRequired Code = "IDENTITY_ORGANISATION_ACCESS_REQUIRED"

	// Membership
	CodeMembershipNotFound                Code = "MEMBERSHIP_NOT_FOUND"
	CodeMembershipInactive                Code = "MEMBERSHIP_INACTIVE"
	CodeMembershipInsufficientPoints      Code = "MEMBERSHIP_INSUFFICIENT_POINTS"
	CodeMembershipRewardUnavailable       Code = "MEMBERSHIP_REWARD_UNAVAILABLE"
	CodeMembershipPointReservationExpired Code = "MEMBERSHIP_POINT_RESERVATION_EXPIRED"

	// User
	CodeUserEmailTaken Code = "USER_EMAIL_TAKEN"

	// Products
	CodeSKUCodeTaken         Code = "SKU_CODE_TAKEN"
	CodePlacingAreaCodeTaken Code = "PLACING_AREA_CODE_TAKEN"
	CodeProductCodeTaken     Code = "PRODUCT_CODE_TAKEN"
	CodeStorageMismatch      Code = "STORAGE_MISMATCH"

	// Inventory
	CodeInsufficientStock Code = "INSUFFICIENT_STOCK"

	// Order
	CodeOrderInvalidTransition Code = "ORDER_INVALID_TRANSITION"
	CodeOrderTerminal          Code = "ORDER_TERMINAL"
	CodeOrderEmpty             Code = "ORDER_EMPTY"

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

// HTTPStatus returns the preferred HTTP status code for an error Code.
func (c Code) HTTPStatus() int {
	switch c {
	case CodeBadRequest, CodeValidation, CodeOrderEmpty, CodeOrderInvalidTransition,
		CodeOrderTerminal, CodeInsufficientStock, CodeStorageMismatch,
		CodeDiscountInactive, CodeDiscountNotStarted, CodeDiscountExpired,
		CodeDiscountExhausted, CodeDiscountMinNotMet, CodeDiscountInapplicable,
		CodeTerminalSignatureFail, CodeTerminalRequestRejected:
		return 400
	case CodeUnauthorized, CodeAuthInvalidCredentials, CodeAuthInvalidToken,
		CodeAuthExpiredToken, CodeAuthWrongPortal, CodeAuthAccountDisabled,
		CodeAuthMFARequired, CodeAuthReauthRequired,
		CodeIdentityAuthIdentityDisabled, CodeIdentityMFARequired:
		return 401
	case CodeForbidden, CodeSecurityPolicyViolation,
		CodeIdentityAccountTypeNotAllowed, CodeIdentityPortalAccessDenied,
		CodeIdentityPortalAccessRevoked, CodeIdentityAccountSuspended,
		CodeIdentityWholesaleOrganisationNotApproved,
		CodeIdentityOrganisationAccessRequired,
		CodeMembershipInactive:
		return 403
	case CodeNotFound, CodeDiscountNotFound, CodeTerminalNotRegistered,
		CodeMembershipNotFound, CodeMembershipRewardUnavailable:
		return 404
	case CodeConflict, CodeUserEmailTaken, CodeSKUCodeTaken,
		CodePlacingAreaCodeTaken, CodeProductCodeTaken, CodeDiscountCodeTaken,
		CodeTerminalBusy, CodeMembershipInsufficientPoints,
		CodeMembershipPointReservationExpired:
		return 409
	case CodeRequestTooLarge:
		return 413
	case CodeTooManyRequests:
		return 429
	case CodeTerminalNotConnected:
		return 424
	case CodeTerminalOutcomeUnknown, CodeTerminalTimeout:
		return 504
	default:
		return 500
	}
}

type Error struct {
	Code    Code              `json:"code"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return string(e.Code) + ": " + e.Message
}

// New builds an *Error with a code and message.
func New(code Code, msg string) *Error {
	return &Error{Code: code, Message: msg}
}

// WithDetails adds a structured details map (used for field-level
// validation errors).
func (e *Error) WithDetails(details map[string]string) *Error {
	e.Details = details
	return e
}
