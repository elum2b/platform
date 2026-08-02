package socket

import (
	etp "github.com/elum-utils/go-etp"
	"github.com/go-playground/validator/v10"
	json "github.com/goccy/go-json"
)

var validate = validator.New()

func Decode(ctx *etp.Context, data any) bool {
	body, err := ctx.Bytes()
	if err != nil {
		return false
	}

	if err := json.Unmarshal(body, data); err != nil {
		return false
	}

	return validate.Struct(data) == nil
}

func Respond(ctx *etp.Context, event string, response any) error {
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	_, err = ctx.Respond(etp.SendOptions{
		Event: event,
		Data:  data,
	})

	return err
}
