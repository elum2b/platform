package apply

import (
	promouser "github.com/elum2b/services/promo/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Request struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"  validate:"required,min=1"`
	Params      string `json:"params"       validate:"required"`
	Code        string `json:"code"         validate:"required"`
	Locale      string `json:"locale,omitempty"`
}

type Response struct {
	Result promouser.ApplyResult `json:"result"`
}

var (
	applyKey         = "promo.apply"
	applyDescription = `
Redeems a promo code for the authenticated application user. Returns the promo
details and redemption status.`
)

// Method exposes the promo apply method.
var Method = adapter.Method[Request, Response]{
	Key:         applyKey,
	Description: applyDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		result, err := services.Promo.User.Apply(
			ctx.Context,
			promouser.ApplyParams{
				Identity: *ctx.Identity,
				Code:     data.Code,
				Locale:   data.Locale,
			},
		)
		if err != nil {
			return Response{}, err
		}

		return Response{Result: result}, nil
	},
}
