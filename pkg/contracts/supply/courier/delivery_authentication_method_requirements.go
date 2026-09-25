package courier

// DeliveryAuthenticationMethodRequirements describes one Supply-supported
// authentication method, without credential values or an authentication flow.
// Method is an open identifier such as "api_token". RequiredFields contains
// credential field names such as "api_base_url" and "api_token", never values.
// A prerequisite may already be supplied by derived configuration; it does not
// necessarily require operator input. Supply owns support and validation.
type DeliveryAuthenticationMethodRequirements struct {
	Method         string   `json:"method"`
	RequiredFields []string `json:"required_fields"`
}
