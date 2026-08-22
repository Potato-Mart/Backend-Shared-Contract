package versioning

import (
	"os"
	"strings"
	"testing"
)

func TestV29ModuleMetadata(t *testing.T) {
	if ModuleName != "Backend-Shared-Contract" || ModuleVersion != "v29.0.2" || MajorVersion != "v29" {
		t.Fatalf("unexpected module metadata: %q %q %q", ModuleName, ModuleVersion, MajorVersion)
	}
}

func TestV29ModulePath(t *testing.T) {
	contents, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	moduleLine := "module github.com/Potato-Mart/Backend-Shared-Contract/v29"
	if !strings.Contains(string(contents), moduleLine) ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v19") ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v20") ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v21") ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v22") ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v23") ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v24") ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v25") ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v26") ||
		strings.Contains(string(contents), "Backend-Shared-Contract/v27") {
		t.Fatalf("go.mod must use the v29 hard-cut module path: %s", contents)
	}
}
