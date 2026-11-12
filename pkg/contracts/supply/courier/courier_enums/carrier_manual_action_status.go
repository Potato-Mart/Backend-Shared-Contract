package courier_enums

// CarrierManualActionStatus is the state of a carrier manual action. Open
// actions show a warning on the order's shipments; Resolved means staff
// recorded that the carrier-side work is done.
type CarrierManualActionStatus string

const (
	CarrierManualActionStatusOpen     CarrierManualActionStatus = "open"
	CarrierManualActionStatusResolved CarrierManualActionStatus = "resolved"
)

func (s CarrierManualActionStatus) IsValid() bool {
	switch s {
	case CarrierManualActionStatusOpen, CarrierManualActionStatusResolved:
		return true
	default:
		return false
	}
}

func (s CarrierManualActionStatus) String() string { return string(s) }
