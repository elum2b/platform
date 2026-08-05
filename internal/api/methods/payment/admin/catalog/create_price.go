package catalog

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateCatalogPriceRequest struct {
	WorkspaceID         string     `json:"workspace_id"            validate:"required,uuid"`
	ProductID           string     `json:"product_id"              validate:"required,max=255"`
	AssetCode           string     `json:"asset_code"              validate:"required,max=255"`
	ListAmountMinor     uint64     `json:"list_amount_minor"       validate:"required,min=0"`
	DiscountAmountMinor uint64     `json:"discount_amount_minor,omitempty"`
	StartsAt            *time.Time `json:"starts_at,omitempty"`
	EndsAt              *time.Time `json:"ends_at,omitempty"`
}

type CreateCatalogPriceResponse struct {
	ID uint64 `json:"id"`
}

var (
	createPriceKey         = "payment.catalog.create_price"
	createPriceDescription = `
Creates a catalog price. Requires the 'payment.catalog.create_price'
permission in the target workspace.`
)

var CreatePrice = adapter.Method[CreateCatalogPriceRequest, CreateCatalogPriceResponse]{
	Key:         createPriceKey,
	Description: createPriceDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(createPriceKey)},
	Handler: func(ctx *adapter.Context, d CreateCatalogPriceRequest) (CreateCatalogPriceResponse, error) {
		id, err := services.Payment.Admin.CreateCatalogPrice(
			ctx.Context,
			padm.CreateCatalogPriceParams{
				WorkspaceID:         d.WorkspaceID,
				ProductID:           d.ProductID,
				AssetCode:           d.AssetCode,
				ListAmountMinor:     d.ListAmountMinor,
				DiscountAmountMinor: d.DiscountAmountMinor,
				StartsAt:            d.StartsAt,
				EndsAt:              d.EndsAt,
			})
		return CreateCatalogPriceResponse{ID: id}, err
	},
}
