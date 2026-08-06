package user

import (
	puser "github.com/elum2b/services/payment/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type IsSubscriptionActiveRequest struct {
	WorkspaceID  string `json:"workspace_id"  validate:"required,uuid"`
	AppID        int64  `json:"app_id"        validate:"required,min=1"`
	PlatformID   int64  `json:"platform_id"   validate:"required,min=1"`
	Params       string `json:"params"        validate:"required"`
	ProductID    string `json:"product_id"    validate:"required,max=255"`
	ProviderCode string `json:"provider_code" validate:"required,max=255"`
}

type IsSubscriptionActiveResponse struct {
	Active bool `json:"active"`
}

var (
	isSubscriptionActiveKey         = "payment.user.is_subscription_active"
	isSubscriptionActiveDescription = `
Checks if the authenticated application user has an active subscription.`
)

var IsSubscriptionActive = adapter.Method[IsSubscriptionActiveRequest, IsSubscriptionActiveResponse]{
	Key:         isSubscriptionActiveKey,
	Description: isSubscriptionActiveDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, d IsSubscriptionActiveRequest) (IsSubscriptionActiveResponse, error) {
		active, err := services.Payment.User.IsSubscriptionActive(
			ctx.Context,
			puser.IsSubscriptionActiveParams{
				Identity:     *ctx.Identity,
				ProductID:    d.ProductID,
				ProviderCode: d.ProviderCode,
			},
		)

		return IsSubscriptionActiveResponse{Active: active}, err
	},
}
