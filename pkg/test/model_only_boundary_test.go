package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// modelBoundaryForbiddenTypeSuffixes identify endpoint and orchestration
// shapes. Domain records may describe a request in their fields, but this
// module must not publish transport DTOs or service workflows as types.
var modelBoundaryForbiddenTypeSuffixes = []string{
	"Request", "Response", "Input", "Output", "Command", "Query",
	"Acknowledgement", "Acknowledgment", "Pagination", "Page", "Result",
	"Migration",
}

// modelBoundaryForbiddenTypeNames are deliberately narrow so domain value
// records such as GeographicContext and PricingContext remain valid while the
// known runtime/transport shapes cannot return in a different package.
var modelBoundaryForbiddenTypeNames = map[string]string{
	"AccessTokenClaims":                  "token claims are owned by Identity runtime",
	"CheckoutCompensationRequestedEvent": "checkout compensation is a backend workflow",
	"CommandEnvelope":                    "command envelopes are service workflow transport",
	"CommandReceipt":                     "command receipts are service workflow transport",
	"InvoiceDeliveryRequestedEvent":      "invoice redelivery is a backend workflow",
	"ProviderOperationContext":           "provider operation state is backend-local",
	"ProviderPayloads":                   "raw provider payloads are backend-local",
	"RequestContext":                     "request context is transport runtime state",
	"ServiceOperation":                   "service operations are workflow transport",
	"VoucherClaimIssuedEvent":            "voucher claim delivery is a backend workflow",
	"WorkflowStatus":                     "workflow state is backend-local",
}

var modelBoundaryForbiddenPackageSegments = map[string]struct{}{
	"adapter": {}, "adapters": {}, "command": {}, "commands": {},
	"controller": {}, "controllers": {}, "dto": {}, "handler": {},
	"handlers": {}, "http": {}, "migration": {}, "migrations": {},
	"persistence": {}, "query": {}, "queries": {}, "repository": {},
	"repositories": {}, "storage": {}, "transport": {}, "workflow": {},
}

var modelBoundaryForbiddenImportPrefixes = []string{
	"database/sql",
	"go.mongodb.org/",
	"gorm.io/",
	"github.com/jackc/pgx",
	"github.com/jmoiron/sqlx",
	"cloud.google.com/go/datastore",
	"cloud.google.com/go/firestore",
	"cloud.google.com/go/pubsub",
	"google.golang.org/grpc",
	"net/http",
}

const eventEnvelopeSource = "contracts/pubsub/envelop/event_envelope.go"

var providerMetadataForbiddenSources = map[string]struct{}{
	"contracts/payments/settlement/settlement.go":         {},
	"contracts/payments/terminal/terminal.go":             {},
	"contracts/payments/terminal/terminal_transaction.go": {},
}

var modelBoundaryJSONTag = regexp.MustCompile(`^json:"[^"]*"$`)

// TestContractIsJSONModelOrEnumOnly keeps this repository limited to
// serializable domain records, open value types, and closed enums. Services
// own transport DTOs, command workflows, persistence records, provider
// diagnostics, runtime helpers, and migration code.
func TestContractIsJSONModelOrEnumOnly(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	contractsRoot := filepath.Join(pkgRoot, "contracts")
	var violations []string

	err := filepath.WalkDir(contractsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if parseErr != nil {
			return parseErr
		}
		normalized := relativePkgPath(t, pkgRoot, path)
		modelBoundaryInspectPackagePath(normalized, &violations)
		modelBoundaryInspectProductionFile(fset, file, normalized, &violations)
		return nil
	})
	if err != nil {
		t.Fatalf("scan contract boundary: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("model-or-enum contract boundary violations:\n%s", strings.Join(violations, "\n"))
	}
}

func modelBoundaryInspectPackagePath(normalized string, violations *[]string) {
	for _, segment := range strings.Split(filepath.ToSlash(filepath.Dir(normalized)), "/") {
		if _, forbidden := modelBoundaryForbiddenPackageSegments[segment]; forbidden {
			*violations = append(*violations, normalized+": forbidden non-model package segment "+segment)
		}
	}
}

