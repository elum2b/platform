package price

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateRequest struct {
	ID                  uint64     `json:"id"                              validate:"required,min=1"`
	WorkspaceID         string     `json:"workspace_id"                    validate:"required,uuid"`
	AssetCode           string     `json:"asset_code,omitempty"`
	ListAmountMinor     uint64     `json:"list_amount_minor,omitempty"`
	DiscountAmountMinor uint64     `json:"discount_amount_minor,omitempty"`
	IsPromotion         bool       `json:"is_promotion,omitempty"`
	StartsAt            *time.Time `json:"starts_at,omitempty"`
	EndsAt              *time.Time `json:"ends_at,omitempty"`
}

type UpdateResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateKey         = "payment.price.update"
	updateDescription = `
Updates a price. Requires the 'payment.price.update'
permission in the target workspace.`
)

var Update = adapter.Method[UpdateRequest, UpdateResponse]{
	Key:         updateKey,
	Description: updateDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(updateKey)},
	Handler: func(ctx *adapter.Context, d UpdateRequest) (UpdateResponse, error) {
		a, err := services.Payment.Admin.UpdatePrice(
			ctx.Context,
			padm.ProductPriceUpdateParams{
				ID:                  d.ID,
				WorkspaceID:         d.WorkspaceID,
				AssetCode:           d.AssetCode,
				ListAmountMinor:     d.ListAmountMinor,
				DiscountAmountMinor: d.DiscountAmountMinor,
				IsPromotion:         d.IsPromotion,
				StartsAt:            d.StartsAt,
				EndsAt:              d.EndsAt,
			})

		return UpdateResponse{Affected: a}, err
	},
}
