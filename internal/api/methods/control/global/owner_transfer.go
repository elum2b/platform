package global

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type OwnerTransferRequest struct {
	AccountID string `json:"account_id" validate:"required,uuid"`
}

var (
	ownerTransferKey         = "control.global.owner.transfer"
	ownerTransferDescription = `
Transfers global platform ownership to another account. Requires the
'control.global.owner.transfer' global permission.`
)

// OwnerTransfer transfers global platform ownership.
var OwnerTransfer = adapter.Method[OwnerTransferRequest, struct{}]{
	Key:         ownerTransferKey,
	Description: ownerTransferDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.owner.transfer"),
	},
	Handler: func(ctx *adapter.Context, data OwnerTransferRequest) (struct{}, error) {
		err := services.Control.Admin.TransferGlobalOwnership(
			ctx.Context,
			ctx.AccountID,
			data.AccountID,
		)

		return struct{}{}, err
	},
}
