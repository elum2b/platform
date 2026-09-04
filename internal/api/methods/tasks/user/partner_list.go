package user

import (
	tuser "github.com/elum2b/services/tasks/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type PartnerListRequest struct {
	WorkspaceID string            `json:"workspace_id"        query:"workspace_id" validate:"required,uuid"`
	AppID       int64             `json:"app_id"              query:"app_id"       validate:"required,min=1"`
	PlatformID  int64             `json:"platform_id"         query:"platform_id"  validate:"required,min=1"`
	Provider    string            `json:"provider"            query:"provider"     validate:"required"`
	GroupKey    string            `json:"group_key"           query:"group_key"    validate:"required,max=255"`
	Platform    string            `json:"platform"            query:"platform"     validate:"required"`
	Locale      string            `json:"locale,omitempty"    query:"locale"`
	Limit       int32             `json:"limit,omitempty"     query:"limit"        validate:"omitempty,min=1,max=100"`
	Variables   map[string]string `json:"variables,omitempty" query:"variables"`
}

type PartnerListResponse struct {
	Tasks []tuser.TaskModel `json:"tasks"`
}

var (
	partnerListKey         = "tasks.user.partner.list"
	partnerListDescription = `
Lists partner tasks for the authenticated application user.`
)

// PartnerList exposes the partner task listing method.
var PartnerList = adapter.Method[PartnerListRequest, PartnerListResponse]{
	Key:         partnerListKey,
	Description: partnerListDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data PartnerListRequest) (PartnerListResponse, error) {
		tasks, err := services.Tasks.User.ListPartner(
			ctx.Context,
			tuser.PartnerListParams{
				Identity:  *ctx.Identity,
				Provider:  data.Provider,
				GroupKey:  data.GroupKey,
				Platform:  data.Platform,
				Locale:    data.Locale,
				Limit:     data.Limit,
				Variables: data.Variables,
			},
		)
		if err != nil {
			return PartnerListResponse{}, err
		}

		return PartnerListResponse{Tasks: tasks}, nil
	},
}
