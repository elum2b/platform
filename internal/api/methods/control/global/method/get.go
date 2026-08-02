package method

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	Key string `json:"key" validate:"required"`
}

type GetResponse struct {
	Key       string                   `json:"key"`
	Service   string                   `json:"service"`
	GroupKey  string                   `json:"group_key"`
	Scope     controladmin.AccessScope `json:"scope"`
	Position  int32                    `json:"position"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

var (
	getKey         = "control.global.method.get"
	getDescription = `
Gets a platform method definition by key. Requires the
'control.global.method.get' global permission.`
)

// Get returns a registered method definition.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.method.get"),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		item, err := services.Control.Admin.GetMethod(ctx.Context, data.Key)
		if err != nil {
			return GetResponse{}, err
		}

		return GetResponse{
			Key: item.Key, Service: item.Service, GroupKey: item.GroupKey,
			Scope: item.Scope, Position: item.Position,
			CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt,
		}, nil
	},
}
