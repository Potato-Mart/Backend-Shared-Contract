package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

var forbiddenTransportImports = map[string]struct{}{
	"net/http":                 {},
	"github.com/gin-gonic/gin": {},
	"github.com/go-chi/chi":    {},
	"github.com/labstack/echo": {},
	"github.com/gofiber/fiber": {},
}

func TestSharedContractHasNoTransportFrameworkImports(t *testing.T) {
	walkGoFiles(t, func(path string, fset *token.FileSet, file *ast.File) {
		for _, spec := range file.Imports {
			importPath, err := strconv.Unquote(spec.Path.Value)
			if err != nil {
				t.Fatalf("%s: invalid import: %v", fset.Position(spec.Pos()), err)
			}
			if _, forbidden := forbiddenTransportImports[importPath]; forbidden {
				t.Errorf("%s imports transport framework %q", fset.Position(spec.Pos()), importPath)
			}
		}
	})
}

func TestSharedContractHasNoTransportTags(t *testing.T) {
	walkGoFiles(t, func(path string, fset *token.FileSet, file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			field, ok := node.(*ast.Field)
			if !ok || field.Tag == nil {
				return true
			}
			tag := strings.ToLower(field.Tag.Value)
			for _, term := range []string{"binding:", "form:", "query:", "header:", "uri:"} {
				if strings.Contains(tag, term) {
					t.Errorf("%s uses transport-specific struct tag %q", fset.Position(field.Tag.Pos()), term)
				}
			}
			return true
		})
	})
}

func walkGoFiles(t *testing.T, inspect func(string, *token.FileSet, *ast.File)) {
	t.Helper()
	pkgRoot := sharedContractPkgRoot(t)
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return err
		}
		inspect(path, fset, file)
		return nil
	})
	if err != nil {
		t.Fatalf("scan shared contract: %v", err)
	}
}
