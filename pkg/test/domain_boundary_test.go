package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/customers/retail"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/identity/account"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/notifications"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/orders/shipping"
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

// TestPermissionKeysRemainOpenCodes keeps each permission catalogue in the
// backend that owns it: workforce keys in Backend-Identity and buyer-portal
// keys in Backend-Customers. The contract may carry the open PermissionKey and
// WholesalePermission strings, but neither domain may reacquire a retired
// constant block or a fixed dotted permission wire literal.
func TestPermissionKeysRemainOpenCodes(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	for _, retiredPath := range []string{
		"contracts/identity/role/role_enums/permission_key.go",
		"contracts/customers/wholesale/wholesale_enums/wholesale_permission.go",
	} {
		if _, statErr := os.Stat(filepath.Join(pkgRoot, filepath.FromSlash(retiredPath))); statErr == nil {
			t.Errorf("retired production path remains: %s", retiredPath)
		} else if !os.IsNotExist(statErr) {
			t.Errorf("inspect retired production path %s: %v", retiredPath, statErr)
		}
	}

	permissionConstantPattern := regexp.MustCompile(`^(PermissionKey|WholesalePermission)[A-Z][A-Za-z0-9_]*$`)
	permissionWirePattern := regexp.MustCompile(`^[a-z][a-z0-9_]*(\.[a-z][a-z0-9_]*)+$`)
	permissionCodeTypes := map[string]struct{}{"PermissionKey": {}, "WholesalePermission": {}}
	permissionScopes := []string{"contracts/identity/role/", "contracts/customers/wholesale/"}
	var violations []string

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
		relativePath := relativePkgPath(t, pkgRoot, path)
		inPermissionScope := false
		for _, scope := range permissionScopes {
			if strings.HasPrefix(relativePath, scope) {
				inPermissionScope = true
				break
			}
		}
		ast.Inspect(file, func(node ast.Node) bool {
			switch typed := node.(type) {
			case *ast.Ident:
				if permissionConstantPattern.MatchString(typed.Name) {
					violations = append(violations, relativePath+": hard-coded permission key identifier "+typed.Name)
				}
			case *ast.GenDecl:
				if typed.Tok != token.CONST {
					return true
				}
				for _, spec := range typed.Specs {
					valueSpec, isValueSpec := spec.(*ast.ValueSpec)
					if !isValueSpec {
						continue
					}
					var declaredType string
					switch typeExpression := valueSpec.Type.(type) {
					case *ast.Ident:
						declaredType = typeExpression.Name
					case *ast.SelectorExpr:
						declaredType = typeExpression.Sel.Name
					}
					if _, isPermissionCode := permissionCodeTypes[declaredType]; !isPermissionCode {
						continue
					}
					for _, name := range valueSpec.Names {
						violations = append(violations, relativePath+": hard-coded "+declaredType+" constant "+name.Name)
					}
				}
			case *ast.BasicLit:
				if !inPermissionScope || typed.Kind != token.STRING {
					return true
				}
				literal, unquoteErr := strconv.Unquote(typed.Value)
				if unquoteErr == nil && permissionWirePattern.MatchString(literal) {
					violations = append(violations, relativePath+": fixed permission wire literal "+literal)
				}
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("scan permission boundary: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("permission catalogues are owned and seeded by Backend-Identity and Backend-Customers, not the contract:\n%s", strings.Join(violations, "\n"))
	}
}
