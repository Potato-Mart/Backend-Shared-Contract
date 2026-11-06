package authorisation

// RoleCode is a wire-stable role key drawn from the vocabulary of one
// audience's role catalogue. The carrying record's portal selects the
// vocabulary: control-portal grants use either a built-in workforce
// role_enums.UserRole key or an owner-defined key, which is never a
// role_enums.RetiredUserRole key; wholesale-portal grants use the closed
// wholesale buyer role keys. The owning service validates a code against its
// own catalogue; consumers treat it as opaque.
type RoleCode string
