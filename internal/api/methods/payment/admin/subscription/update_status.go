package subscription

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateStatusRequest struct {
	WorkspaceID            string  `json:"workspace_id"             query:"workspace_id"             validate:"required,uuid"`
	ProviderCode           string  `json:"provider_code"            query:"provider_code"            validate:"required,max=255"`
	ProviderSubscriptionID string  `json:"provider_subscription_id" query:"provider_subscription_id" validate:"required,max=255"`
	Status                 string  `json:"status"                   query:"status"                   validate:"required"`
	CancelReason           *string `json:"cancel_reason,omitempty"  query:"cancel_reason"`
	EndedAt                *int64  `json:"ended_at,omitempty"       query:"ended_at"`
}

type UpdateStatusResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateStatusKey         = "payment.subscription.update_status"
	updateStatusDescription = `
Updates a subscription status. Requires the
'payment.subscription.update_status' permission in the target workspace.`
)

var UpdateStatus = adapter.Method[UpdateStatusRequest, UpdateStatusResponse]{
	Key:         updateStatusKey,
	Description: updateStatusDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(updateStatusKey)},
	Handler: func(ctx *adapter.Context, d UpdateStatusRequest) (UpdateStatusResponse, error) {
		a, err := services.Payment.Admin.UpdateSubscriptionStatus(
			ctx.Context,
			padm.SubscriptionStatusUpdateParams{
				WorkspaceID:            d.WorkspaceID,
				ProviderCode:           d.ProviderCode,
				ProviderSubscriptionID: d.ProviderSubscriptionID,
				Status:                 d.Status,
				CancelReason:           d.CancelReason,
				EndedAt:                int64ToTime(d.EndedAt),
			})

		return UpdateStatusResponse{Affected: a}, err
	},
}
