package user

import (
	tuser "github.com/elum2b/services/tasks/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type PartnerCheckRequest struct {
	WorkspaceID string            `json:"workspace_id"     validate:"required,uuid"`
	AppID       int64             `json:"app_id"           validate:"required,min=1"`
	PlatformID  int64             `json:"platform_id"      validate:"required,min=1"`
	Params      string            `json:"params"           validate:"required"`
	IssueRef    string            `json:"issue_ref"        validate:"required"`
	Variables   map[string]string `json:"variables,omitempty"`
}

type PartnerCheckResponse struct {
	Result tuser.PartnerCheckOutput `json:"result"`
}

var (
	partnerCheckKey         = "tasks.user.partner.check"
	partnerCheckDescription = `
Checks the completion status of a partner task for the authenticated
application user.`
)

// PartnerCheck exposes the partner task check method.
var PartnerCheck = adapter.Method[PartnerCheckRequest, PartnerCheckResponse]{
	Key:         partnerCheckKey,
	Description: partnerCheckDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data PartnerCheckRequest) (PartnerCheckResponse, error) {
		result, err := services.Tasks.User.CheckPartner(
			ctx.Context,
			tuser.PartnerCheckParams{
				Identity:  *ctx.Identity,
				IssueRef:  data.IssueRef,
				Variables: data.Variables,
			},
		)

		return PartnerCheckResponse{Result: result}, err
	},
}
