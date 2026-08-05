package localization

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CalendarID  string `json:"calendar_id"  validate:"required,max=255"`
	Locale      string `json:"locale"       validate:"required,max=32"`
	Title       string `json:"title"        validate:"required"`
	Description string `json:"description"  validate:"required"`
}

var (
	upsertKey         = "calendar.localization.upsert"
	upsertDescription = `
Creates or updates a calendar localization. Requires the
'calendar.localization.upsert' permission in the target workspace.`
)

// Upsert exposes the calendar localization upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.Calendar.Admin.UpsertLocalization(
			ctx.Context,
			caladmin.SaveLocalizationParams{
				WorkspaceID: data.WorkspaceID,
				CalendarID:  data.CalendarID,
				Locale:      data.Locale,
				Title:       data.Title,
				Description: data.Description,
			},
		)

		return struct{}{}, err
	},
}
