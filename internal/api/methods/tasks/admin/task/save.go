package task

import (
	"encoding/json"
	"time"

	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type SaveRequest struct {
	ID                  uint64          `json:"id,omitempty"`
	WorkspaceID         string          `json:"workspace_id"                   validate:"required,uuid"`
	Key                 string          `json:"key"                            validate:"required,max=255"`
	GroupKey            string          `json:"group_key"                      validate:"required,max=255"`
	SequenceKey         *string         `json:"sequence_key,omitempty"`
	SequencePosition    *uint32         `json:"sequence_position,omitempty"`
	TaskKind            string          `json:"task_kind"                      validate:"required"`
	ActionKey           string          `json:"action_key"                     validate:"required"`
	ActionKind          string          `json:"action_kind"                    validate:"required"`
	ClaimMode           string          `json:"claim_mode"                     validate:"required"`
	StartMode           string          `json:"start_mode"                     validate:"required"`
	TargetCount         uint64          `json:"target_count"`
	ResetUnit           string          `json:"reset_unit"                     validate:"required"`
	ResetEvery          uint32          `json:"reset_every"`
	Position            int32           `json:"position"`
	Payload             json.RawMessage `json:"payload,omitempty"`
	Target              json.RawMessage `json:"target,omitempty"`
	IntegrationKind     *string         `json:"integration_kind,omitempty"`
	IntegrationProvider *string         `json:"integration_provider,omitempty"`
	IntegrationPayload  json.RawMessage `json:"integration_payload,omitempty"`
	ImageURL            *string         `json:"image_url,omitempty"`
	IsVisible           bool            `json:"is_visible"`
	IsActive            bool            `json:"is_active"`
	StartAt             *time.Time      `json:"start_at,omitempty"`
	EndAt               *time.Time      `json:"end_at,omitempty"`
}

type SaveResponse struct {
	ID uint64 `json:"id"`
}

var (
	saveKey         = "tasks.task.save"
	saveDescription = `
Creates or updates a task. Requires the 'tasks.task.save' permission in the
target workspace.`
)

// Save exposes the task upsert method.
var Save = adapter.Method[SaveRequest, SaveResponse]{
	Key:         saveKey,
	Description: saveDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(saveKey),
	},
	Handler: func(ctx *adapter.Context, data SaveRequest) (SaveResponse, error) {
		id, err := services.Tasks.Admin.SaveTask(
			ctx.Context,
			tadmin.SaveTaskParams{
				ID:                  data.ID,
				WorkspaceID:         data.WorkspaceID,
				Key:                 data.Key,
				GroupKey:            data.GroupKey,
				SequenceKey:         data.SequenceKey,
				SequencePosition:    data.SequencePosition,
				TaskKind:            data.TaskKind,
				ActionKey:           data.ActionKey,
				ActionKind:          data.ActionKind,
				ClaimMode:           data.ClaimMode,
				StartMode:           data.StartMode,
				TargetCount:         data.TargetCount,
				ResetUnit:           data.ResetUnit,
				ResetEvery:          data.ResetEvery,
				Position:            data.Position,
				Payload:             data.Payload,
				Target:              data.Target,
				IntegrationKind:     data.IntegrationKind,
				IntegrationProvider: data.IntegrationProvider,
				IntegrationPayload:  data.IntegrationPayload,
				ImageURL:            data.ImageURL,
				IsVisible:           data.IsVisible,
				IsActive:            data.IsActive,
				StartAt:             data.StartAt,
				EndAt:               data.EndAt,
			},
		)

		return SaveResponse{ID: id}, err
	},
}
