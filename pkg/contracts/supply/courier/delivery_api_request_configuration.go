package courier

import "github.com/Potato-Mart/Backend-Shared-Contract/v38/pkg/contracts/supply/courier/courier_enums"

// DeliveryAPIRequestConfiguration is a data-only REST request definition.
// EndpointPath is relative to DeliveryProviderCredentials.APIBaseURL. Headers
// and Query preserve authored key/value order and repeated keys. JSONBody is
// assembled as a top-level JSON object from fields with unique keys. Supply
// owns endpoint validation, secret resolution, and request execution.
type DeliveryAPIRequestConfiguration struct {
	Method       courier_enums.DeliveryAPIHTTPMethod `json:"method"`
	EndpointPath string                              `json:"endpoint_path"`
	Headers      []DeliveryAPIRequestField           `json:"headers,omitempty"`
	Query        []DeliveryAPIRequestField           `json:"query,omitempty"`
	JSONBody     []DeliveryAPIRequestField           `json:"json_body,omitempty"`
}
