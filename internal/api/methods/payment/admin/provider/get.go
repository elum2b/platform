package provider

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	Code string `json:"code" validate:"required,max=255"`
}

type GetResponse struct {
	Provider padm.ProviderModel `json:"provider"`
}

var (
	getKey         = "payment.provider.get"
	getDescription = `
Returns a payment provider. Requires the 'payment.provider.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetProvider(
			ctx.Context, d.Code)

		return GetResponse{Provider: v}, err
	},
}
