package subscription

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID    string `json:"workspace_id"               query:"workspace_id"     validate:"required,uuid"`
	AppID          int64  `json:"app_id,omitempty"           query:"app_id"`
	PlatformID     int64  `json:"platform_id,omitempty"      query:"platform_id"`
	PlatformUserID string `json:"platform_user_id,omitempty" query:"platform_user_id"`
	ProductID      string `json:"product_id,omitempty"       query:"product_id"`
	ProviderCode   string `json:"provider_code,omitempty"    query:"provider_code"`
	Status         string `json:"status,omitempty"           query:"status"`
	Limit          int32  `json:"limit,omitempty"            query:"limit"            validate:"omitempty,min=1,max=100"`
	Offset         int32  `json:"offset,omitempty"           query:"offset"           validate:"min=0"`
}

type ListResponse struct {
	Subscriptions []padm.SubscriptionModel `json:"subscriptions"`
}

var (
	listKey         = "payment.subscription.list"
	listDescription = `
Lists subscriptions. Requires the 'payment.subscription.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListSubscriptions(ctx.Context,
			padm.SubscriptionListParams{
				WorkspaceID:    d.WorkspaceID,
				AppID:          d.AppID,
				PlatformID:     d.PlatformID,
				PlatformUserID: d.PlatformUserID,
				ProductID:      d.ProductID,
				ProviderCode:   d.ProviderCode,
				Status:         d.Status,
				Page: padm.PageParams{
					Limit:  d.Limit,
					Offset: d.Offset,
				},
			})
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Subscriptions: v}, nil
	},
}
