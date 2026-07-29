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

// v21ModelPackageManifest classifies every production package. Adding an
// exported type changes the digest below and requires an explicit manifest
// review instead of silently expanding the shared module.
var v21ModelPackageManifest = map[string]string{
	"common":                     "value",
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

// Reviewed for the v21 delivery, campaign, notification, and wallet cutover:
// v20 BrandID/BrandRef identity remains canonical; delivery schedules and
// preferences are modelled; campaigns link promotions and expose media, typed
// CTA, content/activation revisions, and safe storefront events; notification
// campaign references and push wire values are typed; points policy metadata
// is reusable; quiet-hours and service-local FCM/coupon-preview shapes are absent.
const v21ExportedTypeManifestDigest = "cc3a8493be29692b30813ac17645f3ad5d3469b5fa48351d736477ebe891e15c"

func TestV21ExportedTypesMatchModelManifest(t *testing.T) {
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
		class, classified := v21ModelPackageManifest[packagePath]
		if !classified {
			t.Errorf("%s is not classified in the v21 model manifest", packagePath)
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
		t.Fatalf("read v21 model manifest: %v", err)
	}
	for packagePath := range v21ModelPackageManifest {
		if !seenPackages[packagePath] {
			t.Errorf("manifest package %s has no production source", packagePath)
		}
	}
	sort.Strings(entries)
	sum := sha256.Sum256([]byte(strings.Join(entries, "\n")))
	got := hex.EncodeToString(sum[:])
	if got != v21ExportedTypeManifestDigest {
		t.Fatalf("exported model manifest changed: got %s; classify the change and update the reviewed digest", got)
	}
}
