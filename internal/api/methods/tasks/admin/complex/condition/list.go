package condition

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type ListResponse struct {
	Conditions []tadmin.ComplexConditionModel `json:"conditions"`
}

var (
	listKey         = "tasks.complex_condition.list"
	listDescription = `
Lists complex task conditions in a workspace. Requires the
'tasks.complex_condition.list' permission in the target workspace.`
)

// List exposes the complex condition listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Tasks.Admin.ListComplexConditions(
			ctx.Context,
			data.WorkspaceID,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Conditions: values}, nil
	},
}
