package compliance

import (
	compliance_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/compliance/compliance_enums"
)

type TariffClassification struct {
	Jurisdiction compliance_enums.Jurisdiction `json:"jurisdiction"`
	Code         string                        `json:"code"`
	DutyRate     RateValue                     `json:"duty_rate"`
	GSTRate      *RateValue                    `json:"gst_rate,omitempty"`
	Catalogue    CatalogueReference            `json:"catalogue"`
	Evidence     []EvidenceReference           `json:"evidence,omitempty"`
}
