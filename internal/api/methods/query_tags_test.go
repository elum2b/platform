package methods

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestRequestJSONFieldsHaveMatchingQueryTags(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("determine methods directory")
	}

	methodsDir := filepath.Dir(filename)
	checkFile := func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		return checkRequestFile(t, path)
	}

	err := filepath.WalkDir(methodsDir, checkFile)
	if err != nil {
		t.Fatal(err)
	}
}

func checkRequestFile(t *testing.T, path string) error {
	t.Helper()

	file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		return err
	}

	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}

		for _, spec := range gen.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok || !strings.HasSuffix(typeSpec.Name.Name, "Request") {
				continue
			}

			checkRequestStruct(t, path, typeSpec.Name.Name, structType)
		}
	}

	return nil
}

func checkRequestStruct(
	t *testing.T,
	path, typeName string,
	structType *ast.StructType,
) {
	t.Helper()

	for _, field := range structType.Fields.List {
		if field.Tag == nil {
			continue
		}

		tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
		jsonName := strings.Split(tag.Get("json"), ",")[0]

		if jsonName == "" || jsonName == "-" || tag.Get("query") == jsonName {
			continue
		}

		t.Errorf(
			"%s: %s field with json:%q must have query:%q",
			path,
			typeName,
			jsonName,
			jsonName,
		)
	}
}
