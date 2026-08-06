package user

import (
	"time"

	tuser "github.com/elum2b/services/tasks/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type PartnerStartRequest struct {
	WorkspaceID string            `json:"workspace_id"        validate:"required,uuid"`
	AppID       int64             `json:"app_id"              validate:"required,min=1"`
	PlatformID  int64             `json:"platform_id"         validate:"required,min=1"`
	Params      string            `json:"params"              validate:"required"`
	IssueRef    string            `json:"issue_ref"           validate:"required"`
	Variables   map[string]string `json:"variables,omitempty"`
	Now         time.Time         `json:"now,omitempty"`
}

type PartnerStartResponse struct {
	Result tuser.PartnerStartOutput `json:"result"`
}

var (
	partnerStartKey         = "tasks.user.partner.start"
	partnerStartDescription = `
Starts a partner task for the authenticated application user and returns the
deep link.`
)

// PartnerStart exposes the partner task start method.
var PartnerStart = adapter.Method[PartnerStartRequest, PartnerStartResponse]{
	Key:         partnerStartKey,
	Description: partnerStartDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data PartnerStartRequest) (PartnerStartResponse, error) {
		result, err := services.Tasks.User.StartPartner(
			ctx.Context,
			tuser.PartnerStartParams{
				Identity:  *ctx.Identity,
				IssueRef:  data.IssueRef,
				Variables: data.Variables,
				Now:       data.Now,
			},
		)

		return PartnerStartResponse{Result: result}, err
	},
}
