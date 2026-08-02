package role

import (
	"time"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	RoleID string `json:"role_id" validate:"required,uuid"`
}

type GetResponse struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Position    int32     `json:"position"`
	MemberCount int64     `json:"member_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var (
	getKey         = "control.global.role.get"
	getDescription = `
Gets a global platform role by ID. Requires the
'control.global.role.get' global permission.`
)

// Get returns a global platform role.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.role.get"),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		item, err := services.Control.Admin.GetGlobalRole(
			ctx.Context,
			data.RoleID,
		)
		if err != nil {
			return GetResponse{}, err
		}

		return GetResponse{
			ID: item.ID, Code: item.Code, Title: item.Title,
			Description: item.Description, Position: item.Position,
			MemberCount: item.MemberCount, CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}, nil
	},
}
