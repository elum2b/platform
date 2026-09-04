package order

import (
	elum2b "github.com/elum2b/services"
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListUserRequest struct {
	WorkspaceID    string `json:"workspace_id"     query:"workspace_id"     validate:"required,uuid"`
	AppID          int64  `json:"app_id"           query:"app_id"           validate:"required,min=1"`
	PlatformID     int64  `json:"platform_id"      query:"platform_id"      validate:"required,min=1"`
	PlatformUserID string `json:"platform_user_id" query:"platform_user_id" validate:"required"`
	Status         string `json:"status,omitempty" query:"status"`
	Limit          int32  `json:"limit,omitempty"  query:"limit"            validate:"omitempty,min=1,max=100"`
	Offset         int32  `json:"offset,omitempty" query:"offset"           validate:"min=0"`
}

type ListUserResponse struct {
	Orders []padm.OrderModel `json:"orders"`
}

var (
	listUserKey         = "payment.order.list_user"
	listUserDescription = `
Lists user orders. Requires the 'payment.order.list_user'
permission in the target workspace.`
)

var ListUser = adapter.Method[ListUserRequest, ListUserResponse]{
	Key:         listUserKey,
	Description: listUserDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listUserKey)},
	Handler: func(ctx *adapter.Context, d ListUserRequest) (ListUserResponse, error) {
		v, err := services.Payment.Admin.ListUserOrders(ctx.Context,
			padm.UserOrderListParams{
				Identity: elum2b.Identity{
					WorkspaceID:    d.WorkspaceID,
					AppID:          d.AppID,
					PlatformID:     d.PlatformID,
					PlatformUserID: d.PlatformUserID,
				},
				Status: d.Status,
				Page:   padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		if err != nil {
			return ListUserResponse{}, err
		}

		return ListUserResponse{Orders: v}, nil
	},
}
