package workspace

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type OwnerTransferRequest struct {
	WorkspaceID     string `json:"workspace_id"      validate:"required,uuid"`
	TargetAccountID string `json:"target_account_id" validate:"required,uuid"`
}

type OwnerTransferResponse struct {
	Transferred bool `json:"transferred"`
}

var (
	ownerTransferKey         = "control.workspace.owner.transfer"
	ownerTransferDescription = `
Transfers workspace ownership to another account. Requires the
'control.workspace.owner.transfer' permission in the target workspace.`
)

// OwnerTransfer transfers workspace ownership to another account.
var OwnerTransfer = adapter.Method[OwnerTransferRequest, OwnerTransferResponse]{
	Key:         ownerTransferKey,
	Description: ownerTransferDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.owner.transfer"),
	},
	Handler: func(
		ctx *adapter.Context,
		data OwnerTransferRequest,
	) (OwnerTransferResponse, error) {
		err := services.Control.Admin.TransferWorkspaceOwnership(
			ctx.Context,
			ctx.AccountID,
			data.WorkspaceID,
			data.TargetAccountID,
		)
		if err != nil {
			return OwnerTransferResponse{}, err
		}

		return OwnerTransferResponse{Transferred: true}, nil
	},
}
