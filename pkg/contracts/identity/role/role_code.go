package role

// RoleCode is a wire-stable role key drawn from the vocabulary of one
// audience's role catalogue. The carrying record's portal selects the
// vocabulary: control-portal grants use the closed workforce
// role_enums.UserRole keys and wholesale-portal grants use the closed
// wholesale buyer role keys. The owning service validates a code against its
// own catalogue; consumers treat it as opaque.
type RoleCode string
