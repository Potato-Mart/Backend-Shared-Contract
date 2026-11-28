package ads

// AdEntityRef names one platform object by its platform id and its display
// name, so a read can show "which campaign" without a second lookup.
//
// The id is the platform's own identifier and is opaque: it is matched exactly
// and never parsed. The name is a convenience copy taken at sync time, so it
// can lag a rename on the platform; the id, never the name, is the identity.
type AdEntityRef struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
