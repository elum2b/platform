package calendar

import (
	"time"

	caladmin "github.com/elum2b/services/calendar/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID         string     `json:"workspace_id"          validate:"required,uuid"`
	ID                  string     `json:"id"                    validate:"required,max=255"`
	Type                string     `json:"type"                  validate:"required"`
	Mode                string     `json:"mode"                  validate:"required"`
	IntervalType        string     `json:"interval_type"         validate:"required"`
	IntervalUnit        string     `json:"interval_unit"         validate:"required"`
	IntervalCount       uint32     `json:"interval_count"`
	ResetAfterIntervals uint32     `json:"reset_after_intervals"`
	EndBehavior         string     `json:"end_behavior"          validate:"required"`
	Timezone            string     `json:"timezone"              validate:"required"`
	HideFutureRewards   bool       `json:"hide_future_rewards"`
	IsActive            bool       `json:"is_active"`
	StartAt             *time.Time `json:"start_at,omitempty"`
	EndAt               *time.Time `json:"end_at,omitempty"`
}

type UpsertResponse struct {
	ID       string `json:"id,omitempty"`
	Affected int64  `json:"affected,omitempty"`
}

var (
	upsertKey         = "calendar.upsert"
	upsertDescription = `
Creates or updates a calendar. Requires the 'calendar.upsert' permission in
the target workspace.`
)

// Upsert exposes the calendar upsert method.
var Upsert = adapter.Method[UpsertRequest, UpsertResponse]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (UpsertResponse, error) {
		params := caladmin.SaveCalendarParams{
			WorkspaceID:         data.WorkspaceID,
			ID:                  data.ID,
			Type:                data.Type,
			Mode:                data.Mode,
			IntervalType:        data.IntervalType,
			IntervalUnit:        data.IntervalUnit,
			IntervalCount:       data.IntervalCount,
			ResetAfterIntervals: data.ResetAfterIntervals,
			EndBehavior:         data.EndBehavior,
			Timezone:            data.Timezone,
			HideFutureRewards:   data.HideFutureRewards,
			IsActive:            data.IsActive,
			StartAt:             data.StartAt,
			EndAt:               data.EndAt,
		}

		id, err := services.Calendar.Admin.CreateCalendar(
			ctx.Context,
			params,
		)
		if err != nil {
			affected, uErr := services.Calendar.Admin.UpdateCalendar(
				ctx.Context,
				params,
			)
			if uErr != nil {
				return UpsertResponse{}, err
			}

			return UpsertResponse{Affected: affected}, nil
		}

		return UpsertResponse{ID: id}, nil
	},
}
