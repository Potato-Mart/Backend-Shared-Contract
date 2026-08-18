package event

import "testing"

func TestV29SKUCodeEventsUseSchemaVersion3(t *testing.T) {
	if SKUCodeEventVersion != "3" {
		t.Fatalf("SKU-code event version = %q, want 3", SKUCodeEventVersion)
	}
}
