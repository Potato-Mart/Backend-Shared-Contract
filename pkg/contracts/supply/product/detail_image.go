package product

import common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"

// DetailImage is one product-gallery image. Slice position defines display
// order; AltText and Caption carry optional localized customer-facing copy.
type DetailImage struct {
	MediaID string                 `json:"media_id"`
	URL     string                 `json:"url"`
	AltText []common.LocalizedText `json:"alt_text,omitempty"`
	Caption []common.LocalizedText `json:"caption,omitempty"`
}
