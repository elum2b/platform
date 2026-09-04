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

func TestDecodeGETQueryBinding(t *testing.T) {
	type request struct {
		ID string `json:"id" query:"id" validate:"required,uuid"`

		Name string `json:"name" query:"name" validate:"required"`

		Count int `json:"count" query:"count" validate:"required,min=1"`

		Labels []string `json:"labels" query:"labels" validate:"required,min=2"`
	}

	app := fiber.New()
	app.Get("/", func(ctx fiber.Ctx) error {
		var data request

		if !Decode(ctx, &data) {
			return Error(ctx, serviceerrors.ErrInvalidFields)
		}

		if data.ID != "f47ac10b-58cc-4372-a567-0e02b2c3d479" ||
			data.Name != "test" ||
			data.Count != 2 ||
			len(data.Labels) != 2 ||
			data.Labels[0] != "first" ||
			data.Labels[1] != "second" {
			return Error(ctx, serviceerrors.ErrInvalidFields)
		}

		return ctx.SendStatus(http.StatusOK)
	})

	for _, test := range []struct {
		name   string
		query  string
		status int
	}{
		{
			name:   "valid",
			query:  "?id=f47ac10b-58cc-4372-a567-0e02b2c3d479&name=test&count=2&labels=first&labels=second",
			status: http.StatusOK,
		},
		{
			name:   "invalid integer",
			query:  "?id=f47ac10b-58cc-4372-a567-0e02b2c3d479&name=test&count=invalid&labels=first&labels=second",
			status: http.StatusBadRequest,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			response, err := app.Test(httptest.NewRequestWithContext(
				context.Background(),
				http.MethodGet,
				"/"+test.query,
				http.NoBody,
			))
			if err != nil {
				t.Fatalf("app.Test() error = %v", err)
			}
			defer response.Body.Close()

			if response.StatusCode != test.status {
				t.Errorf(
					"GET binding status = %d, want %d",
					response.StatusCode,
					test.status,
				)
			}
		})
	}
}
