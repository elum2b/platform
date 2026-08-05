package attempt

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateStatusRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
	Status      string `json:"status"       validate:"required"`
}

var (
	updateStatusKey         = "payment.payment_attempt.update_status"
	updateStatusDescription = `
Updates a payment attempt status. Requires the
'payment.payment_attempt.update_status' permission in the target workspace.`
)

var UpdateStatus = adapter.Method[UpdateStatusRequest, struct{}]{
	Key:         updateStatusKey,
	Description: updateStatusDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(updateStatusKey)},
	Handler: func(ctx *adapter.Context, d UpdateStatusRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.UpdatePaymentAttemptStatus(
			ctx.Context,
			padm.AttemptStatusParams{
				WorkspaceID: d.WorkspaceID,
				ID:          d.ID,
				Status:      d.Status,
			})
	},
}
