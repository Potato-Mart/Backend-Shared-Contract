package wish

import (
	"time"
)

// WishProposal is a customer-submitted product idea without customer identity.
// Conversion and fulfilment references are populated only when available.
type WishProposal struct {
	ID                    string            `json:"id"`
	ProductName           string            `json:"product_name"`
	Description           string            `json:"description,omitempty"`
	ReferenceURL          string            `json:"reference_url,omitempty"`
	State                 WishProposalState `json:"state"`
	ConvertedCandidateID  string            `json:"converted_candidate_id,omitempty"`
	CreatedProductSKUCode string            `json:"created_product_sku_code,omitempty"`
	CreatedAt             time.Time         `json:"created_at"`
	UpdatedAt             time.Time         `json:"updated_at"`
}
