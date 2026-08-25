package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

type DeclarationShipment struct {
	DeclarationDate       temporal.Date                      `json:"declaration_date,omitempty"`
	ImportMode            import_compliance_enums.ImportMode `json:"import_mode,omitempty"`
	ShippingMethod        string                             `json:"shipping_method,omitempty"`
	TransportReference    string                             `json:"transport_reference,omitempty"`
	ConsignmentIdentifier string                             `json:"consignment_identifier,omitempty"`
	PortOfLoading         string                             `json:"port_of_loading,omitempty"`
	DestinationPort       string                             `json:"destination_port,omitempty"`
}
