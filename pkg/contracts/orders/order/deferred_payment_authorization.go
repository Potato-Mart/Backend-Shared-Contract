package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/common/security"
	"time"
)

// DeferredPaymentAuthorization is immutable evidence that authenticated staff
// authorized fulfilment before payment at creation. Absence provides no such
// authority. Orders validates role/geography and stamps this snapshot; Supply
// independently verifies authorization and scope before unpaid fulfilment.
type DeferredPaymentAuthorization struct {
	Actor        security.ActorRef `json:"actor"`
	AuthorizedAt time.Time         `json:"authorized_at"`
	Reason       string            `json:"reason,omitempty"`
}
