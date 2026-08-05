package code

import (
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type AddRequest struct {
	WorkspaceID string   `json:"workspace_id" validate:"required,uuid"`
	CPAID       string   `json:"cpa_id"       validate:"required,max=255"`
	Codes       []string `json:"codes"        validate:"required,min=1,dive,required"`
}
type AddResponse struct {
	Added int `json:"added"`
}

var (
	addKey         = "cpa.code.add"
	addDescription = `
Adds personal codes to a CPA offer code pool. Requires the 'cpa.code.add'
permission in the target workspace.`
)

// Add exposes the CPA code addition method.
var Add = adapter.Method[AddRequest, AddResponse]{
	Key:         addKey,
	Description: addDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(addKey)},
	Handler: func(ctx *adapter.Context, data AddRequest) (AddResponse, error) {
		added, err := services.CPA.Admin.AddCodes(
			ctx.Context,
			cpaadmin.AddCodesParams{
				WorkspaceID: data.WorkspaceID,
				CPAID:       data.CPAID,
				Codes:       data.Codes,
			},
		)

		return AddResponse{Added: added}, err
	},
}
