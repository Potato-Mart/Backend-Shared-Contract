package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// Customer is the public projection of a B2C user account linked to auth.users.
// Secret fields such as password hashes and refresh token material never
// appear here — they live only inside the service that manages identity.
// For wholesale/CRM profiles see CustomerProfile in profile.go.
type Customer struct {
	ID                string                      `json:"id"`
	AuthUserID        string                      `json:"auth_user_id,omitempty"`
	Email             string                      `json:"email"`
	Phone             string                      `json:"phone,omitempty"`
	FirstName         string                      `json:"first_name,omitempty"`
	LastName          string                      `json:"last_name,omitempty"`
	DateOfBirth       *time.Time                  `json:"date_of_birth,omitempty"`
	Notes             string                      `json:"notes,omitempty"`
	Source            string                      `json:"source,omitempty"` // "online" | "pos" | "import"
	Role              enums.UserRole              `json:"role"`
	Segment           enums.CustomerSegment       `json:"segment,omitempty"`
	Status            enums.CustomerProfileStatus `json:"status"`
	Points            int                         `json:"points"`
	Tier              enums.CustomerTier          `json:"tier"`
	TierSpend         common.Money                `json:"tier_spend"`
	TotalOrders       int                         `json:"total_orders"`
	TotalSpend        common.Money                `json:"total_spend"`
	AverageOrderSpend common.Money                `json:"average_order_spend"`
	FirstOrderAt      *time.Time                  `json:"first_order_at,omitempty"`
	LastOrderAt       *time.Time                  `json:"last_order_at,omitempty"`
	Tags              []string                    `json:"tags,omitempty"`
	DefaultShipping   *DefaultShipping            `json:"default_shipping,omitempty"`
	ShippingList      []Shipping                  `json:"shipping_list,omitempty"`
	AcceptsMarketing  bool                        `json:"accepts_marketing"`
	IsActive          bool                        `json:"is_active"`
	CreatedAt         time.Time                   `json:"created_at"`
	UpdatedAt         time.Time                   `json:"updated_at"`

	common.DataProtectionFields
}

type DefaultShipping struct {
	DefaultContact *common.Recipient `json:"default_contact,omitempty"`
	DefaultAddress *common.Address   `json:"default_address,omitempty"`
}

type Shipping struct {
	Contact *common.Recipient `json:"contact,omitempty"`
	Address *common.Address   `json:"address,omitempty"`
}
