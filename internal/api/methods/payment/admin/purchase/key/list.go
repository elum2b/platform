package key

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID    string `json:"workspace_id"               query:"workspace_id"     validate:"required,uuid"`
	AppID          int64  `json:"app_id,omitempty"           query:"app_id"`
	ProductID      string `json:"product_id,omitempty"       query:"product_id"`
	Status         string `json:"status,omitempty"           query:"status"`
	PlatformID     int64  `json:"platform_id,omitempty"      query:"platform_id"`
	PlatformUserID string `json:"platform_user_id,omitempty" query:"platform_user_id"`
	Limit          int32  `json:"limit,omitempty"            query:"limit"            validate:"omitempty,min=1,max=100"`
	Offset         int32  `json:"offset,omitempty"           query:"offset"           validate:"min=0"`
}

type ListResponse struct {
	Keys []padm.PurchaseKeyModel `json:"keys"`
}

var (
	listKey         = "payment.purchase_key.list"
	listDescription = `
Lists purchase keys. Requires the 'payment.purchase_key.list'
permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListPurchaseKeys(ctx.Context,
			padm.PurchaseKeyListParams{
				WorkspaceID:    d.WorkspaceID,
				AppID:          d.AppID,
				ProductID:      d.ProductID,
				Status:         d.Status,
				PlatformID:     d.PlatformID,
				PlatformUserID: d.PlatformUserID,
				Page: padm.PageParams{
					Limit:  d.Limit,
					Offset: d.Offset,
				},
			})
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Keys: v}, nil
	},
}
