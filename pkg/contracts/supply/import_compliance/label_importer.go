package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
)

type LabelImporter struct {
	Name    string             `json:"name,omitempty"`
	Address *geography.Address `json:"address,omitempty"`
	Phone   string             `json:"phone,omitempty"`
}