func modelBoundaryInspectProductionFile(fset *token.FileSet, file *ast.File, normalized string, violations *[]string) {
	if filepath.Base(normalized) == "doc.go" {
		if len(file.Imports) != 0 || len(file.Decls) != 0 {
			*violations = append(*violations, normalized+": package documentation files may not contain imports or declarations")
		}
		return
	}

	isEnumFile := modelBoundaryIsLeafEnumFile(normalized)
	jsonAliases := modelBoundaryJSONImportAliases(fset, file, normalized, violations)
	stringTypes := modelBoundaryStringTypes(file)
	enumTypes := modelBoundaryClosedStringEnumTypes(file, stringTypes)
	allowedRawMessages := make(map[token.Pos]struct{})

	for _, declaration := range file.Decls {
		switch typed := declaration.(type) {
		case *ast.FuncDecl:
			if typed.Recv == nil {
				*violations = append(*violations, modelBoundaryPosition(fset, typed.Pos())+": free function "+typed.Name.Name)
				continue
			}
			modelBoundaryValidateIntrinsicMethod(fset, typed, enumTypes, violations)
		case *ast.GenDecl:
			switch typed.Tok {
			case token.TYPE:
				modelBoundaryInspectTypes(fset, typed, isEnumFile, enumTypes, violations)
				modelBoundaryInspectStructFields(fset, typed, normalized, jsonAliases, allowedRawMessages, violations)
			case token.CONST:
				modelBoundaryInspectConstants(fset, typed, isEnumFile, enumTypes, violations)
			case token.VAR:
				*violations = append(*violations, modelBoundaryPosition(fset, typed.Pos())+": variables are runtime state, not shared contracts")
			}
		}
	}

	modelBoundaryInspectRawMessages(fset, file, jsonAliases, allowedRawMessages, violations)
	for _, group := range file.Comments {
		if strings.Contains(group.Text(), "Deprecated:") || strings.Contains(group.Text(), "@deprecated") {
			*violations = append(*violations, normalized+": deprecated production declaration")
		}
	}
}

func modelBoundaryIsLeafEnumFile(normalized string) bool {
	return strings.HasSuffix(filepath.Base(filepath.Dir(normalized)), "_enums")
}

func modelBoundaryJSONImportAliases(fset *token.FileSet, file *ast.File, normalized string, violations *[]string) map[string]struct{} {
	aliases := make(map[string]struct{})
	for _, imported := range file.Imports {
		importPath, err := strconv.Unquote(imported.Path.Value)
		if err != nil {
			*violations = append(*violations, modelBoundaryPosition(fset, imported.Pos())+": invalid import path")
			continue
		}
		for _, forbidden := range modelBoundaryForbiddenImportPrefixes {
			if importPath == forbidden || strings.HasPrefix(importPath, forbidden+"/") {
				*violations = append(*violations, modelBoundaryPosition(fset, imported.Pos())+": non-model dependency "+importPath)
			}
		}
		if importPath != "encoding/json" {
			continue
		}
		if normalized != eventEnvelopeSource {
			*violations = append(*violations, modelBoundaryPosition(fset, imported.Pos())+": encoding/json is reserved for the Pub/Sub event envelope")
			continue
		}
		if imported.Name == nil {
			aliases["json"] = struct{}{}
			continue
		}
		if imported.Name.Name == "_" || imported.Name.Name == "." {
			*violations = append(*violations, modelBoundaryPosition(fset, imported.Name.Pos())+": event envelope must import encoding/json with a named package alias")
			continue
		}
		aliases[imported.Name.Name] = struct{}{}
	}
	return aliases
}

func modelBoundaryStringTypes(file *ast.File) map[string]struct{} {
	stringTypes := make(map[string]struct{})
	for _, declaration := range file.Decls {
		general, ok := declaration.(*ast.GenDecl)
		if !ok || general.Tok != token.TYPE {
			continue
		}
		for _, specification := range general.Specs {
			typeSpec, ok := specification.(*ast.TypeSpec)
			if !ok || typeSpec.Assign.IsValid() {
				continue
			}
			underlying, ok := typeSpec.Type.(*ast.Ident)
			if ok && underlying.Name == "string" {
				stringTypes[typeSpec.Name.Name] = struct{}{}
			}
		}
	}
	return stringTypes
}

// modelBoundaryClosedStringEnumTypes reports string types with typed constant
// members in the same file. Open code types (for example CountryCode) remain
// values, not enums, and are permitted outside an _enums leaf package.
func modelBoundaryClosedStringEnumTypes(file *ast.File, stringTypes map[string]struct{}) map[string]struct{} {
	enumTypes := make(map[string]struct{})
	for _, declaration := range file.Decls {
		general, ok := declaration.(*ast.GenDecl)
		if !ok || general.Tok != token.CONST {
			continue
		}
		for _, specification := range general.Specs {
			valueSpec, ok := specification.(*ast.ValueSpec)
			if !ok {
				continue
			}
			constantType, ok := valueSpec.Type.(*ast.Ident)
			if !ok {
				continue
			}
			if _, stringBacked := stringTypes[constantType.Name]; stringBacked {
				enumTypes[constantType.Name] = struct{}{}
			}
		}
	}
	return enumTypes
}

