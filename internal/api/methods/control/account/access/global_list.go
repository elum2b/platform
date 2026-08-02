package access

import (
	controlinternal "github.com/elum2b/services/control/service/internalapi"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GlobalMethodResponse struct {
	Key      string                      `json:"key"`
	Service  string                      `json:"service"`
	GroupKey string                      `json:"group_key"`
	Scope    controlinternal.AccessScope `json:"scope"`
	Position int32                       `json:"position"`
}

type GlobalListResponse struct {
	Methods []GlobalMethodResponse `json:"methods"`
}

var (
	globalListKey         = "control.account.access.global.list"
	globalListDescription = `
Lists global methods available to the current account.`
)

// GlobalList returns global methods authorized for the account.
var GlobalList = adapter.Method[struct{}, GlobalListResponse]{
	Key:         globalListKey,
	Description: globalListDescription,
	Transports:  adapter.WS | adapter.MCP,
	Handler: func(ctx *adapter.Context, _ struct{}) (GlobalListResponse, error) {
		methods, err := services.Control.Internal.GetAuthorizedGlobalMethods(
			ctx.Context,
			ctx.AccountID,
		)
		if err != nil {
			return GlobalListResponse{}, err
		}

		response := make([]GlobalMethodResponse, 0, len(methods))
		for _, method := range methods {
			response = append(response, GlobalMethodResponse{
				Key: method.Key, Service: method.Service,
				GroupKey: method.GroupKey, Scope: method.Scope,
				Position: method.Position,
			})
		}

		return GlobalListResponse{Methods: response}, nil
	},
}
