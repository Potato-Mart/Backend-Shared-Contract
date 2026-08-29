package wish

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/wish/wish_enums"
)

// WishProposal is a customer-submitted product idea without customer identity.
// Conversion and fulfilment references are populated only when available.
type WishProposal struct {
	ID                   string                       `json:"id"`
	ProductName          string                       `json:"product_name"`
	Description          string                       `json:"description,omitempty"`
	ReferenceURL         string                       `json:"reference_url,omitempty"`
	State                wish_enums.WishProposalState `json:"state"`
	ConvertedCandidateID string                       `json:"converted_candidate_id,omitempty"`
	CreatedSKUCode       string                       `json:"created_sku_code,omitempty"`

	audit.AuditFields
}
