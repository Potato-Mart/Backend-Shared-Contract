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

// v17ModelPackageManifest classifies every production package. Adding an
// exported type changes the digest below and requires an explicit manifest
// review instead of silently expanding the shared module.
var v17ModelPackageManifest = map[string]string{
	"common":                     "value",
	"contracts/analytics":        "record",
	"contracts/campaign":         "entity,event,snapshot,record",
	"contracts/category":         "value",
	"contracts/customers":        "entity,snapshot,record",
	"contracts/identity":         "entity,claims,event,session,record",
	"contracts/importcompliance": "entity,record,snapshot,value",
	"contracts/marketing":        "entity,record",
	"contracts/membership":       "entity,record",
	"contracts/notification":     "entity,event,record",
	"contracts/payments":         "entity,snapshot,record,value",
	"contracts/product":          "entity,snapshot,value",
	"contracts/promotion":        "entity,record,value",
	"contracts/purchase":         "entity,record",
	"contracts/sales":            "entity,snapshot,record",
	"contracts/shared":           "event,record,value",
	"contracts/shipping":         "record,value",
	"contracts/wallet":           "entity,record,snapshot",
	"contracts/warehouse":        "entity,record,snapshot",
	"contracts/wholesale":        "entity,record,snapshot",
	"enums/account":              "enum",
	"enums/apiresponse":          "enum",
	"enums/campaign":             "enum",
	"enums/customer":             "enum",
	"enums/identity":             "enum",
	"enums/importcompliance":     "enum",
	"enums/marketing":            "enum",
	"enums/membership":           "enum",
	"enums/notification":         "enum",
	"enums/payment":              "enum",
	"enums/product":              "enum",
	"enums/promotion":            "enum",
	"enums/purchase":             "enum",
	"enums/sales":                "enum",
	"enums/security":             "enum",
	"enums/shipping":             "enum",
	"enums/wallet":               "enum",
	"enums/warehouse":            "enum",
	"enums/wholesale":            "enum",
	"versioning":                 "module-metadata",
}

// Reviewed addition: import-compliance records and their stable value enums.
const v17ExportedTypeManifestDigest = "2a6fae21d5b4909b8d1e979e80d3a49405e68622602b0f0422579daa4ca91d1b"

func TestV17ExportedTypesMatchModelManifest(t *testing.T) {
	seenPackages := make(map[string]bool)
	var entries []string
	err := filepath.WalkDir(".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		packagePath := filepath.ToSlash(filepath.Dir(path))
		if packagePath == "." {
			return nil
		}
		class, classified := v17ModelPackageManifest[packagePath]
		if !classified {
			t.Errorf("%s is not classified in the v17 model manifest", packagePath)
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
		t.Fatalf("read v17 model manifest: %v", err)
	}
	for packagePath := range v17ModelPackageManifest {
		if !seenPackages[packagePath] {
			t.Errorf("manifest package %s has no production source", packagePath)
		}
	}
	sort.Strings(entries)
	sum := sha256.Sum256([]byte(strings.Join(entries, "\n")))
	got := hex.EncodeToString(sum[:])
	if got != v17ExportedTypeManifestDigest {
		t.Fatalf("exported model manifest changed: got %s; classify the change and update the reviewed digest", got)
	}
}
