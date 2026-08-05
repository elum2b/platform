package localization

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	PromoID     uint64 `json:"promo_id"     validate:"required,min=1"`
}

type ListResponse struct {
	Localizations []promoadmin.LocalizationModel `json:"localizations"`
}

var (
	listKey         = "promo.localization.list"
	listDescription = `
Lists the localizations of a promo. Requires the 'promo.localization.list'
permission in the target workspace.`
)

// List exposes the promo localization listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Promo.Admin.ListLocalizations(
			ctx.Context,
			data.WorkspaceID,
			data.PromoID,
		)

		return ListResponse{Localizations: values}, err
	},
}
