package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification"

// Images groups code-based relationships to managed product media. Render URLs
// are resolved from the media masters by the owning backend.
//
// Every image (cover, each gallery image and each detail image) may carry
// staff-authored alt_text and caption on its ObjectMediaRef; both are optional
// and omitted when empty.
type Images struct {
	Cover   *classification.ObjectMediaRef  `json:"cover,omitempty"`
	Gallery []classification.ObjectMediaRef `json:"gallery,omitempty"`
	Details []classification.ObjectMediaRef `json:"details,omitempty"`
}
