package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// v30RedesignedModelRoots are the domains reconstructed for the v30 cutover.
// The listed files predate this cutover and deliberately remain composite
// records; any newly touched model file in these domains must hold one
// exported struct only.
var v30RedesignedModelRoots = []string{
	"contracts/customers/retail",
	"contracts/insights/analytics",
	"contracts/insights/marketing",
	"contracts/marketing/campaign",
	"contracts/marketing/message",
	"contracts/notifications",
	"contracts/orders/order",
	"contracts/orders/shipping",
	"contracts/pricing/benefit",
	"contracts/pricing/membership",
	"contracts/pricing/pricebook",
	"contracts/pricing/promotion",
	"contracts/pricing/wallet",
	"contracts/review",
	"contracts/supply/product",
}

var v30PreexistingCompositeModelFiles = map[string]struct{}{
	"contracts/insights/analytics/facts.go":         {},
	"contracts/insights/analytics/forecast.go":      {},
	"contracts/orders/order/demand.go":              {},
	"contracts/orders/order/history.go":             {},
	"contracts/orders/order/package_fulfilment.go":  {},
	"contracts/orders/order/preorder_item.go":       {},
	"contracts/orders/shipping/arrival_rule.go":     {},
	"contracts/orders/shipping/delivery_slots.go":   {},
	"contracts/orders/shipping/shipment.go":         {},
	"contracts/supply/product/sales_performance.go": {},
}

func TestV30ProductionFilesContainAtMostOneStringEnum(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	var violations []string
	err := filepath.WalkDir(filepath.Join(pkgRoot, "contracts"), func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		file, parseErr := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		var stringTypes []string
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpec, ok := specification.(*ast.TypeSpec)
				if underlying, isString := typeSpec.Type.(*ast.Ident); ok && isString && underlying.Name == "string" {
					stringTypes = append(stringTypes, typeSpec.Name.Name)
				}
			}
		}
		if len(stringTypes) > 1 {
			violations = append(violations, relativePkgPath(t, pkgRoot, path)+": multiple string enum types "+strings.Join(stringTypes, ", "))
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan enum file hygiene: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("v30 enum file hygiene violations:\n%s", strings.Join(violations, "\n"))
	}
}

func TestV30RedesignedModelFilesContainOneExportedStruct(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	var violations []string
	for _, root := range v30RedesignedModelRoots {
		root := root
		err := filepath.WalkDir(filepath.Join(pkgRoot, filepath.FromSlash(root)), func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") || strings.HasSuffix(filepath.Base(filepath.Dir(path)), "_enums") {
				return nil
			}
			relative := relativePkgPath(t, pkgRoot, path)
			if _, exempt := v30PreexistingCompositeModelFiles[relative]; exempt {
				return nil
			}
			file, parseErr := parser.ParseFile(token.NewFileSet(), path, nil, 0)
			if parseErr != nil {
				return parseErr
			}
			var structs []string
			for _, declaration := range file.Decls {
				general, ok := declaration.(*ast.GenDecl)
				if !ok || general.Tok != token.TYPE {
					continue
				}
				for _, specification := range general.Specs {
					typeSpec, ok := specification.(*ast.TypeSpec)
					if !ok {
						continue
					}
					if _, isStruct := typeSpec.Type.(*ast.StructType); isStruct && typeSpec.Name.IsExported() {
						structs = append(structs, typeSpec.Name.Name)
					}
				}
			}
			if len(structs) > 1 {
				violations = append(violations, relative+": multiple exported structs "+strings.Join(structs, ", "))
			}
			return nil
		})
		if err != nil {
			t.Fatalf("scan model file hygiene in %s: %v", root, err)
		}
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("v30 model file hygiene violations:\n%s", strings.Join(violations, "\n"))
	}
}
