package localization

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CalendarID  string `json:"calendar_id"  validate:"required,max=255"`
	Locale      string `json:"locale"       validate:"required,max=32"`
}

type GetResponse struct {
	Localization caladmin.LocalizationModel `json:"localization"`
}

var (
	getKey         = "calendar.localization.get"
	getDescription = `
Returns a single calendar localization. Requires the
'calendar.localization.get' permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Calendar.Admin.GetLocalization(
			ctx.Context, d.WorkspaceID, d.CalendarID, d.Locale)

		return GetResponse{Localization: v}, err
	},
}
