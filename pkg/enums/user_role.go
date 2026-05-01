package enums

// UserRole identifies which portal a user belongs to.
// A user has exactly one role; admins never place client orders and
// clients cannot access admin-only endpoints.
type UserRole string

const (
	UserRoleAdmin  UserRole = "admin"
	UserRoleClient UserRole = "user"
)

// IsValid reports whether r is a known UserRole value.
func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleAdmin, UserRoleClient:
		return true
	}
	return false
}

func (r UserRole) String() string { return string(r) }
