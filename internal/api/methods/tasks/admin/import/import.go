package importapi

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Request struct {
	WorkspaceID      string               `json:"workspace_id"                validate:"required,uuid"`
	Package          tadmin.ExportPackage `json:"package"                     validate:"required"`
	ConflictStrategy string               `json:"conflict_strategy,omitempty" validate:"omitempty,oneof=fail_on_conflict skip_existing update_existing"`
}

type Response struct {
	Result tadmin.ImportResult `json:"result"`
}

var (
	methodKey         = "tasks.import"
	methodDescription = `
Imports tasks and configuration into a workspace. Requires the 'tasks.import'
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
		value, err := services.Tasks.Admin.Import(
			ctx.Context,
			data.WorkspaceID,
			tadmin.ImportRequest{
				Package:          data.Package,
				ConflictStrategy: data.ConflictStrategy,
			},
		)

		return Response{Result: value}, err
	},
}
