package versioning

import (
	"os"
	"strings"
	"testing"
)

func TestV17ModuleMetadata(t *testing.T) {
	if ModuleName != "Backend-Shared-Contract" || ModuleVersion != "v17.4.0" || MajorVersion != "v17" {
		t.Fatalf("unexpected module metadata: %q %q %q", ModuleName, ModuleVersion, MajorVersion)
	}
}

func TestV17ModulePath(t *testing.T) {
	contents, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	moduleLine := "module github.com/Potato-Mart/Backend-Shared-Contract/v17"
	if !strings.Contains(string(contents), moduleLine) || strings.Contains(string(contents), "Backend-Shared-Contract/v16") {
		t.Fatalf("go.mod must use the v17 hard-cut module path: %s", contents)
	}
}
