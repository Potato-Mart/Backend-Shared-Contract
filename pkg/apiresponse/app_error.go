// Package apiresponse defines the canonical error codes that every
// service in the MeatAndMe backend returns. Holding these codes in a
// shared module means API consumers (web, mobile, future
// microservices) can switch on stable string codes instead of
// matching message strings.
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

	// Auth
	CodeAuthInvalidCredentials Code = "AUTH_INVALID_CREDENTIALS"
	CodeAuthInvalidToken       Code = "AUTH_INVALID_TOKEN"
	CodeAuthExpiredToken       Code = "AUTH_EXPIRED_TOKEN"
	CodeAuthWrongPortal        Code = "AUTH_WRONG_PORTAL"
	CodeAuthAccountDisabled    Code = "AUTH_ACCOUNT_DISABLED"

	// User
	CodeUserEmailTaken Code = "USER_EMAIL_TAKEN"

	// Catalogue
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

	// Payment terminal (EFTPOS) - covers MX51 SCI / Spice / SPI and any
	// other terminal-backed provider. Codes are wire-stable so POS
	// clients can switch on them to drive the recovery / re-pair UX.
	CodeTerminalNotConnected   Code = "TERMINAL_NOT_CONNECTED"   // device offline (MX51: device_not_connected, 424)
	CodeTerminalUnpaired       Code = "TERMINAL_UNPAIRED"        // pairing no longer active (MX51: no_active_pairings_found, 401)
	CodeTerminalOutcomeUnknown Code = "TERMINAL_OUTCOME_UNKNOWN" // result_financial_status = UNKNOWN; needs override flow
	CodeTerminalSignatureFail  Code = "TERMINAL_SIGNATURE_FAIL"  // signature declined by merchant
	CodeTerminalTimeout        Code = "TERMINAL_TIMEOUT"         // polling-loop timeout; override dialog must be shown
	CodePairingNotFound        Code = "PAIRING_NOT_FOUND"        // pairing code unknown (MX51: pairing_not_found, 404)
	CodePairingExpired         Code = "PAIRING_EXPIRED"          // pairing already progressed/expired/revoked (MX51: pairing_not_initial, 422)
	CodePairingForbidden       Code = "PAIRING_FORBIDDEN"        // wrong tenant/key for environment (MX51: pairing_route_forbidden / test_api_key_forbidden_for_live_pairing, 403)
)

// HTTPStatus returns the preferred HTTP status code for an error Code.
func (c Code) HTTPStatus() int {
	switch c {
	case CodeBadRequest, CodeValidation, CodeOrderEmpty, CodeOrderInvalidTransition,
		CodeOrderTerminal, CodeInsufficientStock, CodeStorageMismatch,
		CodeDiscountInactive, CodeDiscountNotStarted, CodeDiscountExpired,
		CodeDiscountExhausted, CodeDiscountMinNotMet, CodeDiscountInapplicable,
		CodeTerminalSignatureFail:
		return 400
	case CodeUnauthorized, CodeAuthInvalidCredentials, CodeAuthInvalidToken,
		CodeAuthExpiredToken, CodeAuthWrongPortal, CodeAuthAccountDisabled,
		CodeTerminalUnpaired:
		return 401
	case CodeForbidden, CodePairingForbidden:
		return 403
	case CodeNotFound, CodeDiscountNotFound, CodePairingNotFound:
		return 404
	case CodeConflict, CodeUserEmailTaken, CodeSKUCodeTaken,
		CodePlacingAreaCodeTaken, CodeProductCodeTaken, CodeDiscountCodeTaken:
		return 409
	case CodePairingExpired:
		return 422
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
