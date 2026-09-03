package main

import (
	"reflect"
	"testing"

	"github.com/invopop/jsonschema"
)

func TestApplyRules(t *testing.T) {
	schema := &jsonschema.Schema{Type: "string"}

	applyRules(
		schema,
		reflect.TypeFor[string](),
		"required,min=2,max=4,uuid,oneof=one two",
	)

	if schema.MinLength == nil || *schema.MinLength != 2 {
		t.Fatalf("minLength = %v, want 2", schema.MinLength)
	}

	if schema.MaxLength == nil || *schema.MaxLength != 4 {
		t.Fatalf("maxLength = %v, want 4", schema.MaxLength)
	}

	if schema.Format != "uuid" {
		t.Fatalf("format = %q, want uuid", schema.Format)
	}

	if !reflect.DeepEqual(schema.Enum, []any{"one", "two"}) {
		t.Fatalf("enum = %#v, want [one two]", schema.Enum)
	}
}

func TestApplyRulesDive(t *testing.T) {
	schema := &jsonschema.Schema{
		Type:  "array",
		Items: &jsonschema.Schema{Type: "string"},
	}

	applyRules(schema, reflect.TypeFor[[]string](), "dive,required,uuid")

	if schema.Items.MinLength == nil || *schema.Items.MinLength != 1 {
		t.Fatalf("items minLength = %v, want 1", schema.Items.MinLength)
	}

	if schema.Items.Format != "uuid" {
		t.Fatalf("items format = %q, want uuid", schema.Items.Format)
	}
}
