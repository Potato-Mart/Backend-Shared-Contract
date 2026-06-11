// Package serviceauth defines the shared types and constants for
// service-to-service (machine-to-machine) authentication across the Potato
// Mart backend. Backend-Management issues short-lived service tokens via a
// client-credentials flow; Operations and Commerce verify them. See ADR
// 0001 in Backend-Management.
//
// The package is deliberately NOT named "internal": a Go directory named
// internal/ is import-restricted to its own module, which would prevent the
// backend services (separate modules) from importing these types.
package serviceauth

import "strings"

// AudienceInternal is the JWT `aud` value that marks a service (machine)
// token. It isolates service tokens from user tokens, which carry a
// different audience and are signed with a different key.
const AudienceInternal = "internal"

// TokenType is the OAuth2 token_type returned by the issuer.
const TokenType = "Bearer"

// Scope is a fine-grained capability a service token may carry. Scopes are
// least-privilege: a client is granted only the scopes its flows need.
type Scope string

const (
	ScopeStockReserve  Scope = "stock:reserve"
	ScopeStockCommit   Scope = "stock:commit"
	ScopeStockRelease  Scope = "stock:release"
	ScopePricingQuote  Scope = "pricing:quote"
	ScopeCustomersRead Scope = "customers:read"
	ScopeSuppliersRead Scope = "suppliers:read"
)

func (s Scope) String() string { return string(s) }

// AllScopes returns every defined scope (useful for registry validation).
func AllScopes() []Scope {
	return []Scope{
		ScopeStockReserve, ScopeStockCommit, ScopeStockRelease,
		ScopePricingQuote, ScopeCustomersRead, ScopeSuppliersRead,
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

// PathToken is the full request path of the service-token endpoint
// (client-credentials grant) served by Backend-Management. Peer services
// POST a ServiceTokenRequest here to obtain a short-lived service token.
const PathToken = "/v1/internal/token"

// ServiceTokenRequest is the client-credentials grant body for
// POST /v1/internal/token, issued by Backend-Management.
type ServiceTokenRequest struct {
	ClientID     string `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
	Scope        string `json:"scope,omitempty"` // space-delimited requested scopes
}

// ServiceTokenResponse is the issued service access token.
type ServiceTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // always "Bearer"
	ExpiresIn   int    `json:"expires_in"` // seconds until expiry
	Scope       string `json:"scope"`      // space-delimited granted scopes
}

// ServiceClaims is the verified claim set carried by a service token,
// surfaced to handlers by the verifier middleware after the signature,
// issuer, audience and expiry have been checked.
type ServiceClaims struct {
	Subject string   // sub: calling service client_id (e.g. "svc-commerce")
	Scopes  []string // granted scopes
}

// HasScope reports whether the claims include the required scope.
func (c ServiceClaims) HasScope(required Scope) bool {
	for _, s := range c.Scopes {
		if s == string(required) {
			return true
		}
	}
	return false
}
