package event

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateProcessingStatusRequest struct {
	WorkspaceID string `json:"workspace_id"      validate:"required,uuid"`
	ID          uint64 `json:"id"                validate:"required,min=1"`
	Status      string `json:"status"            validate:"required"`
	Message     string `json:"message,omitempty"`
}

var (
	updateProcessingStatusKey         = "payment.payment_event.update_processing_status"
	updateProcessingStatusDescription = `
Updates a payment event processing status. Requires the
'payment.payment_event.update_processing_status' permission in the target workspace.`
)

var UpdateProcessingStatus = adapter.Method[UpdateProcessingStatusRequest, struct{}]{
	Key:         updateProcessingStatusKey,
	Description: updateProcessingStatusDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(updateProcessingStatusKey),
	},
	Handler: func(ctx *adapter.Context, d UpdateProcessingStatusRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.UpdatePaymentEventProcessingStatus(
			ctx.Context,
			padm.EventStatusParams{
				WorkspaceID: d.WorkspaceID,
				ID:          d.ID,
				Status:      d.Status,
				Message:     d.Message,
			},
		)
	},
}
