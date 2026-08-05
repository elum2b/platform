package item

import (
	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID    string `json:"workspace_id"               validate:"required,uuid"`
	Type           string `json:"type,omitempty"`
	OnlyNotDeleted bool   `json:"only_not_deleted,omitempty"`
	Limit          int32  `json:"limit,omitempty"            validate:"omitempty,min=1,max=100"`
	Offset         int32  `json:"offset,omitempty"           validate:"min=0"`
}

type ListResponse struct {
	Items []refadmin.ItemModel `json:"items"`
}

var (
	listKey         = "reference.list"
	listDescription = `
Lists reference items in a workspace, with optional type and deletion filter.
Requires the 'reference.list' permission in the target workspace.`
)

// List exposes the reference item listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(listKey),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Reference.Admin.ListItems(
			ctx.Context,
			refadmin.ItemListParams{
				WorkspaceID:    data.WorkspaceID,
				Type:           data.Type,
				OnlyNotDeleted: data.OnlyNotDeleted,
				Page: refadmin.Page{
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
