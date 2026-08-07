package pkg_test

import (
	"path/filepath"
	"testing"
)

func sharedContractPkgRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join(".."))
	if err != nil {
		t.Fatalf("resolve shared contract pkg root: %v", err)
	}
	return filepath.Clean(root)
}

func relativePkgPath(t *testing.T, pkgRoot string, path string) string {
	t.Helper()
	relative, err := filepath.Rel(pkgRoot, path)
	if err != nil {
		t.Fatalf("resolve path %q relative to %q: %v", path, pkgRoot, err)
	}
	return filepath.ToSlash(relative)
}
