package reward

import (
	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpdateRequest struct {
	WorkspaceID string  `json:"workspace_id"   validate:"required,uuid"`
	CalendarID  string  `json:"calendar_id"    validate:"required,max=255"`
	StepID      uint64  `json:"step_id"        validate:"required,min=1"`
	ID          uint64  `json:"id"             validate:"required,min=1"`
	Key         string  `json:"key"            validate:"required,max=255"`
	Type        string  `json:"type"           validate:"required,max=255"`
	Quantity    int64   `json:"quantity"`
	Scale       uint16  `json:"scale"`
	Unit        *string `json:"unit,omitempty"`
	Position    uint32  `json:"position"       validate:"required,min=1"`
}

type UpdateResponse struct {
	Affected int64 `json:"affected"`
}

var (
	updateKey         = "calendar.reward.update"
	updateDescription = `
Updates a reward on a calendar step. Requires the 'calendar.reward.update'
permission in the target workspace.`
)

// Update exposes the calendar reward update method.
var Update = adapter.Method[UpdateRequest, UpdateResponse]{
	Key:         updateKey,
	Description: updateDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(updateKey),
	},
	Handler: func(ctx *adapter.Context, data UpdateRequest) (UpdateResponse, error) {
		affected, err := services.Calendar.Admin.UpdateReward(
			ctx.Context,
			caladmin.SaveRewardParams{
				WorkspaceID: data.WorkspaceID,
				CalendarID:  data.CalendarID,
				StepID:      data.StepID,
				ID:          data.ID,
				Key:         data.Key,
				Type:        data.Type,
				Quantity:    data.Quantity,
				Scale:       data.Scale,
				Unit:        data.Unit,
				Position:    data.Position,
			},
		)

		return UpdateResponse{Affected: affected}, err
	},
}