func modelBoundaryInspectTypes(fset *token.FileSet, declaration *ast.GenDecl, isEnumFile bool, enumTypes map[string]struct{}, violations *[]string) {
	for _, specification := range declaration.Specs {
		typeSpec, ok := specification.(*ast.TypeSpec)
		if !ok {
			continue
		}
		if typeSpec.Assign.IsValid() {
			*violations = append(*violations, modelBoundaryPosition(fset, typeSpec.Pos())+": type alias "+typeSpec.Name.Name)
			continue
		}
		if !typeSpec.Name.IsExported() {
			*violations = append(*violations, modelBoundaryPosition(fset, typeSpec.Pos())+": private production type "+typeSpec.Name.Name)
			continue
		}

		_, isEnum := enumTypes[typeSpec.Name.Name]
		if isEnumFile && !isEnum {
			*violations = append(*violations, modelBoundaryPosition(fset, typeSpec.Pos())+": enum package may only define one closed string enum")
		}
		if !isEnumFile && isEnum {
			*violations = append(*violations, modelBoundaryPosition(fset, typeSpec.Pos())+": closed enum "+typeSpec.Name.Name+" must live in a leaf _enums package")
		}
		if reason := modelBoundaryForbiddenTypeReason(typeSpec); reason != "" {
			*violations = append(*violations, modelBoundaryPosition(fset, typeSpec.Pos())+": "+reason+" "+typeSpec.Name.Name)
		}
	}
}

func modelBoundaryForbiddenTypeReason(typeSpec *ast.TypeSpec) string {
	name := typeSpec.Name.Name
	if reason, forbidden := modelBoundaryForbiddenTypeNames[name]; forbidden {
		return reason
	}
	if strings.Contains(name, "DTO") || strings.Contains(name, "Dto") {
		return "endpoint DTO type"
	}
	for _, suffix := range modelBoundaryForbiddenTypeSuffixes {
		if strings.HasSuffix(name, suffix) {
			return "endpoint DTO/workflow type"
		}
	}
	if _, isStruct := typeSpec.Type.(*ast.StructType); isStruct {
		if strings.Contains(name, "Command") || strings.Contains(name, "Workflow") {
			return "service workflow type"
		}
	}
	return ""
}

func modelBoundaryInspectConstants(fset *token.FileSet, declaration *ast.GenDecl, isEnumFile bool, enumTypes map[string]struct{}, violations *[]string) {
	for _, specification := range declaration.Specs {
		valueSpec, ok := specification.(*ast.ValueSpec)
		if !ok {
			*violations = append(*violations, modelBoundaryPosition(fset, specification.Pos())+": unsupported constant declaration")
			continue
		}
		enumType, typedEnum := valueSpec.Type.(*ast.Ident)
		if !isEnumFile || !typedEnum {
			*violations = append(*violations, modelBoundaryPosition(fset, valueSpec.Pos())+": non-enum constants are not shared contracts")
			continue
		}
		if _, known := enumTypes[enumType.Name]; !known {
			*violations = append(*violations, modelBoundaryPosition(fset, valueSpec.Pos())+": constants must be typed members of a closed string enum")
		}
	}
}

func modelBoundaryInspectStructFields(fset *token.FileSet, declaration *ast.GenDecl, normalized string, jsonAliases map[string]struct{}, allowedRawMessages map[token.Pos]struct{}, violations *[]string) {
	for _, specification := range declaration.Specs {
		typeSpec, ok := specification.(*ast.TypeSpec)
		if !ok {
			continue
		}
		structure, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}
		for _, field := range structure.Fields.List {
			modelBoundaryInspectStructFieldTag(fset, field, violations)
			modelBoundaryInspectProviderMetadataField(fset, field, normalized, violations)
			if selector, rawMessage := modelBoundaryDirectRawMessageSelector(field.Type, jsonAliases); rawMessage && normalized == eventEnvelopeSource && typeSpec.Name.Name == "EventEnvelope" && len(field.Names) == 1 && field.Names[0].Name == "Payload" {
				allowedRawMessages[selector.Pos()] = struct{}{}
			}
		}
	}
}

func modelBoundaryInspectProviderMetadataField(fset *token.FileSet, field *ast.Field, normalized string, violations *[]string) {
	if _, forbiddenSource := providerMetadataForbiddenSources[normalized]; !forbiddenSource {
		return
	}
	for _, name := range field.Names {
		if name.Name == "Metadata" {
			*violations = append(*violations, modelBoundaryPosition(fset, name.Pos())+": provider-facing contracts may not expose untyped metadata")
		}
	}
}

