package item

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID  string  `json:"workspace_id"            query:"workspace_id"  validate:"required,uuid"`
	ProductID    string  `json:"product_id"              query:"product_id"    validate:"required,max=255"`
	ItemID       string  `json:"item_id"                 query:"item_id"       validate:"required,max=255"`
	RewardType   string  `json:"reward_type"             query:"reward_type"   validate:"required"`
	Quantity     int64   `json:"quantity"                query:"quantity"`
	Scale        uint16  `json:"scale"                   query:"scale"`
	DurationUnit *string `json:"duration_unit,omitempty" query:"duration_unit"`
}

var (
	upsertKey         = "payment.product_item.upsert"
	upsertDescription = `
Creates or updates a product item. Requires the 'payment.product_item.upsert'
permission in the target workspace.`
)

var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(upsertKey)},
	Handler: func(ctx *adapter.Context, d UpsertRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.UpsertProductItem(
			ctx.Context,
			padm.ProductItemUpsertParams{
				WorkspaceID:  d.WorkspaceID,
				ProductID:    d.ProductID,
				ItemID:       d.ItemID,
				RewardType:   d.RewardType,
				Quantity:     d.Quantity,
				Scale:        d.Scale,
				DurationUnit: d.DurationUnit,
			})
	},
}
