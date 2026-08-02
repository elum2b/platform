package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	serviceerrors "github.com/elum2b/services/errors"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

func TestRespondEnvelope(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(ctx fiber.Ctx) error {
		return Respond(ctx, map[string]string{"id": "test"})
	})

	response, err := app.Test(httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/",
		http.NoBody,
	))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	var body struct {
		Response map[string]string `json:"response"`
	}

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if body.Response["id"] != "test" {
		t.Fatalf("response.id = %q, want test", body.Response["id"])
	}
}

func TestErrorEnvelope(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(ctx fiber.Ctx) error {
		return Error(ctx, serviceerrors.ErrInvalidFields)
	})

	response, err := app.Test(httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/",
		http.NoBody,
	))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	var body ErrorResponse

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if body.Error.Key != serviceerrors.CodeInvalidFields {
		t.Fatalf(
			"error.key = %q, want %q",
			body.Error.Key,
			serviceerrors.CodeInvalidFields,
		)
	}
}
