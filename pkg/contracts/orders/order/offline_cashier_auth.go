package order

import "github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/orders/order/order_enums"

// OfflineCashierAuth identifies the device-recorded authentication evidence.
// The owning services require the field appropriate to Kind and verify it
// against Identity; this record alone does not prove a login or its time.
type OfflineCashierAuth struct {
	Kind               order_enums.OfflineCashierAuthKind `json:"kind"`
	CredentialRevision int64                              `json:"credential_revision,omitempty"`
	IdentitySessionID  string                             `json:"identity_session_id,omitempty"`
}
