package shipping

import "time"

type Zone struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	States []string `json:"states,omitempty"`
	// Postcodes limits the zone to specific postcodes within States.
	Postcodes []string `json:"postcodes,omitempty"`
	// IsLocal marks the zone as the local metro area (e.g. Metro
	// Melbourne) used for order delivery-region classification.
	IsLocal   bool      `json:"is_local,omitempty"`
	IsActive  bool      `json:"is_active"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
}
