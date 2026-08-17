package wish

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/wish/wish_enums"
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
	CreatedSKUID         string                       `json:"created_sku_id,omitempty"`
	CreatedAt            time.Time                    `json:"created_at"`
	UpdatedAt            time.Time                    `json:"updated_at"`
}
