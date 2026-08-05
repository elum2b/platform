package catalog

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateCatalogPriceRequest struct {
	ID                  uint64     `json:"id"                        validate:"required,min=1"`
	WorkspaceID         string     `json:"workspace_id"              validate:"required,uuid"`
	AssetCode           string     `json:"asset_code,omitempty"`
	ListAmountMinor     uint64     `json:"list_amount_minor,omitempty"`
	DiscountAmountMinor uint64     `json:"discount_amount_minor,omitempty"`
	StartsAt            *time.Time `json:"starts_at,omitempty"`
	EndsAt              *time.Time `json:"ends_at,omitempty"`
}

type UpdateCatalogPriceResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updatePriceKey         = "payment.catalog.update_price"
	updatePriceDescription = `
Updates a catalog price. Requires the 'payment.catalog.update_price'
permission in the target workspace.`
)

var UpdatePrice = adapter.Method[UpdateCatalogPriceRequest, UpdateCatalogPriceResponse]{
	Key:         updatePriceKey,
	Description: updatePriceDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(updatePriceKey)},
	Handler: func(ctx *adapter.Context, d UpdateCatalogPriceRequest) (UpdateCatalogPriceResponse, error) {
		a, err := services.Payment.Admin.UpdateCatalogPrice(
			ctx.Context,
			padm.UpdateCatalogPriceParams{
				ID:                  d.ID,
				WorkspaceID:         d.WorkspaceID,
				AssetCode:           d.AssetCode,
				ListAmountMinor:     d.ListAmountMinor,
				DiscountAmountMinor: d.DiscountAmountMinor,
				StartsAt:            d.StartsAt,
				EndsAt:              d.EndsAt,
			})
		return UpdateCatalogPriceResponse{Affected: a}, err
	},
}
