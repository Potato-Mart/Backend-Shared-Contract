package role_enums

import "testing"

func TestPermissionClassificationWireValuesAreLocked(t *testing.T) {
	classifications := map[PermissionClassification]string{
		PermissionClassificationUI:       "ui",
		PermissionClassificationField:    "field-level",
		PermissionClassificationService:  "service-only",
		PermissionClassificationReserved: "intentionally-reserved",
	}
	if len(classifications) != 4 {
		t.Fatalf("permission classification set has %d values; want 4", len(classifications))
	}
	for classification, wire := range classifications {
		if got := classification.String(); got != wire {
			t.Errorf("PermissionClassification.String() = %q, want %q", got, wire)
		}
		if !classification.IsValid() {
			t.Errorf("PermissionClassification(%q) must validate", wire)
		}
	}
	if PermissionClassification("__invalid__").IsValid() {
		t.Error("unknown permission classification validates")
	}
}
