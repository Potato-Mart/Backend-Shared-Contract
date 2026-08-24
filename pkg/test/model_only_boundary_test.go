package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
)

var modelBoundaryForbiddenTypeSuffixes = []string{
	"Request", "Response", "Input", "Payload", "DTO", "Dto",
	"Command", "Query", "Acknowledgement", "Acknowledgment",
	"Pagination", "Page", "Result",
}

var modelBoundaryForbiddenPackagePrefixes = []string{
	"enums", "logic", "serviceauth",
}

var modelBoundaryApprovedMethods = map[string]struct{}{
	"String": {}, "IsValid": {},
}

var modelBoundaryJSONTag = regexp.MustCompile(`^json:"[^"]*"$`)

func TestV30ContractIsJSONModelOnly(t *testing.T) {
	var violations []string
	pkgRoot := sharedContractPkgRoot(t)
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		normalized := relativePkgPath(t, pkgRoot, path)
		for _, prefix := range modelBoundaryForbiddenPackagePrefixes {
			if normalized == prefix || strings.HasPrefix(normalized, prefix+"/") {
				violations = append(violations, normalized+": forbidden non-model package")
			}
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if parseErr != nil {
			return parseErr
		}
		for _, imported := range file.Imports {
			importPath, unquoteErr := strconv.Unquote(imported.Path.Value)
			if unquoteErr != nil {
				return unquoteErr
			}
			for _, forbidden := range []string{"database/sql", "go.mongodb.org/", "gorm.io/", "github.com/jmoiron/sqlx", "github.com/jackc/pgx", "cloud.google.com/go/firestore", "cloud.google.com/go/datastore"} {
				if importPath == forbidden || strings.HasPrefix(importPath, forbidden) {
					violations = append(violations, modelBoundaryPosition(fset, imported.Pos())+": non-JSON dependency "+importPath)
				}
			}
		}
		file, parseErr = parser.ParseFile(fset, path, nil, parser.ParseComments)
		if parseErr != nil {
			return parseErr
		}
		stringEnumTypes := modelBoundaryStringEnumTypes(file)

		for _, decl := range file.Decls {
			switch typed := decl.(type) {
			case *ast.FuncDecl:
				if typed.Recv == nil {
					violations = append(violations, modelBoundaryPosition(fset, typed.Pos())+": free function "+typed.Name.Name)
					continue
				}
				if _, approved := modelBoundaryApprovedMethods[typed.Name.Name]; !approved {
					violations = append(violations, modelBoundaryPosition(fset, typed.Pos())+": non-intrinsic method "+typed.Name.Name)
					continue
				}
				modelBoundaryValidateIntrinsicMethod(fset, typed, stringEnumTypes, &violations)
			case *ast.GenDecl:
				modelBoundaryInspectDeclaration(fset, typed, stringEnumTypes, &violations)
			}
		}

		for _, group := range file.Comments {
			if strings.Contains(group.Text(), "Deprecated:") || strings.Contains(group.Text(), "@deprecated") {
				violations = append(violations, normalized+": deprecated production declaration")
			}
		}

		ast.Inspect(file, func(node ast.Node) bool {
			field, ok := node.(*ast.Field)
			if !ok {
				return true
			}
			if field.Tag == nil {
				for _, name := range field.Names {
					if name.IsExported() {
						violations = append(violations, modelBoundaryPosition(fset, name.Pos())+": exported JSON field without json tag "+name.Name)
					}
				}
				return true
			}
			tag, unquoteErr := strconv.Unquote(field.Tag.Value)
			if unquoteErr != nil || !modelBoundaryJSONTag.MatchString(tag) {
				violations = append(violations, modelBoundaryPosition(fset, field.Tag.Pos())+": struct tags must contain exactly one json tag")
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("scan v30 contract: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("model-only contract boundary violations:\n%s", strings.Join(violations, "\n"))
	}
}

func modelBoundaryStringEnumTypes(file *ast.File) map[string]struct{} {
	stringTypes := make(map[string]struct{})
	for _, decl := range file.Decls {
		general, ok := decl.(*ast.GenDecl)
		if !ok || general.Tok != token.TYPE {
			continue
		}
		for _, spec := range general.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok || typeSpec.Assign.IsValid() {
				continue
			}
			underlying, ok := typeSpec.Type.(*ast.Ident)
			if ok && underlying.Name == "string" {
				stringTypes[typeSpec.Name.Name] = struct{}{}
			}
		}
	}
	enumTypes := make(map[string]struct{})
	for _, decl := range file.Decls {
		general, ok := decl.(*ast.GenDecl)
		if !ok || general.Tok != token.CONST {
			continue
		}
		for _, spec := range general.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || valueSpec.Type == nil {
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

func modelBoundaryInspectDeclaration(fset *token.FileSet, decl *ast.GenDecl, stringEnumTypes map[string]struct{}, violations *[]string) {
	for _, spec := range decl.Specs {
		switch typed := spec.(type) {
		case *ast.TypeSpec:
			if typed.Assign.IsValid() {
				*violations = append(*violations, modelBoundaryPosition(fset, typed.Pos())+": compatibility/type alias "+typed.Name.Name)
			}
			if !typed.Name.IsExported() {
				continue
			}
			for _, suffix := range modelBoundaryForbiddenTypeSuffixes {
				if strings.HasSuffix(typed.Name.Name, suffix) {
					*violations = append(*violations, modelBoundaryPosition(fset, typed.Pos())+": endpoint DTO type "+typed.Name.Name)
				}
			}
			if typed.Name.Name == "Scope" || strings.HasSuffix(typed.Name.Name, "Scopes") || strings.Contains(typed.Name.Name, "ServiceToken") {
				*violations = append(*violations, modelBoundaryPosition(fset, typed.Pos())+": service-auth catalogue type "+typed.Name.Name)
			}
		case *ast.ValueSpec:
			for index, name := range typed.Names {
				if !name.IsExported() {
					continue
				}
				if decl.Tok == token.VAR {
					*violations = append(*violations, modelBoundaryPosition(fset, name.Pos())+": exported mutable variable "+name.Name)
				}
				if strings.HasPrefix(name.Name, "Path") || strings.HasPrefix(name.Name, "Route") || strings.Contains(name.Name, "Endpoint") {
					*violations = append(*violations, modelBoundaryPosition(fset, name.Pos())+": exported path/route constant "+name.Name)
				}
				// A service-auth scope catalogue is forbidden. A member of a
				// string-backed enum declared in this file is a model value,
				// not a catalogue entry, so a name such as ScopeLevelDepot is
				// allowed — unless it carries the "domain:action" shape a
				// service scope uses, which is forbidden however it is typed.
				scopeNamed := strings.HasPrefix(name.Name, "Scope") || strings.Contains(name.Name, "ServiceScope")
				enumMember := modelBoundaryTypedEnumMember(typed, stringEnumTypes)
				value, isStringLiteral := modelBoundaryStringConstantValue(typed, index)
				if scopeNamed && (!enumMember || strings.Contains(value, ":")) {
					*violations = append(*violations, modelBoundaryPosition(fset, name.Pos())+": exported service-scope constant "+name.Name)
				}
				if isStringLiteral && strings.HasPrefix(value, "/") {
					*violations = append(*violations, modelBoundaryPosition(fset, name.Pos())+": exported route string "+name.Name)
				}
			}
		}
	}
}

// modelBoundaryTypedEnumMember reports whether spec declares constants of a
// string-backed enum type declared in the same file.
func modelBoundaryTypedEnumMember(spec *ast.ValueSpec, stringEnumTypes map[string]struct{}) bool {
	declaredType, ok := spec.Type.(*ast.Ident)
	if !ok {
		return false
	}
	_, enum := stringEnumTypes[declaredType.Name]
	return enum
}

// modelBoundaryStringConstantValue returns the unquoted string value assigned
// at index, and whether that value is a plain string literal.
func modelBoundaryStringConstantValue(spec *ast.ValueSpec, index int) (string, bool) {
	if index >= len(spec.Values) {
		return "", false
	}
	literal, ok := spec.Values[index].(*ast.BasicLit)
	if !ok || literal.Kind != token.STRING {
		return "", false
	}
	value, err := strconv.Unquote(literal.Value)
	if err != nil {
		return "", false
	}
	return value, true
}

func modelBoundaryValidateIntrinsicMethod(fset *token.FileSet, method *ast.FuncDecl, stringEnumTypes map[string]struct{}, violations *[]string) {
	if method.Name.Name != "String" && method.Name.Name != "IsValid" {
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
	if _, approved := stringEnumTypes[receiver.Name]; !approved {
		*violations = append(*violations, modelBoundaryPosition(fset, method.Pos())+": "+method.Name.Name+" receiver "+receiver.Name+" is not a string-backed enum")
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
