package item

import (
	"encoding/json"

	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	WorkspaceID string          `json:"workspace_id" validate:"required,uuid"`
	Key         string          `json:"key"          validate:"required,max=255"`
	Type        string          `json:"type"         validate:"required"`
	Payload     json.RawMessage `json:"payload"      validate:"required"`
	IsActive    bool            `json:"is_active"`
}

var (
	createKey         = "reference.create"
	createDescription = `
Creates a new reference item. Requires the 'reference.create' permission in
the target workspace.`
)

// Create exposes the reference item creation method.
var Create = adapter.Method[CreateRequest, struct{}]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(createKey),
	},
	Handler: func(ctx *adapter.Context, data CreateRequest) (struct{}, error) {
		err := services.Reference.Admin.CreateItem(
			ctx.Context,
			refadmin.SaveItemParams{
				WorkspaceID: data.WorkspaceID,
				Key:         data.Key,
				Type:        data.Type,
				Payload:     data.Payload,
				IsActive:    data.IsActive,
			},
		)

		return struct{}{}, err
	},
}
