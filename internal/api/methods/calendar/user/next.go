package user

import (
	caluser "github.com/elum2b/services/calendar/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type NextRequest struct {
	WorkspaceID string `json:"workspace_id"     query:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"           query:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"      query:"platform_id"  validate:"required,min=1"`
	CalendarRef string `json:"calendar_ref"     query:"calendar_ref" validate:"required,max=255"`
	Locale      string `json:"locale,omitempty" query:"locale"`
}

type NextResponse struct {
	Result caluser.RecordResult `json:"result"`
}

var (
	nextKey         = "calendar.user.next"
	nextDescription = `
Advances the authenticated application user to the next claimable step of a
calendar and grants its rewards.`
)

// Next exposes the calendar user next step method.
var Next = adapter.Method[NextRequest, NextResponse]{
	Key:         nextKey,
	Description: nextDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data NextRequest) (NextResponse, error) {
		result, err := services.Calendar.User.Next(
			ctx.Context,
			caluser.NextParams{
				Identity:    *ctx.Identity,
				CalendarRef: data.CalendarRef,
				Locale:      data.Locale,
			},
		)

		return NextResponse{Result: result}, err
	},
}
