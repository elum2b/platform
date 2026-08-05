package group

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Code        string `json:"code"         validate:"required,max=255"`
}

type GetResponse struct {
	Group padm.ProductGroupModel `json:"group"`
}

var (
	getKey         = "payment.product_group.get"
	getDescription = `
Returns a product group. Requires the 'payment.product_group.get' permission
in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetProductGroup(
			ctx.Context, d.WorkspaceID, d.Code)
		return GetResponse{Group: v}, err
	},
}
