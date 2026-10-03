// Package template defines country-scoped, provider-neutral notification authoring
// records. Notification owns their validation, authorization, storage, rendering,
// translation and publication. These records contain no recipient values,
// credentials, rendered personal content or transport commands.
//
// Schema version, mutable draft revision, immutable published version and
// translation fingerprints describe independent dimensions. See
// package types for wire semantics; owning services enforce consumer invariants.
package template
