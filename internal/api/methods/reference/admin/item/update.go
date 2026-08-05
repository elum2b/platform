package item

import (
	"encoding/json"

	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateRequest struct {
	WorkspaceID string          `json:"workspace_id" validate:"required,uuid"`
	Key         string          `json:"key"          validate:"required,max=255"`
	Payload     json.RawMessage `json:"payload"      validate:"required"`
	IsActive    bool            `json:"is_active"`
}

type UpdateResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateKey         = "reference.update"
	updateDescription = `
Updates a reference item's payload and active state. Requires the
'reference.update' permission in the target workspace.`
)

// Update exposes the reference item update method.
var Update = adapter.Method[UpdateRequest, UpdateResponse]{
	Key:         updateKey,
	Description: updateDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(updateKey),
	},
	Handler: func(ctx *adapter.Context, data UpdateRequest) (UpdateResponse, error) {
		affected, err := services.Reference.Admin.UpdateItem(
			ctx.Context,
			refadmin.UpdateItemParams{
				WorkspaceID: data.WorkspaceID,
				Key:         data.Key,
				Payload:     data.Payload,
				IsActive:    data.IsActive,
			},
		)

		return UpdateResponse{Affected: affected}, err
	},
}
