package wallet

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type GetResponse struct {
	Wallet padm.TONWalletModel `json:"wallet"`
}

var (
	getKey         = "payment.ton_wallet.get"
	getDescription = `
Returns a TON wallet configuration. Requires the 'payment.ton_wallet.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetTONWallet(
			ctx.Context, d.WorkspaceID)

		return GetResponse{Wallet: v}, err
	},
}
