// Command schema writes the platform API method catalog as JSON Schema.
//
// Run manually with: go generate ./schema
//
//go:generate go run .

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/invopop/jsonschema"

	"github.com/elum2b/platform/internal/api/methods"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type catalog struct {
	Methods map[string]methodSchema `json:"methods"`
}

type methodSchema struct {
	Key         string             `json:"key"`
	Description string             `json:"description"`
	Transports  []string           `json:"transports"`
	HTTPMethod  string             `json:"http_method,omitempty"`
	Input       *jsonschema.Schema `json:"input"`
	Output      *jsonschema.Schema `json:"output"`
}

func (catalog *catalog) Add(method adapter.MethodInfo) {
	if _, exists := catalog.Methods[method.Key]; exists {
		panic(fmt.Sprintf("schema: duplicate method key %q", method.Key))
	}

	schema := methodSchema{
		Key:         method.Key,
		Description: method.Description,
		HTTPMethod:  method.HTTPMethod,
		Input:       schemaFor(method.Input),
		Output:      schemaFor(method.Output),
	}
	if method.Transports&adapter.HTTP != 0 {
		schema.Transports = append(schema.Transports, "http")
	}

	if method.Transports&adapter.WS != 0 {
		schema.Transports = append(schema.Transports, "ws")
	}

	if method.Transports&adapter.MCP != 0 {
		schema.Transports = append(schema.Transports, "mcp")
	}

	catalog.Methods[method.Key] = schema
}

func main() {
	output := flag.String("out", "schema.json", "path to the generated schema")

	flag.Parse()

	catalog := &catalog{Methods: make(map[string]methodSchema)}
	methods.Register(adapter.Registry{Catalog: catalog})

	file, err := os.Create(*output)
	if err != nil {
		panic(fmt.Errorf("schema: create %s: %w", *output, err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(catalog); err != nil {
		panic(fmt.Errorf("schema: write %s: %w", filepath.Clean(*output), err))
	}
}

func schemaFor(typ reflect.Type) *jsonschema.Schema {
	schema := jsonschema.ReflectFromType(typ)
	applyValidation(
		schema,
		typ,
		schema.Definitions,
		make(map[reflect.Type]bool),
	)

	return schema
}

func applyValidation(
	schema *jsonschema.Schema,
	typ reflect.Type,
	definitions jsonschema.Definitions,
	visited map[reflect.Type]bool,
) {
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	schema = resolve(schema, definitions)
	if schema == nil {
		return
	}

	switch typ.Kind() {
	case reflect.Struct:
		applyStructValidation(schema, typ, definitions, visited)
	case reflect.Slice, reflect.Array:
		applyValidation(schema.Items, typ.Elem(), definitions, visited)
	case reflect.Map:
		applyValidation(
			schema.AdditionalProperties,
			typ.Elem(),
			definitions,
			visited,
		)
	}
}

func applyStructValidation(
	schema *jsonschema.Schema,
	typ reflect.Type,
	definitions jsonschema.Definitions,
	visited map[reflect.Type]bool,
) {
	if visited[typ] {
		return
	}

	visited[typ] = true

	for field := range typ.Fields() {
		applyFieldValidation(schema, field, definitions, visited)
	}
}

func applyFieldValidation(
	schema *jsonschema.Schema,
	field reflect.StructField,
	definitions jsonschema.Definitions,
	visited map[reflect.Type]bool,
) {
	name, embedded := jsonField(field)
	if name == "" {
		return
	}

	if embedded {
		applyValidation(schema, field.Type, definitions, visited)

		return
	}

	property, exists := schema.Properties.Get(name)
	if !exists {
		return
	}

	tag := field.Tag.Get("validate")
	applyRules(property, field.Type, tag)

	if hasRule(tag, "required") {
		schema.Required = appendUnique(schema.Required, name)
	}

	applyValidation(property, field.Type, definitions, visited)
}

func applyRules(schema *jsonschema.Schema, typ reflect.Type, tag string) {
	if schema == nil || tag == "" {
		return
	}

	rules := strings.Split(tag, ",")
	for index, rule := range rules {
		name, value, _ := strings.Cut(rule, "=")
		switch name {
		case "required":
			applyRequired(schema, typ)
		case "uuid":
			schema.Format = "uuid"
		case "min":
			applyBound(schema, typ, value, true)
		case "max":
			applyBound(schema, typ, value, false)
		case "oneof":
			applyEnum(schema, typ, value)
		case "dive":
			applyDiveRules(schema, indirectType(typ), rules[index+1:])

			return
		}
	}
}

func applyDiveRules(
	schema *jsonschema.Schema,
	typ reflect.Type,
	rules []string,
) {
	tag := strings.Join(rules, ",")

	switch typ.Kind() {
	case reflect.Slice, reflect.Array:
		applyRules(schema.Items, typ.Elem(), tag)
	case reflect.Map:
		applyRules(schema.AdditionalProperties, typ.Elem(), tag)
	}
}

func applyRequired(schema *jsonschema.Schema, typ reflect.Type) {
	switch indirectType(typ).Kind() {
	case reflect.String:
		setMinLength(schema, 1)
	case reflect.Slice, reflect.Array:
		setMinItems(schema, 1)
	case reflect.Map:
		setMinProperties(schema, 1)
	case reflect.Bool:
		schema.Const = true
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64:
		schema.Not = &jsonschema.Schema{Const: 0}
	}
}

func applyBound(
	schema *jsonschema.Schema,
	typ reflect.Type,
	value string,
	minimum bool,
) {
	typ = indirectType(typ)

	switch typ.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		applyCollectionBound(schema, typ.Kind(), value, minimum)
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64:
		applyNumericBound(schema, value, minimum)
	}
}

func applyCollectionBound(
	schema *jsonschema.Schema,
	kind reflect.Kind,
	value string,
	minimum bool,
) {
	bound, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return
	}

	switch kind {
	case reflect.String:
		if minimum {
			setMinLength(schema, bound)
		} else {
			setMaxLength(schema, bound)
		}
	case reflect.Slice, reflect.Array:
		if minimum {
			setMinItems(schema, bound)
		} else {
			setMaxItems(schema, bound)
		}
	case reflect.Map:
		if minimum {
			setMinProperties(schema, bound)
		} else {
			setMaxProperties(schema, bound)
		}
	}
}

func applyNumericBound(schema *jsonschema.Schema, value string, minimum bool) {
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		return
	}

	if minimum {
		schema.Minimum = json.Number(value)

		return
	}

	schema.Maximum = json.Number(value)
}

