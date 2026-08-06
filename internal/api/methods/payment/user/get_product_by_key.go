package user

import (
	puser "github.com/elum2b/services/payment/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetProductByKeyRequest struct {
	WorkspaceID string `json:"workspace_id"         validate:"required,uuid"`
	AppID       int64  `json:"app_id"               validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"          validate:"required,min=1"`
	Params      string `json:"params"               validate:"required"`
	Key         string `json:"key"                  validate:"required,max=255"`
	AssetCode   string `json:"asset_code,omitempty"`
	Locale      string `json:"locale,omitempty"`
}

type GetProductByKeyResponse struct {
	Product puser.ProductModel `json:"product"`
}

var (
	getProductByKeyKey         = "payment.user.get_product_by_key"
	getProductByKeyDescription = `
Returns a product by key for the authenticated application user.`
)

var GetProductByKey = adapter.Method[GetProductByKeyRequest, GetProductByKeyResponse]{
	Key:         getProductByKeyKey,
	Description: getProductByKeyDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, d GetProductByKeyRequest) (GetProductByKeyResponse, error) {
		v, err := services.Payment.User.GetProductByKey(
			ctx.Context,
			puser.GetProductByKeyParams{
				Key:       d.Key,
				AssetCode: d.AssetCode,
				Locale:    d.Locale,
			},
		)
		if err != nil {
			return GetProductByKeyResponse{}, err
		}

		return GetProductByKeyResponse{Product: *v}, nil
	},
}
