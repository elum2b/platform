package catalog

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type AttachProductItemRequest struct {
	WorkspaceID  string  `json:"workspace_id"            validate:"required,uuid"`
	ProductID    string  `json:"product_id"              validate:"required,max=255"`
	ItemID       string  `json:"item_id"                 validate:"required,max=255"`
	RewardType   string  `json:"reward_type"             validate:"required"`
	Quantity     int64   `json:"quantity"`
	Scale        uint16  `json:"scale"`
	DurationUnit *string `json:"duration_unit,omitempty"`
}

var (
	attachProductItemKey         = "payment.catalog.attach_product_item"
	attachProductItemDescription = `
Attaches a product item as part of catalog management. Requires the
'payment.catalog.attach_product_item' permission in the target workspace.`
)

var AttachProductItem = adapter.Method[AttachProductItemRequest, struct{}]{
	Key:         attachProductItemKey,
	Description: attachProductItemDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(attachProductItemKey),
	},
	Handler: func(ctx *adapter.Context, d AttachProductItemRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.AttachProductItem(
			ctx.Context,
			padm.AttachProductItemParams{
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
