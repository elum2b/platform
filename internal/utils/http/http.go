package http

import (
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

type ErrorResponse struct {
	Error ErrorData `json:"error"`
}

type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Decode(ctx fiber.Ctx, data any) bool {
	if err := ctx.Bind().Body(data); err != nil {
		return false
	}

	return validate.Struct(data) == nil
}

func Respond(ctx fiber.Ctx, response any) error {
	return ctx.JSON(response)
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
			Code:    code,
			Message: message,
		},
	})
}
