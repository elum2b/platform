package code

import (
	cpauser "github.com/elum2b/services/cpa/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"  validate:"required,min=1"`
	Params      string `json:"params"       validate:"required"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
}

type GetResponse struct {
	Result cpauser.GetCodeResult `json:"result"`
}

var (
	getKey         = "cpa.user.code.get"
	getDescription = `
Issues or returns the authenticated application user's code for a CPA offer.`
)

// Get exposes the CPA user code method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		result, err := services.CPA.User.GetCode(
			ctx.Context,
			cpauser.GetCodeParams{
				Identity: *ctx.Identity,
				CPAID:    data.CPAID,
			},
		)
		if err != nil {
			return GetResponse{}, err
		}

		return GetResponse{Result: result}, nil
	},
}
