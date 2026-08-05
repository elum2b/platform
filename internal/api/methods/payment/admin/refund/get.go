package refund

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}

type GetResponse struct {
	Refund padm.RefundModel `json:"refund"`
}

var (
	getKey         = "payment.refund.get"
	getDescription = `
Returns a refund. Requires the 'payment.refund.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetRefund(
			ctx.Context,
			padm.RefundRefParams{
				WorkspaceID: d.WorkspaceID,
				ID:          d.ID,
			})
		return GetResponse{Refund: v}, err
	},
}
