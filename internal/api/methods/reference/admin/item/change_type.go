package item

import (
	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ChangeTypeRequest struct {
	WorkspaceID  string `json:"workspace_id" validate:"required,uuid"`
	Key          string `json:"key"          validate:"required,max=255"`
	CurrentType  string `json:"current_type" validate:"required"`
	NewType      string `json:"new_type"     validate:"required"`
	Confirmation string `json:"confirmation" validate:"required"`
}

type ChangeTypeResponse struct {
	Affected int64 `json:"affected"`
}

var (
	changeTypeKey         = "reference.change_type"
	changeTypeDescription = `
Changes the type of a reference item. Requires the 'reference.change_type'
permission in the target workspace. The confirmation field must equal
'CHANGE_REFERENCE_TYPE'.`
)

// ChangeType exposes the reference item type change method.
var ChangeType = adapter.Method[ChangeTypeRequest, ChangeTypeResponse]{
	Key:         changeTypeKey,
	Description: changeTypeDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(changeTypeKey),
	},
	Handler: func(ctx *adapter.Context, data ChangeTypeRequest) (ChangeTypeResponse, error) {
		affected, err := services.Reference.Admin.DangerousChangeType(
			ctx.Context,
			refadmin.DangerousChangeTypeParams{
				WorkspaceID:  data.WorkspaceID,
				Key:          data.Key,
				CurrentType:  data.CurrentType,
				NewType:      data.NewType,
				Confirmation: data.Confirmation,
			},
		)

		return ChangeTypeResponse{Affected: affected}, err
	},
}
