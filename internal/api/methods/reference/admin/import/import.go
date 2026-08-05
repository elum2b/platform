package importapi

import (
	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Request struct {
	WorkspaceID      string                 `json:"workspace_id"                validate:"required,uuid"`
	Package          refadmin.ExportPackage `json:"package"                     validate:"required"`
	ConflictStrategy string                 `json:"conflict_strategy,omitempty" validate:"omitempty,oneof=fail_on_conflict skip_existing update_existing"`
}

type Response struct {
	Result refadmin.ImportResult `json:"result"`
}

var (
	methodKey         = "reference.import"
	methodDescription = `
Imports reference items and localizations into a workspace. Requires the
'reference.import' permission in the target workspace.`
)

// Method exposes the reference import method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Reference.Admin.Import(
			ctx.Context,
			data.WorkspaceID,
			refadmin.ImportRequest{
				Package:          data.Package,
				ConflictStrategy: data.ConflictStrategy,
			},
		)

		return Response{Result: value}, err
	},
}
