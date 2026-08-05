package callback

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type MarkRejectRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
	Reason      string `json:"reason"       validate:"required"`
}

type MarkRejectResponse struct {
	Affected int64 `json:"affected"`
}

var (
	markRejectKey         = "promo.callback.mark_reject"
	markRejectDescription = `
Rejects a promo callback event with a reason. Requires the
'promo.callback.mark_reject' permission in the target workspace.`
)

// MarkReject exposes the promo callback rejection method.
var MarkReject = adapter.Method[MarkRejectRequest, MarkRejectResponse]{
	Key:         markRejectKey,
	Description: markRejectDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(markRejectKey),
	},
	Handler: func(ctx *adapter.Context, data MarkRejectRequest) (MarkRejectResponse, error) {
		affected, err := services.Promo.Admin.MarkCallbackEventReject(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
			data.Reason,
		)

		return MarkRejectResponse{Affected: affected}, err
	},
}
