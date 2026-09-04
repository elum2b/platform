package user

import (
	refuser "github.com/elum2b/services/reference/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string `json:"workspace_id"     query:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"           query:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"      query:"platform_id"  validate:"required,min=1"`
	Locale      string `json:"locale,omitempty" query:"locale"`
	Limit       int32  `json:"limit,omitempty"  query:"limit"        validate:"omitempty,min=1,max=100"`
	Offset      int32  `json:"offset,omitempty" query:"offset"       validate:"min=0"`
}

type ListResponse struct {
	Items []refuser.ItemModel `json:"items"`
}

var (
	listKey         = "reference.user.list"
	listDescription = `
Lists reference items for the authenticated application user.`
)

// List exposes the reference user item listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Reference.User.List(
			ctx.Context,
			refuser.ListParams{
				WorkspaceID: data.WorkspaceID,
				Locale:      data.Locale,
				Page: refuser.Page{
					Limit:  data.Limit,
					Offset: data.Offset,
				},
			},
		)
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{Items: items}, nil
	},
}
