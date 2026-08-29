package retail

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/retail/retail_enums"
)

// RetailCustomerBasicInfo groups stable name, contact, and acquisition fields
// for a retail customer profile.
type RetailCustomerBasicInfo struct {
	Name                   party.PersonName                       `json:"name"`
	Contacts               party.ContactChannels                  `json:"contacts"`
	PreferredContactMethod retail_enums.PreferredContactMethod    `json:"preferred_contact_method,omitempty"`
	DateOfBirth            *time.Time                             `json:"date_of_birth,omitempty"`
	Gender                 retail_enums.CustomerGender            `json:"gender,omitempty"`
	AcquisitionSource      retail_enums.CustomerAcquisitionSource `json:"acquisition_source,omitempty"`
}
