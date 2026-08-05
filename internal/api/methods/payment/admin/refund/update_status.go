package refund

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateStatusRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
	Status      string `json:"status"       validate:"required"`
	Reason      string `json:"reason,omitempty"`
}

type UpdateStatusResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateStatusKey         = "payment.refund.update_status"
	updateStatusDescription = `
Updates a refund status. Requires the 'payment.refund.update_status'
permission in the target workspace.`
)

var UpdateStatus = adapter.Method[UpdateStatusRequest, UpdateStatusResponse]{
	Key:         updateStatusKey,
	Description: updateStatusDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(updateStatusKey)},
	Handler: func(ctx *adapter.Context, d UpdateStatusRequest) (UpdateStatusResponse, error) {
		a, err := services.Payment.Admin.UpdateRefundStatus(
			ctx.Context,
			padm.RefundStatusParams{
				WorkspaceID: d.WorkspaceID,
				ID:          d.ID,
				Status:      d.Status,
				Reason:      d.Reason,
			})
		return UpdateStatusResponse{Affected: a}, err
	},
}
