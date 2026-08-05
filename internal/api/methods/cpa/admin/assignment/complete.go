package assignment

import (
	serviceapi "github.com/elum2b/services"
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CompleteRequest struct {
	WorkspaceID    string `json:"workspace_id"       validate:"required,uuid"`
	CPAID          string `json:"cpa_id"             validate:"required,max=255"`
	AppID          int64  `json:"app_id"             validate:"required,min=1"`
	PlatformID     int64  `json:"platform_id"        validate:"required,min=1"`
	Platform       string `json:"platform,omitempty"`
	PlatformUserID string `json:"platform_user_id"   validate:"required"`
}
type CompleteResponse struct {
	Result cpaadmin.CompleteResult `json:"result"`
}

var (
	completeKey         = "cpa.assignment.complete"
	completeDescription = `
Completes a user's CPA assignment and returns its rewards. Requires the
'cpa.assignment.complete' permission in the target workspace.`
)

// Complete exposes the CPA assignment completion method.
var Complete = adapter.Method[CompleteRequest, CompleteResponse]{
	Key:         completeKey,
	Description: completeDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(completeKey)},
	Handler: func(ctx *adapter.Context, data CompleteRequest) (CompleteResponse, error) {
		result, err := services.CPA.Admin.Complete(
			ctx.Context,
			cpaadmin.CompleteParams{
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

		return CompleteResponse{Result: result}, err
	},
}
