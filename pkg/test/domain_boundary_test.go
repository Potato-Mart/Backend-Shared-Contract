package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/customers/retail"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/identity/account"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/notifications"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/shipping"
)

func TestCustomerAndOrderGeographyUsesCanonicalFields(t *testing.T) {
	assertNoModelFields(t, reflect.TypeOf(retail.RetailCustomer{}), "MarketCode", "CountryCode", "Marketing", "Analytics")
	assertNoModelFields(t, reflect.TypeOf(account.UserProfile{}), "NotificationPreferences")
	for _, model := range []reflect.Type{reflect.TypeOf(order.Cart{}), reflect.TypeOf(order.Order{})} {
		field, found := model.FieldByName("FulfilmentLocation")
		if !found || field.Type != reflect.TypeOf(shipping.FulfilmentLocationSnapshot{}) || field.Tag.Get("json") != "fulfilment_location" {
			t.Errorf("%s must carry the canonical fulfilment_location snapshot", model)
		}
		assertNoModelFields(t, model, "GeographicContext", "Shipping")
	}
}

func TestPublishedNotificationContainsOnlyInAppSafeFields(t *testing.T) {
	assertNoModelFields(t, reflect.TypeOf(notifications.PublishedNotification{}),
		"Recipient", "DestinationCode", "ProviderCode", "Deliveries", "ErrorCode", "ErrorMessage", "Email", "Push", "SMS", "SocialMedia")
}

// TestNotificationProvidersRemainOpenCodes keeps provider implementations and
// fixed identifiers in the owning backend. The notification model may carry an
// open provider_code, but it may not acquire a provider enum or fixed
// Line/Discord wire literal.
func TestNotificationProvidersRemainOpenCodes(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	notificationsRoot := filepath.Join(pkgRoot, "contracts", "notifications")
	var violations []string

	err := filepath.WalkDir(notificationsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		ast.Inspect(file, func(node ast.Node) bool {
			switch typed := node.(type) {
			case *ast.TypeSpec:
				if strings.Contains(typed.Name.Name, "Provider") {
					violations = append(violations, relativePkgPath(t, pkgRoot, path)+": notification provider type "+typed.Name.Name)
				}
			case *ast.BasicLit:
				if typed.Kind != token.STRING {
					return true
				}
				literal, unquoteErr := strconv.Unquote(typed.Value)
				if unquoteErr == nil && (strings.EqualFold(literal, "line") || strings.EqualFold(literal, "discord")) {
					violations = append(violations, relativePkgPath(t, pkgRoot, path)+": fixed social provider literal "+literal)
				}
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("scan notification provider boundary: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("notification provider boundary violations:\n%s", strings.Join(violations, "\n"))
	}
}

func TestRetiredNotificationAndCampaignProductionShapesAreAbsent(t *testing.T) {
	retiredSourceRoots := []string{
		"contracts/customers/campaign",
		"contracts/notifications/backinstock",
		"contracts/notifications/customer",
	}
	pkgRoot := sharedContractPkgRoot(t)
	var violations []string
	for _, root := range retiredSourceRoots {
		root := filepath.Join(pkgRoot, filepath.FromSlash(root))
		if _, statErr := os.Stat(root); os.IsNotExist(statErr) {
			continue
		} else if statErr != nil {
			t.Fatalf("inspect retired source root %s: %v", root, statErr)
		}
		err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if !entry.IsDir() && strings.HasSuffix(path, ".go") {
				violations = append(violations, relativePkgPath(t, pkgRoot, path)+": retired production source remains")
			}
			return nil
		})
		if err != nil {
			t.Fatalf("scan retired source root %s: %v", root, err)
		}
	}

	forbiddenIdentifiers := map[string]struct{}{
		"BackInStock":                            {},
		"CheckoutCompensationRequestedEvent":     {},
		"CustomerConsentChangedEvent":            {},
		"EventTypeCheckoutCompensationRequested": {},
		"EventTypeInvoiceDeliveryRequested":      {},
		"EventTypeVoucherClaimIssued":            {},
		"InvoiceDeliveryRequestedEvent":          {},
		"NotificationTopicGroup":                 {},
		"RetailCustomerMarketingProfile":         {},
		"SKUStatus":                              {},
		"TopicGroupPreferences":                  {},
		"VoucherClaimIssuedEvent":                {},
	}
	err := filepath.WalkDir(filepath.Join(pkgRoot, "contracts"), func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		ast.Inspect(file, func(node ast.Node) bool {
			identifier, isIdentifier := node.(*ast.Ident)
			if !isIdentifier {
				return true
			}
			if _, forbidden := forbiddenIdentifiers[identifier.Name]; forbidden {
				violations = append(violations, relativePkgPath(t, pkgRoot, path)+": forbidden retired identifier "+identifier.Name)
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("scan retired identifiers: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("retired-shape violations:\n%s", strings.Join(violations, "\n"))
	}
}

func assertNoModelFields(t *testing.T, model reflect.Type, names ...string) {
	t.Helper()
	for _, name := range names {
		if _, found := model.FieldByName(name); found {
			t.Errorf("%s retains removed field %s", model, name)
		}
	}
}
