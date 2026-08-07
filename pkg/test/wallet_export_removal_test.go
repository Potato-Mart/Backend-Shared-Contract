package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestV23RemovesWalletExportProductionSymbols(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	for _, removedPath := range []string{
		filepath.Join("contracts", "pricing", "wallet", "export.go"),
		filepath.Join("contracts", "pricing", "wallet", "export_status.go"),
	} {
		absolutePath := filepath.Join(pkgRoot, removedPath)
		if _, err := os.Stat(absolutePath); err == nil {
			t.Errorf("removed wallet-export source still exists: %s", absolutePath)
		}
	}

	fset := token.NewFileSet()
	for _, root := range []string{filepath.Join("contracts", "pricing", "wallet")} {
		absoluteRoot := filepath.Join(pkgRoot, root)
		err := filepath.WalkDir(absoluteRoot, func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
				return nil
			}
			file, err := parser.ParseFile(fset, path, nil, 0)
			if err != nil {
				return err
			}
			ast.Inspect(file, func(node ast.Node) bool {
				switch typed := node.(type) {
				case *ast.TypeSpec:
					if strings.HasPrefix(typed.Name.Name, "WalletExport") {
						t.Errorf("removed wallet-export type remains in %s: %s", path, typed.Name.Name)
					}
				case *ast.ValueSpec:
					for _, name := range typed.Names {
						if strings.HasPrefix(name.Name, "WalletExport") {
							t.Errorf("removed wallet-export value remains in %s: %s", path, name.Name)
						}
					}
				}
				return true
			})
			return nil
		})
		if err != nil {
			t.Fatalf("scan %s for removed wallet-export symbols: %v", root, err)
		}
	}
}
