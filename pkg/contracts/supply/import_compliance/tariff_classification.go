package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

type TariffClassification struct {
	Jurisdiction import_compliance_enums.Jurisdiction `json:"jurisdiction"`
	Code         string                               `json:"code"`
	DutyRate     RateValue                            `json:"duty_rate"`
	GSTRate      *RateValue                           `json:"gst_rate,omitempty"`
	Catalogue    CatalogueReference                   `json:"catalogue"`
	Evidence     []EvidenceReference                  `json:"evidence,omitempty"`
}
