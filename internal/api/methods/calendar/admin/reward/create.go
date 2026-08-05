package reward

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateRequest struct {
	WorkspaceID string  `json:"workspace_id" validate:"required,uuid"`
	CalendarID  string  `json:"calendar_id"  validate:"required,max=255"`
	StepID      uint64  `json:"step_id"      validate:"required,min=1"`
	Key         string  `json:"key"          validate:"required,max=255"`
	Type        string  `json:"type"         validate:"required,max=255"`
	Quantity    int64   `json:"quantity"`
	Scale       uint16  `json:"scale"`
	Unit        *string `json:"unit,omitempty"`
	Position    uint32  `json:"position"     validate:"required,min=1"`
}

type CreateResponse struct {
	ID uint64 `json:"id"`
}

var (
	createKey         = "calendar.reward.create"
	createDescription = `
Creates a reward on a calendar step. Requires the 'calendar.reward.create'
permission in the target workspace.`
)

// Create exposes the calendar reward creation method.
var Create = adapter.Method[CreateRequest, CreateResponse]{
	Key:         createKey,
	Description: createDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(createKey),
	},
	Handler: func(ctx *adapter.Context, data CreateRequest) (CreateResponse, error) {
		id, err := services.Calendar.Admin.CreateReward(
			ctx.Context,
			caladmin.SaveRewardParams{
				WorkspaceID: data.WorkspaceID,
				CalendarID:  data.CalendarID,
				StepID:      data.StepID,
				Key:         data.Key,
				Type:        data.Type,
				Quantity:    data.Quantity,
				Scale:       data.Scale,
				Unit:        data.Unit,
				Position:    data.Position,
			},
		)

		return CreateResponse{ID: id}, err
	},
}
