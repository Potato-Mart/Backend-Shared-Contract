package enums_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/wholesale/wholesale_enums"
	insights_marketing_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/insights/marketing/marketing_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/marketing/campaign/campaign_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/pos/pos_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/benefit/benefit_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/event/event_enums"
)

func TestEnumCoverageIncludesEveryStringEnum(t *testing.T) {
	testSource := enumTestSource(t)
	for _, typeName := range stringEnumTypes(t) {
		if !regexp.MustCompile(`\b` + regexp.QuoteMeta(typeName) + `\b`).MatchString(testSource) {
			t.Errorf("string enum %s is not represented in an enum behavior test", typeName)
		}
	}
	for _, constantName := range stringEnumConstants(t) {
		if !regexp.MustCompile(`\b` + regexp.QuoteMeta(constantName) + `\b`).MatchString(testSource) {
			t.Errorf("string enum constant %s is not represented in an enum behavior test", constantName)
		}
	}
}

func TestAdditionalEnumValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "campaign.CampaignStatus", valid: []stringEnum{campaign_enums.CampaignStatusDraft, campaign_enums.CampaignStatusScheduled, campaign_enums.CampaignStatusActive, campaign_enums.CampaignStatusCompleted, campaign_enums.CampaignStatusArchived}, invalid: campaign_enums.CampaignStatus("__invalid__")},
		{name: "insights.CampaignPredictionStatus", valid: []stringEnum{insights_marketing_enums.CampaignPredictionStatusNotApplicable, insights_marketing_enums.CampaignPredictionStatusReady, insights_marketing_enums.CampaignPredictionStatusWarning}, invalid: insights_marketing_enums.CampaignPredictionStatus("__invalid__")},
		{name: "insights.CampaignPredictionSource", valid: []stringEnum{insights_marketing_enums.CampaignPredictionSourceSameSeries, insights_marketing_enums.CampaignPredictionSourceSimilarEvent, insights_marketing_enums.CampaignPredictionSourceLast14DaysDoubled}, invalid: insights_marketing_enums.CampaignPredictionSource("__invalid__")},
		{name: "wholesale.WholesaleApplicationStatus", valid: []stringEnum{wholesale_enums.WholesaleApplicationStatusPending, wholesale_enums.WholesaleApplicationStatusApproved, wholesale_enums.WholesaleApplicationStatusRejected}, invalid: wholesale_enums.WholesaleApplicationStatus("__invalid__")},
		{name: "wholesale.GroupOrderManagerApplicationStatus", valid: []stringEnum{wholesale_enums.GroupOrderManagerApplicationStatusPending, wholesale_enums.GroupOrderManagerApplicationStatusApproved, wholesale_enums.GroupOrderManagerApplicationStatusRejected, wholesale_enums.GroupOrderManagerApplicationStatusCancelled}, invalid: wholesale_enums.GroupOrderManagerApplicationStatus("__invalid__")},
		{name: "order.CustomerOrderBucket", valid: []stringEnum{order_enums.CustomerOrderBucketCurrent, order_enums.CustomerOrderBucketCompleted, order_enums.CustomerOrderBucketCancelled, order_enums.CustomerOrderBucketRefunded}, invalid: order_enums.CustomerOrderBucket("__invalid__")},
		{name: "pos.SessionStatus", valid: []stringEnum{pos_enums.SessionStatusOpen, pos_enums.SessionStatusClosed}, invalid: pos_enums.SessionStatus("__invalid__")},
		{name: "pos.CashMovementKind", valid: []stringEnum{pos_enums.CashMovementKindCashIn, pos_enums.CashMovementKindCashOut, pos_enums.CashMovementKindFloatAdjustment, pos_enums.CashMovementKindDrop}, invalid: pos_enums.CashMovementKind("__invalid__")},
		{name: "payment.CustomerPaymentAllocationKind", valid: []stringEnum{payment_enums.CustomerPaymentAllocationKindExternalPayment, payment_enums.CustomerPaymentAllocationKindWalletBalance, payment_enums.CustomerPaymentAllocationKindWalletPoints, payment_enums.CustomerPaymentAllocationKindGiftCard, payment_enums.CustomerPaymentAllocationKindVoucher, payment_enums.CustomerPaymentAllocationKindOther}, invalid: payment_enums.CustomerPaymentAllocationKind("__invalid__")},
		{name: "payment.PaymentCompleteness", valid: []stringEnum{payment_enums.PaymentCompletenessComplete, payment_enums.PaymentCompletenessPartial, payment_enums.PaymentCompletenessUnavailable}, invalid: payment_enums.PaymentCompleteness("__invalid__")},
		{name: "payment.PaymentSummaryComponent", valid: []stringEnum{payment_enums.PaymentSummaryComponentAllocationHistory, payment_enums.PaymentSummaryComponentPaymentTotals, payment_enums.PaymentSummaryComponentRedemptionTimestamps, payment_enums.PaymentSummaryComponentPointsEarned}, invalid: payment_enums.PaymentSummaryComponent("__invalid__")},
		{name: "benefit.OwnerType", valid: []stringEnum{benefit_enums.OwnerTypeRetailCustomer, benefit_enums.OwnerTypeWholesaleOrganisation}, invalid: benefit_enums.OwnerType("__invalid__")},
		{name: "membership.QualifyingSpendReason", valid: []stringEnum{membership_enums.QualifyingSpendReasonOrderPaid, membership_enums.QualifyingSpendReasonRefund}, invalid: membership_enums.QualifyingSpendReason("__invalid__")},
		{name: "membership.TierProgressReason", valid: []stringEnum{membership_enums.TierProgressReasonNoActiveTiers, membership_enums.TierProgressReasonManualQualification, membership_enums.TierProgressReasonUnsupportedMetric, membership_enums.TierProgressReasonCurrencyMismatch, membership_enums.TierProgressReasonMembershipNotAssigned}, invalid: membership_enums.TierProgressReason("__invalid__")},
		{name: "membership.TierBenefitKind", valid: []stringEnum{membership_enums.TierBenefitKindQualifyingSpend, membership_enums.TierBenefitKindPointsMultiplier, membership_enums.TierBenefitKindDiscountPercent, membership_enums.TierBenefitKindFreeShippingThreshold, membership_enums.TierBenefitKindBirthdayBonusPoints}, invalid: membership_enums.TierBenefitKind("__invalid__")},
		{name: "event.EventTopic", valid: []stringEnum{event_enums.EventTopicOrderEvents, event_enums.EventTopicPaymentEvents, event_enums.EventTopicRefundEvents, event_enums.EventTopicStockEvents, event_enums.EventTopicFulfilmentEvents, event_enums.EventTopicCustomerEvents, event_enums.EventTopicProductStats, event_enums.EventTopicStorefrontEvents, event_enums.EventTopicCatalogEvents}, invalid: event_enums.EventTopic("__invalid__")},
		{name: "security.MediaVisibility", valid: []stringEnum{security_enums.MediaVisibilityPublic, security_enums.MediaVisibilityPrivate}, invalid: security_enums.MediaVisibility("__invalid__")},
		{name: "packaging.PackageHandlingUnit", valid: []stringEnum{packaging_enums.PackageHandlingUnitEach, packaging_enums.PackageHandlingUnitCase}, invalid: packaging_enums.PackageHandlingUnit("__invalid__")},
	})
}

