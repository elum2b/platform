package group

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID    string  `json:"workspace_id"              validate:"required,uuid"`
	Code           string  `json:"code"                      validate:"required,max=255"`
	TitleKey       *string `json:"title_key,omitempty"`
	DescriptionKey *string `json:"description_key,omitempty"`
	Position       int32   `json:"position"`
	IsActive       bool    `json:"is_active"`
}

var (
	upsertKey         = "payment.product_group.upsert"
	upsertDescription = `
Creates or updates a product group. Requires the
'payment.product_group.upsert' permission in the target workspace.`
)

var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(upsertKey)},
	Handler: func(ctx *adapter.Context, d UpsertRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.UpsertProductGroup(
			ctx.Context,
			padm.ProductGroupUpsertParams{
				WorkspaceID: d.WorkspaceID, Code: d.Code,
				TitleKey: d.TitleKey, DescriptionKey: d.DescriptionKey,
				Position: d.Position, IsActive: d.IsActive,
			},
		)
	},
}
