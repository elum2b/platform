package counter

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID    string `json:"workspace_id"               query:"workspace_id"     validate:"required,uuid"`
	AppID          int64  `json:"app_id,omitempty"           query:"app_id"`
	ProductID      string `json:"product_id,omitempty"       query:"product_id"`
	PlatformID     int64  `json:"platform_id,omitempty"      query:"platform_id"`
	PlatformUserID string `json:"platform_user_id,omitempty" query:"platform_user_id"`
	Limit          int32  `json:"limit,omitempty"            query:"limit"            validate:"omitempty,min=1,max=100"`
	Offset         int32  `json:"offset,omitempty"           query:"offset"           validate:"min=0"`
}

type ListResponse struct {
	Counters []padm.ProductLimitCounterModel `json:"counters"`
}

var (
	listKey         = "payment.product_limit_counter.list"
	listDescription = `
Lists product limit counters. Requires the
'payment.product_limit_counter.list' permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, d ListRequest) (ListResponse, error) {
		v, err := services.Payment.Admin.ListProductLimitCounters(ctx.Context,
			padm.ProductLimitCounterListParams{
				WorkspaceID:    d.WorkspaceID,
				AppID:          d.AppID,
				ProductID:      d.ProductID,
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

		return ListResponse{Counters: v}, nil
	},
}
