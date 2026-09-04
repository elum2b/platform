package user

import (
	caluser "github.com/elum2b/services/calendar/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetProgressRequest struct {
	WorkspaceID string `json:"workspace_id" query:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"       query:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"  query:"platform_id"  validate:"required,min=1"`
	CalendarID  string `json:"calendar_id"  query:"calendar_id"  validate:"required,max=255"`
}

type GetProgressResponse struct {
	Progress *caluser.ProgressModel `json:"progress"`
}

var (
	getProgressKey         = "calendar.user.progress.get"
	getProgressDescription = `
Returns the application user's progress on a calendar.`
)

// GetProgress exposes the calendar user progress retrieval method.
var GetProgress = adapter.Method[GetProgressRequest, GetProgressResponse]{
	Key:         getProgressKey,
	Description: getProgressDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data GetProgressRequest) (GetProgressResponse, error) {
		progress, err := services.Calendar.User.GetProgress(
			ctx.Context,
			caluser.GetProgressParams{
				Identity:   *ctx.Identity,
				CalendarID: data.CalendarID,
			},
		)

		return GetProgressResponse{Progress: progress}, err
	},
}
