package role

import (
	"time"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type Item struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Position    int32     `json:"position"`
	MemberCount int64     `json:"member_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListResponse struct {
	Roles []Item `json:"roles"`
}

var (
	listKey         = "control.global.role.list"
	listDescription = `
Lists global platform roles. Requires the
'control.global.role.list' global permission.`
)

// List returns global platform roles.
var List = adapter.Method[struct{}, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.role.list"),
	},
	Handler: func(ctx *adapter.Context, _ struct{}) (ListResponse, error) {
		items, err := services.Control.Admin.ListGlobalRoles(ctx.Context)
		if err != nil {
			return ListResponse{}, err
		}

		roles := make([]Item, 0, len(items))
		for _, item := range items {
			roles = append(roles, Item{
				ID: item.ID, Code: item.Code, Title: item.Title,
				Description: item.Description, Position: item.Position,
				MemberCount: item.MemberCount, CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			})
		}

		return ListResponse{Roles: roles}, nil
	},
}
