package offer

import (
	cpauser "github.com/elum2b/services/cpa/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"     query:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"           query:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"      query:"platform_id"  validate:"required,min=1"`
	Locale      string `json:"locale,omitempty" query:"locale"`
}

type ListResponse struct {
	Offers []cpauser.OfferModel `json:"offers"`
}

var (
	listKey         = "cpa.user.offer.list"
	listDescription = `
Lists active CPA offers available to the authenticated application user.`
)

// List exposes the CPA user offer listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		offers, err := services.CPA.User.ListActive(
			ctx.Context,
			cpauser.ListActiveParams{
				Identity: *ctx.Identity,
				Locale:   data.Locale,
			},
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Offers: offers}, nil
	},
}
