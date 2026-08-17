package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/pubsub/event/event_enums"
)

// v27EventSchemaVersion2Payloads is the reviewed set of routed Pub/Sub payloads
// that publish event_version "2". It was derived by walking the transitive JSON
// key set of every payload type in contracts/pubsub/event against the v26 tree:
// a payload is listed when a nested type reached through any depth of
// embedding, slice, or pointer changed a JSON key, which is why
// OrderPackingProjection is present even though its own fields are unchanged.
//
// RELEASE_NOTES.md declares this table the source of truth for consumers, and
// consumers reject a wrong event_version rather than decoding it leniently. A
// payload missing here would be silently decoded with an empty sku_id or
// dead-lettered in production, so the document and this slice are pinned to
// each other.
//
// v28.0.0 does not move the schema version. Its payload additions — the
// optional market_id, country_code, and depot_code geography — are all
// omitempty fields that a version-2 consumer decodes unchanged, so this table
// and the v27.0.0 section that publishes it stay authoritative. Publishers
// must not bump event_version for them; a bump would break every consumer
// that rejects a version it does not recognize, for no decode benefit.
var v27EventSchemaVersion2Payloads = []string{
	"CatalogBaseCostChangedEvent",
	"CatalogListingChangedEvent",
	"InventoryDateMarkThresholdEvent",
	"InventoryLotReceivedEvent",
	"InventoryPackageConvertedEvent",
	"InventoryQualityAssessedEvent",
	"InventoryReservationChangedEvent",
	"InventorySaleCommittedEvent",
	"InventoryStockBucketChangedEvent",
	"OrderFact",
	"OrderPackingProjection",
	"OrderPaidEvent",
	"ProductSalesRollup",
	"RefundCompletedEvent",
	"RefundFact",
	"StockLocationAvailabilityChangedEvent",
	"StockStagingChangedEvent",
}

var releaseNotesEventRow = regexp.MustCompile(
	"(?m)^\\| *([0-9]+) *\\| *`([^`]+)` *\\| *`([^`]+)` *\\| *`([^`]+)` *\\|",
)

// TestV27ReleaseNotesEventTableMatchesTheReviewedPayloadSet keeps the published
// event schema version 2 table and the reviewed payload set from drifting
// apart. Editing either side alone fails here.
func TestV27ReleaseNotesEventTableMatchesTheReviewedPayloadSet(t *testing.T) {
	section := v27ReleaseNotesSection(t)
	matches := releaseNotesEventRow.FindAllStringSubmatch(section, -1)
	if len(matches) == 0 {
		t.Fatal("RELEASE_NOTES.md v27.0.0 section contains no event schema version 2 rows")
	}

	var payloads []string
	for index, match := range matches {
		number, err := strconv.Atoi(match[1])
		if err != nil {
			t.Fatalf("row %d has a non-numeric index %q", index+1, match[1])
		}
		if number != index+1 {
			t.Errorf("event table row %d is numbered %d; numbering must stay contiguous from 1", index+1, number)
		}
		payloads = append(payloads, match[2])

		eventType := event_enums.EventType(match[3])
		if !eventType.IsValid() {
			t.Errorf("row %d names event type %q, which is not a valid contract EventType", number, match[3])
		}
		topic := event_enums.EventTopic(match[4])
		if !topic.IsValid() {
			t.Errorf("row %d names topic %q, which is not a valid contract EventTopic", number, match[4])
		}
	}

	got := append([]string(nil), payloads...)
	sort.Strings(got)
	want := append([]string(nil), v27EventSchemaVersion2Payloads...)
	sort.Strings(want)
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("RELEASE_NOTES.md event table payloads =\n  %v\nreviewed set =\n  %v", got, want)
	}
}

// TestV27EventSchemaVersion2PayloadsAreDeclaredRoutedTypes proves every listed
// payload is a real exported type in the routed event package, so a typo or a
// renamed payload cannot survive in the published table.
func TestV27EventSchemaVersion2PayloadsAreDeclaredRoutedTypes(t *testing.T) {
	declared := v27RoutedEventPayloadTypes(t)
	for _, payload := range v27EventSchemaVersion2Payloads {
		if _, ok := declared[payload]; !ok {
			t.Errorf("event schema version 2 payload %s is not an exported type in contracts/pubsub/event", payload)
		}
	}

	// The retired v26 payload must not reappear in the routed package or in
	// the published table.
	if _, ok := declared["PackagePricingAvailabilityChangedEvent"]; ok {
		t.Error("retired PackagePricingAvailabilityChangedEvent must not be declared in contracts/pubsub/event")
	}
	for _, payload := range v27EventSchemaVersion2Payloads {
		if payload == "PackagePricingAvailabilityChangedEvent" {
			t.Error("retired PackagePricingAvailabilityChangedEvent must not appear in the event schema version 2 table")
		}
	}
}

func v27ReleaseNotesSection(t *testing.T) string {
	t.Helper()
	path := filepath.Join(filepath.Dir(sharedContractPkgRoot(t)), "RELEASE_NOTES.md")
	contents, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read RELEASE_NOTES.md: %v", err)
	}
	text := string(contents)
	start := strings.Index(text, "## v27.0.0 ")
	if start < 0 {
		t.Fatal("RELEASE_NOTES.md has no v27.0.0 release section")
	}
	rest := text[start+len("## v27.0.0 "):]
	if end := strings.Index(rest, "\n## v"); end >= 0 {
		return rest[:end]
	}
	return rest
}

func v27RoutedEventPayloadTypes(t *testing.T) map[string]struct{} {
	t.Helper()
	root := filepath.Join(sharedContractPkgRoot(t), "contracts", "pubsub", "event")
	declared := make(map[string]struct{})
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		if filepath.Dir(path) != root {
			return nil
		}
		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpec, ok := specification.(*ast.TypeSpec)
				if ok && typeSpec.Name.IsExported() {
					declared[typeSpec.Name.Name] = struct{}{}
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan routed event payload types: %v", err)
	}
	return declared
}
