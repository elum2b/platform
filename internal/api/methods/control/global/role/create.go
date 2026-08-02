package role

import (
	"time"

	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	ID          string `json:"id,omitempty"          validate:"omitempty,uuid"`
	Code        string `json:"code"                  validate:"required,max=255"`
	Title       string `json:"title"                 validate:"required,max=255"`
	Description string `json:"description,omitempty" validate:"max=1000"`
	Position    int32  `json:"position,omitempty"`
}

type CreateResponse struct {
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
	createKey         = "control.global.role.create"
	createDescription = `
Creates a global platform role. Requires the
'control.global.role.create' global permission.`
)

// Create creates a global platform role.
var Create = adapter.Method[CreateRequest, CreateResponse]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.role.create"),
	},
	Handler: func(ctx *adapter.Context, data CreateRequest) (CreateResponse, error) {
		item, err := services.Control.Admin.CreateGlobalRole(
			ctx.Context,
			controladmin.CreateRoleParams{
				ActorID: ctx.AccountID, ID: data.ID, Code: data.Code,
				Title: data.Title, Description: data.Description,
				Position: data.Position,
			},
		)
		if err != nil {
			return CreateResponse{}, err
		}

		return CreateResponse{
			ID: item.ID, Code: item.Code, Title: item.Title,
			Description: item.Description, Position: item.Position,
			MemberCount: item.MemberCount, CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}, nil
	},
}
