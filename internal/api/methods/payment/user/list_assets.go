package user

import (
	puser "github.com/elum2b/services/payment/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListAssetsRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"  validate:"required,min=1"`
	Params      string `json:"params"       validate:"required"`
}

type ListAssetsResponse struct {
	Assets []puser.AssetModel `json:"assets"`
}

var (
	listAssetsKey         = "payment.user.list_assets"
	listAssetsDescription = `
Lists available assets for the authenticated application user.`
)

var ListAssets = adapter.Method[ListAssetsRequest, ListAssetsResponse]{
	Key:         listAssetsKey,
	Description: listAssetsDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, d ListAssetsRequest) (ListAssetsResponse, error) {
		v, err := services.Payment.User.ListAssets(
			ctx.Context,
			puser.ListAssetsParams{},
		)
		if err != nil {
			return ListAssetsResponse{}, err
		}

		return ListAssetsResponse{Assets: v}, nil
	},
}
