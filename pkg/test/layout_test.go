package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

const contractImportPrefix = "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/"

var expectedEnumPackages = []string{
	"contracts/common/commerce/commerce_enums",
	"contracts/common/geography/geography_enums",
	"contracts/common/identity/identity_enums",
	"contracts/common/packaging/packaging_enums",
	"contracts/common/security/security_enums",
	"contracts/customers/retail/retail_enums",
	"contracts/customers/wholesale/wholesale_enums",
	"contracts/identity/access/access_enums",
	"contracts/identity/account/account_enums",
	"contracts/identity/role/role_enums",
	"contracts/insights/analytics/analytics_enums",
	"contracts/insights/marketing/marketing_enums",
	"contracts/marketing/campaign/campaign_enums",
	"contracts/marketing/message/message_enums",
	"contracts/notifications/notification_enums",
	"contracts/orders/order/order_enums",
	"contracts/orders/pos/pos_enums",
	"contracts/orders/shipping/shipping_enums",
	"contracts/payments/payment/payment_enums",
	"contracts/payments/settlement/settlement_enums",
	"contracts/payments/terminal/terminal_enums",
	"contracts/pricing/benefit/benefit_enums",
	"contracts/pricing/market/market_enums",
	"contracts/pricing/membership/membership_enums",
	"contracts/pricing/pricebook/pricebook_enums",
	"contracts/pricing/promotion/promotion_enums",
	"contracts/pricing/quote/quote_enums",
	"contracts/pricing/wallet/wallet_enums",
	"contracts/pubsub/event/event_enums",
	"contracts/supply/classification/classification_enums",
	"contracts/supply/import_compliance/import_compliance_enums",
	"contracts/supply/listing/listing_enums",
	"contracts/supply/product/product_enums",
	"contracts/supply/purchase/purchase_enums",
	"contracts/customers/review/review_enums",
	"contracts/supply/warehouse/warehouse_enums",
	"contracts/supply/wish/wish_enums",
}

// TestPkgTreeContainsOnlyContractsAndRepositoryGates keeps release
// metadata, workflows, and runtime helpers out of pkg. Production contract
// source belongs under contracts; repository-wide test gates belong under
// test. Everything else is backend-local or repository documentation.
func TestPkgTreeContainsOnlyContractsAndRepositoryGates(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	var violations []string

	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		relative := relativePkgPath(t, pkgRoot, path)
		if strings.HasPrefix(relative, "contracts/") || strings.HasPrefix(relative, "test/") {
			return nil
		}
		violations = append(violations, relative+": Go source must live under contracts or test")
		return nil
	})
	if err != nil {
		t.Fatalf("scan pkg tree: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("pkg-tree boundary violations:\n%s", strings.Join(violations, "\n"))
	}
}

