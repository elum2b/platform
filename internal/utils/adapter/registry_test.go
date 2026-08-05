package adapter

import (
	"reflect"
	"testing"
)

func TestGroupInheritsMiddlewareInOrder(t *testing.T) {
	registry := Registry{}
	order := make([]string, 0, 3)

	parent := registry.Group(func(*Context) error {
		order = append(order, "parent")

		return nil
	})
	child := parent.Group(func(*Context) error {
		order = append(order, "child")

		return nil
	})
	child.Use(func(*Context) error {
		order = append(order, "used")

		return nil
	})

	_, middleware := child.registration()
	for _, handler := range middleware {
		if err := handler(&Context{}); err != nil {
			t.Fatalf("middleware error = %v", err)
		}
	}

	want := []string{"parent", "child", "used"}
	if !reflect.DeepEqual(order, want) {
		t.Fatalf("middleware order = %v, want %v", order, want)
	}
}
