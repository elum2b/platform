package subscription

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetByProviderIDRequest struct {
	WorkspaceID            string `json:"workspace_id"            validate:"required,uuid"`
	ProviderCode           string `json:"provider_code"           validate:"required,max=255"`
	ProviderSubscriptionID string `json:"provider_subscription_id" validate:"required,max=255"`
}

type GetByProviderIDResponse struct {
	Subscription padm.SubscriptionModel `json:"subscription"`
}

var (
	getByProviderIDKey         = "payment.subscription.get_by_provider_id"
	getByProviderIDDescription = `
Returns a subscription by provider ID. Requires the
'payment.subscription.get_by_provider_id' permission in the target workspace.`
)

var GetByProviderID = adapter.Method[GetByProviderIDRequest, GetByProviderIDResponse]{
	Key:         getByProviderIDKey,
	Description: getByProviderIDDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getByProviderIDKey)},
	Handler: func(ctx *adapter.Context, d GetByProviderIDRequest) (GetByProviderIDResponse, error) {
		v, err := services.Payment.Admin.GetSubscriptionByProviderID(
			ctx.Context,
			padm.SubscriptionProviderRefParams{
				WorkspaceID:            d.WorkspaceID,
				ProviderCode:           d.ProviderCode,
				ProviderSubscriptionID: d.ProviderSubscriptionID,
			})
		return GetByProviderIDResponse{Subscription: v}, err
	},
}
