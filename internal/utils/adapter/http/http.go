package http

import (
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

type SuccessResponse struct {
	Response any `json:"response"`
}

type ErrorResponse struct {
	Error ErrorData `json:"error"`
}

type ErrorData struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

// RegisterFunc registers HTTP methods on the control router.
type RegisterFunc func(fiber.Router)

// Init registers the internal HTTP API used exclusively for authentication.
func Init(app fiber.Router, register RegisterFunc) {
	internal := app.Group("/internal")
	control := internal.Group("/control")

	register(control)
}

func Decode(ctx fiber.Ctx, data any) bool {
	if err := ctx.Bind().All(data); err != nil {
		return false
	}

	return validate.Struct(data) == nil
}

func Respond(ctx fiber.Ctx, response any) error {
	return ctx.JSON(SuccessResponse{Response: response})
}

func Error(ctx fiber.Ctx, err error) error {
	code := serviceerrors.CodeOf(err)
	message := serviceerrors.MessageOf(err)
	status := fiber.StatusInternalServerError

	switch code {
	case serviceerrors.CodeInvalidFields:
		status = fiber.StatusBadRequest
	case serviceerrors.CodeUnauthorized:
		status = fiber.StatusUnauthorized
	case serviceerrors.CodeForbidden:
		status = fiber.StatusForbidden
	case serviceerrors.CodeNotFound:
		status = fiber.StatusNotFound
	case serviceerrors.CodeConflict, serviceerrors.CodeDuplicate:
		status = fiber.StatusConflict
	case serviceerrors.CodeFailedPrecondition:
		status = fiber.StatusPreconditionFailed
	case serviceerrors.CodeRateLimit:
		status = fiber.StatusTooManyRequests
	case serviceerrors.CodeTimeout:
		status = fiber.StatusGatewayTimeout
	case serviceerrors.CodeNotReady, serviceerrors.CodeUnavailable:
		status = fiber.StatusServiceUnavailable
	case serviceerrors.CodeUnsupported:
		status = fiber.StatusNotImplemented
	}

	if code == "" {
		code = serviceerrors.CodeInternalError
	}

	if message == "" {
		message = "internal error"
	}

	return ctx.Status(status).JSON(ErrorResponse{
		Error: ErrorData{
			Key:     code,
			Message: message,
		},
	})
}
