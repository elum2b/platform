package user

import (
	caluser "github.com/elum2b/services/calendar/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id"     query:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"           query:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"      query:"platform_id"  validate:"required,min=1"`
	Ref         string `json:"ref"              query:"ref"          validate:"required,max=255"`
	Locale      string `json:"locale,omitempty" query:"locale"`
}

type GetResponse struct {
	Calendar caluser.CalendarModel `json:"calendar"`
}

var (
	getKey         = "calendar.user.get"
	getDescription = `
Returns a calendar with its steps and rewards for the authenticated
application user.`
)

// Get exposes the calendar user retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		cal, err := services.Calendar.User.GetCalendar(
			ctx.Context,
			caluser.GetCalendarParams{
				Identity: *ctx.Identity,
				Ref:      data.Ref,
				Locale:   data.Locale,
			},
		)

		return GetResponse{Calendar: cal}, err
	},
}
