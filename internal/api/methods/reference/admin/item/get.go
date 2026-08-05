package item

import (
	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Key         string `json:"key"          validate:"required,max=255"`
}

type GetResponse struct {
	Item refadmin.ItemModel `json:"item"`
}

var (
	getKey         = "reference.get"
	getDescription = `
Returns a reference item with its localizations. Requires the 'reference.get'
permission in the target workspace.`
)

// Get exposes the reference item retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		item, err := services.Reference.Admin.GetItem(
			ctx.Context,
			data.WorkspaceID,
			data.Key,
		)

		return GetResponse{Item: item}, err
	},
}
