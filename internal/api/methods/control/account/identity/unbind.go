package identity

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UnbindRequest struct {
	Provider string `json:"provider" validate:"required"`
}

type UnbindResponse struct {
	Affected int64 `json:"affected"`
}

var (
	unbindKey         = "control.account.identity.unbind"
	unbindDescription = `
Removes an external identity from the authenticated account.`
)

// Unbind removes an identity from the authenticated account.
var Unbind = adapter.Method[UnbindRequest, UnbindResponse]{
	Key:         unbindKey,
	Description: unbindDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data UnbindRequest) (UnbindResponse, error) {
		affected, err := services.Control.Admin.UnbindIdentity(
			ctx.Context,
			ctx.AccountID,
			data.Provider,
		)
		if err != nil {
			return UnbindResponse{}, err
		}

		return UnbindResponse{Affected: affected}, nil
	},
}
