package versioning

import "testing"

func TestV16ModuleMetadata(t *testing.T) {
	if ModuleName != "Backend-Shared-Contract" || ModuleVersion != "v16.0.0" || MajorVersion != "v16" {
		t.Fatalf("unexpected module metadata: %q %q %q", ModuleName, ModuleVersion, MajorVersion)
	}
}
