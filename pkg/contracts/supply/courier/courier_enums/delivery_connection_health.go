package courier_enums

// DeliveryConnectionHealth is sanitized connection-test evidence. Unknown
// includes never checked; a configured credential does not imply healthy.
type DeliveryConnectionHealth string

const (
	DeliveryConnectionHealthUnknown   DeliveryConnectionHealth = "unknown"
	DeliveryConnectionHealthHealthy   DeliveryConnectionHealth = "healthy"
	DeliveryConnectionHealthUnhealthy DeliveryConnectionHealth = "unhealthy"
)

func (h DeliveryConnectionHealth) IsValid() bool {
	switch h {
	case DeliveryConnectionHealthUnknown, DeliveryConnectionHealthHealthy, DeliveryConnectionHealthUnhealthy:
		return true
	default:
		return false
	}
}

func (h DeliveryConnectionHealth) String() string { return string(h) }