func modelBoundaryInspectStructFieldTag(fset *token.FileSet, field *ast.Field, violations *[]string) {
	if len(field.Names) == 0 {
		return // embedded shared value records retain their own JSON fields.
	}
	for _, name := range field.Names {
		if !name.IsExported() {
			*violations = append(*violations, modelBoundaryPosition(fset, name.Pos())+": private struct field "+name.Name)
			continue
		}
		if field.Tag == nil {
			*violations = append(*violations, modelBoundaryPosition(fset, name.Pos())+": exported JSON field without json tag "+name.Name)
		}
	}
	if field.Tag == nil {
		return
	}
	rawTag, err := strconv.Unquote(field.Tag.Value)
	if err != nil || !modelBoundaryJSONTag.MatchString(rawTag) {
		*violations = append(*violations, modelBoundaryPosition(fset, field.Tag.Pos())+": struct tags must contain exactly one json tag")
		return
	}
	if reflect.StructTag(rawTag).Get("json") == "-" {
		*violations = append(*violations, modelBoundaryPosition(fset, field.Tag.Pos())+": exported fields may not use json:\"-\"")
	}
}

func modelBoundaryDirectRawMessageSelector(expression ast.Expr, jsonAliases map[string]struct{}) (*ast.SelectorExpr, bool) {
	selector, ok := expression.(*ast.SelectorExpr)
	if !ok || selector.Sel.Name != "RawMessage" {
		return nil, false
	}
	identifier, ok := selector.X.(*ast.Ident)
	if !ok {
		return nil, false
	}
	_, imported := jsonAliases[identifier.Name]
	return selector, imported
}

func modelBoundaryInspectRawMessages(fset *token.FileSet, file *ast.File, jsonAliases map[string]struct{}, allowedRawMessages map[token.Pos]struct{}, violations *[]string) {
	ast.Inspect(file, func(node ast.Node) bool {
		selector, ok := node.(*ast.SelectorExpr)
		if !ok || selector.Sel.Name != "RawMessage" {
			return true
		}
		identifier, ok := selector.X.(*ast.Ident)
		if !ok {
			return true
		}
		if _, imported := jsonAliases[identifier.Name]; !imported {
			return true
		}
		if _, allowed := allowedRawMessages[selector.Pos()]; !allowed {
			*violations = append(*violations, modelBoundaryPosition(fset, selector.Pos())+": json.RawMessage is allowed only for EventEnvelope.Payload in "+eventEnvelopeSource)
		}
		return true
	})
}

func modelBoundaryValidateIntrinsicMethod(fset *token.FileSet, method *ast.FuncDecl, enumTypes map[string]struct{}, violations *[]string) {
	if method.Name.Name != "String" && method.Name.Name != "IsValid" {
		*violations = append(*violations, modelBoundaryPosition(fset, method.Pos())+": non-intrinsic method "+method.Name.Name)
		return
	}
	if method.Recv == nil || len(method.Recv.List) != 1 {
		*violations = append(*violations, modelBoundaryPosition(fset, method.Pos())+": "+method.Name.Name+" must have one enum receiver")
		return
	}
	receiver, ok := method.Recv.List[0].Type.(*ast.Ident)
	if !ok {
		*violations = append(*violations, modelBoundaryPosition(fset, method.Pos())+": "+method.Name.Name+" requires a value receiver on a string-backed enum")
		return
	}
	if _, approved := enumTypes[receiver.Name]; !approved {
		*violations = append(*violations, modelBoundaryPosition(fset, method.Pos())+": "+method.Name.Name+" receiver "+receiver.Name+" is not a closed string enum")
		return
	}
	if method.Type.Params != nil && len(method.Type.Params.List) != 0 {
		*violations = append(*violations, modelBoundaryPosition(fset, method.Pos())+": "+method.Name.Name+" must not accept arguments")
		return
	}
	if method.Type.Results == nil || len(method.Type.Results.List) != 1 || len(method.Type.Results.List[0].Names) > 1 {
		*violations = append(*violations, modelBoundaryPosition(fset, method.Pos())+": "+method.Name.Name+" must return one value")
		return
	}
	want := "string"
	if method.Name.Name == "IsValid" {
		want = "bool"
	}
	result, ok := method.Type.Results.List[0].Type.(*ast.Ident)
	if !ok || result.Name != want {
		*violations = append(*violations, modelBoundaryPosition(fset, method.Pos())+": "+method.Name.Name+" must return "+want)
	}
}

func modelBoundaryPosition(fset *token.FileSet, pos token.Pos) string {
	return fset.Position(pos).String()
}
