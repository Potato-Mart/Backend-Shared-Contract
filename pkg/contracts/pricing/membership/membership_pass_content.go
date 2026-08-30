package membership

import "time"

// MembershipPassContent is the provider-neutral snapshot used to issue a
// customer-owned membership pass. Google save-JWT responses, Apple pass
// packages, signing policy, routes, and media types remain backend-owned.
type MembershipPassContent struct {
	CustomerNumber  string                `json:"customer_number"`
	TierKey         string                `json:"tier_key,omitempty"`
	AvailablePoints int                   `json:"available_points"`
	Barcode         MembershipPassBarcode `json:"barcode"`
	GeneratedAt     time.Time             `json:"generated_at"`
}
