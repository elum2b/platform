package user

import (
	"time"

	tuser "github.com/elum2b/services/tasks/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type StartRequest struct {
	WorkspaceID string    `json:"workspace_id"  query:"workspace_id" validate:"required,uuid"`
	AppID       int64     `json:"app_id"        query:"app_id"       validate:"required,min=1"`
	PlatformID  int64     `json:"platform_id"   query:"platform_id"  validate:"required,min=1"`
	TaskRef     string    `json:"task_ref"      query:"task_ref"     validate:"required"`
	Now         time.Time `json:"now,omitempty" query:"now"`
}

type StartResponse struct {
	Result tuser.StartTaskResult `json:"result"`
}

var (
	startKey         = "tasks.user.start"
	startDescription = `
Starts a task for the authenticated application user.`
)

// Start exposes the tasks user start method.
var Start = adapter.Method[StartRequest, StartResponse]{
	Key:         startKey,
	Description: startDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data StartRequest) (StartResponse, error) {
		result, err := services.Tasks.User.StartTask(
			ctx.Context,
			tuser.StartTaskParams{
				Identity: *ctx.Identity,
				TaskRef:  data.TaskRef,
				Now:      data.Now,
			},
		)

		return StartResponse{Result: result}, err
	},
}
