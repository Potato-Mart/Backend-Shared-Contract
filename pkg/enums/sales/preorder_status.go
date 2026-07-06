package salesenum

// PreorderStatus is the lifecycle of a durable customer preorder record.
// Endpoint command payloads remain owned by the service exposing the API.
type PreorderStatus string

const (
	PreorderStatusRequested PreorderStatus = "requested"
	PreorderStatusAccepted  PreorderStatus = "accepted"
	PreorderStatusRejected  PreorderStatus = "rejected"
	PreorderStatusCancelled PreorderStatus = "cancelled"
	PreorderStatusConverted PreorderStatus = "converted"
	PreorderStatusFulfilled PreorderStatus = "fulfilled"
	PreorderStatusExpired   PreorderStatus = "expired"
)

func (s PreorderStatus) IsValid() bool {
	switch s {
	case PreorderStatusRequested, PreorderStatusAccepted, PreorderStatusRejected,
		PreorderStatusCancelled, PreorderStatusConverted, PreorderStatusFulfilled,
		PreorderStatusExpired:
		return true
	}
	return false
}

func (s PreorderStatus) String() string { return string(s) }
