package status

import (
	cpauser "github.com/elum2b/services/cpa/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" query:"workspace_id" validate:"required,uuid"`
	AppID       int64  `json:"app_id"       query:"app_id"       validate:"required,min=1"`
	PlatformID  int64  `json:"platform_id"  query:"platform_id"  validate:"required,min=1"`
	CPAID       string `json:"cpa_id"       query:"cpa_id"       validate:"required,max=255"`
}

type GetResponse struct {
	Assignment *cpauser.AssignmentModel `json:"assignment"`
}

var (
	getKey         = "cpa.user.status.get"
	getDescription = `
Returns the authenticated application user's assignment for a CPA offer.`
)

// Get exposes the CPA user assignment status method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		assignment, err := services.CPA.User.GetStatus(
			ctx.Context,
			cpauser.GetStatusParams{
				Identity: *ctx.Identity,
				CPAID:    data.CPAID,
			},
		)
		if err != nil {
			return GetResponse{}, err
		}

		return GetResponse{Assignment: assignment}, nil
	},
}
