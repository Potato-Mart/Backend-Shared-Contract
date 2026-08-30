package geography

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography/geography_enums"

// AdministrativeAreaRef carries an authoritative subdivision code and
// optional display metadata.
type AdministrativeAreaRef struct {
	Code SubdivisionCode                        `json:"code"`
	Name string                                 `json:"name,omitempty"`
	Type geography_enums.AdministrativeAreaType `json:"type,omitempty"`
}
