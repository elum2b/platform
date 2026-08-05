package assignment

import (
	serviceapi "github.com/elum2b/services"
	cpaadmin "github.com/elum2b/services/cpa/service/admin"
	cpauser "github.com/elum2b/services/cpa/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID    string `json:"workspace_id"       validate:"required,uuid"`
	CPAID          string `json:"cpa_id"             validate:"required,max=255"`
	AppID          int64  `json:"app_id"             validate:"required,min=1"`
	PlatformID     int64  `json:"platform_id"        validate:"required,min=1"`
	Platform       string `json:"platform,omitempty"`
	PlatformUserID string `json:"platform_user_id"   validate:"required"`
}
type GetResponse struct {
	Assignment *cpaadmin.AssignmentModel `json:"assignment"`
}

var (
	getKey         = "cpa.assignment.get"
	getDescription = `
Returns a user's assignment for a CPA offer. Requires the 'cpa.assignment.get'
permission in the target workspace.`
)

// Get exposes the CPA assignment retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		value, err := services.CPA.Admin.GetUserAssignment(
			ctx.Context,
			cpauser.GetStatusParams{
				Identity: serviceapi.Identity{
					WorkspaceID:    data.WorkspaceID,
					AppID:          data.AppID,
					PlatformID:     data.PlatformID,
					Platform:       data.Platform,
					PlatformUserID: data.PlatformUserID,
				},
				CPAID: data.CPAID,
			},
		)

		return GetResponse{Assignment: value}, err
	},
}
