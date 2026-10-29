package pkg_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pricing/wallet/wallet_enums"
)

func TestRewardHasNoTierIssuanceConfiguration(t *testing.T) {
	rewardType := reflect.TypeOf(membership.Reward{})
	for _, field := range []string{"TriggerTierKey", "IssueOnTierAchievement"} {
		if _, exists := rewardType.FieldByName(field); exists {
			t.Errorf("retired tier-issuance field %s remains on Reward", field)
		}
	}
}

func TestRejectsRetiredInstrumentKinds(t *testing.T) {
	if wallet_enums.RewardType("VOUCHER").IsValid() ||
		wallet_enums.WalletInstrumentType("voucher").IsValid() ||
		payment_enums.CustomerPaymentAllocationKind("voucher").IsValid() {
		t.Fatal("retired instrument kind remains valid")
	}
	if !wallet_enums.RewardTypeCoupon.IsValid() || !wallet_enums.WalletInstrumentTypeCoupon.IsValid() ||
		!wallet_enums.RewardTypeGiftCard.IsValid() || !payment_enums.CustomerPaymentAllocationKindGiftCard.IsValid() {
		t.Fatal("supported coupon or gift-card kind was retired")
	}
}

func TestProductionContractsHaveNoVoucherSurface(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	err := filepath.WalkDir(filepath.Join(pkgRoot, "contracts"), func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.Contains(strings.ToLower(string(contents)), "voucher") {
			t.Errorf("retired instrument remains in production contract %s", relativePkgPath(t, pkgRoot, path))
		}
		return nil
	})
	if err != nil {
		t.Fatalf("inspect current contract surfaces: %v", err)
	}
}
