package key

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateStatusRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
	Status      string `json:"status"       validate:"required"`
}

type UpdateStatusResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateStatusKey         = "payment.purchase_key.update_status"
	updateStatusDescription = `
Updates a purchase key status. Requires the
'payment.purchase_key.update_status' permission in the target workspace.`
)

var UpdateStatus = adapter.Method[UpdateStatusRequest, UpdateStatusResponse]{
	Key:         updateStatusKey,
	Description: updateStatusDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(updateStatusKey)},
	Handler: func(ctx *adapter.Context, d UpdateStatusRequest) (UpdateStatusResponse, error) {
		a, err := services.Payment.Admin.UpdatePurchaseKeyStatus(
			ctx.Context, d.WorkspaceID, d.ID, d.Status)

		return UpdateStatusResponse{Affected: a}, err
	},
}
