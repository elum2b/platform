package localization

import (
	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ItemKey     string `json:"item_key"     validate:"required,max=255"`
	Locale      string `json:"locale"       validate:"required,max=32"`
}

type GetResponse struct {
	Localization refadmin.LocalizationModel `json:"localization"`
}

var (
	getKey         = "reference.localization.get"
	getDescription = `
Returns a single reference item localization. Requires the
'reference.localization.get' permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Reference.Admin.GetLocalization(
			ctx.Context, d.WorkspaceID, d.ItemKey, d.Locale)

		return GetResponse{Localization: v}, err
	},
}