func TestRetiredReviewPackageIsAbsent(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	retiredReviewRoot := filepath.Join(pkgRoot, "contracts", "review")
	if _, err := os.Stat(retiredReviewRoot); err == nil {
		t.Errorf("retired Review package directory must not exist: %s", retiredReviewRoot)
	} else if !os.IsNotExist(err) {
		t.Fatalf("inspect retired Review package directory: %v", err)
	}

	canonicalReviewRoot := filepath.Join(pkgRoot, "contracts", "customers", "review")
	if info, err := os.Stat(canonicalReviewRoot); err != nil {
		t.Fatalf("canonical Review package directory is missing: %v", err)
	} else if !info.IsDir() {
		t.Fatalf("canonical Review package path is not a directory: %s", canonicalReviewRoot)
	}

	retiredImportPrefix := contractImportPrefix + "contracts/review"
	var violations []string
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			return nil
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if parseErr != nil {
			return parseErr
		}
		for _, imported := range file.Imports {
			importPath := strings.Trim(imported.Path.Value, `"`)
			if importPath == retiredImportPrefix || strings.HasPrefix(importPath, retiredImportPrefix+"/") {
				violations = append(violations, relativePkgPath(t, pkgRoot, path)+": imports retired Review package "+importPath)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan Review package imports: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("retired Review package violations:\n%s", strings.Join(violations, "\n"))
	}
}

func TestCommonAndEnumPackageLayout(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	contractsRoot := filepath.Join(pkgRoot, "contracts")
	retiredShared := filepath.Join(contractsRoot, "common", "shared")
	if _, err := os.Stat(retiredShared); !os.IsNotExist(err) {
		t.Fatalf("retired common/shared directory must not exist: %v", err)
	}

	expected := make(map[string]struct{}, len(expectedEnumPackages))
	for _, packagePath := range expectedEnumPackages {
		expected[packagePath] = struct{}{}
	}
	found := make(map[string]struct{}, len(expected))
	var violations []string
	err := filepath.WalkDir(contractsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if !strings.HasSuffix(entry.Name(), "_enums") {
				return nil
			}
			// Local tooling can leave an empty directory after a package removal.
			// Only a directory containing Go source is a contract enum package.
			sources, globErr := filepath.Glob(filepath.Join(path, "*.go"))
			if globErr != nil {
				return globErr
			}
			if len(sources) == 0 {
				return nil
			}
			packagePath := "contracts/" + relativePkgPath(t, contractsRoot, path)
			found[packagePath] = struct{}{}
			if _, ok := expected[packagePath]; !ok {
				violations = append(violations, "unexpected enum package "+packagePath)
			}
			return nil
		}
		if !strings.HasSuffix(entry.Name(), ".go") {
			return nil
		}

		contents, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		if strings.Contains(string(contents), "/v23/") || strings.Contains(string(contents), "common/shared") {
			violations = append(violations, relativePkgPath(t, pkgRoot, path)+": retired module or common/shared source path")
		}
		if !strings.HasSuffix(filepath.Base(filepath.Dir(path)), "_enums") {
			return nil
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, contents, 0)
		if parseErr != nil {
			return parseErr
		}
		if got, want := file.Name.Name, filepath.Base(filepath.Dir(path)); got != want {
			violations = append(violations, relativePkgPath(t, pkgRoot, path)+": package "+got+", want "+want)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan package layout: %v", err)
	}
	for packagePath := range expected {
		if _, ok := found[packagePath]; !ok {
			violations = append(violations, "missing required enum package "+packagePath)
		}
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("package layout violations:\n%s", strings.Join(violations, "\n"))
	}
}

func TestEnumTypesExposeIntrinsicValidation(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	contractsRoot := filepath.Join(pkgRoot, "contracts")
	type methods map[string]map[string]struct{}
	byPackage := make(map[string]methods)
	var violations []string

	err := filepath.WalkDir(contractsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") || !strings.HasSuffix(filepath.Base(filepath.Dir(path)), "_enums") {
			return nil
		}
		packagePath := "contracts/" + relativePkgPath(t, contractsRoot, filepath.Dir(path))
		if byPackage[packagePath] == nil {
			byPackage[packagePath] = methods{}
		}
		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		stringTypes := make(map[string]struct{})
		finiteTypes := make(map[string]struct{})
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, spec := range general.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if underlying, isString := typeSpec.Type.(*ast.Ident); ok && isString && underlying.Name == "string" {
					stringTypes[typeSpec.Name.Name] = struct{}{}
				}
			}
		}
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.CONST {
				continue
			}
			for _, spec := range general.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				if enumType, ok := valueSpec.Type.(*ast.Ident); ok {
					if _, stringBacked := stringTypes[enumType.Name]; stringBacked {
						finiteTypes[enumType.Name] = struct{}{}
					}
				}
			}
		}
		for _, declaration := range file.Decls {
			function, ok := declaration.(*ast.FuncDecl)
			if !ok || function.Recv == nil || len(function.Recv.List) != 1 {
				continue
			}
			receiver, ok := function.Recv.List[0].Type.(*ast.Ident)
			if !ok {
				continue
			}
			if _, stringBacked := stringTypes[receiver.Name]; stringBacked {
				if byPackage[packagePath][receiver.Name] == nil {
					byPackage[packagePath][receiver.Name] = make(map[string]struct{})
				}
				byPackage[packagePath][receiver.Name][function.Name.Name] = struct{}{}
			}
		}
		for enumType := range finiteTypes {
			for _, method := range []string{"String", "IsValid"} {
				if _, ok := byPackage[packagePath][enumType][method]; !ok {
					violations = append(violations, packagePath+"."+enumType+" is missing "+method)
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan enum methods: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("enum validation violations:\n%s", strings.Join(violations, "\n"))
	}
}

func TestCommonPackageDependencyGraph(t *testing.T) {
	expected := map[string][]string{
		"contracts/common/audit":                     {},
		"contracts/common/commerce/commerce_enums":   {},
		"contracts/common/device":                    {},
		"contracts/common/geography":                 {"contracts/common/geography/geography_enums"},
		"contracts/common/geography/geography_enums": {},
		"contracts/common/geometry":                  {},
		"contracts/common/identity":                  {},
		"contracts/common/identity/identity_enums":   {},
		"contracts/common/localization":              {},
		"contracts/common/measurement":               {},
		"contracts/common/metadata":                  {},
		"contracts/common/money":                     {},
		"contracts/common/packaging":                 {"contracts/common/measurement", "contracts/common/packaging/packaging_enums"},
		"contracts/common/packaging/packaging_enums": {},
		"contracts/common/party":                     {"contracts/common/geography", "contracts/common/metadata", "contracts/common/money", "contracts/common/security"},
		"contracts/common/security":                  {"contracts/common/audit", "contracts/common/metadata", "contracts/common/security/security_enums", "contracts/identity/role/role_enums"},
		"contracts/common/security/security_enums":   {},
		"contracts/common/temporal":                  {},
	}

	pkgRoot := sharedContractPkgRoot(t)
	commonRoot := filepath.Join(pkgRoot, "contracts", "common")
	actual := make(map[string]map[string]struct{}, len(expected))
	err := filepath.WalkDir(commonRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		packagePath := "contracts/common/" + relativePkgPath(t, commonRoot, filepath.Dir(path))
		if actual[packagePath] == nil {
			actual[packagePath] = make(map[string]struct{})
		}
		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if parseErr != nil {
			return parseErr
		}
		for _, imported := range file.Imports {
			importPath := strings.Trim(imported.Path.Value, `"`)
			if strings.HasPrefix(importPath, contractImportPrefix) {
				actual[packagePath][strings.TrimPrefix(importPath, contractImportPrefix)] = struct{}{}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("read common dependency graph: %v", err)
	}

	var violations []string
	for packagePath, allowed := range expected {
		got := stringSetKeys(actual[packagePath])
		want := append([]string(nil), allowed...)
		sort.Strings(want)
		if strings.Join(got, ",") != strings.Join(want, ",") {
			violations = append(violations, packagePath+" imports ["+strings.Join(got, ",")+"], want ["+strings.Join(want, ",")+"]")
		}
	}
	for packagePath := range actual {
		if _, known := expected[packagePath]; !known {
			violations = append(violations, "unclassified common package "+packagePath)
		}
	}
	if cycle := dependencyCycle(actual); len(cycle) > 0 {
		violations = append(violations, "common package dependency cycle: "+strings.Join(cycle, " -> "))
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("common dependency graph violations:\n%s", strings.Join(violations, "\n"))
	}
}

func stringSetKeys(values map[string]struct{}) []string {
	keys := make([]string, 0, len(values))
	for value := range values {
		keys = append(keys, value)
	}
	sort.Strings(keys)
	return keys
}

func dependencyCycle(graph map[string]map[string]struct{}) []string {
	const (
		unvisited = iota
		visiting
		visited
	)
	states := make(map[string]int, len(graph))
	var stack []string
	var visit func(string) []string
	visit = func(node string) []string {
		states[node] = visiting
		stack = append(stack, node)
		for _, dependency := range stringSetKeys(graph[node]) {
			if _, internal := graph[dependency]; !internal {
				continue
			}
			if states[dependency] == visiting {
				for index, candidate := range stack {
					if candidate == dependency {
						return append(append([]string(nil), stack[index:]...), dependency)
					}
				}
			}
			if states[dependency] == unvisited {
				if cycle := visit(dependency); len(cycle) > 0 {
					return cycle
				}
			}
		}
		stack = stack[:len(stack)-1]
		states[node] = visited
		return nil
	}
	for _, node := range graphNodes(graph) {
		if states[node] == unvisited {
			if cycle := visit(node); len(cycle) > 0 {
				return cycle
			}
		}
	}
	return nil
}

func graphNodes(graph map[string]map[string]struct{}) []string {
	nodes := make([]string, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}
	sort.Strings(nodes)
	return nodes
}
