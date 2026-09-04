package user

import (
	"time"

	tuser "github.com/elum2b/services/tasks/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ClaimRequest struct {
	WorkspaceID string    `json:"workspace_id"  query:"workspace_id" validate:"required,uuid"`
	AppID       int64     `json:"app_id"        query:"app_id"       validate:"required,min=1"`
	PlatformID  int64     `json:"platform_id"   query:"platform_id"  validate:"required,min=1"`
	TaskRef     string    `json:"task_ref"      query:"task_ref"     validate:"required"`
	OperationID string    `json:"operation_id"  query:"operation_id" validate:"required"`
	Now         time.Time `json:"now,omitempty" query:"now"`
}

type ClaimResponse struct {
	Result tuser.ClaimResult `json:"result"`
}

var (
	claimKey         = "tasks.user.claim"
	claimDescription = `
Claims a completed task for the authenticated application user.`
)

// Claim exposes the tasks user claim method.
var Claim = adapter.Method[ClaimRequest, ClaimResponse]{
	Key:         claimKey,
	Description: claimDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data ClaimRequest) (ClaimResponse, error) {
		result, err := services.Tasks.User.Claim(
			ctx.Context,
			tuser.ClaimParams{
				Identity:    *ctx.Identity,
				TaskRef:     data.TaskRef,
				OperationID: data.OperationID,
				Now:         data.Now,
			},
		)

		return ClaimResponse{Result: result}, err
	},
}
