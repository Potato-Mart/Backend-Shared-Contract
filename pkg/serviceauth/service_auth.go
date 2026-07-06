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

import serviceauthenum "github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/enums/serviceauth"

// AudienceInternal is the JWT `aud` value that marks a service (machine)
// token. It isolates service tokens from user tokens, which carry a
// different audience and are signed with a different key.
const AudienceInternal = "internal"

// TokenType is the OAuth2 token_type returned by the issuer.
const TokenType = "Bearer"

// PathToken is the full request path of the service-token endpoint
// (client-credentials grant) served by Backend-Management.
const PathToken = "/v1/internal/token"

// ServiceClaims is the verified claim set carried by a service token,
// surfaced to handlers by the verifier middleware after the signature,
// issuer, audience and expiry have been checked.
type ServiceClaims struct {
	Subject string   // sub: calling service client_id (e.g. "svc-commerce")
	Scopes  []string // granted scopes
}

// HasScope reports whether the claims include the required scope.
func (c ServiceClaims) HasScope(required serviceauthenum.Scope) bool {
	for _, s := range c.Scopes {
		if s == string(required) {
			return true
		}
	}
	return false
}
