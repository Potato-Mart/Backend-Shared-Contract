package courier_enums

// DeliveryAPIHTTPMethod is one of the methods supported by manually authored
// courier request configurations.
type DeliveryAPIHTTPMethod string

const (
	DeliveryAPIHTTPMethodGet  DeliveryAPIHTTPMethod = "GET"
	DeliveryAPIHTTPMethodPost DeliveryAPIHTTPMethod = "POST"
)

func (m DeliveryAPIHTTPMethod) IsValid() bool {
	switch m {
	case DeliveryAPIHTTPMethodGet, DeliveryAPIHTTPMethodPost:
		return true
	default:
		return false
	}
}

func (m DeliveryAPIHTTPMethod) String() string { return string(m) }
