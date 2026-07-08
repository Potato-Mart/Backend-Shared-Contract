package wholesale

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/common"
	wholesaleenum "github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/enums/wholesale"
)

// WholesaleApplicationRequest is the public trade-account application payload.
// The backend creates the customer identity plus pending wholesale records from
// this single request; approval remains an admin action.
type WholesaleApplicationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`

	ContactName  string `json:"contact_name"`
	ContactPhone string `json:"contact_phone"`
	ContactRole  string `json:"contact_role,omitempty"`

	BusinessName string `json:"business_name"`
	TradingName  string `json:"trading_name,omitempty"`
	LegalName    string `json:"legal_name,omitempty"`
	ABN          string `json:"abn,omitempty"`
	BusinessType string `json:"business_type,omitempty"`

	BusinessAddress *common.ContactAddress `json:"business_address,omitempty"`
	DeliveryAddress *common.ContactAddress `json:"delivery_address,omitempty"`

	ExpectedOrderVolume   string   `json:"expected_order_volume,omitempty"`
	StorageRequirements   []string `json:"storage_requirements,omitempty"`
	PaymentPreference     string   `json:"payment_preference,omitempty"`
	PreferredDeliveryDays []string `json:"preferred_delivery_days,omitempty"`
	Notes                 string   `json:"notes,omitempty"`
}

// WholesaleApplicationResponse is returned after a public application submit.
type WholesaleApplicationResponse struct {
	State wholesaleenum.WholesaleApplicationState `json:"state"`

	UserID                    string `json:"user_id,omitempty"`
	WholesaleOrganisationCode string `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string `json:"organisation_access_id,omitempty"`

	OrganisationStatus wholesaleenum.WholesaleOrganisationStatus `json:"organisation_status,omitempty"`
	AccessStatus       wholesaleenum.OrganisationAccessStatus    `json:"access_status,omitempty"`

	EmailVerificationRequired bool   `json:"email_verification_required"`
	Message                   string `json:"message,omitempty"`
}

// WholesaleApplicationStatusResponse reports the current user's trade-account
// application and approval status.
type WholesaleApplicationStatusResponse struct {
	WholesaleApplicationResponse

	BusinessName string `json:"business_name,omitempty"`
	TradingName  string `json:"trading_name,omitempty"`
	ContactName  string `json:"contact_name,omitempty"`
	Email        string `json:"email,omitempty"`
}

// WholesaleApplicationDecisionRequest records an admin approval/rejection note.
type WholesaleApplicationDecisionRequest struct {
	Reason string `json:"reason,omitempty"`
}

// WholesaleApplicationDecisionResponse is returned by admin approval actions.
type WholesaleApplicationDecisionResponse struct {
	WholesaleApplicationStatusResponse
}

// WholesaleApplicationReviewListItem is the compact admin review projection for
// one trade-account application. It keeps all lifecycle ids/statuses together
// with the business/contact fields needed by the review queue.
type WholesaleApplicationReviewListItem struct {
	WholesaleApplicationStatusResponse

	BusinessType        string `json:"business_type,omitempty"`
	ContactPhone        string `json:"contact_phone,omitempty"`
	ContactRole         string `json:"contact_role,omitempty"`
	ABN                 string `json:"abn,omitempty"`
	ExpectedOrderVolume string `json:"expected_order_volume,omitempty"`
	PaymentPreference   string `json:"payment_preference,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// WholesaleApplicationReviewDetail is the full admin review projection for one
// organisation application, including application notes and delivery metadata.
type WholesaleApplicationReviewDetail struct {
	WholesaleApplicationReviewListItem

	BusinessAddress       *common.ContactAddress `json:"business_address,omitempty"`
	DeliveryAddress       *common.ContactAddress `json:"delivery_address,omitempty"`
	StorageRequirements   []string               `json:"storage_requirements,omitempty"`
	PreferredDeliveryDays []string               `json:"preferred_delivery_days,omitempty"`
	Notes                 string                 `json:"notes,omitempty"`
}
