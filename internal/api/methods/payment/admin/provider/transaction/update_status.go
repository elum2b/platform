package transaction

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateStatusRequest struct {
	WorkspaceID string `json:"workspace_id"      validate:"required,uuid"`
	ID          uint64 `json:"id"                validate:"required,min=1"`
	Status      string `json:"status"            validate:"required"`
	Message     string `json:"message,omitempty"`
}

type UpdateStatusResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateStatusKey         = "payment.provider_transaction.update_status"
	updateStatusDescription = `
Updates a provider transaction status. Requires the
'payment.provider_transaction.update_status' permission in the target workspace.`
)

var UpdateStatus = adapter.Method[UpdateStatusRequest, UpdateStatusResponse]{
	Key:         updateStatusKey,
	Description: updateStatusDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(updateStatusKey)},
	Handler: func(ctx *adapter.Context, d UpdateStatusRequest) (UpdateStatusResponse, error) {
		a, err := services.Payment.Admin.UpdateProviderTransactionStatus(
			ctx.Context, d.WorkspaceID, d.ID, d.Status, d.Message)

		return UpdateStatusResponse{Affected: a}, err
	},
}
