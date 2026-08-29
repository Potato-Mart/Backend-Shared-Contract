package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
)

type RFIInspectionLocation struct {
	BusinessNameAndAANumber string             `json:"business_name_and_aa_number,omitempty"`
	PremiseAddress          *geography.Address `json:"premise_address,omitempty"`
	OpeningHours            string             `json:"opening_hours,omitempty"`
	ContactName             string             `json:"contact_name,omitempty"`
	ContactPhone            string             `json:"contact_phone,omitempty"`
	PrivateResidence        bool               `json:"private_residence"`
}
