package access

import (
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	Locale string                   `json:"locale,omitempty" validate:"max=32"`
	Scope  controladmin.AccessScope `json:"scope,omitempty"  validate:"omitempty,oneof=global workspace"`
}

type Access struct {
	Key         string                   `json:"key"`
	Scope       controladmin.AccessScope `json:"scope"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
}

type Group struct {
	Key         string   `json:"key"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Accesses    []Access `json:"accesses"`
}

type Item struct {
	Service     string  `json:"service"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Groups      []Group `json:"groups"`
}

type ListResponse struct {
	Services []Item `json:"services"`
}

var (
	listKey         = "control.global.access.list"
	listDescription = `
Lists the localized access catalog for global or workspace methods. Requires the
'control.global.access.list' global permission.`
)

// List returns the access catalog.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.access.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListAccess(
			ctx.Context,
			data.Locale,
			data.Scope,
		)
		if err != nil {
			return ListResponse{}, err
		}

		response := make([]Item, 0, len(items))
		for _, item := range items {
			groups := make([]Group, 0, len(item.Groups))
			for _, group := range item.Groups {
				accesses := make([]Access, 0, len(group.Accesses))
				for _, access := range group.Accesses {
					accesses = append(accesses, Access{
						Key: access.Key, Scope: access.Scope,
						Title: access.Title, Description: access.Desc,
					})
				}

				groups = append(groups, Group{
					Key: group.Key, Title: group.Title,
					Description: group.Description, Accesses: accesses,
				})
			}

			response = append(response, Item{
				Service: item.Service, Title: item.Title,
				Description: item.Description, Groups: groups,
			})
		}

		return ListResponse{Services: response}, nil
	},
}
