package counter

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID    string `json:"workspace_id"               validate:"required,uuid"`
	AppID          int64  `json:"app_id"                     validate:"required,min=1"`
	PlatformID     int64  `json:"platform_id"                validate:"required,min=1"`
	ProductID      string `json:"product_id"                 validate:"required,max=255"`
	CounterScope   string `json:"counter_scope"              validate:"required,oneof=global user"`
	PlatformUserID string `json:"platform_user_id,omitempty"`
	WindowStart    int64  `json:"window_start"               validate:"required,min=0"`
	WindowEnd      int64  `json:"window_end"                 validate:"required,min=0"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "payment.product_limit_counter.delete"
	deleteDescription = `
Deletes product limit counters. Requires the
'payment.product_limit_counter.delete' permission in the target workspace.`
)

var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(deleteKey)},
	Handler: func(ctx *adapter.Context, d DeleteRequest) (DeleteResponse, error) {
		a, err := services.Payment.Admin.DeleteProductLimitCounter(
			ctx.Context,
			padm.ProductLimitCounterDeleteParams{
				WorkspaceID:    d.WorkspaceID,
				AppID:          d.AppID,
				PlatformID:     d.PlatformID,
				ProductID:      d.ProductID,
				CounterScope:   d.CounterScope,
				PlatformUserID: d.PlatformUserID,
				WindowStart:    time.Unix(d.WindowStart, 0),
				WindowEnd:      time.Unix(d.WindowEnd, 0),
			})

		return DeleteResponse{Affected: a}, err
	},
}
