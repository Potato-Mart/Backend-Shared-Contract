package versioning

import (
	"os"
	"strings"
	"testing"
)

func TestV22ModuleMetadata(t *testing.T) {
	if ModuleName != "Backend-Shared-Contract" || ModuleVersion != "v22.1.0" || MajorVersion != "v22" {
		t.Fatalf("unexpected module metadata: %q %q %q", ModuleName, ModuleVersion, MajorVersion)
	}
}

func TestV22ModulePath(t *testing.T) {
	contents, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	moduleLine := "module github.com/Potato-Mart/Backend-Shared-Contract/v22"
	if !strings.Contains(string(contents), moduleLine) ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v19") ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v20") ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v21") {
		t.Fatalf("go.mod must use the v22 hard-cut module path: %s", contents)
	}
}
