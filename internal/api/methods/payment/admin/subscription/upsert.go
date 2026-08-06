package subscription

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID            string  `json:"workspace_id"               validate:"required,uuid"`
	AppID                  int64   `json:"app_id"                     validate:"required,min=1"`
	PlatformID             int64   `json:"platform_id"                validate:"required,min=1"`
	PlatformUserID         string  `json:"platform_user_id"           validate:"required"`
	InternalUserID         *int64  `json:"internal_user_id,omitempty"`
	ProductID              string  `json:"product_id"                 validate:"required,max=255"`
	OrderID                *int64  `json:"order_id,omitempty"`
	AttemptID              *int64  `json:"attempt_id,omitempty"`
	ProviderCode           string  `json:"provider_code"              validate:"required,max=255"`
	ProviderSubscriptionID string  `json:"provider_subscription_id"   validate:"required,max=255"`
	Status                 string  `json:"status"                     validate:"required"`
	CancelReason           *string `json:"cancel_reason,omitempty"`
	StartedAt              int64   `json:"started_at"                 validate:"required,min=0"`
	EndedAt                *int64  `json:"ended_at,omitempty"`
}

type UpsertResponse struct {
	ID uint64 `json:"id"`
}

var (
	upsertKey         = "payment.subscription.upsert"
	upsertDescription = `
Creates or updates a subscription. Requires the 'payment.subscription.upsert'
permission in the target workspace.`
)

var Upsert = adapter.Method[UpsertRequest, UpsertResponse]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(upsertKey)},
	Handler: func(ctx *adapter.Context, d UpsertRequest) (UpsertResponse, error) {
		id, err := services.Payment.Admin.UpsertSubscription(
			ctx.Context,
			padm.SubscriptionUpsertParams{
				WorkspaceID:            d.WorkspaceID,
				AppID:                  d.AppID,
				PlatformID:             d.PlatformID,
				PlatformUserID:         d.PlatformUserID,
				InternalUserID:         d.InternalUserID,
				ProductID:              d.ProductID,
				OrderID:                d.OrderID,
				AttemptID:              d.AttemptID,
				ProviderCode:           d.ProviderCode,
				ProviderSubscriptionID: d.ProviderSubscriptionID,
				Status:                 d.Status,
				CancelReason:           d.CancelReason,
				StartedAt:              time.Unix(d.StartedAt, 0),
				EndedAt:                int64ToTime(d.EndedAt),
			})

		return UpsertResponse{ID: id}, err
	},
}

func int64ToTime(ts *int64) *time.Time {
	if ts == nil {
		return nil
	}

	t := time.Unix(*ts, 0)

	return &t
}
