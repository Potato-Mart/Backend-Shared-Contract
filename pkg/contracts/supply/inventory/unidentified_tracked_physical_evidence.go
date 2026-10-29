package inventory

import (
	"encoding/json"
	"errors"
)

// UnidentifiedTrackedPhysicalEvidence records a physical quantity whose
// tracked unit and physical source have not been established. The enclosing
// allocation reference is reservation history, not a claim that the reserved
// unit or bucket was observed. It cannot imply a lot, return source or cost.
// Supply owns quantity accounting and later evidence-backed reconciliation.
type UnidentifiedTrackedPhysicalEvidence struct{}

// UnmarshalJSON rejects any claimed unit, lot or source embedded in this arm.
func (*UnidentifiedTrackedPhysicalEvidence) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil || fields == nil || len(fields) != 0 {
		return errors.New("unidentified tracked evidence must be an empty object")
	}
	return nil
}
