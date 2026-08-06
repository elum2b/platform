package user

import (
	"time"

	tuser "github.com/elum2b/services/tasks/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListActiveRequest struct {
	WorkspaceID string    `json:"workspace_id"        validate:"required,uuid"`
	AppID       int64     `json:"app_id"              validate:"required,min=1"`
	PlatformID  int64     `json:"platform_id"         validate:"required,min=1"`
	Params      string    `json:"params"              validate:"required"`
	Locale      string    `json:"locale,omitempty"`
	GroupKey    string    `json:"group_key,omitempty"`
	Now         time.Time `json:"now,omitempty"`
}

type ListActiveResponse struct {
	Groups []tuser.TaskGroupModel `json:"groups"`
}

var (
	listActiveKey         = "tasks.user.list_active"
	listActiveDescription = `
Lists active tasks grouped for the authenticated application user.`
)

// ListActive exposes the tasks user active listing method.
var ListActive = adapter.Method[ListActiveRequest, ListActiveResponse]{
	Key:         listActiveKey,
	Description: listActiveDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data ListActiveRequest) (ListActiveResponse, error) {
		groups, err := services.Tasks.User.ListActive(
			ctx.Context,
			tuser.ListActiveParams{
				Identity: *ctx.Identity,
				Locale:   data.Locale,
				GroupKey: data.GroupKey,
				Now:      data.Now,
			},
		)
		if err != nil {
			return ListActiveResponse{}, err
		}

		return ListActiveResponse{Groups: groups}, nil
	},
}
