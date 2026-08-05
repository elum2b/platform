package importapi

import (
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Request struct {
	WorkspaceID      string                 `json:"workspace_id"                validate:"required,uuid"`
	Package          cpaadmin.ExportPackage `json:"package"                     validate:"required"`
	ConflictStrategy string                 `json:"conflict_strategy,omitempty" validate:"omitempty,oneof=fail_on_conflict skip_existing update_existing"`
}

type Response struct {
	Result cpaadmin.ImportResult `json:"result"`
}

var (
	methodKey         = "cpa.import"
	methodDescription = `
Imports CPA offers and configuration into a workspace. Requires the
'cpa.import' permission in the target workspace.`
)

// Method exposes the CPA import method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(methodKey)},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.CPA.Admin.Import(
			ctx.Context,
			data.WorkspaceID,
			cpaadmin.ImportRequest{
				Package:          data.Package,
				ConflictStrategy: data.ConflictStrategy,
			},
		)

		return Response{Result: value}, err
	},
}
