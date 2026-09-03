package resource

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type OrderRequest struct {
	WorkspaceID      string `json:"workspace_id"                 validate:"required,uuid"`
	ItemKey          string `json:"item_key"                     validate:"required,max=255"`
	ResourceKey      string `json:"resource_key"                 validate:"required,max=255"`
	AfterResourceKey string `json:"after_resource_key,omitempty" validate:"omitempty,max=255"`
}

var (
	insertAfterKey         = "reference.resource.insert_after"
	insertAfterDescription = `
Attaches an unattached resource after another item resource. An empty
'after_resource_key' places it first. Requires the
'reference.resource.insert_after' permission in the target workspace.`
)

// InsertAfter attaches a resource at a semantic position in an item.
var InsertAfter = adapter.Method[OrderRequest, struct{}]{
	Key:         insertAfterKey,
	Description: insertAfterDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(insertAfterKey),
	},
	Handler: func(ctx *adapter.Context, data OrderRequest) (struct{}, error) {
		err := services.Reference.Resource.InsertAfter(
			ctx.Context,
			data.WorkspaceID,
			data.ItemKey,
			data.ResourceKey,
			data.AfterResourceKey,
		)

		return struct{}{}, err
	},
}

var (
	moveAfterKey         = "reference.resource.move_after"
	moveAfterDescription = `
Moves an attached resource after another item resource. An empty
'after_resource_key' places it first. Requires the
'reference.resource.move_after' permission in the target workspace.`
)

// MoveAfter moves a resource to a semantic position in an item.
var MoveAfter = adapter.Method[OrderRequest, struct{}]{
	Key:         moveAfterKey,
	Description: moveAfterDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(moveAfterKey),
	},
	Handler: func(ctx *adapter.Context, data OrderRequest) (struct{}, error) {
		err := services.Reference.Resource.MoveAfter(
			ctx.Context,
			data.WorkspaceID,
			data.ItemKey,
			data.ResourceKey,
			data.AfterResourceKey,
		)

		return struct{}{}, err
	},
}
