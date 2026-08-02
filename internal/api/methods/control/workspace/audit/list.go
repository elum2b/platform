package audit

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string    `json:"workspace_id"        validate:"required,uuid"`
	Limit       int32     `json:"limit,omitempty"     validate:"omitempty,min=1,max=100"`
	CursorAt    time.Time `json:"cursor_at,omitempty"`
	CursorID    string    `json:"cursor_id,omitempty"`
}

type ListResponse struct {
	Items []controladmin.AuditEventModel `json:"items"`
}

var (
	listKey         = "control.workspace.audit.list"
	listDescription = `
Lists audit entries for a workspace. Requires the
'control.workspace.audit.list' permission in the target workspace.`
)

// List returns workspace audit entries.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.audit.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListWorkspaceAudit(
			ctx.Context,
			data.WorkspaceID,
			controladmin.Page{
				Limit:    data.Limit,
				CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Items: items}, nil
	},
}
