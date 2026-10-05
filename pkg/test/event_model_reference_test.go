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

	event_enums "github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pubsub/routing"
)

// documentedEventPayloads is the reviewed payload set in the current event reference.
var documentedEventPayloads = []string{
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

var eventModelRow = regexp.MustCompile(
	"(?m)^\\| *([0-9]+) *\\| *`([^`]+)` *\\| *`([^`]+)` *\\| *`([^`]+)` *\\|",
)

// TestEventModelReferenceMatchesReviewedPayloadSet keeps model names and routing enums aligned.
func TestEventModelReferenceMatchesReviewedPayloadSet(t *testing.T) {
	section := eventModelReferenceSection(t)
	matches := eventModelRow.FindAllStringSubmatch(section, -1)
	if len(matches) == 0 {
		t.Fatal("docs/development/README.md event payload reference section contains no rows")
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
	want := append([]string(nil), documentedEventPayloads...)
	sort.Strings(want)
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("docs/development/README.md event table payloads =\n  %v\nreviewed set =\n  %v", got, want)
	}
}

// TestDocumentedEventPayloadsAreDeclaredRoutedTypes proves every listed
// payload is a real exported type in the routed event package, so a typo or a
// renamed payload cannot survive in the published table.
func TestDocumentedEventPayloadsAreDeclaredRoutedTypes(t *testing.T) {
	declared := routedEventPayloadTypes(t)
	for _, payload := range documentedEventPayloads {
		if _, ok := declared[payload]; !ok {
			t.Errorf("documented payload %s is not an exported type in contracts/pubsub", payload)
		}
	}

	// The retired payload must not reappear in the routed package or in the
	// published table.
	if _, ok := declared["PackagePricingAvailabilityChangedEvent"]; ok {
		t.Error("retired PackagePricingAvailabilityChangedEvent must not be declared in contracts/pubsub")
	}
	for _, payload := range documentedEventPayloads {
		if payload == "PackagePricingAvailabilityChangedEvent" {
			t.Error("retired PackagePricingAvailabilityChangedEvent must not appear in the event payload reference table")
		}
	}
}

func eventModelReferenceSection(t *testing.T) string {
	t.Helper()
	path := filepath.Join(filepath.Dir(sharedContractPkgRoot(t)), "docs", "development", "README.md")
	contents, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read docs/development/README.md: %v", err)
	}
	text := string(contents)
	const heading = "## Event payload reference"
	if count := strings.Count(text, heading); count != 1 {
		t.Fatalf("docs/development/README.md has %d %q headings, want exactly one", count, heading)
	}
	start := strings.Index(text, heading)
	rest := text[start+len(heading):]
	if end := strings.Index(rest, "\n## "); end >= 0 {
		return rest[:end]
	}
	return rest
}

func routedEventPayloadTypes(t *testing.T) map[string]struct{} {
	t.Helper()
	root := filepath.Join(sharedContractPkgRoot(t), "contracts", "pubsub")
	declared := make(map[string]struct{})
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		relative, relErr := filepath.Rel(root, filepath.Dir(path))
		if relErr != nil {
			return relErr
		}
		if relative == "envelope" || relative == "routing" || strings.HasPrefix(relative, "envelope"+string(filepath.Separator)) || strings.HasPrefix(relative, "routing"+string(filepath.Separator)) {
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
