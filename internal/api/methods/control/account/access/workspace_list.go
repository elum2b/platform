package access

import (
	controlinternal "github.com/elum2b/services/control/service/internalapi"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type WorkspaceListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type WorkspaceMethodResponse struct {
	Key      string                      `json:"key"`
	Service  string                      `json:"service"`
	GroupKey string                      `json:"group_key"`
	Scope    controlinternal.AccessScope `json:"scope"`
	Position int32                       `json:"position"`
}

type WorkspaceListResponse struct {
	Methods []WorkspaceMethodResponse `json:"methods"`
}

var (
	workspaceListKey         = "control.account.access.workspace.list"
	workspaceListDescription = `
Lists methods available to the current account in a workspace.`
)

// WorkspaceList returns workspace methods authorized for the account.
var WorkspaceList = adapter.Method[WorkspaceListRequest, WorkspaceListResponse]{
	Key:         workspaceListKey,
	Description: workspaceListDescription,
	Transports:  adapter.WS | adapter.MCP,
	Handler: func(
		ctx *adapter.Context,
		data WorkspaceListRequest,
	) (WorkspaceListResponse, error) {
		methods, err := services.Control.Internal.GetAuthorizedWorkspaceMethods(
			ctx.Context,
			ctx.AccountID,
			data.WorkspaceID,
		)
		if err != nil {
			return WorkspaceListResponse{}, err
		}

		response := make([]WorkspaceMethodResponse, 0, len(methods))
		for _, method := range methods {
			response = append(response, WorkspaceMethodResponse{
				Key: method.Key, Service: method.Service,
				GroupKey: method.GroupKey, Scope: method.Scope,
				Position: method.Position,
			})
		}

		return WorkspaceListResponse{Methods: response}, nil
	},
}
