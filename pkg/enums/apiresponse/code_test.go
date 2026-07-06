package apiresponseenum

import "testing"

func TestCodeIsValidAndString(t *testing.T) {
	valid := []Code{
		CodeInternal, CodeBadRequest, CodeValidation, CodeNotFound,
		CodeConflict, CodeUnauthorized, CodeForbidden, CodeTooManyRequests,
		CodeRequestTooLarge, CodeAuthInvalidCredentials, CodeAuthInvalidToken,
		CodeAuthExpiredToken, CodeAuthWrongPortal, CodeAuthAccountDisabled,
		CodeAuthMFARequired, CodeAuthReauthRequired, CodeSecurityPolicyViolation,
		CodeIdentityAccountTypeNotAllowed, CodeIdentityPortalAccessDenied,
		CodeIdentityPortalAccessRevoked, CodeIdentityAccountSuspended,
		CodeIdentityAuthIdentityDisabled, CodeIdentityMFARequired,
		CodeIdentityWholesaleOrganisationNotApproved,
		CodeIdentityOrganisationAccessRequired, CodeMembershipNotFound,
		CodeMembershipInactive, CodeMembershipInsufficientPoints,
		CodeMembershipRewardUnavailable, CodeMembershipPointReservationExpired,
		CodeUserEmailTaken, CodeSKUCodeTaken, CodePlacingAreaCodeTaken,
		CodeProductCodeTaken, CodeStorageMismatch, CodeInsufficientStock,
		CodeOrderInvalidTransition, CodeOrderTerminal, CodeOrderEmpty,
		CodeDiscountNotFound, CodeDiscountInactive, CodeDiscountNotStarted,
		CodeDiscountExpired, CodeDiscountExhausted, CodeDiscountMinNotMet,
		CodeDiscountCodeTaken, CodeDiscountInapplicable,
		CodeTerminalNotConnected, CodeTerminalNotRegistered,
		CodeTerminalOutcomeUnknown, CodeTerminalSignatureFail,
		CodeTerminalTimeout, CodeTerminalRequestRejected, CodeTerminalBusy,
	}

	for _, code := range valid {
		if !code.IsValid() {
			t.Fatalf("%s should be valid", code)
		}
		if got := code.String(); got != string(code) {
			t.Fatalf("%s.String() = %q, want %q", code, got, string(code))
		}
	}

	if Code("__invalid__").IsValid() {
		t.Fatal("invalid code should not be valid")
	}
}

func TestCodeHTTPStatus(t *testing.T) {
	tests := []struct {
		code Code
		want int
	}{
		{CodeValidation, 400},
		{CodeAuthInvalidToken, 401},
		{CodeForbidden, 403},
		{CodeNotFound, 404},
		{CodeConflict, 409},
		{CodeRequestTooLarge, 413},
		{CodeTooManyRequests, 429},
		{CodeTerminalNotConnected, 424},
		{CodeTerminalTimeout, 504},
		{CodeInternal, 500},
		{Code("__unknown__"), 500},
	}

	for _, tt := range tests {
		if got := tt.code.HTTPStatus(); got != tt.want {
			t.Fatalf("%s.HTTPStatus() = %d, want %d", tt.code, got, tt.want)
		}
	}
}
