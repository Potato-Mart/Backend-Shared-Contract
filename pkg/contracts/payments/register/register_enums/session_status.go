package register_enums

// SessionStatus is the lifecycle state of a register's daily trading session.
type SessionStatus string

const (
	SessionStatusOpen   SessionStatus = "open"
	SessionStatusClosed SessionStatus = "closed"
)

func (s SessionStatus) IsValid() bool {
	switch s {
	case SessionStatusOpen, SessionStatusClosed:
		return true
	default:
		return false
	}
}

func (s SessionStatus) String() string { return string(s) }
