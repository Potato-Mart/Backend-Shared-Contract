package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/device"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/order/order_enums"
)

// SourceDevice captures request and POS attribution for the order source.
type SourceDevice struct {
	Type    order_enums.OrderSourceDeviceType `json:"type,omitempty"`
	LocalID string                            `json:"local_id,omitempty"`
	Name    string                            `json:"name,omitempty"`

	// POS carries first-class in-store attribution when the order originates
	// at a point of sale.
	POS *POSAttribution `json:"pos,omitempty"`

	// Metadata stores source-specific details that should not become first-class
	// contract fields yet, for example app_version, operator_id,
	// forwarded_for, device_model, or network_interface.
	Metadata metadata.Metadata `json:"metadata,omitempty"`

	// DeviceRecord carries shared fingerprint/request attributes such as
	// device_key, ip_address, user_agent, os, and browser.
	device.DeviceRecord
}
