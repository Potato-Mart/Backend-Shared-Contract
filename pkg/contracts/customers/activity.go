package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// CustomerActivity is a single entry in the customer interaction timeline.
// It records orders, calls, complaints, tier changes, refunds, and any
// other touchpoints for a customer profile.
type CustomerActivity struct {
	ID                string                     `json:"id"`
	CustomerProfileID string                     `json:"customer_profile_id"`
	ActivityType      enums.CustomerActivityType `json:"activity_type"`
	Title             string                     `json:"title"`
	Body              string                     `json:"body,omitempty"`
	Amount            *float64                   `json:"amount,omitempty"` // e.g. refund amount or points delta
	Metadata          common.Metadata            `json:"metadata,omitempty"`
	OccurredAt        time.Time                  `json:"occurred_at"`
	CreatedBy         string                     `json:"created_by,omitempty"`
	CreatedAt         time.Time                  `json:"created_at"`
}

// CustomerIdentity is a single cross-channel identifier belonging to a
// customer profile. Multiple identities (phone, email, LINE, etc.) can
// point to the same profile, enabling cross-channel customer matching.
type CustomerIdentity struct {
	ID                string                     `json:"id"`
	CustomerProfileID string                     `json:"customer_profile_id"`
	Kind              enums.CustomerIdentityKind `json:"kind"`
	Value             string                     `json:"value"`
	Label             string                     `json:"label,omitempty"`
	Verified          bool                       `json:"verified"`
	CreatedAt         time.Time                  `json:"created_at"`
}