func enumTestSource(t *testing.T) string {
	t.Helper()
	pkgRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("resolve pkg test root: %v", err)
	}
	var source []string
	err = filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		contents, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		source = append(source, string(contents))
		return nil
	})
	if err != nil {
		t.Fatalf("read enum tests: %v", err)
	}
	return strings.Join(source, "\n")
}

func stringEnumTypes(t *testing.T) []string {
	t.Helper()
	contractsRoot, err := filepath.Abs(filepath.Join("..", "..", "contracts"))
	if err != nil {
		t.Fatalf("resolve contracts root: %v", err)
	}
	var types []string
	err = filepath.WalkDir(contractsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || entry.Name() == "" || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") || !strings.HasSuffix(filepath.Dir(path), "_enums") {
			return nil
		}
		contents, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		for _, match := range regexp.MustCompile(`(?m)^type\s+([A-Za-z0-9_]+)\s+string\s*$`).FindAllStringSubmatch(string(contents), -1) {
			types = append(types, match[1])
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan enum types: %v", err)
	}
	sort.Strings(types)
	return types
}

func stringEnumConstants(t *testing.T) []string {
	t.Helper()
	contractsRoot, err := filepath.Abs(filepath.Join("..", "..", "contracts"))
	if err != nil {
		t.Fatalf("resolve contracts root: %v", err)
	}
	constantPattern := regexp.MustCompile(`(?m)^\s*([A-Za-z0-9_]+)\s+[A-Za-z0-9_]+\s*=\s*"`)
	var constants []string
	err = filepath.WalkDir(contractsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") || !strings.HasSuffix(filepath.Dir(path), "_enums") {
			return nil
		}
		contents, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		for _, match := range constantPattern.FindAllStringSubmatch(string(contents), -1) {
			constants = append(constants, match[1])
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan enum constants: %v", err)
	}
	sort.Strings(constants)
	return constants
}
