package user

import (
	refuser "github.com/elum2b/services/reference/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id"     query:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"           query:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"      query:"platform_id"  validate:"required,min=1"`
	Key         string `json:"key"              query:"key"          validate:"required,max=255"`
	Locale      string `json:"locale,omitempty" query:"locale"`
}

type GetResponse struct {
	Item refuser.ItemModel `json:"item"`
}

var (
	getKey         = "reference.user.get"
	getDescription = `
Returns a reference item for the authenticated application user.`
)

// Get exposes the reference user item retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		item, err := services.Reference.User.Get(
			ctx.Context,
			refuser.GetParams{
				WorkspaceID: data.WorkspaceID,
				Key:         data.Key,
				Locale:      data.Locale,
			},
		)

		return GetResponse{Item: item}, err
	},
}