func applyEnum(schema *jsonschema.Schema, typ reflect.Type, value string) {
	for item := range strings.FieldsSeq(value) {
		switch indirectType(typ).Kind() {
		case reflect.String:
			schema.Enum = append(schema.Enum, item)
		case reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			if parsed, err := strconv.ParseInt(item, 10, 64); err == nil {
				schema.Enum = append(schema.Enum, parsed)
			}
		case reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64:
			if parsed, err := strconv.ParseUint(item, 10, 64); err == nil {
				schema.Enum = append(schema.Enum, parsed)
			}
		case reflect.Float32, reflect.Float64:
			if parsed, err := strconv.ParseFloat(item, 64); err == nil {
				schema.Enum = append(schema.Enum, parsed)
			}
		case reflect.Bool:
			if parsed, err := strconv.ParseBool(item); err == nil {
				schema.Enum = append(schema.Enum, parsed)
			}
		}
	}
}

func jsonField(field reflect.StructField) (string, bool) {
	if field.PkgPath != "" {
		return "", false
	}

	parts := strings.Split(field.Tag.Get("json"), ",")
	if parts[0] == "-" {
		return "", false
	}

	if field.Anonymous && parts[0] == "" {
		return field.Name, true
	}

	if parts[0] != "" {
		return parts[0], false
	}

	return field.Name, false
}

func resolve(
	schema *jsonschema.Schema,
	definitions jsonschema.Definitions,
) *jsonschema.Schema {
	const prefix = "#/$defs/"

	if schema != nil && strings.HasPrefix(schema.Ref, prefix) {
		return definitions[strings.TrimPrefix(schema.Ref, prefix)]
	}

	return schema
}

func indirectType(typ reflect.Type) reflect.Type {
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	return typ
}

func hasRule(tag string, target string) bool {
	for rule := range strings.SplitSeq(tag, ",") {
		name, _, _ := strings.Cut(rule, "=")
		if name == target {
			return true
		}
	}

	return false
}

func appendUnique(values []string, value string) []string {
	for _, existing := range values {
		if existing == value {
			return values
		}
	}

	return append(values, value)
}

func setMinLength(schema *jsonschema.Schema, value uint64) {
	if schema.MinLength == nil || *schema.MinLength < value {
		schema.MinLength = &value
	}
}
func setMaxLength(schema *jsonschema.Schema, value uint64) {
	if schema.MaxLength == nil || *schema.MaxLength > value {
		schema.MaxLength = &value
	}
}
func setMinItems(schema *jsonschema.Schema, value uint64) {
	if schema.MinItems == nil || *schema.MinItems < value {
		schema.MinItems = &value
	}
}
func setMaxItems(schema *jsonschema.Schema, value uint64) {
	if schema.MaxItems == nil || *schema.MaxItems > value {
		schema.MaxItems = &value
	}
}
func setMinProperties(schema *jsonschema.Schema, value uint64) {
	if schema.MinProperties == nil || *schema.MinProperties < value {
		schema.MinProperties = &value
	}
}
func setMaxProperties(schema *jsonschema.Schema, value uint64) {
	if schema.MaxProperties == nil || *schema.MaxProperties > value {
		schema.MaxProperties = &value
	}
}
