package apiresponse

import "testing"

func TestCodeIsValidAndString(t *testing.T) {
	valid := []Code{
		CodeInternal, CodeBadRequest, CodeValidation, CodeNotFound,
		CodeConflict, CodeUnauthorized, CodeForbidden, CodeTooManyRequests,
		CodeRequestTooLarge, CodeServiceUnavailable, CodeAuthInvalidCredentials, CodeAuthInvalidToken,
		CodeAuthExpiredToken, CodeAuthWrongPortal, CodeAuthAccountDisabled,
		CodeAuthMFARequired, CodeAuthReauthRequired, CodeEmailVerificationRequired,
		CodeSecurityPolicyViolation,
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
	if CodeCartNotActive.String() != "CART_NOT_ACTIVE" ||
		CodeEmailVerificationRequired.String() != "EMAIL_VERIFICATION_REQUIRED" {
		t.Fatal("v23 mobile recovery error codes changed")
	}
}
