package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/classification"

// Images groups code-only relationships to managed product media. Render URLs
// are resolved from the media masters by the owning backend.
type Images struct {
	Cover   *classification.ObjectMediaRef  `json:"cover,omitempty"`
	Gallery []classification.ObjectMediaRef `json:"gallery,omitempty"`
	Details []classification.ObjectMediaRef `json:"details,omitempty"`
}
