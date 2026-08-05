package importapi

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Request struct {
	WorkspaceID      string             `json:"workspace_id"                validate:"required,uuid"`
	Package          padm.ExportPackage `json:"package"                     validate:"required"`
	ConflictStrategy string             `json:"conflict_strategy,omitempty" validate:"omitempty,oneof=fail_on_conflict skip_existing update_existing"`
}

type Response struct {
	Result padm.ImportResult `json:"result"`
}

var (
	methodKey         = "payment.import"
	methodDescription = `
Imports payment configuration into a workspace. Requires the 'payment.import'
permission in the target workspace.`
)

var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Payment.Admin.Import(
			ctx.Context,
			data.WorkspaceID,
			padm.ImportRequest{
				Package:          data.Package,
				ConflictStrategy: data.ConflictStrategy,
			},
		)

		return Response{Result: value}, err
	},
}
