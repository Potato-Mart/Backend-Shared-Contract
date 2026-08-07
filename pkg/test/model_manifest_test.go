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

// v22ModelPackageManifest classifies every production package. Adding an
// exported type changes the digest below and requires an explicit manifest
// review instead of silently expanding the shared module.
var v22ModelPackageManifest = map[string]string{
	"common":                     "enum,value",
	"contracts/analytics":        "record",
	"contracts/benefit":          "value",
	"contracts/campaign":         "entity,event,snapshot,record",
	"contracts/category":         "value",
	"contracts/customers":        "entity,event,snapshot,record",
	"contracts/events":           "event",
	"contracts/favourite":        "entity,record,value",
	"contracts/identity":         "entity,claims,event,session,record",
	"contracts/importcompliance": "entity,record,snapshot,value",
	"contracts/marketing":        "entity,record",
	"contracts/membership":       "entity,record,value",
	"contracts/notification":     "entity,event,record",
	"contracts/payments":         "entity,event,snapshot,record,value",
	"contracts/pos":              "entity,record,snapshot",
	"contracts/product":          "entity,event,snapshot,value",
	"contracts/promotion":        "entity,event,record,value",
	"contracts/purchase":         "entity,record",
	"contracts/review":           "entity,record,value",
	"contracts/sales":            "entity,event,snapshot,record",
	"contracts/shared":           "event,record,value",
	"contracts/shipping":         "record,value",
	"contracts/wallet":           "entity,record,snapshot,value",
	"contracts/warehouse":        "entity,event,record,snapshot",
	"contracts/wholesale":        "entity,record,snapshot",
	"contracts/wish":             "entity,record,value",
	"enums/account":              "enum",
	"enums/apiresponse":          "enum",
	"enums/benefit":              "enum",
	"enums/campaign":             "enum",
	"enums/customer":             "enum",
	"enums/events":               "enum",
	"enums/favourite":            "enum",
	"enums/geography":            "enum",
	"enums/identity":             "enum",
	"enums/importcompliance":     "enum",
	"enums/marketing":            "enum",
	"enums/membership":           "enum",
	"enums/notification":         "enum",
	"enums/payment":              "enum",
	"enums/pos":                  "enum",
	"enums/product":              "enum",
	"enums/promotion":            "enum",
	"enums/purchase":             "enum",
	"enums/review":               "enum",
	"enums/sales":                "enum",
	"enums/security":             "enum",
	"enums/shipping":             "enum",
	"enums/wallet":               "enum",
	"enums/warehouse":            "enum",
	"enums/wholesale":            "enum",
	"enums/wish":                 "enum",
	"versioning":                 "module-metadata",
}

// Reviewed for the v22.2.0 JSON-only surface: typed geography, package
// identity and composition, depot-qualified inventory, consolidated group
// fulfilment, geographic scope/context, revisioned inventory events, and the
// public storefront commercial/expiry projections.
const v22ExportedTypeManifestDigest = "a99fcc9ab563e1f5893070af2d102898e6ee828ad7cf84741fb270e6c243e2e8"

func TestV22ExportedTypesMatchModelManifest(t *testing.T) {
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
		class, classified := v22ModelPackageManifest[packagePath]
		if !classified {
			t.Errorf("%s is not classified in the v22 model manifest", packagePath)
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
		t.Fatalf("read v22 model manifest: %v", err)
	}
	for packagePath := range v22ModelPackageManifest {
		if !seenPackages[packagePath] {
			t.Errorf("manifest package %s has no production source", packagePath)
		}
	}
	sort.Strings(entries)
	sum := sha256.Sum256([]byte(strings.Join(entries, "\n")))
	got := hex.EncodeToString(sum[:])
	if got != v22ExportedTypeManifestDigest {
		t.Fatalf("exported model manifest changed: got %s; classify the change and update the reviewed digest", got)
	}
}
