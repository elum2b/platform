package user

import (
	caluser "github.com/elum2b/services/calendar/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RecordRequest struct {
	WorkspaceID string `json:"workspace_id" query:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"       query:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"  query:"platform_id"  validate:"required,min=1"`
	CalendarRef string `json:"calendar_ref" query:"calendar_ref" validate:"required,max=255"`
	OperationID string `json:"operation_id" query:"operation_id" validate:"required"`
}

type RecordResponse struct {
	Result caluser.RecordResult `json:"result"`
}

var (
	recordKey         = "calendar.user.record"
	recordDescription = `
Records a calendar operation for the authenticated application user.`
)

// Record exposes the calendar user record method.
var Record = adapter.Method[RecordRequest, RecordResponse]{
	Key:         recordKey,
	Description: recordDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data RecordRequest) (RecordResponse, error) {
		result, err := services.Calendar.User.Record(
			ctx.Context,
			caluser.RecordParams{
				Identity:    *ctx.Identity,
				CalendarRef: data.CalendarRef,
				OperationID: data.OperationID,
			},
		)

		return RecordResponse{Result: result}, err
	},
}
