package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
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
	"apiresponse", "logic", "serviceauth", "enums/serviceauth",
	"contracts/pricing", "contracts/stockops",
}

var modelBoundaryApprovedMethods = map[string]struct{}{
	"String": {}, "IsValid": {},
	"MarshalJSON": {}, "UnmarshalJSON": {},
	"MarshalText": {}, "UnmarshalText": {},
	"MarshalBSON": {}, "UnmarshalBSON": {},
	"MarshalBSONValue": {}, "UnmarshalBSONValue": {},
}

func TestV17ContractIsModelOnly(t *testing.T) {
	var violations []string
	err := filepath.WalkDir(".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		normalized := filepath.ToSlash(path)
		for _, prefix := range modelBoundaryForbiddenPackagePrefixes {
			if normalized == prefix || strings.HasPrefix(normalized, prefix+"/") {
				violations = append(violations, normalized+": forbidden non-model package")
			}
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}

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
				modelBoundaryValidateIntrinsicMethod(fset, typed, &violations)
			case *ast.GenDecl:
				modelBoundaryInspectDeclaration(fset, typed, &violations)
			}
		}

		ast.Inspect(file, func(node ast.Node) bool {
			field, ok := node.(*ast.Field)
			if !ok || field.Tag == nil {
				return true
			}
			tag := strings.ToLower(field.Tag.Value)
			for _, transportTag := range []string{"binding:", "form:", "query:", "header:", "uri:"} {
				if strings.Contains(tag, transportTag) {
					violations = append(violations, modelBoundaryPosition(fset, field.Tag.Pos())+": transport struct tag "+transportTag)
				}
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("scan v17 contract: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("model-only contract boundary violations:\n%s", strings.Join(violations, "\n"))
	}
}

func modelBoundaryInspectDeclaration(fset *token.FileSet, decl *ast.GenDecl, violations *[]string) {
	for _, spec := range decl.Specs {
		switch typed := spec.(type) {
		case *ast.TypeSpec:
			if !typed.Name.IsExported() {
				continue
			}
			if typed.Assign.IsValid() {
				*violations = append(*violations, modelBoundaryPosition(fset, typed.Pos())+": exported compatibility/type alias "+typed.Name.Name)
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
				if strings.HasPrefix(name.Name, "Scope") || strings.Contains(name.Name, "ServiceScope") {
					*violations = append(*violations, modelBoundaryPosition(fset, name.Pos())+": exported service-scope constant "+name.Name)
				}
				if index >= len(typed.Values) {
					continue
				}
				literal, ok := typed.Values[index].(*ast.BasicLit)
				if !ok || literal.Kind != token.STRING {
					continue
				}
				value, unquoteErr := strconv.Unquote(literal.Value)
				if unquoteErr == nil && strings.HasPrefix(value, "/") {
					*violations = append(*violations, modelBoundaryPosition(fset, name.Pos())+": exported route string "+name.Name)
				}
			}
		}
	}
}

func modelBoundaryValidateIntrinsicMethod(fset *token.FileSet, method *ast.FuncDecl, violations *[]string) {
	if method.Name.Name != "String" && method.Name.Name != "IsValid" {
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
