package promo

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}

type GetResponse struct {
	Promo promoadmin.PromoModel `json:"promo"`
}

var (
	getKey         = "promo.get"
	getDescription = `
Returns a promo with its localizations and rewards. Requires the 'promo.get'
permission in the target workspace.`
)

// Get exposes the promo retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		promo, err := services.Promo.Admin.GetPromo(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
		)

		return GetResponse{Promo: promo}, err
	},
}
