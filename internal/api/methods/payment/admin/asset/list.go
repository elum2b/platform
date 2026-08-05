package asset

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListResponse struct {
	Assets []padm.AssetModel `json:"assets"`
}

var (
	listKey         = "payment.asset.list"
	listDescription = `
Lists payment assets. Requires the 'payment.asset.list'
permission in the target workspace.`
)

var List = adapter.Method[struct{}, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, _ struct{}) (ListResponse, error) {
		v, err := services.Payment.Admin.ListAssets(ctx.Context)
		if err != nil {
			return ListResponse{}, err
		}
		return ListResponse{Assets: v}, nil
	},
}
