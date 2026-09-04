package user

import (
	puser "github.com/elum2b/services/payment/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetUSDTPriceRequest struct {
	WorkspaceID string `json:"workspace_id" query:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"       query:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"  query:"platform_id"  validate:"required,min=1"`
	AssetCode   string `json:"asset_code"   query:"asset_code"   validate:"required,max=255"`
}

type GetUSDTPriceResponse struct {
	Price puser.USDTPriceModel `json:"price"`
}

var (
	getUSDTPriceKey         = "payment.user.get_usdt_price"
	getUSDTPriceDescription = `
Returns the USDT price for an asset for the authenticated application user.`
)

var GetUSDTPrice = adapter.Method[GetUSDTPriceRequest, GetUSDTPriceResponse]{
	Key:         getUSDTPriceKey,
	Description: getUSDTPriceDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, d GetUSDTPriceRequest) (GetUSDTPriceResponse, error) {
		v, err := services.Payment.User.GetUSDTPrice(
			ctx.Context,
			puser.GetUSDTPriceParams{
				AssetCode: d.AssetCode,
			},
		)
		if err != nil {
			return GetUSDTPriceResponse{}, err
		}

		return GetUSDTPriceResponse{Price: *v}, nil
	},
}
