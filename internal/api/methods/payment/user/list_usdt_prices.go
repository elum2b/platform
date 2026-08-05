package user

import (
	puser "github.com/elum2b/services/payment/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListUSDTPricesRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"  validate:"required,min=1"`
	Params      string `json:"params"       validate:"required"`
}

type ListUSDTPricesResponse struct {
	Prices []puser.USDTPriceModel `json:"prices"`
}

var (
	listUSDTPricesKey         = "payment.user.list_usdt_prices"
	listUSDTPricesDescription = `
Lists USDT prices for the authenticated application user.`
)

var ListUSDTPrices = adapter.Method[ListUSDTPricesRequest, ListUSDTPricesResponse]{
	Key:         listUSDTPricesKey,
	Description: listUSDTPricesDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, d ListUSDTPricesRequest) (ListUSDTPricesResponse, error) {
		v, err := services.Payment.User.ListUSDTPrices(
			ctx.Context,
			puser.ListUSDTPricesParams{},
		)
		if err != nil {
			return ListUSDTPricesResponse{}, err
		}

		return ListUSDTPricesResponse{Prices: v}, nil
	},
}
