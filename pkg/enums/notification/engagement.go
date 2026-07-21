package notification

// EngagementAction identifies a customer engagement fact about a delivered
// notification.
type EngagementAction string

const (
	EngagementActionDelivered EngagementAction = "delivered"
	EngagementActionOpened    EngagementAction = "opened"
	EngagementActionClicked   EngagementAction = "clicked"
)

func (a EngagementAction) IsValid() bool {
	switch a {
	case EngagementActionDelivered, EngagementActionOpened, EngagementActionClicked:
		return true
	default:
		return false
	}
}

func (a EngagementAction) String() string { return string(a) }
