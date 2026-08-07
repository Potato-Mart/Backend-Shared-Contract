package pkg_test

import (
	"crypto/sha256"
	"encoding/hex"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// v23ModelPackageManifest classifies every production package. Adding an
// exported type changes the digest below and requires an explicit manifest
// review instead of silently expanding the shared module.
var v23ModelPackageManifest = map[string]string{
	"contracts/common/apiresponse":        "enum",
	"contracts/common/geography":          "enum,value",
	"contracts/common/security":           "enum,event,record,value",
	"contracts/common/shared":             "event,record,value",
	"contracts/customers/campaign":        "entity,event,record,snapshot,enum",
	"contracts/customers/retail":          "entity,event,record,snapshot,enum",
	"contracts/customers/wholesale":       "entity,record,snapshot,enum",
	"contracts/identity/access":           "claims,event,record,session,enum",
	"contracts/identity/account":          "entity,event,record,enum",
	"contracts/identity/role":             "entity,event,record,enum",
	"contracts/insights/analytics":        "record",
	"contracts/insights/marketing":        "entity,record,enum",
	"contracts/notifications/backinstock": "entity,event,record,enum",
	"contracts/notifications/customer":    "entity,event,record,enum",
	"contracts/orders/order":              "entity,event,snapshot,record,enum",
	"contracts/orders/pos":                "entity,record,snapshot,enum",
	"contracts/orders/shipping":           "record,value,enum",
	"contracts/payments/payment":          "entity,event,snapshot,record,value,enum",
	"contracts/payments/settlement":       "entity,event,snapshot,record,value,enum",
	"contracts/payments/terminal":         "entity,event,snapshot,record,value,enum",
	"contracts/pricing/benefit":           "value,enum",
	"contracts/pricing/membership":        "entity,record,value,enum",
	"contracts/pricing/promotion":         "entity,event,record,value,enum",
	"contracts/pricing/wallet":            "entity,record,snapshot,value,enum",
	"contracts/pubsub/envelop":            "record",
	"contracts/pubsub/event":              "event,record,enum",
	"contracts/supply/category":           "value,enum",
	"contracts/supply/favourite":          "entity,record,value,enum",
	"contracts/supply/importcompliance":   "entity,record,snapshot,value,enum",
	"contracts/supply/product":            "entity,event,snapshot,value,enum",
	"contracts/supply/purchase":           "entity,record,enum",
	"contracts/supply/review":             "entity,record,value,enum",
	"contracts/supply/warehouse":          "entity,event,record,snapshot,value,enum",
	"contracts/supply/wish":               "entity,record,value,enum",
	"versioning":                          "module-metadata",
}

// Reviewed for the v23.0.0 JSON-only surface after the domain-layout
// hard-cutover. The digest captures the complete exported model manifest.
// identity and composition, depot-qualified inventory, consolidated group
// fulfilment, geographic scope/context, revisioned inventory events, and the
// public storefront commercial/expiry projections.
const v23ExportedTypeManifestDigest = "bd4b50d6b2a7b44904a8a6e87283c9fcdcd7be5430ef3299051392a990ae70c8"

func TestV23ExportedTypesMatchModelManifest(t *testing.T) {
	seenPackages := make(map[string]bool)
	var entries []string
	pkgRoot := sharedContractPkgRoot(t)
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		packagePath := filepath.ToSlash(filepath.Dir(filepath.FromSlash(relativePkgPath(t, pkgRoot, path))))
		if packagePath == "." {
			return nil
		}
		class, classified := v23ModelPackageManifest[packagePath]
		if !classified {
			t.Errorf("%s is not classified in the v23 model manifest", packagePath)
			return nil
		}
		seenPackages[packagePath] = true

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		for _, decl := range file.Decls {
			general, ok := decl.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, spec := range general.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if ok && typeSpec.Name.IsExported() {
					entries = append(entries, packagePath+"|"+class+"|"+typeSpec.Name.Name)
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("read v23 model manifest: %v", err)
	}
	for packagePath := range v23ModelPackageManifest {
		if !seenPackages[packagePath] {
			t.Errorf("manifest package %s has no production source", packagePath)
		}
	}
	sort.Strings(entries)
	sum := sha256.Sum256([]byte(strings.Join(entries, "\n")))
	got := hex.EncodeToString(sum[:])
	if got != v23ExportedTypeManifestDigest {
		t.Fatalf("exported model manifest changed: got %s; classify the change and update the reviewed digest", got)
	}
}
