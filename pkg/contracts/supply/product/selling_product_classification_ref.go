package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security"
)

// SellingProductClassificationRef is a render-safe classification projection.
// Media resolves a brand logo, collection icon, or category image where the
// corresponding classification master supplies one.
type SellingProductClassificationRef struct {
	Code  string                       `json:"code"`
	Name  []localization.LocalizedName `json:"name"`
	Slug  string                       `json:"slug"`
	Media *security.ObjectMedia        `json:"media,omitempty"`
}
