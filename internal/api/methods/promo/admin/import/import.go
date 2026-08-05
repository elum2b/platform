package importapi

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Request struct {
	WorkspaceID      string                   `json:"workspace_id"                validate:"required,uuid"`
	Package          promoadmin.ExportPackage `json:"package"                     validate:"required"`
	ConflictStrategy string                   `json:"conflict_strategy,omitempty" validate:"omitempty,oneof=fail_on_conflict skip_existing update_existing"`
}

type Response struct {
	Result promoadmin.ImportResult `json:"result"`
}

var (
	methodKey         = "promo.import"
	methodDescription = `
Imports promos and configuration into a workspace. Requires the 'promo.import'
permission in the target workspace.`
)

// Method exposes the promo import method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Promo.Admin.Import(
			ctx.Context,
			data.WorkspaceID,
			promoadmin.ImportRequest{
				Package:          data.Package,
				ConflictStrategy: data.ConflictStrategy,
			},
		)

		return Response{Result: value}, err
	},
}
