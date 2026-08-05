package localization

import (
	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ItemKey     string `json:"item_key"     validate:"required,max=255"`
}

type ListResponse struct {
	Localizations []refadmin.LocalizationModel `json:"localizations"`
}

var (
	listKey         = "reference.localization.list"
	listDescription = `
Lists the localizations of a reference item. Requires the
'reference.localization.list' permission in the target workspace.`
)

// List exposes the reference localization listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Reference.Admin.ListLocalizations(
			ctx.Context,
			data.WorkspaceID,
			data.ItemKey,
		)

		return ListResponse{Localizations: values}, err
	},
}
