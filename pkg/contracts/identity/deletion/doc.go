// Package deletion defines the internal, model-only account-deletion
// coordination boundary. It is deliberately independent of HTTP, database,
// provider, and service-specific implementation details so that all
// participating services can use the same idempotent command and receipt
// shapes.
//
// These contracts are for trusted service-to-service communication only.
// They carry an opaque request identifier and the internal canonical user
// identifier, but never email addresses, phone numbers, account profiles,
// database identifiers, provider credentials, or raw obligation details.
package deletion
