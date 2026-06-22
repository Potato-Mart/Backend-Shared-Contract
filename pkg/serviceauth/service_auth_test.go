package serviceauth

import "testing"

func TestScopeJoinParseRoundTrip(t *testing.T) {
	joined := JoinScopes([]Scope{ScopeStockReserve, ScopeStockCommit})
	if joined != "stock:reserve stock:commit" {
		t.Fatalf("JoinScopes = %q", joined)
	}
	got := ParseScopes(joined)
	if len(got) != 2 || got[0] != "stock:reserve" || got[1] != "stock:commit" {
		t.Fatalf("ParseScopes = %v", got)
	}
}

func TestScopeIsValid(t *testing.T) {
	if !ScopeStockReserve.IsValid() {
		t.Fatal("stock:reserve should be valid")
	}
	if !ScopeMembershipRedeem.IsValid() {
		t.Fatal("membership:redeem should be valid")
	}
	if Scope("bogus:scope").IsValid() {
		t.Fatal("bogus:scope should be invalid")
	}
	if len(AllScopes()) == 0 {
		t.Fatal("AllScopes must not be empty")
	}
}

func TestServiceClaimsHasScope(t *testing.T) {
	c := ServiceClaims{Subject: "svc-commerce", Scopes: ParseScopes("stock:reserve pricing:quote")}
	if !c.HasScope(ScopeStockReserve) {
		t.Fatal("expected stock:reserve present")
	}
	if c.HasScope(ScopeStockCommit) {
		t.Fatal("did not expect stock:commit")
	}
}
