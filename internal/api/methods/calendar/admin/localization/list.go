package localization

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CalendarID  string `json:"calendar_id"  validate:"required,max=255"`
}

type ListResponse struct {
	Localizations []caladmin.LocalizationModel `json:"localizations"`
}

var (
	listKey         = "calendar.localization.list"
	listDescription = `
Lists the localizations of a calendar. Requires the
'calendar.localization.list' permission in the target workspace.`
)

// List exposes the calendar localization listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Calendar.Admin.ListLocalizations(
			ctx.Context,
			data.WorkspaceID,
			data.CalendarID,
		)

		return ListResponse{Localizations: values}, err
	},
}
