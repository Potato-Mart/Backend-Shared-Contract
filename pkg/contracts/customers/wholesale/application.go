package wholesale

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/party"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/customers/wholesale/wholesale_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/supply/warehouse/warehouse_enums"
)

// WholesaleApplicantSnapshot freezes the applicant identity reviewed by an
// administrator without making Identity data part of the application key.
type WholesaleApplicantSnapshot struct {
	UserID        string `json:"user_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone,omitempty"`
	Role          string `json:"role,omitempty"`
	EmailVerified bool   `json:"email_verified"`
}

// WholesaleApplication is the durable source of truth for a wholesale access
// request. Decisions are compare-and-set from pending.
type WholesaleApplication struct {
	ID                        string                                     `json:"id"`
	ApplicationNumber         string                                     `json:"application_number"`
	Applicant                 WholesaleApplicantSnapshot                 `json:"applicant"`
	Organisation              party.OrganisationDetail                   `json:"organisation"`
	ExpectedMonthlyOrders     int                                        `json:"expected_monthly_orders,omitempty"`
	ExpectedMonthlySpend      *money.Money                               `json:"expected_monthly_spend,omitempty"`
	StorageRequirements       []warehouse_enums.StorageType              `json:"storage_requirements,omitempty"`
	PaymentPreference         string                                     `json:"payment_preference,omitempty"`
	PreferredDeliveryDays     []int                                      `json:"preferred_delivery_days,omitempty"`
	Notes                     string                                     `json:"notes,omitempty"`
	Status                    wholesale_enums.WholesaleApplicationStatus `json:"status"`
	WholesaleOrganisationCode string                                     `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string                                     `json:"organisation_access_id,omitempty"`
	SubmittedAt               time.Time                                  `json:"submitted_at"`
	Approval                  *audit.LifecycleAction                     `json:"approval,omitempty"`
	Rejection                 *audit.LifecycleAction                     `json:"rejection,omitempty"`

	audit.AuditFields
}
