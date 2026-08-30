package pkg_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/retail"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/group"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/wholesale"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/identity/access"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/identity/account"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/identity/authorisation"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notification/preference"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/shipping"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/merchant"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/receipt"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/register"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/settlement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/coupon"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/balance"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/giftcard"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/ledger"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/reservation"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/reward"
	notificationevent "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/notification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/review"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/wish"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/forecasting"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/inventory"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/procurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse"
)

type typePolicy struct {
	model           reflect.Type
	requireAudit    bool
	requireProtected bool
}

// v33TypePolicyRegistry is the reviewed policy for mutable persisted roots,
// PII-bearing records, immutable ledgers, snapshots, projections, and events.
// It is intentionally explicit: a new root must be reviewed instead of
// inheriting audit or privacy metadata by package convention.
var v33TypePolicyRegistry = map[string]typePolicy{
	"notification preference": {reflect.TypeOf(preference.NotificationPreferences{}), true, true},
	"payment":                 {reflect.TypeOf(payment.Payment{}), true, true},
	"outbound shipment":       {reflect.TypeOf(fulfilment.OutboundShipment{}), true, true},
	"merchant legal profile":  {reflect.TypeOf(merchant.MerchantLegalProfile{}), true, true},
	"settlement":              {reflect.TypeOf(settlement.Settlement{}), true, true},
	"brand":                   {reflect.TypeOf(classification.Brand{}), true, false},
	"stock location":          {reflect.TypeOf(warehouse.StockLocation{}), true, false},
	"wish proposal":           {reflect.TypeOf(wish.WishProposal{}), true, false},
	"wish candidate":          {reflect.TypeOf(wish.WishCandidate{}), true, false},
	"wish ballot":             {reflect.TypeOf(wish.WishBallot{}), true, false},
	"checkout reservation":    {reflect.TypeOf(reservation.CheckoutBenefitReservation{}), true, true},
	"point reservation":       {reflect.TypeOf(reservation.PointReservation{}), true, true},
	"coupon assignment":       {reflect.TypeOf(coupon.CouponAssignment{}), true, true},
	"reward redemption":       {reflect.TypeOf(reward.RewardRedemption{}), true, true},
	"shipping zone":           {reflect.TypeOf(shipping.Zone{}), true, false},
	"shipping rate":           {reflect.TypeOf(shipping.Rate{}), true, false},
	"arrival blacklist":       {reflect.TypeOf(shipping.ShippingArrivalBlacklist{}), true, false},
	"SKU demand forecast":     {reflect.TypeOf(forecasting.SKUDemandForecast{}), true, false},
	"retail customer":         {reflect.TypeOf(retail.RetailCustomer{}), true, true},
	"group manager application": {reflect.TypeOf(group.GroupOrderManagerApplication{}), true, true},
	"wholesale organisation":  {reflect.TypeOf(wholesale.WholesaleOrganisation{}), true, true},
	"wholesale application":   {reflect.TypeOf(wholesale.WholesaleApplication{}), true, true},
	"organisation access":     {reflect.TypeOf(wholesale.OrganisationAccess{}), true, true},
	"user account":            {reflect.TypeOf(account.UserAccount{}), true, true},
	"retail account profile":  {reflect.TypeOf(account.RetailCustomerAccountProfile{}), true, true},
	"wholesale account profile": {reflect.TypeOf(account.WholesaleCustomerAccountProfile{}), true, true},
	"user profile":            {reflect.TypeOf(account.UserProfile{}), true, true},
	"auth identity":           {reflect.TypeOf(account.AuthIdentity{}), true, true},
	"portal access":           {reflect.TypeOf(access.PortalAccess{}), true, true},
	"role assignment":         {reflect.TypeOf(authorisation.RoleAssignment{}), true, true},
	"membership account":      {reflect.TypeOf(membership.MembershipAccount{}), true, true},
	"customer wallet":         {reflect.TypeOf(balance.CustomerWallet{}), false, true},
	"login session":           {reflect.TypeOf(access.LoginSession{}), false, true},
	"user device":             {reflect.TypeOf(account.UserDevice{}), false, true},
	"qualifying spend ledger": {reflect.TypeOf(membership.QualifyingSpendLedgerEntry{}), false, true},
	"point ledger":            {reflect.TypeOf(ledger.PointLedgerEntry{}), false, true},
	"gift card transaction":   {reflect.TypeOf(giftcard.GiftCardTransaction{}), false, true},
	"carrying cost movement":  {reflect.TypeOf(procurement.CarryingCostMovement{}), false, true},
	"stock movement":          {reflect.TypeOf(inventory.StockMovement{}), false, true},
	"cash movement":           {reflect.TypeOf(register.CashMovement{}), false, true},
	"receipt snapshot":        {reflect.TypeOf(receipt.ReceiptSnapshot{}), false, true},
	"register session":        {reflect.TypeOf(register.RegisterSession{}), true, true},
	"published review":        {reflect.TypeOf(review.PublishedReview{}), false, false},
	"preference event":        {reflect.TypeOf(notificationevent.NotificationPreferencesChangedEvent{}), false, false},
}

