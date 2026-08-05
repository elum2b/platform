package report

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID    string `json:"workspace_id"     validate:"required,uuid"`
	AppID          int64  `json:"app_id,omitempty"`
	PlatformID     int64  `json:"platform_id,omitempty"`
	PlatformUserID string `json:"platform_user_id,omitempty"`
	Status         string `json:"status,omitempty"`
	ProductID      string `json:"product_id,omitempty"`
	ProviderCode   string `json:"provider_code,omitempty"`
	AssetCode      string `json:"asset_code,omitempty"`
	CreatedFrom    *int64 `json:"created_from,omitempty"`
	CreatedUntil   *int64 `json:"created_until,omitempty"`
	MinAmountMinor uint64 `json:"min_amount_minor,omitempty"`
	MaxAmountMinor uint64 `json:"max_amount_minor,omitempty"`
	Sort           string `json:"sort,omitempty"`
	Direction      string `json:"direction,omitempty"`
	Limit          int32  `json:"limit,omitempty"      validate:"omitempty,min=1,max=100"`
	Offset         int32  `json:"offset,omitempty"     validate:"min=0"`
}

type GetResponse struct {
	Report padm.PaymentReport `json:"report"`
}

var (
	getKey         = "payment.report.get"
	getDescription = `
Returns a payment report. Requires the 'payment.report.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetPaymentReport(
			ctx.Context,
			padm.PaymentReportParams{
				WorkspaceID:    d.WorkspaceID,
				AppID:          d.AppID,
				PlatformID:     d.PlatformID,
				PlatformUserID: d.PlatformUserID,
				Status:         d.Status,
				ProductID:      d.ProductID,
				ProviderCode:   d.ProviderCode,
				AssetCode:      d.AssetCode,
				CreatedFrom:    int64ToTimePtr(d.CreatedFrom),
				CreatedUntil:   int64ToTimePtr(d.CreatedUntil),
				MinAmountMinor: d.MinAmountMinor,
				MaxAmountMinor: d.MaxAmountMinor,
				Sort:           padm.PaymentSortField(d.Sort),
				Direction:      padm.SortDirection(d.Direction),
				Page:           padm.PageParams{Limit: d.Limit, Offset: d.Offset},
			})
		return GetResponse{Report: v}, err
	},
}

func int64ToTimePtr(ts *int64) *time.Time {
	if ts == nil {
		return nil
	}
	t := time.Unix(*ts, 0)
	return &t
}
