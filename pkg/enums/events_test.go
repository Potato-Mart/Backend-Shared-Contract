package enums_test

import (
	"testing"

	eventsenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/events"
)

func TestVoucherClaimIssuedEventType(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name:    "events.EventTypeVoucherClaimIssued",
			valid:   []stringEnum{eventsenum.EventTypeVoucherClaimIssued},
			invalid: eventsenum.EventType("__invalid__"),
		},
	})
	if got := eventsenum.EventTypeVoucherClaimIssued.String(); got != "voucher.claim_issued" {
		t.Fatalf("voucher claim event type = %q", got)
	}
}
