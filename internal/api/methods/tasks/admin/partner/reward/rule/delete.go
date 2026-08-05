package rule

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type DeleteRequest struct {
	WorkspaceID  string `json:"workspace_id"  validate:"required,uuid"`
	Provider     string `json:"provider"      validate:"required"`
	GroupKey     string `json:"group_key"     validate:"required,max=255"`
	ExternalType string `json:"external_type" validate:"required"`
	RewardKey    string `json:"reward_key"    validate:"required,max=255"`
}

type DeleteResponse struct {
	Affected int64 `json:"affected"`
}

var (
	rewardRuleDeleteKey         = "tasks.partner.reward_rule.delete"
	rewardRuleDeleteDescription = `
Deletes a partner reward rule. Requires the
'tasks.partner.reward_rule.delete' permission in the target workspace.`
)

// Delete exposes the partner reward rule deletion method.
var Delete = adapter.Method[DeleteRequest, DeleteResponse]{
	Key:         rewardRuleDeleteKey,
	Description: rewardRuleDeleteDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(rewardRuleDeleteKey),
	},
	Handler: func(ctx *adapter.Context, data DeleteRequest) (DeleteResponse, error) {
		affected, err := services.Tasks.Admin.DeletePartnerRewardRule(
			ctx.Context,
			data.WorkspaceID,
			data.Provider,
			data.GroupKey,
			data.ExternalType,
			data.RewardKey,
		)

		return DeleteResponse{Affected: affected}, err
	},
}
