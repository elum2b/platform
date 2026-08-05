package price

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	WorkspaceID         string     `json:"workspace_id"           validate:"required,uuid"`
	ProductID           string     `json:"product_id"             validate:"required,max=255"`
	AssetCode           string     `json:"asset_code"             validate:"required,max=255"`
	ListAmountMinor     uint64     `json:"list_amount_minor"      validate:"required,min=0"`
	DiscountAmountMinor uint64     `json:"discount_amount_minor,omitempty"`
	IsPromotion         bool       `json:"is_promotion"`
	StartsAt            *time.Time `json:"starts_at,omitempty"`
	EndsAt              *time.Time `json:"ends_at,omitempty"`
}

type CreateResponse struct {
	ID uint64 `json:"id"`
}

var (
	createKey         = "payment.price.create"
	createDescription = `
Creates a price. Requires the 'payment.price.create'
permission in the target workspace.`
)

var Create = adapter.Method[CreateRequest, CreateResponse]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(createKey)},
	Handler: func(ctx *adapter.Context, d CreateRequest) (CreateResponse, error) {
		id, err := services.Payment.Admin.CreatePrice(
			ctx.Context,
			padm.ProductPriceCreateParams{
				WorkspaceID:         d.WorkspaceID,
				ProductID:           d.ProductID,
				AssetCode:           d.AssetCode,
				ListAmountMinor:     d.ListAmountMinor,
				DiscountAmountMinor: d.DiscountAmountMinor,
				IsPromotion:         d.IsPromotion,
				StartsAt:            d.StartsAt,
				EndsAt:              d.EndsAt,
			})
		return CreateResponse{ID: id}, err
	},
}
