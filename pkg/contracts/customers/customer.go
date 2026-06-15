package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

// Customer is the public projection of a B2C user account linked to auth.users.
// Secret fields such as password hashes and refresh token material never
// appear here — they live only inside the service that manages identity.
// For wholesale/CRM profiles see CompanyCustomer in company_customer.go.
type Customer struct {
	ID              string          `json:"id"`
	AuthUserID      string          `json:"auth_user_id,omitempty"`
	Email           string          `json:"email"`
	Phone           string          `json:"phone,omitempty"`
	FirstName       string          `json:"first_name,omitempty"`
	LastName        string          `json:"last_name,omitempty"`
	CustomerProfile CustomerProfile `json:"customer_profile"`
	DateOfBirth     *time.Time      `json:"date_of_birth,omitempty"`
	Source          enums.OrderType `json:"source,omitempty"` // acquisition channel: "online" | "pos" | "import"

	// ── Grouped state shared with CompanyCustomer ─────────────────────
	CRM              CRM              `json:"crm"`
	Loyalty          LoyaltyStatus    `json:"loyalty"`
	OrderStats       OrderStats       `json:"order_stats"`
	MarketingConsent MarketingConsent `json:"marketing_consent"`
	RFM              *RFMMetrics      `json:"rfm,omitempty"`
	Referral         *Referral        `json:"referral,omitempty"`

	// ── Shipping ──────────────────────────────────────────────────────
	DefaultShipping *common.ContactAddress  `json:"default_shipping,omitempty"`
	ShippingList    []common.ContactAddress `json:"shipping_list,omitempty"`

	common.AuditFields          `bson:",inline"`
	common.DataProtectionFields `bson:",inline"`
}
