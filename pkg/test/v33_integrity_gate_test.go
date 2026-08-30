package pkg_test

import (
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// TestV33DomainOwnershipPackagesArePopulated locks the reviewed ownership
// taxonomy. Each listed package must contain production Go source; therefore
// adding an empty taxonomy directory does not create a contract package.
func TestV33DomainOwnershipPackagesArePopulated(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	requiredPackages := []string{
		"contracts/identity/access", "contracts/identity/account", "contracts/identity/authorisation",
		"contracts/customers/retail", "contracts/customers/wholesale", "contracts/customers/group", "contracts/customers/preference",
		"contracts/orders/cart", "contracts/orders/order", "contracts/orders/group_order", "contracts/orders/fulfilment", "contracts/orders/shipping", "contracts/orders/buyer", "contracts/orders/subscription",
		"contracts/payments/payment", "contracts/payments/provider", "contracts/payments/merchant", "contracts/payments/receipt", "contracts/payments/register", "contracts/payments/settlement", "contracts/payments/terminal",
		"contracts/pricing/pricebook", "contracts/pricing/quote", "contracts/pricing/promotion", "contracts/pricing/benefit", "contracts/pricing/coupon", "contracts/pricing/membership", "contracts/pricing/market", "contracts/pricing/special",
		"contracts/pricing/wallet/balance", "contracts/pricing/wallet/ledger", "contracts/pricing/wallet/points", "contracts/pricing/wallet/giftcard", "contracts/pricing/wallet/reward", "contracts/pricing/wallet/reservation",
		"contracts/notification/core", "contracts/notification/email", "contracts/notification/sms", "contracts/notification/push", "contracts/notification/preference", "contracts/notification/delivery",
		"contracts/insights/analytics", "contracts/insights/sales", "contracts/insights/customer",
		"contracts/marketing/campaign", "contracts/marketing/audience", "contracts/marketing/message",
		"contracts/supply/catalogue/classification", "contracts/supply/catalogue/product", "contracts/supply/catalogue/listing", "contracts/supply/catalogue/review", "contracts/supply/catalogue/wish", "contracts/supply/catalogue/favourite",
		"contracts/supply/inventory", "contracts/supply/warehouse", "contracts/supply/procurement", "contracts/supply/compliance", "contracts/supply/fulfilment", "contracts/supply/forecasting",
		"contracts/pubsub/envelope", "contracts/pubsub/routing", "contracts/pubsub/orders", "contracts/pubsub/payments", "contracts/pubsub/supply", "contracts/pubsub/customers", "contracts/pubsub/pricing", "contracts/pubsub/notification",
	}
	for _, packagePath := range requiredPackages {
		matches, err := filepath.Glob(filepath.Join(pkgRoot, filepath.FromSlash(packagePath), "*.go"))
		if err != nil {
			t.Fatalf("list %s source: %v", packagePath, err)
		}
		for _, match := range matches {
			if !strings.HasSuffix(match, "_test.go") {
				goto populated
			}
		}
		t.Errorf("required v33 domain package %s has no production source", packagePath)
	populated:
	}

	for _, forbiddenPackage := range []string{
		"contracts/payments/invoice", "contracts/payments/refund", "contracts/payments/finance",
	} {
		matches, err := filepath.Glob(filepath.Join(pkgRoot, filepath.FromSlash(forbiddenPackage), "*.go"))
		if err != nil {
			t.Fatalf("list %s source: %v", forbiddenPackage, err)
		}
		if len(matches) != 0 {
			t.Errorf("unowned v33 package %s must not contain contract source", forbiddenPackage)
		}
	}
}

// TestV33ContractImportGraphIsAcyclic provides a source-level ownership gate
// in addition to Go's compiler check, so an illegal cross-domain cycle has a
// concise package-path diagnostic.
func TestV33ContractImportGraphIsAcyclic(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	contractsRoot := filepath.Join(pkgRoot, "contracts")
	graph := make(map[string]map[string]struct{})

	err := filepath.WalkDir(contractsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		packagePath := "contracts/" + relativePkgPath(t, contractsRoot, filepath.Dir(path))
		if graph[packagePath] == nil {
			graph[packagePath] = make(map[string]struct{})
		}
		file, parseErr := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
		if parseErr != nil {
			return parseErr
		}
		for _, imported := range file.Imports {
			importPath := strings.Trim(imported.Path.Value, `"`)
			contractPath := strings.TrimPrefix(importPath, contractImportPrefix)
			if contractPath == importPath || !strings.HasPrefix(contractPath, "contracts/") {
				continue
			}
			graph[packagePath][contractPath] = struct{}{}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan contract imports: %v", err)
	}

	state := make(map[string]uint8, len(graph))
	stack := make([]string, 0, len(graph))
	var visit func(string)
	visit = func(packagePath string) {
		if state[packagePath] == 1 {
			cycleStart := 0
			for index, member := range stack {
				if member == packagePath {
					cycleStart = index
					break
				}
			}
			cycle := append(append([]string(nil), stack[cycleStart:]...), packagePath)
			t.Fatalf("v33 contract import cycle: %s", strings.Join(cycle, " -> "))
		}
		if state[packagePath] == 2 {
			return
		}
		state[packagePath] = 1
		stack = append(stack, packagePath)
		dependencies := make([]string, 0, len(graph[packagePath]))
		for dependency := range graph[packagePath] {
			dependencies = append(dependencies, dependency)
		}
		sort.Strings(dependencies)
		for _, dependency := range dependencies {
			visit(dependency)
		}
		stack = stack[:len(stack)-1]
		state[packagePath] = 2
	}

	packages := make([]string, 0, len(graph))
	for packagePath := range graph {
		packages = append(packages, packagePath)
	}
	sort.Strings(packages)
	for _, packagePath := range packages {
		visit(packagePath)
	}
}
