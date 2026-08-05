package group

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	Code        string `json:"code"         validate:"required,max=255"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	deleteKey         = "payment.product_group.delete"
	deleteDescription = `
Deletes a product group. Requires the 'payment.product_group.delete'
permission in the target workspace.`
)

var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         deleteKey,
	Description: deleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(deleteKey)},
	Handler: func(ctx *adapter.Context, d DeleteRequest) (DeleteResponse, error) {
		a, err := services.Payment.Admin.DeleteProductGroup(
			ctx.Context, d.WorkspaceID, d.Code)
		return DeleteResponse{Affected: a}, err
	},
}
