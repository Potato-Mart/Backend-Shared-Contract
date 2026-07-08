package serviceauth

import (
	"testing"

	serviceauthenum "github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/enums/serviceauth"
)

func TestScopeJoinParseRoundTrip(t *testing.T) {
	joined := serviceauthenum.JoinScopes([]serviceauthenum.Scope{serviceauthenum.ScopeStockReserve, serviceauthenum.ScopeStockCommit})
	if joined != "stock:reserve stock:commit" {
		t.Fatalf("JoinScopes = %q", joined)
	}
	got := serviceauthenum.ParseScopes(joined)
	if len(got) != 2 || got[0] != "stock:reserve" || got[1] != "stock:commit" {
		t.Fatalf("ParseScopes = %v", got)
	}
}

func TestScopeIsValid(t *testing.T) {
	if !serviceauthenum.ScopeStockReserve.IsValid() {
		t.Fatal("stock:reserve should be valid")
	}
	if !serviceauthenum.ScopeMembershipRedeem.IsValid() {
		t.Fatal("membership:redeem should be valid")
	}
	if !serviceauthenum.ScopeProductsRead.IsValid() {
		t.Fatal("products:read should be valid")
	}
	if !serviceauthenum.ScopeWholesaleTermsRead.IsValid() {
		t.Fatal("wholesale:terms:read should be valid")
	}
	if serviceauthenum.Scope("bogus:scope").IsValid() {
		t.Fatal("bogus:scope should be invalid")
	}
	if len(serviceauthenum.AllScopes()) == 0 {
		t.Fatal("AllScopes must not be empty")
	}
}

func TestServiceClaimsHasScope(t *testing.T) {
	c := ServiceClaims{Subject: "svc-commerce", Scopes: serviceauthenum.ParseScopes("stock:reserve pricing:quote")}
	if !c.HasScope(serviceauthenum.ScopeStockReserve) {
		t.Fatal("expected stock:reserve present")
	}
	if c.HasScope(serviceauthenum.ScopeStockCommit) {
		t.Fatal("did not expect stock:commit")
	}
}
