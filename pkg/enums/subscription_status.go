package enums

// SubscriptionStatus is the lifecycle of a customer subscription.
type SubscriptionStatus string

const (
	SubscriptionStatusActive    SubscriptionStatus = "ACTIVE"
	SubscriptionStatusPaused    SubscriptionStatus = "PAUSED"
	SubscriptionStatusCancelled SubscriptionStatus = "CANCELLED"
)

func (s SubscriptionStatus) IsValid() bool {
	switch s {
	case SubscriptionStatusActive, SubscriptionStatusPaused, SubscriptionStatusCancelled:
		return true
	}
	return false
}

func (s SubscriptionStatus) String() string { return string(s) }
