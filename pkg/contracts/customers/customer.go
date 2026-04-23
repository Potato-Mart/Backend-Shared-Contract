package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/pkg/enums"
)

// Customer is the public projection of a user account. Secret fields such
// as password hashes and refresh token material never appear here —
// they live only inside the service that manages identity.
type Customer struct {
	ID               string             `json:"id"`
	AuthUserID       string             `json:"auth_user_id,omitempty"`
	Email            string             `json:"email"`
	Phone            string             `json:"phone,omitempty"`
	FirstName        string             `json:"first_name,omitempty"`
	LastName         string             `json:"last_name,omitempty"`
	Role             enums.UserRole     `json:"role"`
	CustomerType     enums.CustomerType `json:"customer_type,omitempty"`
	Active           bool               `json:"active"`
	Points           int                `json:"points"`
	Tier             enums.CustomerTier `json:"tier"`
	TierSpend        float64            `json:"tier_spend"`
	TotalOrders      int                `json:"total_orders"`
	TotalSpend       float64            `json:"total_spend"`
	LastOrderAt      *time.Time         `json:"last_order_at,omitempty"`
	Tags             []string           `json:"tags,omitempty"`
	DefaultShipping  *DefaultShipping   `json:"default_shipping,omitempty"`
	ShippingList     []Shipping         `json:"shipping_list,omitempty"`
	AcceptsMarketing bool               `json:"accepts_marketing"`
	IsActive         bool               `json:"is_active"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

type DefaultShipping struct {
	DefaultContact *common.Recipient `json:"default_contact,omitempty"`
	DefaultAddress *common.Address   `json:"default_address,omitempty"`
}

type Shipping struct {
	Contact *common.Recipient `json:"contact,omitempty"`
	Address *common.Address   `json:"address,omitempty"`
}
