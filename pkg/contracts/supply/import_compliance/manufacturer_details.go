package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
)

type ManufacturerDetails struct {
	Name    string             `json:"name"`
	Address *geography.Address `json:"address,omitempty"`
	Phone   string             `json:"phone,omitempty"`
}
