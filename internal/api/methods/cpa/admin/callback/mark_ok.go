package callback

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type MarkOKRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}
type MarkOKResponse struct {
	Affected int64 `json:"affected"`
}

var (
	markOKKey         = "cpa.callback.mark_ok"
	markOKDescription = `
Marks a CPA callback event as successfully delivered. Requires the
'cpa.callback.mark_ok' permission in the target workspace.`
)

// MarkOK exposes the CPA callback success method.
var MarkOK = adapter.Method[MarkOKRequest, MarkOKResponse]{
	Key:         markOKKey,
	Description: markOKDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(markOKKey)},
	Handler: func(ctx *adapter.Context, data MarkOKRequest) (MarkOKResponse, error) {
		affected, err := services.CPA.Admin.MarkCallbackEventOK(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
		)

		return MarkOKResponse{Affected: affected}, err
	},
}
