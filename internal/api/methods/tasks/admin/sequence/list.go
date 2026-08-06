package sequence

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type ListResponse struct {
	Sequences []tadmin.SequenceModel `json:"sequences"`
}

var (
	listKey         = "tasks.sequence.list"
	listDescription = `
Lists task sequences in a workspace. Requires the 'tasks.sequence.list'
permission in the target workspace.`
)

// List exposes the task sequence listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		seqs, err := services.Tasks.Admin.ListSequences(
			ctx.Context,
			data.WorkspaceID,
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Sequences: seqs}, nil
	},
}
