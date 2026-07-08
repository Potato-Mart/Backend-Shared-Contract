package serviceauthenum

import "strings"

// Scope is a fine-grained capability a service token may carry. Scopes are
// least-privilege: a client is granted only the scopes its flows need.
type Scope string

const (
	ScopeStockReserve       Scope = "stock:reserve"
	ScopeStockCommit        Scope = "stock:commit"
	ScopeStockRelease       Scope = "stock:release"
	ScopePricingQuote       Scope = "pricing:quote"
	ScopeProductsRead       Scope = "products:read"
	ScopeCustomersRead      Scope = "customers:read"
	ScopeSuppliersRead      Scope = "suppliers:read"
	ScopeMembershipRead     Scope = "membership:read"
	ScopeMembershipPoints   Scope = "membership:points"
	ScopeMembershipRedeem   Scope = "membership:redeem"
	ScopeWholesaleTermsRead Scope = "wholesale:terms:read"
	ScopePromotionGrant     Scope = "promotion:grant"
	ScopeRestockNotify      Scope = "restock:notify"
)

func (s Scope) String() string { return string(s) }

// AllScopes returns every defined scope (useful for registry validation).
func AllScopes() []Scope {
	return []Scope{
		ScopeStockReserve, ScopeStockCommit, ScopeStockRelease,
		ScopePricingQuote, ScopeProductsRead, ScopeCustomersRead, ScopeSuppliersRead,
		ScopeMembershipRead, ScopeMembershipPoints, ScopeMembershipRedeem,
		ScopeWholesaleTermsRead, ScopePromotionGrant, ScopeRestockNotify,
	}
}

// IsValid reports whether s is a known scope.
func (s Scope) IsValid() bool {
	for _, k := range AllScopes() {
		if k == s {
			return true
		}
	}
	return false
}

// ParseScopes splits a space-delimited scope string into trimmed tokens.
func ParseScopes(s string) []string { return strings.Fields(s) }

// JoinScopes renders scopes as a space-delimited string (the wire format).
func JoinScopes(scopes []Scope) string {
	parts := make([]string, len(scopes))
	for i, s := range scopes {
		parts[i] = string(s)
	}
	return strings.Join(parts, " ")
}
