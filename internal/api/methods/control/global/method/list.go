package method

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Item struct {
	Key       string                   `json:"key"`
	Service   string                   `json:"service"`
	GroupKey  string                   `json:"group_key"`
	Scope     controladmin.AccessScope `json:"scope"`
	Position  int32                    `json:"position"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

type ListResponse struct {
	Methods []Item `json:"methods"`
}

var (
	listKey         = "control.global.method.list"
	listDescription = `
Lists all registered platform method definitions. Requires the
'control.global.method.list' global permission.`
)

// List returns registered method definitions.
var List = adapter.Method[struct{}, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.method.list"),
	},
	Handler: func(ctx *adapter.Context, _ struct{}) (ListResponse, error) {
		items, err := services.Control.Admin.ListMethods(ctx.Context)
		if err != nil {
			return ListResponse{}, err
		}

		methods := make([]Item, 0, len(items))
		for _, item := range items {
			methods = append(methods, Item{
				Key: item.Key, Service: item.Service, GroupKey: item.GroupKey,
				Scope: item.Scope, Position: item.Position,
				CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt,
			})
		}

		return ListResponse{Methods: methods}, nil
	},
}
