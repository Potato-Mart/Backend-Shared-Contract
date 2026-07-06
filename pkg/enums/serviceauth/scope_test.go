package serviceauthenum

import "testing"

func TestScopeHelpers(t *testing.T) {
	joined := JoinScopes([]Scope{ScopeStockReserve, ScopePricingQuote})
	if joined != "stock:reserve pricing:quote" {
		t.Fatalf("JoinScopes = %q", joined)
	}

	parsed := ParseScopes(joined)
	if len(parsed) != 2 || parsed[0] != string(ScopeStockReserve) || parsed[1] != string(ScopePricingQuote) {
		t.Fatalf("ParseScopes = %#v", parsed)
	}

	for _, scope := range AllScopes() {
		if !scope.IsValid() {
			t.Fatalf("%s should be valid", scope)
		}
		if scope.String() != string(scope) {
			t.Fatalf("%s.String() = %q, want %q", scope, scope.String(), string(scope))
		}
	}

	if Scope("__invalid__").IsValid() {
		t.Fatal("invalid scope should not be valid")
	}
}