func TestV33AuditAndPrivacyTypePolicy(t *testing.T) {
	for name, policy := range v33TypePolicyRegistry {
		hasAudit := hasDirectEmbeddedField(policy.model, reflect.TypeOf(audit.AuditFields{}))
		if hasAudit != policy.requireAudit {
			t.Errorf("%s AuditFields = %v, want %v", name, hasAudit, policy.requireAudit)
		}
		hasProtection := hasDirectEmbeddedField(policy.model, reflect.TypeOf(security.DataProtectionFields{}))
		if hasProtection != policy.requireProtected {
			t.Errorf("%s DataProtectionFields = %v, want %v", name, hasProtection, policy.requireProtected)
		}
	}
}

func TestLifecycleAndConsentEvidencePolicy(t *testing.T) {
	lifecycle := reflect.TypeOf(audit.LifecycleAction{})
	for _, fieldName := range []string{"By", "At"} {
		field, ok := lifecycle.FieldByName(fieldName)
		if !ok || field.Tag.Get("json") == "" || strings.Contains(field.Tag.Get("json"), "omitempty") {
			t.Errorf("LifecycleAction.%s must be required", fieldName)
		}
	}
	if field, _ := lifecycle.FieldByName("At"); field.Type != reflect.TypeOf(time.Time{}) {
		t.Errorf("LifecycleAction.At = %s, want time.Time", field.Type)
	}

	consent := reflect.TypeOf(preference.NotificationChannelConsent{})
	for _, fieldName := range []string{"Actor", "Source", "PolicyVersion", "RequestID", "ChangedAt"} {
		field, ok := consent.FieldByName(fieldName)
		if !ok || field.Tag.Get("json") == "" || strings.Contains(field.Tag.Get("json"), "omitempty") {
			t.Errorf("NotificationChannelConsent.%s must be required", fieldName)
		}
	}
	if field, _ := consent.FieldByName("Actor"); field.Type != reflect.TypeOf(security.ActorRef{}) {
		t.Errorf("NotificationChannelConsent.Actor = %s, want security.ActorRef", field.Type)
	}

	event := reflect.TypeOf(notificationevent.NotificationPreferencesChangedEvent{})
	var fields []string
	for index := 0; index < event.NumField(); index++ {
		fields = append(fields, event.Field(index).Name)
	}
	sort.Strings(fields)
	want := []string{"AccountID", "ChangedChannels", "ChangedTopicCodes", "CustomerNumber", "PreferencesRevision", "UserID"}
	if !reflect.DeepEqual(fields, want) {
		t.Errorf("NotificationPreferencesChangedEvent fields = %v, want %v", fields, want)
	}
}

func TestCarryingCostMovementHasImmutableEvidence(t *testing.T) {
	model := reflect.TypeOf(procurement.CarryingCostMovement{})
	for _, fieldName := range []string{"Actor", "RequestID", "CorrelationID"} {
		if _, ok := model.FieldByName(fieldName); !ok {
			t.Errorf("CarryingCostMovement missing %s evidence", fieldName)
		}
	}
	if hasDirectEmbeddedField(model, reflect.TypeOf(audit.AuditFields{})) {
		t.Error("CarryingCostMovement must remain an immutable ledger without AuditFields")
	}
}

func hasDirectEmbeddedField(model, target reflect.Type) bool {
	for index := 0; index < model.NumField(); index++ {
		field := model.Field(index)
		if field.Anonymous && field.Type == target {
			return true
		}
	}
	return false
}
