package operations

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/classification/classification_enums"
)

// OutboundContainerPlan describes one outbound shipping container and its
// package-aware contents.
type OutboundContainerPlan struct {
	ID                string                           `json:"id"`
	ContainerCode     string                           `json:"container_code"`
	StorageType       classification_enums.StorageType `json:"storage_type"`
	Contents          []OutboundContainerContent       `json:"contents,omitempty"`
	IsManuallyPlanned bool                             `json:"is_manually_planned"`
	UpdatedAt         time.Time                        `json:"updated_at"`
}
