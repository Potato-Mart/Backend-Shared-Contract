package versioning

import (
	"os"
	"strings"
	"testing"
)

func TestV19ModuleMetadata(t *testing.T) {
	if ModuleName != "Backend-Shared-Contract" || ModuleVersion != "v19.0.0" || MajorVersion != "v19" {
		t.Fatalf("unexpected module metadata: %q %q %q", ModuleName, ModuleVersion, MajorVersion)
	}
}

func TestV19ModulePath(t *testing.T) {
	contents, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	moduleLine := "module github.com/Potato-Mart/Backend-Shared-Contract/v19"
	if !strings.Contains(string(contents), moduleLine) || strings.Contains(string(contents), "Backend-Shared-Contract/v18") {
		t.Fatalf("go.mod must use the v19 hard-cut module path: %s", contents)
	}
}
