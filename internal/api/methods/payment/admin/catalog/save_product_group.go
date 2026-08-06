package catalog

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type SaveProductGroupRequest struct {
	WorkspaceID    string  `json:"workspace_id"              validate:"required,uuid"`
	Code           string  `json:"code"                      validate:"required,max=255"`
	TitleKey       *string `json:"title_key,omitempty"`
	DescriptionKey *string `json:"description_key,omitempty"`
	Position       int32   `json:"position"`
	IsActive       bool    `json:"is_active"`
}

var (
	saveProductGroupKey         = "payment.catalog.save_product_group"
	saveProductGroupDescription = `
Saves a product group as part of catalog management. Requires the
'payment.catalog.save_product_group' permission in the target workspace.`
)

var SaveProductGroup = adapter.Method[SaveProductGroupRequest, struct{}]{
	Key:         saveProductGroupKey,
	Description: saveProductGroupDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(saveProductGroupKey),
	},
	Handler: func(ctx *adapter.Context, d SaveProductGroupRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.SaveProductGroup(
			ctx.Context,
			padm.SaveProductGroupParams{
				WorkspaceID:    d.WorkspaceID,
				Code:           d.Code,
				TitleKey:       d.TitleKey,
				DescriptionKey: d.DescriptionKey,
				Position:       d.Position,
				IsActive:       d.IsActive,
			})
	},
}
