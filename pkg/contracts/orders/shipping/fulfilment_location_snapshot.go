package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/orders/shipping/shipping_enums"
)

// FulfilmentLocationSnapshot freezes the address or depot that resolved a
// cart's commercial market and eligibility. Delivery uses DeliveryAddress and
// no SelectedDepotCode; pickup and in-store carry use SelectedDepotCode and no
// DeliveryAddress. Changing this snapshot requires the cart to be repriced.
type FulfilmentLocationSnapshot struct {
	Intent              shipping_enums.FulfilmentIntent `json:"intent"`
	DeliveryAddress     *party.ContactAddress           `json:"delivery_address,omitempty"`
	SelectedDepotCode   string                          `json:"selected_depot_code,omitempty"`
	GeographicContext   geography.GeographicContext     `json:"geographic_context"`
	LocationFingerprint string                          `json:"location_fingerprint"`
	CapturedAt          time.Time                       `json:"captured_at"`
}
