package localization

import (
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
}
type ListResponse struct {
	Localizations []cpaadmin.LocalizationModel `json:"localizations"`
}

var (
	listKey         = "cpa.localization.list"
	listDescription = `
Lists the localizations of a CPA offer. Requires the 'cpa.localization.list'
permission in the target workspace.`
)

// List exposes the CPA localization listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.CPA.Admin.ListLocalizations(
			ctx.Context,
			data.WorkspaceID,
			data.CPAID,
		)

		return ListResponse{Localizations: values}, err
	},
}
