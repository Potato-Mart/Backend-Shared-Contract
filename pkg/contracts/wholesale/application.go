package wholesale

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
	wholesaleenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/wholesale"
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
	ID                        string                                   `json:"id"`
	ApplicationNumber         string                                   `json:"application_number"`
	Applicant                 WholesaleApplicantSnapshot               `json:"applicant"`
	Organisation              common.OrganisationDetail                `json:"organisation"`
	ExpectedMonthlyOrders     int                                      `json:"expected_monthly_orders,omitempty"`
	ExpectedMonthlySpend      *common.Money                            `json:"expected_monthly_spend,omitempty"`
	StorageRequirements       []warehouseenum.StorageType              `json:"storage_requirements,omitempty"`
	PaymentPreference         string                                   `json:"payment_preference,omitempty"`
	PreferredDeliveryDays     []int                                    `json:"preferred_delivery_days,omitempty"`
	Notes                     string                                   `json:"notes,omitempty"`
	Status                    wholesaleenum.WholesaleApplicationStatus `json:"status"`
	WholesaleOrganisationCode string                                   `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string                                   `json:"organisation_access_id,omitempty"`
	SubmittedAt               time.Time                                `json:"submitted_at"`
	Approval                  *common.LifecycleAction                  `json:"approval,omitempty"`
	Rejection                 *common.LifecycleAction                  `json:"rejection,omitempty"`

	common.AuditFields
}
