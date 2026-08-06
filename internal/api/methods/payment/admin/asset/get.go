package asset

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	Code string `json:"code" validate:"required,max=255"`
}

type GetResponse struct {
	Asset padm.AssetModel `json:"asset"`
}

var (
	getKey         = "payment.asset.get"
	getDescription = `
Returns a payment asset. Requires the 'payment.asset.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetAsset(
			ctx.Context, d.Code)

		return GetResponse{Asset: v}, err
	},
}
