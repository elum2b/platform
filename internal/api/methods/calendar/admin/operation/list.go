package operation

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"     validate:"required,uuid"`
	CalendarID  string `json:"calendar_id"      validate:"required,max=255"`
	Limit       int32  `json:"limit,omitempty"  validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty" validate:"min=0"`
}

type ListResponse struct {
	Operations []caladmin.OperationModel `json:"operations"`
}

var (
	listKey         = "calendar.operation.list"
	listDescription = `
Lists operations of a calendar. Requires the 'calendar.operation.list'
permission in the target workspace.`
)

// List exposes the calendar operation listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Calendar.Admin.ListOperations(
			ctx.Context,
			data.WorkspaceID,
			data.CalendarID,
			caladmin.Page{Limit: data.Limit, Offset: data.Offset},
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Operations: values}, nil
	},
}
