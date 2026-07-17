package versioning

import (
	"os"
	"strings"
	"testing"
)

func TestV18ModuleMetadata(t *testing.T) {
	if ModuleName != "Backend-Shared-Contract" || ModuleVersion != "v18.0.0" || MajorVersion != "v18" {
		t.Fatalf("unexpected module metadata: %q %q %q", ModuleName, ModuleVersion, MajorVersion)
	}
}

func TestV18ModulePath(t *testing.T) {
	contents, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	moduleLine := "module github.com/Potato-Mart/Backend-Shared-Contract/v18"
	if !strings.Contains(string(contents), moduleLine) || strings.Contains(string(contents), "Backend-Shared-Contract/v17") {
		t.Fatalf("go.mod must use the v18 hard-cut module path: %s", contents)
	}
}
