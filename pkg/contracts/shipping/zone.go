package shipping

import "time"

type Zone struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	States    []string  `json:"states,omitempty"`
	Postcodes []string  `json:"postcodes,omitempty"`
	IsActive  bool      `json:"is_active"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
}
