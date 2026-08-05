package user

import (
	"time"

	caluser "github.com/elum2b/services/calendar/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListActiveRequest struct {
	WorkspaceID string    `json:"workspace_id"     validate:"required,uuid"`
	AppID       int64     `json:"app_id"           validate:"required,min=1"`
	PlatformID  int64     `json:"platform_id"      validate:"required,min=1"`
	Params      string    `json:"params"           validate:"required"`
	Locale      string    `json:"locale,omitempty"`
	Now         time.Time `json:"now,omitempty"`
}

type ListActiveResponse struct {
	Calendars []caluser.ActiveCalendarModel `json:"calendars"`
}

var (
	listActiveKey         = "calendar.user.list_active"
	listActiveDescription = `
Lists active calendars available to the authenticated application user.`
)

// ListActive exposes the calendar user active listing method.
var ListActive = adapter.Method[ListActiveRequest, ListActiveResponse]{
	Key:         listActiveKey,
	Description: listActiveDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data ListActiveRequest) (ListActiveResponse, error) {
		calendars, err := services.Calendar.User.ListActive(
			ctx.Context,
			caluser.ListActiveParams{
				WorkspaceID: data.WorkspaceID,
				Locale:      data.Locale,
				Now:         data.Now,
			},
		)
		if err != nil {
			return ListActiveResponse{}, err
		}

		return ListActiveResponse{Calendars: calendars}, nil
	},
}
