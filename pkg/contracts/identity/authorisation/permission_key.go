package authorisation

// PermissionKey is a wire-stable workforce RBAC permission identifier such as
// "user.read" or "role.write". It is an open typed string: the concrete key
// catalogue, its validation, and the retired-key deny list are configuration
// owned and seeded by Backend-Identity, not a closed contract enum.
type PermissionKey string
