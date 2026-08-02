package identity

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	authutils "github.com/elum2b/platform/internal/utils/auth"
)

type BindRequest authutils.IdentityRequest

type BindResponse struct {
	Bound bool `json:"bound"`
}

var (
	bindKey         = "control.account.identity.bind"
	bindDescription = `
Binds a verified external identity to the authenticated account.`
)

// Bind binds a verified identity to the authenticated account.
var Bind = adapter.Method[BindRequest, BindResponse]{
	Key:         bindKey,
	Description: bindDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data BindRequest) (BindResponse, error) {
		identity, err := authutils.ResolveIdentity(
			ctx.Context,
			authutils.IdentityRequest(data),
		)
		if err != nil {
			return BindResponse{}, err
		}

		if err := services.Control.Admin.BindIdentity(
			ctx.Context,
			ctx.AccountID,
			identity,
		); err != nil {
			return BindResponse{}, err
		}

		return BindResponse{Bound: true}, nil
	},
}
