package importapi

import (
	"bytes"

	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/archive"
)

type Request struct {
	WorkspaceID      string            `json:"workspace_id"                validate:"required,uuid"`
	Archive          []byte            `json:"archive"                     validate:"required"`
	ConflictStrategy string            `json:"conflict_strategy,omitempty" validate:"omitempty,oneof=fail_on_conflict skip_existing update_existing"`
	Secrets          map[string]string `json:"secrets,omitempty"`
}

type Response struct {
	Job tadmin.ArchiveJob `json:"job"`
}

var (
	methodKey         = "tasks.import"
	methodDescription = `
Queues an archive import of tasks and configuration into a workspace. Requires the 'tasks.import'
permission in the target workspace.`
)

// Method exposes the tasks import method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Tasks.Admin.QueueArchiveImport(
			ctx.Context,
			tadmin.QueueArchiveImportParams{
				WorkspaceID: data.WorkspaceID,
				FileName:    archive.FileName("tasks"),
				ImportRequest: tadmin.ImportRequest{
					ConflictStrategy: data.ConflictStrategy,
					Secrets:          data.Secrets,
				},
				Archive: bytes.NewReader(data.Archive),
			},
		)

		return Response{Job: value}, err
	},
}
